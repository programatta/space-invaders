package internal

import "github.com/hajimehoshi/ebiten/v2"

type Ufo struct {
	sprite Sprite
	posX   float32
	posY   float32
	remove bool
}

func NewUfo(posX, posY float32, sprite Sprite) *Ufo {
	return &Ufo{sprite: sprite, posX: posX, posY: posY}
}

func (u *Ufo) Update() {
	if !u.remove {
		u.posX++
		if u.posX >= float32(DesignWidth) {
			u.posX = -100
		}
	}
}

func (u *Ufo) Draw(screen *ebiten.Image) {
	if !u.remove {
		opUfo := &ebiten.DrawImageOptions{}
		opUfo.GeoM.Translate(float64(u.posX), float64(u.posY))
		screen.DrawImage(u.sprite.Image, opUfo)
	}
}

// Implementación de la interface Collider.
func (u *Ufo) Rect() (float32, float32, float32, float32) {
	width := float32(u.sprite.Image.Bounds().Dx())
	height := float32(u.sprite.Image.Bounds().Dy())
	return u.posX, u.posY, width, height
}

func (u *Ufo) OnCollide() {
	u.remove = true
}
