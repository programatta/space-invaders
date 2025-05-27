package play

import (
	"fmt"
	"image/color"
	"math/rand"
	"os"
	"slices"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/programatta/spaceinvaders/internal/config"
	"github.com/programatta/spaceinvaders/internal/sound"
	"github.com/programatta/spaceinvaders/internal/sprite"
	"github.com/programatta/spaceinvaders/internal/states"
	"github.com/programatta/spaceinvaders/internal/states/play/common"
	"github.com/programatta/spaceinvaders/internal/states/play/enemy"
	"github.com/programatta/spaceinvaders/internal/states/play/explosion"
	"github.com/programatta/spaceinvaders/internal/states/play/player"
)

type PlayState struct {
	spriteCreator     *sprite.SpriteCreator
	soundEffects      *sound.SoundEffects
	nextState         states.StateId
	textFace          *text.GoTextFace
	ufo               *enemy.Ufo
	enemies           []*enemy.Alien
	bullets           []*Bullet
	bunkers           []*player.Bunker
	explosions        []explosion.Explosioner
	cannon            *player.Cannon
	cannonCount       uint8
	score             uint32
	enemiesCurrentDir float32
	newDirection      float32
	alienFireTime     float32
	gameOverTime      float32
	level             *Level
	innerStateId      playInnerStateId
	pauseScreen       *ebiten.Image
}

func NewPlayState(spriteCreator *sprite.SpriteCreator, soundEffects *sound.SoundEffects, textFace *text.GoTextFace) *PlayState {
	playState := &PlayState{}

	//obtenemos los sprites
	ufoSprite, _ := spriteCreator.SpriteByName("ufo")
	cannonSprite, _ := spriteCreator.SpriteByName("cannon")

	ufo := enemy.NewUfo(-20, 15, ufoSprite)
	cannon := player.NewCannon(float32(0), float32(config.DesignHeight-10), cannonSprite, playState)

	playState.spriteCreator = spriteCreator
	playState.soundEffects = soundEffects
	playState.textFace = textFace
	playState.ufo = ufo
	playState.cannon = cannon
	playState.pauseScreen = nil

	return playState
}

/*
Implementación interface State
*/
func (ps *PlayState) Start() {
	ps.nextState = states.Play
	ps.cannon.Reset()
	ps.cannonCount = 3
	ps.score = 0
	ps.gameOverTime = 0
	ps.innerStateId = starting
	ps.level = NewLevel()
	ps.reset()
}

func (ps *PlayState) ProcessEvents() {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		os.Exit(0)
	}

	switch ps.innerStateId {
	case playing:
		ps.processKeyEventPlaying()
	case pause:
		ps.processKeyEventPause()
	}
}

func (ps *PlayState) Update() {
	switch ps.innerStateId {
	case starting:
		ps.innerStateId = playing
	case playing:
		ps.updatePlaying()
	case pauseRequest:
		ps.updatePause()
	case gameOver:
		ps.updateGameOver()
	}
}

func (ps *PlayState) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x03, 0x04, 0x5e, 0xFF})

	switch ps.innerStateId {
	case playing:
		fallthrough
	case pauseRequest:
		ps.drawPlaying(screen)
		if ps.innerStateId == pauseRequest {
			ps.pauseScreen = ebiten.NewImageFromImage(screen)
		}
	case pause:
		ps.drawPause(screen, ps.pauseScreen)
	case gameOver:
		ps.drawGameOver(screen)
	}
}

func (ps *PlayState) NextState() states.StateId {
	return ps.nextState
}

/*
Implementación interface Notifier
*/
func (ps *PlayState) OnChangeDirection(newDirection float32) {
	ps.newDirection = newDirection
}

func (ps *PlayState) OnCreateCannonBullet(posX, posY float32, color color.Color) {
	bulletSprite, _ := ps.spriteCreator.SpriteByName("bullet")
	bullet := NewBullet(posX, posY, bulletSprite, color, -1)
	ps.bullets = append(ps.bullets, bullet)
	go ps.soundEffects.PlayShoot()
}

func (ps *PlayState) OnCreateAlienBullet(posX, posY float32, color color.Color) {
	bulletSprite, _ := ps.spriteCreator.SpriteByName("bullet")
	bullet := NewBullet(posX, posY, bulletSprite, color, 1)
	ps.bullets = append(ps.bullets, bullet)
}

