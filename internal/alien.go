package internal

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Alien struct {
	sprites       []Sprite
	currentSprite uint
	posX          float32
	posY          float32
	currentDirX   float32
	lastDirX      float32
	time          float32
	notifier      Notifier
	remove        bool
}

func NewAlien(posX, posY float32, sprite1, sprite2 Sprite, notifier Notifier) *Alien {
	sprites := []Sprite{sprite1, sprite2}
	return &Alien{sprites: sprites, posX: posX, posY: posY, currentSprite: 0, time: 0, notifier: notifier}
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

	a.time += dt
	if a.time >= 0.35 {
		a.posX += speed * dt * a.currentDirX
		a.currentSprite = (a.currentSprite + 1) % 2
		a.time = 0
	}

	if a.posX+float32(a.sprites[a.currentSprite].Image.Bounds().Dx()) >= float32(DesignWidth) {
		a.notifier.OnChangeDirection(-1)
		a.posX = float32(DesignWidth) - float32(a.sprites[a.currentSprite].Image.Bounds().Dx())
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

// Implementación de la interface Collider.
func (a *Alien) Rect() (float32, float32, float32, float32) {
	width := float32(a.sprites[a.currentSprite].Image.Bounds().Dx())
	height := float32(a.sprites[a.currentSprite].Image.Bounds().Dy())
	return a.posX, a.posY, width, height
}

func (a *Alien) OnCollide() {
	a.remove = true
}

const speed float32 = 200
