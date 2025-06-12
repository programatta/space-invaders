package play

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/programatta/spaceinvaders/internal/states"
)

type PlayState struct {
	nextState states.StateId
}

func NewPlayState() *PlayState {
	return &PlayState{}
}

func (ps *PlayState) ProcessEvents() {}

func (ps *PlayState) Update() {}

func (ps *PlayState) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x03, 0x04, 0x5e, 0xFF})
}

func (ps *PlayState) NextState() states.StateId {
	return ps.nextState
}
