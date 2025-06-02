package internal

import (
	"image/color"
	"slices"

	"github.com/hajimehoshi/ebiten/v2"
)

type Notifier interface {
	OnCreateCannonBullet(posX, posY float32)
}

type Collider interface {
	Rect() (float32, float32, float32, float32)
	OnCollide()
}

type Game struct {
	spriteCreator *SpriteCreator
	cannon        *Cannon
	bullets       []*Bullet
	bunkers       []*Bunker
	ufo           *Ufo
	enemies       []*Alien
}

func NewGame() *Game {
	spriteCreator := NewSpriteCreator()

	game := &Game{}
	game.spriteCreator = spriteCreator

	spriteCannon, _ := spriteCreator.SpriteByName("cannon")
	game.cannon = NewCannon(float32(0), float32(DesignHeight-10), spriteCannon, game)

	spriteBunker, _ := spriteCreator.SpriteByName("bunker")
	bunker1 := NewBunker(float32(27), float32(DesignHeight-40), spriteBunker)

	space := float32(bunker1.sprite.Image.Bounds().Dx())
	bunker2 := NewBunker(27+space+20, float32(DesignHeight-40), spriteBunker)
	bunker3 := NewBunker(27+2*(space+20), float32(DesignHeight-40), spriteBunker)
	bunker4 := NewBunker(27+3*(space+20), float32(DesignHeight-40), spriteBunker)
	game.bunkers = []*Bunker{bunker1, bunker2, bunker3, bunker4}

	ufoSprite, _ := spriteCreator.SpriteByName("ufo")
	ufo := NewUfo(-20, 5, ufoSprite)
	game.ufo = ufo

	enemies := createEnemies(spriteCreator)
	game.enemies = enemies
	return game
}

// Implementación de la interface esperada por ebiten.
func (g *Game) Update() error {
	g.cannon.ProcessKeyEvents()

	g.cannon.Update()
	for _, bullet := range g.bullets {
		bullet.Update()
	}

	g.ufo.Update()
	for _, enemy := range g.enemies {
		enemy.Update()
	}

	if len(g.bullets) > 0 {
		g.bullets = slices.DeleteFunc(g.bullets, func(bullet *Bullet) bool {
			return bullet.CanRemove()
		})
	}

	//Colisiones.
	for _, bullet := range g.bullets {
		for _, bunker := range g.bunkers {
			if g.checkCollision(bullet, bunker) {
				if bunker.DoDamage(bullet.posX, bullet.posY) {
					bullet.OnCollide()
				}
			}
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x03, 0x04, 0x5e, 0xFF})

	g.ufo.Draw(screen)

	for _, enemy := range g.enemies {
		enemy.Draw(screen)
	}

	g.cannon.Draw(screen)
	for _, bullet := range g.bullets {
		bullet.Draw(screen)
	}

	for _, bunker := range g.bunkers {
		bunker.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return DesignWidth, DesignHeight
}

// Implementación de la interface Notifier
func (g *Game) OnCreateCannonBullet(posX, posY float32) {
	spriteBullet, _ := g.spriteCreator.SpriteByName("bullet")
	bullet := NewBullet(posX, posY, spriteBullet)
	g.bullets = append(g.bullets, bullet)
}

func (g *Game) checkCollision(sourceObj, targetObj Collider) bool {
	sx0, sy0, sw, sh := sourceObj.Rect()
	tx0, ty0, tw, th := targetObj.Rect()

	hasCollision := sx0 < tx0+tw && sx0+sw > tx0 && sy0 < ty0+th && sh+sy0 > ty0

	return hasCollision
}

func createEnemies(spriteCreator *SpriteCreator) []*Alien {
	enemies := []*Alien{}

	squids := createSquids(11, 1, 11, 35, spriteCreator)
	enemies = append(enemies, squids...)

	crabs := createCrabs(11, 2, 10, 50, spriteCreator)
	enemies = append(enemies, crabs...)

	octopuses := createOctopuses(11, 2, 9, 80, spriteCreator)
	enemies = append(enemies, octopuses...)

	return enemies
}

func createCrabs(count, rows uint8, initX, initY float32, spriteCreator *SpriteCreator) []*Alien {
	sprite1, _ := spriteCreator.SpriteByName("crab1")
	sprite2, _ := spriteCreator.SpriteByName("crab2")
	crabs := []*Alien{}

	posX := initX
	posY := initY
	for i := range count * rows {
		crab := NewAlien(posX, posY, sprite1, sprite2)
		crabs = append(crabs, crab)
		posX += float32(sprite1.Image.Bounds().Dx() + 6)
		if i > 0 && (i+1)%count == 0 {
			posX = initX
			posY += float32(sprite1.Image.Bounds().Dy() + 5)
		}
	}
	return crabs
}

func createOctopuses(count, rows uint8, initX, initY float32, spriteCreator *SpriteCreator) []*Alien {
	sprite1, _ := spriteCreator.SpriteByName("octopus1")
	sprite2, _ := spriteCreator.SpriteByName("octopus2")
	octopuses := []*Alien{}

	posX := initX
	posY := initY
	for i := range count * rows {
		octopus := NewAlien(posX, posY, sprite1, sprite2)
		octopuses = append(octopuses, octopus)
		posX += float32(sprite1.Image.Bounds().Dx() + 5)
		if i > 0 && (i+1)%count == 0 {
			posX = initX
			posY += float32(sprite1.Image.Bounds().Dy() + 5)
		}
	}
	return octopuses
}

func createSquids(count, rows uint8, initX, initY float32, spriteCreator *SpriteCreator) []*Alien {
	sprite1, _ := spriteCreator.SpriteByName("squid1")
	sprite2, _ := spriteCreator.SpriteByName("squid2")
	squids := []*Alien{}

	posX := initX
	posY := initY
	for i := range count * rows {
		squid := NewAlien(posX, posY, sprite1, sprite2)
		squids = append(squids, squid)
		posX += float32(sprite1.Image.Bounds().Dx() + 9)
		if i > 0 && (i+1)%count == 0 {
			posX = initX
			posY += float32(sprite1.Image.Bounds().Dy())
		}
	}
	return squids
}

const dt float32 = float32(1.0 / 60)

const WindowWidth int = 642
const WindowHeight int = 642
const DesignWidth int = WindowWidth / 3
const DesignHeight int = WindowHeight / 3
