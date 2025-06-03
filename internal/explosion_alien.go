package internal

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type ExplosionAlien struct {
	sprite     Sprite
	posX       float32
	posY       float32
	colorRed   float32
	colorGreen float32
	colorBlue  float32
	colorAlpha float32
	time       float32
	remove     bool
}

func NewExplosion(posX, posY float32, sprite Sprite, color color.Color) *ExplosionAlien {
	red, green, blue, alpha := color.RGBA()
	colorRed := float32(red)
	colorGreen := float32(green)
	colorBlue := float32(blue)
	colorAlpha := float32(alpha)

	return &ExplosionAlien{sprite: sprite, posX: posX, posY: posY, colorRed: colorRed, colorGreen: colorGreen, colorBlue: colorBlue, colorAlpha: colorAlpha}
}

func (ea *ExplosionAlien) CanRemove() bool {
	return ea.remove
}

func (ea *ExplosionAlien) Update() {
	ea.time += dt
	if ea.time >= 0.35 {
		ea.time = 0
		ea.remove = true
	}
}

func (ea *ExplosionAlien) Draw(screen *ebiten.Image) {
	opExplosionAlien := &ebiten.DrawImageOptions{}
	opExplosionAlien.GeoM.Translate(float64(ea.posX), float64(ea.posY))
	opExplosionAlien.ColorScale.Scale(ea.colorRed, ea.colorGreen, ea.colorBlue, ea.colorAlpha)
	screen.DrawImage(ea.sprite.Image, opExplosionAlien)
}
