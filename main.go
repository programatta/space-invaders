package main

import "github.com/hajimehoshi/ebiten/v2"

type Game struct{}

// Implementación de la interface esperada por ebiten.
func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Juego")

	game := &Game{}
	err := ebiten.RunGame(game)
	if err != nil {
		panic(err)
	}
}
