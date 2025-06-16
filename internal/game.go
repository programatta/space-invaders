package internal

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/programatta/spaceinvaders/internal/config"
	"github.com/programatta/spaceinvaders/internal/sounds"
	"github.com/programatta/spaceinvaders/internal/sprite"
	"github.com/programatta/spaceinvaders/internal/states"
	"github.com/programatta/spaceinvaders/internal/states/loader"
	"github.com/programatta/spaceinvaders/internal/states/play"
	"github.com/programatta/spaceinvaders/internal/states/presentation"
	"github.com/programatta/spaceinvaders/internal/utils"
)

type Game struct {
	currentState   states.State
	currentStateId states.StateId
	states         map[states.StateId]states.State
}

func NewGame() *Game {
	textFace := utils.LoadEmbeddedFont(8)
	spriteCreator := sprite.NewSpriteCreator()
	soundEffects := sounds.NewSoundEffects()

	game := &Game{}

	game.states = make(map[states.StateId]states.State)
	game.states[states.Presentation] = presentation.NewPresentationState(spriteCreator, textFace)
	game.states[states.Play] = play.NewPlayState(spriteCreator, textFace, soundEffects)

	game.currentState = loader.NewLoaderState(spriteCreator, textFace)
	game.currentStateId = states.Loader

	return game
}

// Implementación de la interface esperada por ebiten.
func (g *Game) Update() error {
	next := g.currentState.NextState()
	if next != g.currentStateId {
		g.currentState = g.states[next]
		g.currentStateId = next
		g.currentState.Start()
	}
	g.currentState.ProcessEvents()
	g.currentState.Update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.currentState.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return config.DesignWidth, config.DesignHeight
}
