package internal

import "github.com/hajimehoshi/ebiten/v2"

type Ufo struct {
	sprite Sprite
	posX   float32
	posY   float32
}

func NewUfo(posX, posY float32, sprite Sprite) *Ufo {
	return &Ufo{sprite: sprite, posX: posX, posY: posY}
}

func (u *Ufo) Update() {
	u.posX++
	if u.posX >= float32(DesignWidth) {
		u.posX = -100
	}
}

func (u *Ufo) Draw(screen *ebiten.Image) {
	opUfo := &ebiten.DrawImageOptions{}
	opUfo.GeoM.Translate(float64(u.posX), float64(u.posY))
	screen.DrawImage(u.sprite.Image, opUfo)
}
