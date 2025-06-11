package enemy

import (
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/programatta/spaceinvaders/internal/config"
	"github.com/programatta/spaceinvaders/internal/sprite"
)

type Ufo struct {
	sprite sprite.Sprite
	posX   float32
	posY   float32
	remove bool
	active bool
	speed  float32
}

func NewUfo(posX, posY float32, sprite sprite.Sprite) *Ufo {
	return &Ufo{sprite: sprite, posX: posX, posY: posY, active: true, speed: 60}
}

func (u *Ufo) Update() {
	if !u.remove {
		u.posX += u.speed * config.Dt
		if u.posX >= float32(config.DesignWidth) {
			u.resetPosition()
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

func (u *Ufo) Score() uint16 {
	scores := []uint16{150, 175, 200, 225, 250, 275, 300, 325, 350}
	pos := rand.IntN(len(scores))
	return scores[pos]
}

func (u *Ufo) IsActive() bool {
	return u.active
}

// Implementación de la interface Collider.
func (u *Ufo) Rect() (float32, float32, float32, float32) {
	width := float32(u.sprite.Image.Bounds().Dx())
	height := float32(u.sprite.Image.Bounds().Dy())
	return u.posX, u.posY, width, height
}

func (u *Ufo) OnCollide() {
	u.remove = true
	u.active = false
}

func (u *Ufo) Position() (float32, float32) {
	return u.posX, u.posY
}

func (u *Ufo) Reset() {
	u.resetPosition()
	u.remove = false
	u.active = true
}

func (u *Ufo) resetPosition() {
	pos := rand.IntN(500) + 200
	u.posX = -float32(pos)
}
