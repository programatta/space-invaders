package presentation

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/programatta/spaceinvaders/internal/config"
	"github.com/programatta/spaceinvaders/internal/sprite"
	"github.com/programatta/spaceinvaders/internal/states"
)

type PresentationState struct {
	spriteCreator *sprite.SpriteCreator
	textFace      *text.GoTextFace
	nextState     states.StateId
	time          float32
	uiTitleText   string
	textIndex     int
}

func NewPresentationState(spriteCreator *sprite.SpriteCreator, textFace *text.GoTextFace) *PresentationState {
	return &PresentationState{
		spriteCreator: spriteCreator,
		textFace:      textFace,
		nextState:     states.Presentation,
	}
}

func (ps *PresentationState) ProcessEvents() {}

func (ps *PresentationState) Update() {
	ps.time += config.Dt
	if ps.time >= 0.12 {
		if ps.textIndex <= len(title) {
			ps.uiTitleText = title[0:ps.textIndex]
			ps.textIndex++
		} else {
			ps.textIndex = 0
		}
		ps.time = 0
	}
}

func (ps *PresentationState) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x03, 0x04, 0x5e, 0xFF})

	widthText, _ := text.Measure(title, ps.textFace, 0)
	titleX := float64(config.DesignWidth/2) - widthText/2
	titleY := float64(config.DesignHeight/2 - 60)

	textOp := &text.DrawOptions{}
	textOp.GeoM.Translate(titleX, titleY)
	textOp.ColorScale.ScaleWithColor(color.White)
	text.Draw(screen, ps.uiTitleText, ps.textFace, textOp)
}

func (ps *PresentationState) NextState() states.StateId {
	return ps.nextState
}

const title string = "SPACE INVADERS"
