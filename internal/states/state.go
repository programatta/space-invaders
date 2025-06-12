package states

import "github.com/hajimehoshi/ebiten/v2"

type StateId int

const (
	Loader StateId = iota
	Presentation
	Play
)

// State define el comportamiento de un estado de juego
type State interface {
	ProcessEvents()
	Update()
	Draw(screen *ebiten.Image)
	NextState() StateId
}
