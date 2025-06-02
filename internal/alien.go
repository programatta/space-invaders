package internal

import "github.com/hajimehoshi/ebiten/v2"

type Alien struct {
	sprites []Sprite
	posX    float32
	posY    float32
}

func NewAlien(posX, posY float32, sprite1, sprite2 Sprite) *Alien {
	sprites := []Sprite{sprite1, sprite2}
	return &Alien{sprites: sprites, posX: posX, posY: posY}
}

func (a *Alien) Update() {

}

func (a *Alien) Draw(screen *ebiten.Image) {
	spriteOptions := &ebiten.DrawImageOptions{}
	spriteOptions.GeoM.Translate(float64(a.posX), float64(a.posY))

	screen.DrawImage(a.sprites[0].Image, spriteOptions)
}
