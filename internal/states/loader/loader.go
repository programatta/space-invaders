package loader

import (
	"fmt"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/programatta/spaceinvaders/internal/config"
	"github.com/programatta/spaceinvaders/internal/sprite"
	"github.com/programatta/spaceinvaders/internal/states"
)

type LoaderState struct {
	spriteCreator *sprite.SpriteCreator
	textFace      *text.GoTextFace
	version       string
	time          float32
	nextState     states.StateId
}

func NewLoaderState(spriteCreator *sprite.SpriteCreator, textFace *text.GoTextFace, version string) *LoaderState {
	return &LoaderState{
		spriteCreator: spriteCreator,
		textFace:      textFace,
		version:       version,
		time:          0,
		nextState:     states.Loader,
	}
}

/*
Implementación interface State
*/
func (ls *LoaderState) Start()         {}
func (ls *LoaderState) ProcessEvents() {}

func (ls *LoaderState) Update() {
	ls.time += config.Dt
	if ls.time >= viewChangeDelay {
		ls.time = 0
		ls.nextState = states.Presentation
	}
}

func (ls *LoaderState) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x03, 0x04, 0x5e, 0xFF})

	crabSprite, _ := ls.spriteCreator.SpriteByName("crab1")
	drawOp := &ebiten.DrawImageOptions{}
	drawOp.GeoM.Translate(20, 10)
	drawOp.GeoM.Scale(5.0, 5.0)
	drawOp.GeoM.Rotate(rotationInRads)
	drawOp.ColorScale.ScaleAlpha(0.16)
	screen.DrawImage(crabSprite.Image, drawOp)

	uiTitleText := "SPACE INVADERS"
	titleX := float64(config.DesignWidth/2 - 95)
	titleY := float64(config.DesignHeight/2 - 65)

	titleOp := &text.DrawOptions{}
	titleOp.GeoM.Translate(titleX, titleY)
	titleOp.GeoM.Scale(2.0, 2.0)
	titleOp.ColorScale.ScaleWithColor(color.White)
	text.Draw(screen, uiTitleText, ls.textFace, titleOp)

	uiVersionText := fmt.Sprintf("V%s", ls.version)
	versionX := float64(config.DesignWidth/2 + 40)
	versionY := float64(config.DesignHeight/2 - 4)

	versionOp := &text.DrawOptions{}
	versionOp.GeoM.Translate(versionX, versionY)
	versionOp.ColorScale.ScaleWithColor(color.White)
	text.Draw(screen, uiVersionText, ls.textFace, versionOp)

	uiTechText := "Powered by Golang & Ebiten"
	techX := float64(config.DesignWidth/2 - 77)
	techY := float64(config.DesignHeight/2 + 50)

	techOp := &text.DrawOptions{}
	techOp.GeoM.Translate(techX, techY)
	techOp.ColorScale.ScaleWithColor(color.White)
	text.Draw(screen, uiTechText, ls.textFace, techOp)
}

func (ls *LoaderState) NextState() states.StateId {
	return ls.nextState
}

const viewChangeDelay float32 = 1.35 //en segundos.
const rotationInRads float64 = 45 * math.Pi / 180
