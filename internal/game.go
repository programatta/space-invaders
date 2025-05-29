package internal

import (
	"image/color"
	"slices"

	"github.com/hajimehoshi/ebiten/v2"
)

type Notifier interface {
	OnCreateCannonBullet(posX, posY float32)
}

type Game struct {
	cannon  *Cannon
	bullets []*Bullet
	bunkers []*Bunker
}

func NewGame() *Game {
	spriteCannon := SpriteFromArray(spriteDataCannon, 1, color.RGBA{0, 255, 0, 255})
	spriteBunker := SpriteFromArray(spriteDataBunker, 1, color.RGBA{0, 255, 0, 255})

	game := &Game{}
	game.cannon = NewCannon(float32(0), float32(DesignHeight-10), spriteCannon, game)

	bunker1 := NewBunker(float32(27), float32(DesignHeight-40), spriteBunker)

	space := float32(bunker1.sprite.Bounds().Dx())
	bunker2 := NewBunker(27+space+20, float32(DesignHeight-40), spriteBunker)
	bunker3 := NewBunker(27+2*(space+20), float32(DesignHeight-40), spriteBunker)
	bunker4 := NewBunker(27+3*(space+20), float32(DesignHeight-40), spriteBunker)
	game.bunkers = []*Bunker{bunker1, bunker2, bunker3, bunker4}

	return game
}

// Implementación de la interface esperada por ebiten.
func (g *Game) Update() error {
	g.cannon.ProcessKeyEvents()

	g.cannon.Update()
	for _, bullet := range g.bullets {
		bullet.Update()
	}
	if len(g.bullets) > 0 {
		g.bullets = slices.DeleteFunc(g.bullets, func(bullet *Bullet) bool {
			return bullet.CanRemove()
		})
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x03, 0x04, 0x5e, 0xFF})

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
	spriteBullet := SpriteFromArray(spriteDataBullet, 1, color.RGBA{0, 255, 0, 255})
	bullet := NewBullet(posX, posY, spriteBullet)
	g.bullets = append(g.bullets, bullet)
}

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

var spriteDataBullet = [][]int{
	{1},
	{1},
}

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

const dt float32 = float32(1.0 / 60)

const WindowWidth int = 642
const WindowHeight int = 642
const DesignWidth int = WindowWidth / 3
const DesignHeight int = WindowHeight / 3
