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
	nextState            states.StateId
	time                 float32
	uiTitleText          string
	uiMisteryText        string
	uiSquidText          string
	uiCrabText           string
	uiOctopusText        string
	textIndex            int
	currentStep          int
	pressSpaceTime       float32
	reloadScreenTime     float32
	showPressSpaceToPlay bool
	innerState           presentationInnerStateId
}

func NewPresentationState(spriteCreator *sprite.SpriteCreator, textFace *text.GoTextFace) *PresentationState {
	return &PresentationState{
		spriteCreator: spriteCreator,
		textFace:      textFace,
	}
}

func (ps *PresentationState) Start() {
	ps.nextState = states.Presentation
	ps.innerState = showScores
	ps.time = 0
	ps.pressSpaceTime = 0
	ps.uiTitleText = ""
	ps.reloadScreenTime = 0
	ps.currentStep = 0
	ps.textIndex = 0
}

func (ps *PresentationState) ProcessEvents() {
	if inpututil.IsKeyJustReleased(ebiten.KeySpace) {
		ps.nextState = states.Play
	}

	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		os.Exit(0)
	}
}

func (ps *PresentationState) Update() {
	switch ps.innerState {
	case showScores:
		ps.time += config.Dt
		if ps.time >= stepDelay {
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
				ps.innerState = showToPlay
			}
			ps.time = 0
		}
	case showToPlay:
		ps.pressSpaceTime += config.Dt
		if ps.pressSpaceTime >= blinkDelay {
			ps.pressSpaceTime = 0
			ps.showPressSpaceToPlay = !ps.showPressSpaceToPlay
		}

		ps.reloadScreenTime += config.Dt
		if ps.reloadScreenTime > reloadDelay {
			ps.uiTitleText = ""
			ps.reloadScreenTime = 0
			ps.currentStep = 0
			ps.innerState = showScores
		}
	}
}

func (ps *PresentationState) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x03, 0x04, 0x5e, 0xFF})

	var lineX float64 = 0
	var lineY float64 = 0
	if ps.currentStep >= 0 {
		widthText, _ := text.Measure(title, ps.textFace, 0)
		lineX = float64(config.DesignWidth/2) - widthText/2
		lineY = float64(config.DesignHeight/2 - 60)

		ps.drawText(screen, ps.uiTitleText, lineX, lineY, color.White)
		lineY += 20
	}
	if ps.currentStep >= 1 {
		ufoSprite, _ := ps.spriteCreator.SpriteByName("ufo")
		ps.drawIcon(screen, ufoSprite, lineX, lineY)
		ps.drawText(screen, ps.uiMisteryText, lineX+20, lineY+1, ufoSprite.Color)
		lineY += 20
	}
	if ps.currentStep >= 2 {
		squidSprite, _ := ps.spriteCreator.SpriteByName("squid1")
		ps.drawIcon(screen, squidSprite, lineX+4, lineY)
		ps.drawText(screen, ps.uiSquidText, lineX+20, lineY+1, squidSprite.Color)
		lineY += 20
	}
	if ps.currentStep >= 3 {
		crabSprite, _ := ps.spriteCreator.SpriteByName("crab1")
		ps.drawIcon(screen, crabSprite, lineX+3, lineY)
		ps.drawText(screen, ps.uiCrabText, lineX+20, lineY+1, crabSprite.Color)
		lineY += 20
	}
	if ps.currentStep >= 4 {
		octopusSprite, _ := ps.spriteCreator.SpriteByName("octopus1")
		ps.drawIcon(screen, octopusSprite, lineX+3, lineY)
		ps.drawText(screen, ps.uiOctopusText, lineX+20, lineY+1, octopusSprite.Color)
	}
	if ps.currentStep >= 5 {
		if ps.showPressSpaceToPlay {
			widthText, _ := text.Measure(pressToPlay, ps.textFace, 0)
			playX := float64(config.DesignWidth/2) - widthText/2
			playY := float64(config.DesignHeight/2 + 50)
			ps.drawText(screen, pressToPlay, playX, playY, color.White)
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
const squid string = "= 30 POINTS"
const crab string = "= 20 POINTS"
const octopus string = "= 10 POINTS"
const pressToPlay string = "PRESS SPACE TO PLAY"

const stepDelay float32 = 0.12  //en segundos.
const reloadDelay float32 = 5.0 //en segundos.
const blinkDelay float32 = 0.5  //en segundos.
