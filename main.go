package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct{}

// Implementación de la interface esperada por ebiten.
func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x03, 0x04, 0x5e, 0xFF})
	vector.DrawFilledRect(screen, 100, 100, 120, 120, color.RGBA{0, 255, 0, 255}, true)

	//sprite cañón.
	opCannon := &ebiten.DrawImageOptions{}
	opCannon.GeoM.Translate(300, 300)
	screen.DrawImage(spriteCannon, opCannon)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

var spriteCannon *ebiten.Image

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Juego")

	spriteCannon = SpriteFromArray(spriteDataCannon, 1, color.RGBA{0, 255, 0, 255})

	game := &Game{}
	err := ebiten.RunGame(game)
	if err != nil {
		panic(err)
	}
}

func SpriteFromArray(data [][]int, pixelSize int, colorOn color.Color) *ebiten.Image {
	h := len(data)
	w := len(data[0])
	img := ebiten.NewImage(w*pixelSize, h*pixelSize)

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if data[y][x] == 1 {
				rect := ebiten.NewImage(pixelSize, pixelSize)
				rect.Fill(colorOn)

				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(x*pixelSize), float64(y*pixelSize))
				img.DrawImage(rect, op)
			}
		}
	}
	return img
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
