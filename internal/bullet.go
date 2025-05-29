package internal

import "github.com/hajimehoshi/ebiten/v2"

type Bullet struct {
	sprite *ebiten.Image
	posX   float32
	posY   float32
}

func NewBullet(posX, posY float32, sprite *ebiten.Image) *Bullet {
	return &Bullet{
		sprite: sprite,
		posX:   posX,
		posY:   posY,
	}
}

func (b *Bullet) Update() {
	if b.posY+float32(b.sprite.Bounds().Dy()) > 0 {
		b.posY -= 1
	}
}

func (b *Bullet) Draw(screen *ebiten.Image) {
	opBullet := &ebiten.DrawImageOptions{}
	opBullet.GeoM.Translate(float64(b.posX), float64(b.posY))
	screen.DrawImage(b.sprite, opBullet)
}
