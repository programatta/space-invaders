package play

import (
	"fmt"

	"github.com/programatta/spaceinvaders/internal/config"
	"github.com/programatta/spaceinvaders/internal/sprite"
	"github.com/programatta/spaceinvaders/internal/states/play/common"
	"github.com/programatta/spaceinvaders/internal/states/play/enemy"
	"github.com/programatta/spaceinvaders/internal/states/play/player"
)

func createBunkers(spriteCreator *sprite.SpriteCreator) []*player.Bunker {
	bunkerSprite, _ := spriteCreator.SpriteByName("bunker")

	var posX float32 = 27
	bunkers := []*player.Bunker{}
	for range 4 {
		bunker := player.NewBunker(posX, float32(config.DesignHeight-40), bunkerSprite)
		bunkers = append(bunkers, bunker)
		posX += float32(bunkerSprite.Image.Bounds().Dx()) + 20
	}
	return bunkers
}

func createEnemies(spriteCreator *sprite.SpriteCreator, notifier common.Notifier) []*enemy.Alien {
	enemies := []*enemy.Alien{}

	squids := createAlienFormation("squid", 11, 1, 9, 5, 30, 11, 35, spriteCreator, notifier)
	enemies = append(enemies, squids...)

	crabs := createAlienFormation("crab", 11, 2, 6, 5, 20, 10, 50, spriteCreator, notifier)
	enemies = append(enemies, crabs...)

	octopuses := createAlienFormation("octopus", 11, 2, 5, 5, 10, 9, 80, spriteCreator, notifier)
	enemies = append(enemies, octopuses...)
	return enemies
}

func createAlienFormation(alienName string, count, rows, offsetX, offsetY, points uint8, initX, initY float32, spriteCreator *sprite.SpriteCreator, notifier common.Notifier) []*enemy.Alien {
	sprite1, _ := spriteCreator.SpriteByName(fmt.Sprintf("%s1", alienName))
	sprite2, _ := spriteCreator.SpriteByName(fmt.Sprintf("%s2", alienName))
	aliens := []*enemy.Alien{}

	posX := initX
	posY := initY
	for i := range count * rows {
		alien := enemy.NewAlien(posX, posY, sprite1, sprite2, points, config.AlienMoveDelay, notifier)
		aliens = append(aliens, alien)
		posX += float32(sprite1.Image.Bounds().Dx() + int(offsetX))
		if i > 0 && (i+1)%count == 0 {
			posX = initX
			posY += float32(sprite1.Image.Bounds().Dy() + int(offsetY))
		}
	}
	return aliens
}
