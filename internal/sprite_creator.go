package internal

import "image/color"

type SpriteCreator struct {
	sprites map[string]Sprite
}

func NewSpriteCreator() *SpriteCreator {
	spriteCreator := &SpriteCreator{}

	spriteCreator.sprites = make(map[string]Sprite)
	spriteCreator.createCrabSprites()
	spriteCreator.createOctopusSprites()
	spriteCreator.createSquidSprites()
	spriteCreator.createUfoSprite()
	spriteCreator.createCannonSprite()
	spriteCreator.createBunkerSprite()
	spriteCreator.createBulletSprite()

	return spriteCreator
}

func (sc *SpriteCreator) SpriteByName(name string) (Sprite, error) {
	return sc.sprites[name], nil
}

func (sc *SpriteCreator) createCrabSprites() {
	var spriteDataCrab1 = [][]int{
		{0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0},
		{0, 0, 1, 1, 1, 1, 1, 1, 1, 0, 0},
		{0, 1, 1, 0, 1, 1, 1, 0, 1, 1, 0},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 0, 1, 1, 1, 1, 1, 1, 1, 0, 1},
		{1, 0, 1, 0, 0, 0, 0, 0, 1, 0, 1},
		{0, 0, 0, 1, 1, 0, 1, 1, 0, 0, 0},
	}

	var spriteDataCrab2 = [][]int{
		{0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0},
		{1, 0, 0, 1, 0, 0, 0, 1, 0, 0, 1},
		{1, 0, 1, 1, 1, 1, 1, 1, 1, 0, 1},
		{1, 1, 1, 0, 1, 1, 1, 0, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
		{0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0},
	}

	sprite1 := SpriteFromArray(spriteDataCrab1, 1, color.RGBA{0, 255, 0, 255})
	sprite2 := SpriteFromArray(spriteDataCrab2, 1, color.RGBA{0, 255, 0, 255})
	sc.sprites["crab1"] = Sprite{Image: sprite1, Color: color.RGBA{0, 255, 0, 255}, Data: spriteDataCrab1}
	sc.sprites["crab2"] = Sprite{Image: sprite2, Color: color.RGBA{0, 255, 0, 255}, Data: spriteDataCrab2}
}

func (sc *SpriteCreator) createOctopusSprites() {
	var spriteDataOctopus1 = [][]int{
		{0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0},
		{0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{0, 0, 1, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0},
		{0, 0, 1, 1, 0, 0, 0, 0, 1, 1, 0, 0},
	}

	var spriteDataOctopus2 = [][]int{
		{0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0},
		{0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{0, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 0},
		{0, 0, 1, 1, 0, 1, 1, 0, 1, 1, 0, 0},
		{1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1},
	}
	sprite1 := SpriteFromArray(spriteDataOctopus1, 1, color.RGBA{255, 255, 0, 255})
	sprite2 := SpriteFromArray(spriteDataOctopus2, 1, color.RGBA{255, 255, 0, 255})
	sc.sprites["octopus1"] = Sprite{Image: sprite1, Color: color.RGBA{255, 255, 0, 255}, Data: spriteDataOctopus1}
	sc.sprites["octopus2"] = Sprite{Image: sprite2, Color: color.RGBA{255, 255, 0, 255}, Data: spriteDataOctopus2}
}

func (sc *SpriteCreator) createSquidSprites() {
	var spriteDataSquid1 = [][]int{
		{0, 0, 0, 1, 1, 0, 0, 0},
		{0, 0, 1, 1, 1, 1, 0, 0},
		{0, 1, 1, 1, 1, 1, 1, 0},
		{1, 1, 0, 1, 1, 0, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1},
		{0, 0, 1, 0, 0, 1, 0, 0},
		{0, 1, 0, 1, 1, 0, 1, 0},
		{1, 0, 1, 0, 0, 1, 0, 1},
	}

	var spriteDataSquid2 = [][]int{
		{0, 0, 0, 1, 1, 0, 0, 0},
		{0, 0, 1, 1, 1, 1, 0, 0},
		{0, 1, 1, 1, 1, 1, 1, 0},
		{1, 1, 0, 1, 1, 0, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1},
		{0, 1, 0, 1, 1, 0, 1, 0},
		{1, 0, 0, 0, 0, 0, 0, 1},
		{0, 1, 0, 0, 0, 0, 1, 0},
	}

	sprite1 := SpriteFromArray(spriteDataSquid1, 1, color.RGBA{255, 0, 255, 255})
	sprite2 := SpriteFromArray(spriteDataSquid2, 1, color.RGBA{255, 0, 255, 255})
	sc.sprites["squid1"] = Sprite{Image: sprite1, Color: color.RGBA{255, 0, 255, 255}, Data: spriteDataSquid1}
	sc.sprites["squid2"] = Sprite{Image: sprite2, Color: color.RGBA{255, 0, 255, 255}, Data: spriteDataSquid2}
}

func (sc *SpriteCreator) createUfoSprite() {
	var spriteDataUFO = [][]int{
		{0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0},
		{0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0},
		{0, 1, 1, 0, 1, 1, 0, 1, 1, 0, 1, 1, 0, 1, 1, 0},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{0, 0, 1, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0},
	}
	sprite := SpriteFromArray(spriteDataUFO, 1, color.RGBA{255, 0, 0, 255})
	sc.sprites["ufo"] = Sprite{Image: sprite, Color: color.RGBA{255, 0, 0, 255}, Data: spriteDataUFO}
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
