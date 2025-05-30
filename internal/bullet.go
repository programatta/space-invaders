package internal

import "github.com/hajimehoshi/ebiten/v2"

type Bullet struct {
	sprite Sprite
	posX   float32
	posY   float32
	remove bool
}

func NewBullet(posX, posY float32, sprite Sprite) *Bullet {
	return &Bullet{
		sprite: sprite,
		posX:   posX,
		posY:   posY,
	}
}

func (b *Bullet) Update() {
	b.posY -= 1
	if b.posY-float32(b.sprite.Image.Bounds().Dy()) < 0 {
		b.remove = true
	}
}

func (b *Bullet) Draw(screen *ebiten.Image) {
	opBullet := &ebiten.DrawImageOptions{}
	opBullet.GeoM.Translate(float64(b.posX), float64(b.posY))
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
