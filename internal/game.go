package internal

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/programatta/spaceinvaders/internal/config"
	"github.com/programatta/spaceinvaders/internal/sound"
	"github.com/programatta/spaceinvaders/internal/sprite"
	"github.com/programatta/spaceinvaders/internal/states"
	"github.com/programatta/spaceinvaders/internal/states/loader"
	"github.com/programatta/spaceinvaders/internal/states/play"
	"github.com/programatta/spaceinvaders/internal/states/presentation"
	"github.com/programatta/spaceinvaders/internal/utils"
)

// Game contiene el estado en curso y su identificador, y el conjunto de
// estados que se usan durante el juego.
type Game struct {
	currentState   states.State
	currentStateId states.StateId
	states         map[states.StateId]states.State
}

// NewGame carga los recursos del juego y crea los estados de juego (loader,
// presentation y play) dejando el estado loader como predeterminado.
//
// Devielve la instancia del juego.
func NewGame(version string) *Game {
	// * Usaremos la versión para mostrarla en el estado correcto.
	fmt.Printf("\nSpace invaders v%s\n\n", version)

	game := &Game{}

	textFace := utils.LoadEmbeddedFont(8)
	spriteCreator := sprite.NewSpriteCreator()
	soundEffects := sound.NewSoundEffects()

	game.states = make(map[states.StateId]states.State)
	game.states[states.Presentation] = presentation.NewPresentationState(spriteCreator, textFace)
	game.states[states.Play] = play.NewPlayState(spriteCreator, soundEffects, textFace)

	game.currentState = loader.NewLoaderState(spriteCreator, textFace, version)
	game.currentStateId = states.Loader
	return game
}

// ----------------------------------------------------------------------------
// Implementa Ebiten Game Interface
// ----------------------------------------------------------------------------

// Update realiza el cambio de estado si es necesario y permite procesar
// eventos y actualizar su lógica.
func (g *Game) Update() error {
	next := g.currentState.NextState()
	if next != g.currentStateId {
		g.currentState = g.states[next]
		g.currentState.Start()
		g.currentStateId = next
	}
	g.currentState.ProcessEvents()
	g.currentState.Update()

	return nil
}

// Draw dibuja el estado actual.
func (g *Game) Draw(screen *ebiten.Image) {
	g.currentState.Draw(screen)
}

// Layout determina el tamaño del canvas
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return config.DesignWidth, config.DesignHeight
}
