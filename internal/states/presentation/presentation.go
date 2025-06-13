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
	uiMisteryText string
	uiSquidText   string
	uiCrabText    string
	uiOctopusText string
	textIndex     int
	currentStep   int
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
		if ps.currentStep == 0 {
			ps.uiTitleText = ps.typeWriteText(title)
		}
		if ps.currentStep == 1 {
			ps.uiMisteryText = ps.typeWriteText(mistery)
		}
		if ps.currentStep == 2 {
			ps.uiSquidText = ps.typeWriteText(squid)
		}
		if ps.currentStep == 3 {
			ps.uiCrabText = ps.typeWriteText(crab)
		}
		if ps.currentStep == 4 {
			ps.uiOctopusText = ps.typeWriteText(octopus)
		}
		if ps.currentStep == 5 {
			ps.currentStep = 0
		}
		ps.time = 0
	}
}

func (ps *PresentationState) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x03, 0x04, 0x5e, 0xFF})

	if ps.currentStep >= 0 {
		widthText, _ := text.Measure(title, ps.textFace, 0)
		titleX := float64(config.DesignWidth/2) - widthText/2
		titleY := float64(config.DesignHeight/2 - 60)

		ps.drawText(screen, ps.uiTitleText, titleX, titleY, color.White)
	}
	if ps.currentStep >= 1 {
		misteryTextX := float64(config.DesignWidth/2 - 40)
		misteryTextY := float64(config.DesignHeight/2 - 40)

		ps.drawText(screen, ps.uiMisteryText, misteryTextX, misteryTextY, color.White)
	}

	if ps.currentStep >= 2 {
		widthText, _ := text.Measure(squid, ps.textFace, 0)
		squidTextX := float64(config.DesignWidth/2) - widthText/2
		squidTextY := float64(config.DesignHeight/2 - 20)

		ps.drawText(screen, ps.uiSquidText, squidTextX, squidTextY, color.White)
	}
	if ps.currentStep >= 3 {
		widthText, _ := text.Measure(crab, ps.textFace, 0)
		crabTextX := float64(config.DesignWidth/2) - widthText/2
		crabTextY := float64(config.DesignHeight / 2)

		ps.drawText(screen, ps.uiCrabText, crabTextX, crabTextY, color.White)
	}
	if ps.currentStep >= 4 {
		widthText, _ := text.Measure(octopus, ps.textFace, 0)
		octopusTextX := float64(config.DesignWidth/2) - widthText/2
		octopusTextY := float64(config.DesignHeight/2 + 20)

		ps.drawText(screen, ps.uiOctopusText, octopusTextX, octopusTextY, color.White)
	}
}

func (ps *PresentationState) NextState() states.StateId {
	return ps.nextState
}

func (ps *PresentationState) typeWriteText(text string) string {
	tmpText := text
	if ps.textIndex <= len(text) {
		tmpText = text[0:ps.textIndex]
		ps.textIndex++
	} else {
		ps.currentStep++
		ps.textIndex = 0
	}
	return tmpText
}

func (ps *PresentationState) drawText(screen *ebiten.Image, textstr string, posX, posY float64, color color.Color) {
	textOp := &text.DrawOptions{}
	textOp.GeoM.Translate(posX, posY)
	textOp.ColorScale.ScaleWithColor(color)
	text.Draw(screen, textstr, ps.textFace, textOp)
}

const title string = "SPACE INVADERS"
const mistery string = "= ? MISTERY"
const squid string = "= 30 POINTS"
const crab string = "= 20 POINTS"
const octopus string = "= 10 POINTS"