func (ps *PlayState) OnResetCannon() {
	if ps.cannonCount == 0 {
		ps.innerStateId = gameOver
	} else {
		ps.cannon.Reset()
	}
}

func (ps *PlayState) OnResetUfo() {
	ps.ufo.Reset()
}

// -----------------------------------------------------------------------------
// Sección de procesamiento de eventos por estado.
// -----------------------------------------------------------------------------

func (ps *PlayState) processKeyEventPlaying() {
	if inpututil.IsKeyJustPressed(ebiten.KeyP) {
		ps.innerStateId = pauseRequest
	}

	//Para pruebas de paso de nivel.
	if inpututil.IsKeyJustReleased(ebiten.KeyL) {
		if len(ps.enemies) > 0 {
			ps.enemies = slices.DeleteFunc(ps.enemies, func(alien *enemy.Alien) bool {
				return true
			})
		}
	}
	ps.cannon.ProcessKeyEvents()
}

func (ps *PlayState) processKeyEventPause() {
	if inpututil.IsKeyJustPressed(ebiten.KeyP) {
		ps.innerStateId = playing
		ps.pauseScreen.Deallocate()
		ps.pauseScreen = nil
	}
}

// -----------------------------------------------------------------------------
// Sección de actualización por estado.
// -----------------------------------------------------------------------------

func (ps *PlayState) updatePlaying() {
	ps.cannon.Update()

	ps.alienFireTime += config.Dt
	if ps.alienFireTime > ps.level.Current().alienFireDelay {
		if len(ps.enemies) > 0 {
			pos := rand.Intn(len(ps.enemies))
			ps.enemies[pos].Fire()
		}
		ps.alienFireTime = 0
	}

	var enemyIncrementSpeed float32 = 0
	if len(ps.enemies) == 1 {
		enemyIncrementSpeed = 7
	} else if 2 <= len(ps.enemies) && len(ps.enemies) <= 5 {
		enemyIncrementSpeed = 5
	} else if 6 <= len(ps.enemies) && len(ps.enemies) <= 10 {
		enemyIncrementSpeed = 3
	}

	for _, enemy := range ps.enemies {
		enemy.IncrementSpeed(enemyIncrementSpeed)
	}

	for _, bullet := range ps.bullets {
		bullet.Update()
	}

	ps.ufo.Update()
	ufoX, _ := ps.ufo.Position()
	if ufoX >= 5 {
		go ps.soundEffects.PlayUfo()
	}

	for _, enemy := range ps.enemies {
		enemy.ChangeDirection(ps.enemiesCurrentDir)
		enemy.Update()
	}

	for _, explosion := range ps.explosions {
		explosion.Update()
	}

	//Colisiones balas.
	for _, bullet := range ps.bullets {
		if bullet.dirY < 0 {
			//Bala de cañon
			for _, bunker := range ps.bunkers {
				if common.CheckCollision(bullet, bunker) {
					if bunker.DoDamage(bullet.posX, bullet.posY, -1) {
						bullet.OnCollide()
					}
				}
			}
			for _, enemy := range ps.enemies {
				if common.CheckCollision(bullet, enemy) {
					bullet.OnCollide()
					enemy.OnCollide()

					alienExplosionSprite, _ := ps.spriteCreator.SpriteByName("alienExplosion")
					enemyX, enemyY := enemy.Position()
					explosion := explosion.NewExplosion(enemyX, enemyY, alienExplosionSprite, enemy.Color())
					ps.explosions = append(ps.explosions, explosion)
					ps.score += uint32(enemy.Score())
					go ps.soundEffects.PlayAlienKilled()
				}
			}
			if ps.ufo.IsActive() && common.CheckCollision(bullet, ps.ufo) {
				bullet.OnCollide()
				ps.ufo.OnCollide()

				ufoExplosionSprite, _ := ps.spriteCreator.SpriteByName("ufoExplosion")
				ufoX, ufoY := ps.ufo.Position()
				explosionUfo := explosion.NewExplosionUfo(ufoX, ufoY, ufoExplosionSprite, ps.textFace, ps.ufo.Score(), ps)
				ps.explosions = append(ps.explosions, explosionUfo)
				ps.score += uint32(ps.ufo.Score())

			}
		} else {
			//Bala de alien.
			for _, bunker := range ps.bunkers {
				if common.CheckCollision(bullet, bunker) {
					if bunker.DoDamage(bullet.posX, bullet.posY, 1) {
						bullet.OnCollide()
					}
				}
			}
			if ps.cannon.IsActive() && common.CheckCollision(bullet, ps.cannon) {
				cannonExplosion1Sprite, _ := ps.spriteCreator.SpriteByName("cannonExplosion1")
				cannonExplosion2Sprite, _ := ps.spriteCreator.SpriteByName("cannonExplosion2")
				posX, posY, _, _ := ps.cannon.Rect()
				explosionCannon := explosion.NewExplosionCannon(posX, posY, cannonExplosion1Sprite, cannonExplosion2Sprite, ps)
				ps.explosions = append(ps.explosions, explosionCannon)
				if ps.cannonCount > 0 {
					ps.cannonCount--
					ps.cannon.OnCollide()
				}
				bullet.OnCollide()
				go ps.soundEffects.PlayCannonExplosion()
			}
		}
	}

	//Colisines alien con bunker, cañon o llenando al suelo.
	for _, enemy := range ps.enemies {
		for _, bunker := range ps.bunkers {
			if common.CheckCollision(enemy, bunker) {
				bunker.OnCollide()
				break
			}
		}
		_, enemyY, _, enemyH := enemy.Rect()
		if enemyY+enemyH >= float32(config.DesignHeight) {
			ps.innerStateId = gameOver
			return
		}

		if common.CheckCollision(enemy, ps.cannon) {
			ps.innerStateId = gameOver
			return
		}
	}

	//Borramos los elementos marcados como borrados.
	if len(ps.bullets) > 0 {
		ps.bullets = slices.DeleteFunc(ps.bullets, func(bullet *Bullet) bool {
			return bullet.CanRemove()
		})
	}

	if len(ps.explosions) > 0 {
		ps.explosions = slices.DeleteFunc(ps.explosions, func(explosion explosion.Explosioner) bool {
			return explosion.CanRemove()
		})
	}

	if len(ps.enemies) > 0 {
		ps.enemies = slices.DeleteFunc(ps.enemies, func(alien *enemy.Alien) bool {
			return alien.CanRemove()
		})
	} else {
		ps.nextLevel()
		ps.reset()
		ps.cannon.Reset()
	}

	if len(ps.bunkers) > 0 {
		ps.bunkers = slices.DeleteFunc(ps.bunkers, func(bunker *player.Bunker) bool {
			return bunker.CanRemove()
		})
	}

	if ps.newDirection != ps.enemiesCurrentDir {
		ps.enemiesCurrentDir = ps.newDirection
	}
}

