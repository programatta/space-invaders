package internal

import "github.com/hajimehoshi/ebiten/v2"

type Bunker struct {
	sprite *ebiten.Image
	posX   float32
	posY   float32
}

func NewBunker(posX, posY float32, sprite *ebiten.Image) *Bunker {
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
	screen.DrawImage(b.sprite, opBunker)
}
