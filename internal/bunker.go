package internal

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Bunker struct {
	sprite Sprite
	posX   float32
	posY   float32
}

func NewBunker(posX, posY float32, sprite Sprite) *Bunker {
	spriteDataBunker := make([][]int, len(sprite.Data))
	for i := range sprite.Data {
		spriteDataBunker[i] = make([]int, len(sprite.Data[i]))
		copy(spriteDataBunker[i], sprite.Data[i])
	}

	sprite.Data = spriteDataBunker

	return &Bunker{
		sprite: sprite,
		posX:   posX,
		posY:   posY,
	}
}

func (b *Bunker) Update() {

}

func (b *Bunker) Draw(screen *ebiten.Image) {
	opBunker := &ebiten.DrawImageOptions{}
	opBunker.GeoM.Translate(float64(b.posX), float64(b.posY))
	screen.DrawImage(b.sprite.Image, opBunker)
}

func (b *Bunker) DoDamage(damageX, damageY float32) bool {
	damage := false
	logX := int((damageX - b.posX))
	logY := int((damageY - b.posY))

	if 0 <= logY && logY < b.sprite.Image.Bounds().Dy() {
		if b.sprite.Data[logY][logX] != 0 {
			b.sprite.Data[logY][logX] = 0
			if logY-1 >= 0 {
				b.sprite.Data[logY-1][logX] = 0
			}
			b.sprite.Image = SpriteFromArray(b.sprite.Data, 1, color.RGBA{0, 255, 0, 255})
			damage = true
		}
	}
	return damage
}

// Implementación de la interface Collider.
func (b *Bunker) Rect() (float32, float32, float32, float32) {
	width := float32(b.sprite.Image.Bounds().Dx())
	height := float32(b.sprite.Image.Bounds().Dy())
	return b.posX, b.posY, width, height
}

func (b *Bunker) OnCollide() {

}
