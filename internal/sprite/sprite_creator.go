package sprite

import (
	"image/color"

	"github.com/programatta/spaceinvaders/internal/config"
	"github.com/programatta/spaceinvaders/internal/utils"
)

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
	spriteCreator.createAlienExplosion()
	spriteCreator.createUfoExplosion()
	spriteCreator.createCannonExplosions()

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

	sprite1 := utils.SpriteFromArray(spriteDataCrab1, config.PixelSize, color.RGBA{0, 255, 0, 255})
	sprite2 := utils.SpriteFromArray(spriteDataCrab2, config.PixelSize, color.RGBA{0, 255, 0, 255})
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
	sprite1 := utils.SpriteFromArray(spriteDataOctopus1, config.PixelSize, color.RGBA{255, 255, 0, 255})
	sprite2 := utils.SpriteFromArray(spriteDataOctopus2, config.PixelSize, color.RGBA{255, 255, 0, 255})
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

	sprite1 := utils.SpriteFromArray(spriteDataSquid1, config.PixelSize, color.RGBA{255, 0, 255, 255})
	sprite2 := utils.SpriteFromArray(spriteDataSquid2, config.PixelSize, color.RGBA{255, 0, 255, 255})
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
	sprite := utils.SpriteFromArray(spriteDataUFO, config.PixelSize, color.RGBA{255, 0, 0, 255})
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
	sprite := utils.SpriteFromArray(spriteDataCannon, config.PixelSize, color.RGBA{0, 255, 255, 255})
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
	sprite := utils.SpriteFromArray(spriteDataBunker, config.PixelSize, color.RGBA{0, 255, 0, 255})
	sc.sprites["bunker"] = Sprite{Image: sprite, Color: color.RGBA{0, 255, 0, 255}, Data: spriteDataBunker}
}

func (sc *SpriteCreator) createBulletSprite() {
	var spriteDataBullet = [][]int{
		{1},
		{1},
	}
	sprite := utils.SpriteFromArray(spriteDataBullet, config.PixelSize, color.RGBA{255, 255, 255, 255})
	sc.sprites["bullet"] = Sprite{Image: sprite, Color: color.RGBA{255, 255, 255, 255}, Data: spriteDataBullet}
}

func (sc *SpriteCreator) createAlienExplosion() {
	var spriteDataExplosionAlien = [][]int{
		{0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0},
		{0, 1, 0, 0, 1, 0, 1, 0, 0, 1, 0},
		{0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0},
		{1, 1, 0, 0, 0, 0, 0, 0, 0, 1, 1},
		{0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0},
		{0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 0, 1, 0, 0, 1, 0},
		{0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0},
	}
	sprite := utils.SpriteFromArray(spriteDataExplosionAlien, config.PixelSize, color.RGBA{255, 255, 255, 255})
	sc.sprites["alienExplosion"] = Sprite{Image: sprite, Color: color.RGBA{255, 255, 255, 255}, Data: spriteDataExplosionAlien}
}

func (sc *SpriteCreator) createUfoExplosion() {
	var spriteDataExplosionUFO = [][]int{
		{0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1},
		{0, 0, 0, 1, 0, 1, 0, 1, 0, 0, 0, 0, 1, 0, 1, 0},
		{1, 0, 1, 0, 0, 0, 1, 1, 0, 1, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 1, 0, 1, 1, 0, 1, 1, 1, 0, 1, 0},
		{0, 0, 1, 0, 1, 0, 1, 1, 1, 0, 1, 0, 1, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 1, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0},
	}
	sprite := utils.SpriteFromArray(spriteDataExplosionUFO, config.PixelSize, color.RGBA{255, 0, 0, 255})
	sc.sprites["ufoExplosion"] = Sprite{Image: sprite, Color: color.RGBA{255, 0, 0, 255}, Data: spriteDataExplosionUFO}
}

func (sc *SpriteCreator) createCannonExplosions() {
	var spriteDataExplosionCannon1 = [][]int{
		{0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0},
		{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0},
		{0, 0, 0, 1, 0, 0, 0, 1, 1, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 1, 1, 0, 0, 1, 0},
		{0, 0, 1, 0, 0, 1, 1, 0, 0, 0, 1, 0, 1},
		{0, 0, 0, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0},
		{0, 0, 1, 1, 0, 1, 1, 1, 1, 0, 0, 1, 0},
	}

	var spriteDataExplosionCannon2 = [][]int{
		{0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 1, 0, 1, 0, 0, 0},
		{0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0},
		{1, 0, 0, 1, 0, 1, 1, 0, 0, 1, 0, 0, 0},
		{0, 0, 1, 1, 1, 1, 1, 0, 0, 0, 1, 0, 0},
		{0, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 0, 1},
	}
	sprite1 := utils.SpriteFromArray(spriteDataExplosionCannon1, config.PixelSize, color.RGBA{0, 255, 255, 255})
	sprite2 := utils.SpriteFromArray(spriteDataExplosionCannon2, config.PixelSize, color.RGBA{0, 255, 255, 255})
	sc.sprites["cannonExplosion1"] = Sprite{Image: sprite1, Color: color.RGBA{0, 255, 255, 255}, Data: spriteDataExplosionCannon1}
	sc.sprites["cannonExplosion2"] = Sprite{Image: sprite2, Color: color.RGBA{0, 255, 255, 255}, Data: spriteDataExplosionCannon2}
}
