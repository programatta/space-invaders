package main

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/programatta/spaceinvaders/internal"
)

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Juego")

	game := internal.NewGame()
	err := ebiten.RunGame(game)
	if err != nil {
		panic(err)
	}
}