func (ps *PlayState) updatePause() {
	if ps.pauseScreen != nil {
		ps.innerStateId = pause
	}
}

func (ps *PlayState) updateGameOver() {
	ps.gameOverTime += config.Dt
	if ps.gameOverTime >= gameOverDelay {
		ps.gameOverTime = 0
		ps.nextState = states.Presentation
	}
}

// -----------------------------------------------------------------------------
// Sección de dibujo de pantalla por estado.
// -----------------------------------------------------------------------------

func (ps *PlayState) drawPlaying(screen *ebiten.Image) {
	uiCannonCountText := fmt.Sprintf("LIVES:%1d", ps.cannonCount)
	op := &text.DrawOptions{}
	op.GeoM.Translate(10, 6)
	op.ColorScale.ScaleWithColor(color.White)
	text.Draw(screen, uiCannonCountText, ps.textFace, op)

	uiScoreText := fmt.Sprintf("SCORE:%06d", ps.score)
	op = &text.DrawOptions{}
	op.GeoM.Translate(float64(config.DesignWidth-len(uiCannonCountText)*12), 6)
	op.ColorScale.ScaleWithColor(color.White)
	text.Draw(screen, uiScoreText, ps.textFace, op)

	ps.ufo.Draw(screen)
	for _, enemy := range ps.enemies {
		enemy.Draw(screen)
	}

	for _, explosion := range ps.explosions {
		explosion.Draw(screen)
	}

	for _, bullet := range ps.bullets {
		bullet.Draw(screen)
	}

	for _, bunker := range ps.bunkers {
		bunker.Draw(screen)
	}

	ps.cannon.Draw(screen)
}

