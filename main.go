package main

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/programatta/spaceinvaders/internal"
	"github.com/programatta/spaceinvaders/internal/config"
)

func main() {
	ebiten.SetWindowSize(config.WindowWidth, config.WindowHeight)
	ebiten.SetWindowTitle("Juego")

	game := internal.NewGame()
	err := ebiten.RunGame(game)
	if err != nil {
		panic(err)
	}
}
