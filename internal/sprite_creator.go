package internal

import "image/color"

type SpriteCreator struct {
	sprites map[string]Sprite
}

func NewSpriteCreator() *SpriteCreator {
	spriteCreator := &SpriteCreator{}

	spriteCreator.sprites = make(map[string]Sprite)

	spriteCreator.createCannonSprite()
	spriteCreator.createBunkerSprite()
	spriteCreator.createBulletSprite()

	return spriteCreator
}

func (sc *SpriteCreator) SpriteByName(name string) (Sprite, error) {
	return sc.sprites[name], nil
}

func (sc *SpriteCreator) createCannonSprite() {
	var spriteDataCannon = [][]int{
		{0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0},
		{0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	}
	sprite := SpriteFromArray(spriteDataCannon, 1, color.RGBA{0, 255, 255, 255})
	sc.sprites["cannon"] = Sprite{Image: sprite, Color: color.RGBA{0, 255, 255, 255}, Data: spriteDataCannon}
}

func (sc *SpriteCreator) createBunkerSprite() {
	var spriteDataBunker = [][]int{
		{0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0},
		{0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	}
	sprite := SpriteFromArray(spriteDataBunker, 1, color.RGBA{0, 255, 0, 255})
	sc.sprites["bunker"] = Sprite{Image: sprite, Color: color.RGBA{0, 255, 0, 255}, Data: spriteDataBunker}
}

func (sc *SpriteCreator) createBulletSprite() {
	var spriteDataBullet = [][]int{
		{1},
		{1},
	}
	sprite := SpriteFromArray(spriteDataBullet, 1, color.RGBA{255, 255, 255, 255})
	sc.sprites["bullet"] = Sprite{Image: sprite, Color: color.RGBA{255, 255, 255, 255}, Data: spriteDataBullet}
}
