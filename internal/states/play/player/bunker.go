package player

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/programatta/spaceinvaders/internal/config"
	"github.com/programatta/spaceinvaders/internal/sprite"
	"github.com/programatta/spaceinvaders/internal/utils"
)

type Bunker struct {
	sprite sprite.Sprite
	posX   float32
	posY   float32
	remove bool
}

// NewBunker crea un nuevo objeto bunker en la posición indicada por parámetro
// además de generar la imagen a partir de un array.
func NewBunker(posX, posY float32, sprite sprite.Sprite) *Bunker {
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
		remove: false,
	}
}

func (b *Bunker) Rect() (float32, float32, float32, float32) {
	width := float32(b.sprite.Image.Bounds().Dx())
	height := float32(b.sprite.Image.Bounds().Dy())
	return b.posX, b.posY, width, height
}

func (b *Bunker) OnCollide() {
	b.remove = true
}

func (b *Bunker) CanRemove() bool {
	return b.remove
}

func (b *Bunker) Draw(screen *ebiten.Image) {
	opBunker := &ebiten.DrawImageOptions{}
	opBunker.GeoM.Translate(float64(b.posX), float64(b.posY))
	screen.DrawImage(b.sprite.Image, opBunker)
}

func (b *Bunker) DoDamage(damageX, damageY float32, dir int) bool {
	damage := false
	logX := int((damageX - b.posX)) / config.PixelSize
	logY := int((damageY - b.posY)) / config.PixelSize

	if 0 <= logY && logY < b.sprite.Image.Bounds().Dy() {
		if b.sprite.Data[logY][logX] != 0 {
			b.sprite.Data[logY][logX] = 0

			if dir > 0 {
				if logY+1 < b.sprite.Image.Bounds().Dy()-1 {
					b.sprite.Data[logY+1][logX] = 0
				}
			} else {
				if logY-1 >= 0 {
					b.sprite.Data[logY-1][logX] = 0
				}
			}
			b.sprite.Image = utils.SpriteFromArray(b.sprite.Data, config.PixelSize, b.sprite.Color)
			damage = true
		}
	}
	return damage
}
