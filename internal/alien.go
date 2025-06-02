package internal

import "github.com/hajimehoshi/ebiten/v2"

type Alien struct {
	sprites       []Sprite
	currentSprite uint
	posX          float32
	posY          float32
	time          float32
}

func NewAlien(posX, posY float32, sprite1, sprite2 Sprite) *Alien {
	sprites := []Sprite{sprite1, sprite2}
	return &Alien{sprites: sprites, posX: posX, posY: posY, currentSprite: 0, time: 0}
}

func (a *Alien) Update() {
	a.time += dt
	if a.time >= 0.35 {
		a.currentSprite = (a.currentSprite + 1) % 2
		a.time = 0
	}
}

func (a *Alien) Draw(screen *ebiten.Image) {
	spriteOptions := &ebiten.DrawImageOptions{}
	spriteOptions.GeoM.Translate(float64(a.posX), float64(a.posY))

	screen.DrawImage(a.sprites[a.currentSprite].Image, spriteOptions)
}
