package main

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/programatta/spaceinvaders/internal"
)

func main() {
	ebiten.SetWindowSize(internal.WindowWidth, internal.WindowHeight)
	ebiten.SetWindowTitle("Juego")

	game := internal.NewGame()
	err := ebiten.RunGame(game)
	if err != nil {
		panic(err)
	}
}
