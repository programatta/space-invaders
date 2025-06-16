package play

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/programatta/spaceinvaders/internal/config"
	"github.com/programatta/spaceinvaders/internal/sprite"
)

type Bullet struct {
	sprite     sprite.Sprite
	posX       float32
	posY       float32
	dirY       float32
	colorRed   float32
	colorGreen float32
	colorBlue  float32
	colorAlpha float32
	remove     bool
}

func NewBullet(posX, posY float32, sprite sprite.Sprite, color color.Color, dirY float32) *Bullet {
	red, green, blue, alpha := color.RGBA()
	colorRed := float32(red)
	colorGreen := float32(green)
	colorBlue := float32(blue)
	colorAlpha := float32(alpha)
	return &Bullet{
		sprite:     sprite,
		posX:       posX,
		posY:       posY,
		dirY:       dirY,
		colorRed:   colorRed,
		colorGreen: colorGreen,
		colorBlue:  colorBlue,
		colorAlpha: colorAlpha,
	}
}

func (b *Bullet) Update() {
	b.posY += b.dirY

	if b.posY-float32(b.sprite.Image.Bounds().Dy()) < 0 || b.posY > float32(config.DesignHeight) {
		b.remove = true
	}
}

func (b *Bullet) Draw(screen *ebiten.Image) {
	opBullet := &ebiten.DrawImageOptions{}
	opBullet.GeoM.Translate(float64(b.posX), float64(b.posY))
	opBullet.ColorScale.Scale(b.colorRed, b.colorGreen, b.colorBlue, b.colorAlpha)
	screen.DrawImage(b.sprite.Image, opBullet)
}

func (b *Bullet) CanRemove() bool {
	return b.remove
}

// Implementación de la interface Collider.
func (b *Bullet) Rect() (float32, float32, float32, float32) {
	width := float32(b.sprite.Image.Bounds().Dx())
	height := float32(b.sprite.Image.Bounds().Dy())
	return b.posX, b.posY, width, height
}

func (b *Bullet) OnCollide() {
	b.remove = true
}
