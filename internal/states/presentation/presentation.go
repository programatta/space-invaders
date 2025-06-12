package presentation

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/programatta/spaceinvaders/internal/states"
)

type PresentationState struct {
	nextState states.StateId
}

func NewPresentationState() *PresentationState {
	return &PresentationState{}
}

func (ps *PresentationState) ProcessEvents() {}

func (ps *PresentationState) Update() {}

func (ps *PresentationState) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x03, 0x04, 0x5e, 0xFF})
}

func (ps *PresentationState) NextState() states.StateId {
	return ps.nextState
}