func (ps *PlayState) drawPause(screen, pauseGameScreen *ebiten.Image) {
	opPauseGameScreen := &ebiten.DrawImageOptions{}
	opPauseGameScreen.GeoM.Translate(0, 0)

	rect := screen.Bounds()
	pauseImg := ebiten.NewImage(rect.Dx(), rect.Dy())
	pauseImg.Fill(color.RGBA{0x0A, 0x0A, 0x0A, 0xAA})

	opPauseImg := &ebiten.DrawImageOptions{}
	opPauseImg.GeoM.Translate(0, 0)

	uiPauseText := "PAUSE"
	opPause := &text.DrawOptions{}
	opPause.GeoM.Scale(2.0, 2.0)
	opPause.GeoM.Translate(float64(config.DesignWidth/2-len(uiPauseText)/2*13), float64(config.DesignWidth/2)-12)
	opPause.ColorScale.ScaleWithColor(color.White)
	text.Draw(pauseImg, uiPauseText, ps.textFace, opPause)

	screen.DrawImage(pauseGameScreen, opPauseGameScreen)
	screen.DrawImage(pauseImg, opPauseImg)
}

func (ps *PlayState) drawGameOver(screen *ebiten.Image) {
	uiGameOverText := "GAME OVER"
	gameOverX := float64(config.DesignWidth/2 - 24)
	gameOverY := float64(config.DesignHeight/2 - 30)

	op := &text.DrawOptions{}
	op.GeoM.Translate(gameOverX, gameOverY)
	op.ColorScale.ScaleWithColor(color.White)
	text.Draw(screen, uiGameOverText, ps.textFace, op)

	uiYourScoreText := fmt.Sprintf("YOUR SCORE:%06d", ps.score)
	yourScoreX := float64(config.DesignWidth/2 - 50)
	yourScoreY := float64(config.DesignHeight/2 - 15)
	op = &text.DrawOptions{}
	op.GeoM.Translate(yourScoreX, yourScoreY)
	op.ColorScale.ScaleWithColor(color.White)
	text.Draw(screen, uiYourScoreText, ps.textFace, op)
}

func (ps *PlayState) reset() {
	bunkers := createBunkers(ps.spriteCreator)
	enemies := createEnemies(ps.spriteCreator, ps.level.Current(), ps)

	ps.enemies = enemies
	ps.bullets = []*Bullet{}
	ps.bunkers = bunkers
	ps.explosions = []explosion.Explosioner{}
	ps.enemiesCurrentDir = 1
	ps.newDirection = 1
	ps.alienFireTime = 0
}

func (ps *PlayState) nextLevel() {
	hasLevels := ps.level.Next()
	if !hasLevels {
		ps.nextState = states.Presentation
	}
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

func createEnemies(spriteCreator *sprite.SpriteCreator, configLevel ConfigLevel, notifier common.Notifier) []*enemy.Alien {
	enemies := []*enemy.Alien{}

	squids := createAlien("squid", 11, 1, 9, 5, 30, 11, 35, spriteCreator, configLevel, notifier)
	enemies = append(enemies, squids...)

	crabs := createAlien("crab", 11, 2, 6, 5, 20, 10, 50, spriteCreator, configLevel, notifier)
	enemies = append(enemies, crabs...)

	octopuses := createAlien("octopus", 11, 2, 5, 5, 10, 9, 80, spriteCreator, configLevel, notifier)
	enemies = append(enemies, octopuses...)
	return enemies
}

func createAlien(alienName string, count, rows, offsetX, offsetY, points uint8, initX, initY float32, spriteCreator *sprite.SpriteCreator, configLevel ConfigLevel, notifier common.Notifier) []*enemy.Alien {
	sprite1, _ := spriteCreator.SpriteByName(fmt.Sprintf("%s1", alienName))
	sprite2, _ := spriteCreator.SpriteByName(fmt.Sprintf("%s2", alienName))
	aliens := []*enemy.Alien{}

	posX := initX
	posY := initY
	for i := range count * rows {
		alien := enemy.NewAlien(posX, posY, sprite1, sprite2, points, configLevel.alienMoveDelay, notifier)
		aliens = append(aliens, alien)
		posX += float32(sprite1.Image.Bounds().Dx() + int(offsetX))
		if i > 0 && (i+1)%count == 0 {
			posX = initX
			posY += float32(sprite1.Image.Bounds().Dy() + int(offsetY))
		}
	}
	return aliens

}

const gameOverDelay float32 = 3.0 //en segs.
