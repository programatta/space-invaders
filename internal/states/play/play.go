package play

import (
	"fmt"
	"image/color"
	"math/rand"
	"slices"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/programatta/spaceinvaders/internal/config"
	"github.com/programatta/spaceinvaders/internal/sounds"
	"github.com/programatta/spaceinvaders/internal/sprite"
	"github.com/programatta/spaceinvaders/internal/states"
	"github.com/programatta/spaceinvaders/internal/states/play/common"
	"github.com/programatta/spaceinvaders/internal/states/play/enemy"
	"github.com/programatta/spaceinvaders/internal/states/play/explosion"
	"github.com/programatta/spaceinvaders/internal/states/play/player"
)

type PlayState struct {
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
	explosions        []explosion.Explosioner
	alienFireTime     float32
	innerStateId      playInnerStateId
	nextState         states.StateId
}

func NewPlayState(spriteCreator *sprite.SpriteCreator, textFace *text.GoTextFace, soundEffects *sounds.SoundEffects) *PlayState {

	playState := &PlayState{}

	playState.spriteCreator = spriteCreator
	playState.textFace = textFace
	playState.soundEffects = soundEffects

	spriteCannon, _ := spriteCreator.SpriteByName("cannon")
	playState.cannon = player.NewCannon(float32(0), float32(config.DesignHeight-10), spriteCannon, playState)

	bunkers := createBunkers(spriteCreator)
	playState.bunkers = bunkers

	ufoSprite, _ := spriteCreator.SpriteByName("ufo")
	ufo := enemy.NewUfo(-20, 15, ufoSprite)
	playState.ufo = ufo

	enemies := createEnemies(spriteCreator, playState)
	playState.enemies = enemies
	playState.enemiesCurrentDir = 1
	playState.newDirection = 1
	playState.cannonCount = 3
	playState.score = 0
	playState.innerStateId = playing

	playState.nextState = states.Play

	return playState
}

func (ps *PlayState) ProcessEvents() {
	switch ps.innerStateId {
	case playing:
		ps.processKeyEventPlaying()
	case gameOver:
		ps.processKeyEventGameOver()
	}
}

func (ps *PlayState) Update() {
	if ps.innerStateId == playing {
		ps.updatePlaying()
	}
}

func (ps *PlayState) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x03, 0x04, 0x5e, 0xFF})
	switch ps.innerStateId {
	case playing:
		ps.drawPlaying(screen)
	case gameOver:
		ps.drawGameOver(screen)
	}
}

func (ps *PlayState) NextState() states.StateId {
	return ps.nextState
}

// -----------------------------------------------------------------------------
// Implementación de la interface Notifier
// -----------------------------------------------------------------------------
func (ps *PlayState) OnCreateCannonBullet(posX, posY float32, color color.Color) {
	spriteBullet, _ := ps.spriteCreator.SpriteByName("bullet")
	bullet := NewBullet(posX, posY, spriteBullet, color, -1)
	ps.bullets = append(ps.bullets, bullet)
	ps.soundEffects.PlayShoot()
}

func (ps *PlayState) OnCreateAlienBullet(posX, posY float32, color color.Color) {
	spriteBullet, _ := ps.spriteCreator.SpriteByName("bullet")
	bullet := NewBullet(posX, posY, spriteBullet, color, 1)
	ps.bullets = append(ps.bullets, bullet)
}

func (ps *PlayState) OnChangeDirection(newDirection float32) {
	ps.newDirection = newDirection
}

func (ps *PlayState) OnResetUfo() {
	ps.ufo.Reset()
}

func (ps *PlayState) OnResetCannon() {
	if ps.cannonCount == 0 {
		ps.innerStateId = gameOver
	}
	ps.cannon.Reset()
}

// -----------------------------------------------------------------------------
// Sección de procesamiento de eventos por estado.
// -----------------------------------------------------------------------------

func (ps *PlayState) processKeyEventPlaying() {
	ps.cannon.ProcessKeyEvents()
}

func (ps *PlayState) processKeyEventGameOver() {
	if inpututil.IsKeyJustReleased(ebiten.KeySpace) {
		// resetear el juego.
		ps.reset()
		ps.cannon.Reset()
		ps.cannonCount = 3
		ps.score = 0
		ps.innerStateId = playing
	}
}

func (ps *PlayState) reset() {
	bunkers := createBunkers(ps.spriteCreator)
	enemies := createEnemies(ps.spriteCreator, ps)

	ps.enemies = enemies
	ps.bullets = []*Bullet{}
	ps.bunkers = bunkers
	ps.explosions = []explosion.Explosioner{}
	ps.enemiesCurrentDir = 1
	ps.newDirection = 1
	ps.alienFireTime = 0
}

// -----------------------------------------------------------------------------
// Sección de actualización por estado.
// -----------------------------------------------------------------------------

