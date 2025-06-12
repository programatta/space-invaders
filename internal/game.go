package internal

import (
	"fmt"
	"image/color"
	"math/rand"
	"slices"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/programatta/spaceinvaders/internal/common"
	"github.com/programatta/spaceinvaders/internal/config"
	"github.com/programatta/spaceinvaders/internal/enemy"
	"github.com/programatta/spaceinvaders/internal/explosion"
	"github.com/programatta/spaceinvaders/internal/player"
	"github.com/programatta/spaceinvaders/internal/sounds"
	"github.com/programatta/spaceinvaders/internal/sprite"
	"github.com/programatta/spaceinvaders/internal/utils"
)

type Manageer interface {
	Update()
	Draw(screen *ebiten.Image)
}

type Eraser interface {
	CanRemove() bool
}

type Explosioner interface {
	Manageer
	Eraser
}

type Collider interface {
	Rect() (float32, float32, float32, float32)
	OnCollide()
}

type Game struct {
	spriteCreator     *sprite.SpriteCreator
	textFace          *text.GoTextFace
	soundEffects      *sounds.SoundEffects
	cannon            *player.Cannon
	cannonCount       uint8
	score             uint32
	bullets           []*Bullet
	bunkers           []*player.Bunker
	ufo               *enemy.Ufo
	enemies           []*enemy.Alien
	enemiesCurrentDir float32
	newDirection      float32
	explosions        []Explosioner
	alienFireTime     float32
	innerStateId      playInnerStateId
}

func NewGame() *Game {
	spriteCreator := sprite.NewSpriteCreator()
	textFace := utils.LoadEmbeddedFont(8)
	soundEffects := sounds.NewSoundEffects()

	game := &Game{}
	game.spriteCreator = spriteCreator
	game.textFace = textFace

	spriteCannon, _ := spriteCreator.SpriteByName("cannon")
	game.cannon = player.NewCannon(float32(0), float32(config.DesignHeight-10), spriteCannon, game)

	bunkers := createBunkers(spriteCreator)
	game.bunkers = bunkers

	ufoSprite, _ := spriteCreator.SpriteByName("ufo")
	ufo := enemy.NewUfo(-20, 15, ufoSprite)
	game.ufo = ufo

	enemies := createEnemies(spriteCreator, game)
	game.enemies = enemies
	game.enemiesCurrentDir = 1
	game.newDirection = 1
	game.cannonCount = 3
	game.score = 0
	game.innerStateId = playing
	game.soundEffects = soundEffects
	return game
}

