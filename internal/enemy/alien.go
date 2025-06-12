package enemy

import (
	"image/color"
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/programatta/spaceinvaders/internal/common"
	"github.com/programatta/spaceinvaders/internal/config"
	"github.com/programatta/spaceinvaders/internal/sprite"
)

type Alien struct {
	sprites        []sprite.Sprite
	currentSprite  uint
	posX           float32
	posY           float32
	score          uint8
	alienMoveDelay float32
	currentDirX    float32
	lastDirX       float32
	time           float32
	notifier       common.Notifier
	remove         bool
	currentDelay   float32
}

func NewAlien(posX, posY float32, sprite1, sprite2 sprite.Sprite, score uint8, alienMoveDelay float32, notifier common.Notifier) *Alien {
	sprites := []sprite.Sprite{sprite1, sprite2}
	return &Alien{sprites: sprites, posX: posX, posY: posY, score: score, alienMoveDelay: alienMoveDelay, currentSprite: 0, time: 0, currentDelay: alienMoveDelay, notifier: notifier}
}

func (a *Alien) Position() (float32, float32) {
	return a.posX, a.posY
}

func (a *Alien) Color() color.Color {
	return a.sprites[a.currentSprite].Color
}

func (a *Alien) ChangeDirection(currentDir float32) {
	a.currentDirX = currentDir
}

func (a *Alien) Update() {
	if a.lastDirX != a.currentDirX {
		if a.lastDirX != 0 {
			a.posY += 5
		}
		a.lastDirX = a.currentDirX
	}

	a.time += config.Dt
	if a.time >= a.currentDelay {
		a.posX += config.AlienSpeed * config.Dt * a.currentDirX
		a.currentSprite = (a.currentSprite + 1) % 2
		a.time = 0
	}

	if a.posX+float32(a.sprites[a.currentSprite].Image.Bounds().Dx()) >= float32(config.DesignWidth) {
		a.notifier.OnChangeDirection(-1)
		a.posX = float32(config.DesignWidth) - float32(a.sprites[a.currentSprite].Image.Bounds().Dx())
	} else if a.posX <= 0 {
		a.notifier.OnChangeDirection(1)
		a.posX = 0
	}
}

func (a *Alien) Draw(screen *ebiten.Image) {
	spriteOptions := &ebiten.DrawImageOptions{}
	spriteOptions.GeoM.Translate(float64(a.posX), float64(a.posY))

	screen.DrawImage(a.sprites[a.currentSprite].Image, spriteOptions)
}

func (a *Alien) CanRemove() bool {
	return a.remove
}

func (a *Alien) IncrementSpeed(incrementSpeed float32) {
	if incrementSpeed > 0 {
		a.currentDelay = a.alienMoveDelay / incrementSpeed
	}
}

// Implementación de la interface Collider.
func (a *Alien) Rect() (float32, float32, float32, float32) {
	width := float32(a.sprites[a.currentSprite].Image.Bounds().Dx())
	height := float32(a.sprites[a.currentSprite].Image.Bounds().Dy())
	return a.posX, a.posY, width, height
}

func (a *Alien) OnCollide() {
	a.remove = true
}

func (a *Alien) Fire() {
	difX := 0
	dado := rand.IntN(3)
	if dado == 0 {
		difX = -1
	}
	if dado == 2 {
		difX = 1
	}

	bulletX := a.posX + float32(a.sprites[a.currentSprite].Image.Bounds().Dx()/2+difX)
	bulletY := a.posY + float32(a.sprites[a.currentSprite].Image.Bounds().Dy())
	a.notifier.OnCreateAlienBullet(bulletX, bulletY, a.sprites[a.currentSprite].Color)
}

func (a *Alien) Score() uint8 {
	return a.score
}
