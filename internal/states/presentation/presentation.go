package presentation

import (
	"image/color"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/programatta/spaceinvaders/internal/config"
	"github.com/programatta/spaceinvaders/internal/sprite"
	"github.com/programatta/spaceinvaders/internal/states"
)

type PresentationState struct {
	spriteCreator        *sprite.SpriteCreator
	textFace             *text.GoTextFace
	time                 float32
	timeBetweenStep      float32
	nextState            states.StateId
	uiTitleText          string
	uiMisteryText        string
	uiSquidText          string
	uiCrabText           string
	uiOctopusText        string
	textIndex            int
	currentStep          int
	canPlay              bool
	reloadScreenTime     float32
	pressSpaceTime       float32
	showPressSpaceToPlay bool
}

func NewPresentationState(spriteCreator *sprite.SpriteCreator, textFace *text.GoTextFace) *PresentationState {
	return &PresentationState{
		spriteCreator:        spriteCreator,
		textFace:             textFace,
		time:                 0,
		timeBetweenStep:      stepDelay,
		nextState:            states.Presentation,
		uiTitleText:          "",
		uiMisteryText:        "",
		uiSquidText:          "",
		uiCrabText:           "",
		uiOctopusText:        "",
		textIndex:            0,
		currentStep:          0,
		canPlay:              false,
		reloadScreenTime:     0,
		pressSpaceTime:       0,
		showPressSpaceToPlay: true,
	}
}

/*
Implementación interface State
*/

func (ps *PresentationState) Start() {
	ps.time = 0
	ps.timeBetweenStep = stepDelay
	ps.nextState = states.Presentation
	ps.uiTitleText = ""
	ps.uiMisteryText = ""
	ps.uiSquidText = ""
	ps.uiCrabText = ""
	ps.uiOctopusText = ""
	ps.textIndex = 0
	ps.currentStep = 0
	ps.reloadScreenTime = 0
	ps.pressSpaceTime = 0
	ps.showPressSpaceToPlay = true
}

func (ps *PresentationState) ProcessEvents() {
	if ps.canPlay && inpututil.IsKeyJustReleased(ebiten.KeySpace) {
		ps.nextState = states.Play
	}

	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		os.Exit(0)
	}
}

func (ps *PresentationState) Update() {
	ps.time += config.Dt
	if ps.time >= ps.timeBetweenStep {
		if ps.currentStep == 0 {
			ps.uiTitleText = ps.typeWriteText(title)
		}
		if ps.currentStep == 1 {
			ps.uiMisteryText = ps.typeWriteText(mistery)
		}
		if ps.currentStep == 2 {
			ps.uiSquidText = ps.typeWriteText(squild)
		}
		if ps.currentStep == 3 {
			ps.uiCrabText = ps.typeWriteText(crab)
		}
		if ps.currentStep == 4 {
			ps.uiOctopusText = ps.typeWriteText(octopus)
		}
		if ps.currentStep == 5 {
			ps.timeBetweenStep = 0
			ps.canPlay = true
			ps.reloadScreenTime += config.Dt
			ps.pressSpaceTime += config.Dt
			if ps.pressSpaceTime >= blinkDelay {
				ps.pressSpaceTime = 0
				ps.showPressSpaceToPlay = !ps.showPressSpaceToPlay
			}

			if ps.reloadScreenTime > reloadDelay {
				ps.timeBetweenStep = stepDelay
				ps.uiTitleText = ""
				ps.reloadScreenTime = 0
				ps.currentStep = 0
			}
		}
		ps.time = 0
	}
}

func (ps *PresentationState) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x03, 0x04, 0x5e, 0xFF})

	if ps.currentStep >= 0 {
		titleX := float64(config.DesignWidth/2 - 40)
		titleY := float64(config.DesignHeight/2 - 60)
		ps.drawText(screen, ps.uiTitleText, titleX, titleY, color.White)
	}

	if ps.currentStep >= 1 {
		misteryTextX := float64(config.DesignWidth/2 - 40)
		misteryTextY := float64(config.DesignHeight/2 - 40)

		ufoSprite, _ := ps.spriteCreator.SpriteByName("ufo")
		ps.drawIcon(screen, ufoSprite, misteryTextX, misteryTextY)
		ps.drawText(screen, ps.uiMisteryText, misteryTextX+20, misteryTextY+1, ufoSprite.Color)
	}

	if ps.currentStep >= 2 {
		squidTextX := float64(config.DesignWidth/2 - 36)
		squidTextY := float64(config.DesignHeight/2 - 20)

		squidSprite, _ := ps.spriteCreator.SpriteByName("squid1")
		ps.drawIcon(screen, squidSprite, squidTextX, squidTextY)
		ps.drawText(screen, ps.uiSquidText, squidTextX+16, squidTextY+1, squidSprite.Color)
	}

	if ps.currentStep >= 3 {
		crabTextX := float64(config.DesignWidth/2 - 37)
		crabTextY := float64(config.DesignHeight / 2)

		crabSprite, _ := ps.spriteCreator.SpriteByName("crab1")
		ps.drawIcon(screen, crabSprite, crabTextX, crabTextY)
		ps.drawText(screen, ps.uiCrabText, crabTextX+17, crabTextY+1, crabSprite.Color)
	}

	if ps.currentStep >= 4 {
		octopusTextX := float64(config.DesignWidth/2 - 37)
		octopusTextY := float64(config.DesignHeight/2 + 20)

		octopusSprite, _ := ps.spriteCreator.SpriteByName("octopus1")
		ps.drawIcon(screen, octopusSprite, octopusTextX, octopusTextY)
		ps.drawText(screen, ps.uiOctopusText, octopusTextX+17, octopusTextY+1, octopusSprite.Color)
	}

	if ps.currentStep >= 5 {
		if ps.showPressSpaceToPlay {
			playX := float64(config.DesignWidth/2 - 55)
			playY := float64(config.DesignHeight/2 + 50)
			ps.drawText(screen, "PRESS SPACE TO PLAY", playX, playY, color.White)
		}
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

func (ps *PresentationState) drawIcon(screen *ebiten.Image, sprite sprite.Sprite, posX, posY float64) {
	drawOp := &ebiten.DrawImageOptions{}
	drawOp.GeoM.Translate(posX, posY)
	screen.DrawImage(sprite.Image, drawOp)
}

const title string = "SPACE INVADERS"
const mistery string = "= ? MISTERY"
const squild string = "= 30 POINTS"
const crab string = "= 20 POINTS"
const octopus string = "= 10 POINTS"

const stepDelay float32 = 0.12  //en segundos.
const reloadDelay float32 = 5.0 //en segundos.
const blinkDelay float32 = 0.5  //en segundos.