// Implementación de la interface esperada por ebiten.
func (g *Game) Update() error {
	switch g.innerStateId {
	case playing:
		g.processKeyEventPlaying()
		g.updatePlaying()
	case gameOver:
		g.processKeyEventGameOver()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x03, 0x04, 0x5e, 0xFF})

	switch g.innerStateId {
	case playing:
		g.drawPlaying(screen)
	case gameOver:
		g.drawGameOver(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return config.DesignWidth, config.DesignHeight
}

// Implementación de la interface Notifier
func (g *Game) OnCreateCannonBullet(posX, posY float32, color color.Color) {
	spriteBullet, _ := g.spriteCreator.SpriteByName("bullet")
	bullet := NewBullet(posX, posY, spriteBullet, color, -1)
	g.bullets = append(g.bullets, bullet)
	g.soundEffects.PlayShoot()
}

func (g *Game) OnCreateAlienBullet(posX, posY float32, color color.Color) {
	spriteBullet, _ := g.spriteCreator.SpriteByName("bullet")
	bullet := NewBullet(posX, posY, spriteBullet, color, 1)
	g.bullets = append(g.bullets, bullet)
}

func (g *Game) OnChangeDirection(newDirection float32) {
	g.newDirection = newDirection
}

func (g *Game) OnResetUfo() {
	g.ufo.Reset()
}

func (g *Game) OnResetCannon() {
	if g.cannonCount == 0 {
		g.innerStateId = gameOver
	}
	g.cannon.Reset()
}

// -----------------------------------------------------------------------------
// Sección de procesamiento de eventos por estado.
// -----------------------------------------------------------------------------

func (g *Game) processKeyEventPlaying() {
	g.cannon.ProcessKeyEvents()
}

func (g *Game) processKeyEventGameOver() {
	if inpututil.IsKeyJustReleased(ebiten.KeySpace) {
		// resetear el juego.
		g.reset()
		g.cannon.Reset()
		g.cannonCount = 3
		g.score = 0
		g.innerStateId = playing
	}
}

// -----------------------------------------------------------------------------
// Sección de actualización por estado.
// -----------------------------------------------------------------------------

func (g *Game) updatePlaying() {
	g.cannon.Update()

	g.alienFireTime += config.Dt
	if g.alienFireTime > 0.400 {
		if len(g.enemies) > 0 {
			pos := rand.Intn(len(g.enemies))
			g.enemies[pos].Fire()
		}
		g.alienFireTime = 0
	}

	for _, bullet := range g.bullets {
		bullet.Update()
	}

	g.ufo.Update()
	ufoX, _ := g.ufo.Position()
	if ufoX >= 5 {
		g.soundEffects.PlayUfo()
	}

	var enemyIncrementSpeed float32 = 0
	if len(g.enemies) == 1 {
		enemyIncrementSpeed = 7
	} else if 2 <= len(g.enemies) && len(g.enemies) <= 5 {
		enemyIncrementSpeed = 5
	} else if 6 <= len(g.enemies) && len(g.enemies) <= 10 {
		enemyIncrementSpeed = 3
	}

	for _, enemy := range g.enemies {
		enemy.IncrementSpeed(enemyIncrementSpeed)
		enemy.ChangeDirection(g.enemiesCurrentDir)
		enemy.Update()
	}

	for _, explosion := range g.explosions {
		explosion.Update()
	}

	//Colisiones.
	for _, bullet := range g.bullets {
		if bullet.dirY < 0 {
			//Bala de cañon
			for _, bunker := range g.bunkers {
				if g.checkCollision(bullet, bunker) {
					if bunker.DoDamage(bullet.posX, bullet.posY, -1) {
						bullet.OnCollide()
					}
				}
			}
			for _, enemy := range g.enemies {
				if g.checkCollision(bullet, enemy) {
					bullet.OnCollide()
					enemy.OnCollide()

					alienExplosionSprite, _ := g.spriteCreator.SpriteByName("alienExplosion")
					enemyX, enemyY := enemy.Position()
					explosion := explosion.NewExplosion(enemyX, enemyY, alienExplosionSprite, enemy.Color())
					g.explosions = append(g.explosions, explosion)
					g.score += uint32(enemy.Score())
					g.soundEffects.PlayAlienKilled()
				}
			}
			if g.ufo.IsActive() && g.checkCollision(bullet, g.ufo) {
				bullet.OnCollide()
				g.ufo.OnCollide()

				ufoExplosionSprite, _ := g.spriteCreator.SpriteByName("ufoExplosion")
				ufoX, ufoY := g.ufo.Position()
				ufoScore := g.ufo.Score()
				explosionUfo := explosion.NewExplosionUfo(ufoX, ufoY, ufoExplosionSprite, g.textFace, ufoScore, g)
				g.explosions = append(g.explosions, explosionUfo)
				g.score += uint32(ufoScore)
			}
		} else {
			//Bala de alien.
			for _, bunker := range g.bunkers {
				if g.checkCollision(bullet, bunker) {
					if bunker.DoDamage(bullet.posX, bullet.posY, 1) {
						bullet.OnCollide()
					}
				}
			}
			if g.cannon.IsActive() && g.checkCollision(bullet, g.cannon) {
				cannonExplosion1Sprite, _ := g.spriteCreator.SpriteByName("cannonExplosion1")
				cannonExplosion2Sprite, _ := g.spriteCreator.SpriteByName("cannonExplosion2")
				posX, posY := g.cannon.Position()
				explosionCannon := explosion.NewExplosionCannon(posX, posY, cannonExplosion1Sprite, cannonExplosion2Sprite, g)
				g.explosions = append(g.explosions, explosionCannon)
				if g.cannonCount > 0 {
					g.cannonCount--
					g.cannon.OnCollide()
				}
				bullet.OnCollide()
				g.soundEffects.PlayCannonExplosion()
			}
		}
	}

	//Colisines alien con bunker
	for _, enemy := range g.enemies {
		for _, bunker := range g.bunkers {
			if g.checkCollision(enemy, bunker) {
				bunker.OnCollide()
				break
			}
		}
	}

	if len(g.bullets) > 0 {
		g.bullets = slices.DeleteFunc(g.bullets, func(bullet *Bullet) bool {
			return bullet.CanRemove()
		})
	}

	if len(g.enemies) > 0 {
		g.enemies = slices.DeleteFunc(g.enemies, func(alien *enemy.Alien) bool {
			return alien.CanRemove()
		})
	} else {
		g.reset()
		g.cannon.Reset()
	}

	if len(g.explosions) > 0 {
		g.explosions = slices.DeleteFunc(g.explosions, func(explosion Explosioner) bool {
			return explosion.CanRemove()
		})
	}

	if len(g.bunkers) > 0 {
		g.bunkers = slices.DeleteFunc(g.bunkers, func(bunker *player.Bunker) bool {
			return bunker.CanRemove()
		})
	}

	if g.newDirection != g.enemiesCurrentDir {
		g.enemiesCurrentDir = g.newDirection
	}
}

// -----------------------------------------------------------------------------
// Sección de dibujo de pantalla por estado.
// -----------------------------------------------------------------------------

func (g *Game) drawPlaying(screen *ebiten.Image) {
	uiCannonCountText := fmt.Sprintf("LIVES:%1d", g.cannonCount)
	op := &text.DrawOptions{}
	op.GeoM.Translate(10, 6)
	op.ColorScale.ScaleWithColor(color.White)
	text.Draw(screen, uiCannonCountText, g.textFace, op)

	uiScoreText := fmt.Sprintf("SCORE:%06d", g.score)
	widthText, _ := text.Measure(uiScoreText, g.textFace, 0)
	op = &text.DrawOptions{}
	op.GeoM.Translate(float64(config.DesignWidth)-widthText-10, 6)
	op.ColorScale.ScaleWithColor(color.White)
	text.Draw(screen, uiScoreText, g.textFace, op)

	g.ufo.Draw(screen)

	for _, enemy := range g.enemies {
		enemy.Draw(screen)
	}

	for _, explosion := range g.explosions {
		explosion.Draw(screen)
	}

	g.cannon.Draw(screen)
	for _, bullet := range g.bullets {
		bullet.Draw(screen)
	}

	for _, bunker := range g.bunkers {
		bunker.Draw(screen)
	}

}

func (g *Game) drawGameOver(screen *ebiten.Image) {
	uiGameOverText := "GAME OVER"
	widthText, _ := text.Measure(uiGameOverText, g.textFace, 0)
	gameOverX := float64(config.DesignWidth/2) - widthText/2
	gameOverY := float64(config.DesignHeight/2 - 30)

	op := &text.DrawOptions{}
	op.GeoM.Translate(gameOverX, gameOverY)
	op.ColorScale.ScaleWithColor(color.White)
	text.Draw(screen, uiGameOverText, g.textFace, op)

	uiYourScoreText := fmt.Sprintf("YOUR SCORE:%06d", g.score)
	widthText, _ = text.Measure(uiYourScoreText, g.textFace, 0)
	yourScoreX := float64(config.DesignWidth/2) - widthText/2
	yourScoreY := float64(config.DesignHeight/2 - 15)

	op = &text.DrawOptions{}
	op.GeoM.Translate(yourScoreX, yourScoreY)
	op.ColorScale.ScaleWithColor(color.White)
	text.Draw(screen, uiYourScoreText, g.textFace, op)
}

func (g *Game) checkCollision(sourceObj, targetObj Collider) bool {
	sx0, sy0, sw, sh := sourceObj.Rect()
	tx0, ty0, tw, th := targetObj.Rect()

	hasCollision := sx0 < tx0+tw && sx0+sw > tx0 && sy0 < ty0+th && sh+sy0 > ty0

	return hasCollision
}

func (g *Game) reset() {
	bunkers := createBunkers(g.spriteCreator)
	enemies := createEnemies(g.spriteCreator, g)

	g.enemies = enemies
	g.bullets = []*Bullet{}
	g.bunkers = bunkers
	g.explosions = []Explosioner{}
	g.enemiesCurrentDir = 1
	g.newDirection = 1
	g.alienFireTime = 0
}

func createBunkers(spriteCreator *sprite.SpriteCreator) []*player.Bunker {
	bunkerSprite, _ := spriteCreator.SpriteByName("bunker")

	var posX float32 = 27
	bunkers := []*player.Bunker{}
	for range 4 {
		bunker := player.NewBunker(posX, float32(config.DesignHeight-40), bunkerSprite)
		bunkers = append(bunkers, bunker)
		posX += float32(bunkerSprite.Image.Bounds().Dx()) + 20
	}
	return bunkers
}

func createEnemies(spriteCreator *sprite.SpriteCreator, notifier common.Notifier) []*enemy.Alien {
	enemies := []*enemy.Alien{}

	squids := createAlienFormation("squid", 11, 1, 9, 5, 30, 11, 35, spriteCreator, notifier)
	enemies = append(enemies, squids...)

	crabs := createAlienFormation("crab", 11, 2, 6, 5, 20, 10, 50, spriteCreator, notifier)
	enemies = append(enemies, crabs...)

	octopuses := createAlienFormation("octopus", 11, 2, 5, 5, 10, 9, 80, spriteCreator, notifier)
	enemies = append(enemies, octopuses...)
	return enemies
}

func createAlienFormation(alienName string, count, rows, offsetX, offsetY, points uint8, initX, initY float32, spriteCreator *sprite.SpriteCreator, notifier common.Notifier) []*enemy.Alien {
	sprite1, _ := spriteCreator.SpriteByName(fmt.Sprintf("%s1", alienName))
	sprite2, _ := spriteCreator.SpriteByName(fmt.Sprintf("%s2", alienName))
	aliens := []*enemy.Alien{}

	posX := initX
	posY := initY
	for i := range count * rows {
		alien := enemy.NewAlien(posX, posY, sprite1, sprite2, points, config.AlienMoveDelay, notifier)
		aliens = append(aliens, alien)
		posX += float32(sprite1.Image.Bounds().Dx() + int(offsetX))
		if i > 0 && (i+1)%count == 0 {
			posX = initX
			posY += float32(sprite1.Image.Bounds().Dy() + int(offsetY))
		}
	}
	return aliens
}
