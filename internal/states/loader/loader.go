package loader

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/programatta/spaceinvaders/internal/states"
)

type LoaderState struct {
	nextState states.StateId
}

func NewLoaderState() *LoaderState {
	return &LoaderState{
		nextState: states.Loader,
	}
}

func (ls *LoaderState) ProcessEvents() {}

func (ls *LoaderState) Update() {}

func (ls *LoaderState) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x03, 0x04, 0x5e, 0xFF})
	ebitenutil.DebugPrint(screen, "Loader")
}

func (ls *LoaderState) NextState() states.StateId {
	return ls.nextState
}
