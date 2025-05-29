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
}

func NewGame() *Game {
	spriteCannon := SpriteFromArray(spriteDataCannon, 1, color.RGBA{0, 255, 0, 255})

	game := &Game{}
	game.cannon = NewCannon(100, 150, spriteCannon, game)

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
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth / 3, outsideHeight / 3
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

const dt float32 = float32(1.0 / 60)
