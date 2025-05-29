package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/programatta/spaceinvaders/internal"
)

type Game struct {
	cannon *internal.Cannon
	bullet *internal.Bullet
}

// Implementación de la interface esperada por ebiten.
func (g *Game) Update() error {
	g.cannon.ProcessKeyEvents()

	g.cannon.Update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x03, 0x04, 0x5e, 0xFF})

	g.cannon.Draw(screen)
	g.bullet.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth / 3, outsideHeight / 3
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Juego")

	spriteCannon := internal.SpriteFromArray(spriteDataCannon, 1, color.RGBA{0, 255, 0, 255})
	spriteBullet := internal.SpriteFromArray(spriteDataBullet, 1, color.RGBA{0, 255, 0, 255})

	game := &Game{}
	game.cannon = internal.NewCannon(100, 150, spriteCannon)
	game.bullet = internal.NewBullet(106, 146, spriteBullet)

	err := ebiten.RunGame(game)
	if err != nil {
		panic(err)
	}
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