func (ps *PlayState) updatePlaying() {
	ps.cannon.Update()

	ps.alienFireTime += config.Dt
	if ps.alienFireTime > 0.400 {
		if len(ps.enemies) > 0 {
			pos := rand.Intn(len(ps.enemies))
			ps.enemies[pos].Fire()
		}
		ps.alienFireTime = 0
	}

	for _, bullet := range ps.bullets {
		bullet.Update()
	}

	ps.ufo.Update()
	ufoX, _ := ps.ufo.Position()
	if ufoX >= 5 {
		ps.soundEffects.PlayUfo()
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
		enemy.ChangeDirection(ps.enemiesCurrentDir)
		enemy.Update()
	}

	for _, explosion := range ps.explosions {
		explosion.Update()
	}

	//Colisiones.
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
					ps.soundEffects.PlayAlienKilled()
				}
			}
			if ps.ufo.IsActive() && common.CheckCollision(bullet, ps.ufo) {
				bullet.OnCollide()
				ps.ufo.OnCollide()

				ufoExplosionSprite, _ := ps.spriteCreator.SpriteByName("ufoExplosion")
				ufoX, ufoY := ps.ufo.Position()
				ufoScore := ps.ufo.Score()
				explosionUfo := explosion.NewExplosionUfo(ufoX, ufoY, ufoExplosionSprite, ps.textFace, ufoScore, ps)
				ps.explosions = append(ps.explosions, explosionUfo)
				ps.score += uint32(ufoScore)
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
				posX, posY := ps.cannon.Position()
				explosionCannon := explosion.NewExplosionCannon(posX, posY, cannonExplosion1Sprite, cannonExplosion2Sprite, ps)
				ps.explosions = append(ps.explosions, explosionCannon)
				if ps.cannonCount > 0 {
					ps.cannonCount--
					ps.cannon.OnCollide()
				}
				bullet.OnCollide()
				ps.soundEffects.PlayCannonExplosion()
			}
		}
	}

	//Colisines alien con bunker
	for _, enemy := range ps.enemies {
		for _, bunker := range ps.bunkers {
			if common.CheckCollision(enemy, bunker) {
				bunker.OnCollide()
				break
			}
		}
	}

	if len(ps.bullets) > 0 {
		ps.bullets = slices.DeleteFunc(ps.bullets, func(bullet *Bullet) bool {
			return bullet.CanRemove()
		})
	}

	if len(ps.enemies) > 0 {
		ps.enemies = slices.DeleteFunc(ps.enemies, func(alien *enemy.Alien) bool {
			return alien.CanRemove()
		})
	} else {
		ps.reset()
		ps.cannon.Reset()
	}

	if len(ps.explosions) > 0 {
		ps.explosions = slices.DeleteFunc(ps.explosions, func(explosion explosion.Explosioner) bool {
			return explosion.CanRemove()
		})
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
	widthText, _ := text.Measure(uiScoreText, ps.textFace, 0)
	op = &text.DrawOptions{}
	op.GeoM.Translate(float64(config.DesignWidth)-widthText-10, 6)
	op.ColorScale.ScaleWithColor(color.White)
	text.Draw(screen, uiScoreText, ps.textFace, op)

	ps.ufo.Draw(screen)

	for _, enemy := range ps.enemies {
		enemy.Draw(screen)
	}

	for _, explosion := range ps.explosions {
		explosion.Draw(screen)
	}

	ps.cannon.Draw(screen)
	for _, bullet := range ps.bullets {
		bullet.Draw(screen)
	}

	for _, bunker := range ps.bunkers {
		bunker.Draw(screen)
	}

}

func (ps *PlayState) drawGameOver(screen *ebiten.Image) {
	uiGameOverText := "GAME OVER"
	widthText, _ := text.Measure(uiGameOverText, ps.textFace, 0)
	gameOverX := float64(config.DesignWidth/2) - widthText/2
	gameOverY := float64(config.DesignHeight/2 - 30)

	op := &text.DrawOptions{}
	op.GeoM.Translate(gameOverX, gameOverY)
	op.ColorScale.ScaleWithColor(color.White)
	text.Draw(screen, uiGameOverText, ps.textFace, op)

	uiYourScoreText := fmt.Sprintf("YOUR SCORE:%06d", ps.score)
	widthText, _ = text.Measure(uiYourScoreText, ps.textFace, 0)
	yourScoreX := float64(config.DesignWidth/2) - widthText/2
	yourScoreY := float64(config.DesignHeight/2 - 15)

	op = &text.DrawOptions{}
	op.GeoM.Translate(yourScoreX, yourScoreY)
	op.ColorScale.ScaleWithColor(color.White)
	text.Draw(screen, uiYourScoreText, ps.textFace, op)
}
