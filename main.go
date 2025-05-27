package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/programatta/spaceinvaders/internal"
	"github.com/programatta/spaceinvaders/internal/config"
)

// Capturamos la versión desde -ldflags "-X main.Version=$(VERSION)" desde el makefile.
var Version = "dev"

func main() {
	ebiten.SetWindowSize(config.WindowWidth, config.WindowHeight)
	ebiten.SetWindowTitle("Space Invades")

	game := internal.NewGame(Version)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
