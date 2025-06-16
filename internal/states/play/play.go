package play

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/programatta/spaceinvaders/internal/sounds"
	"github.com/programatta/spaceinvaders/internal/sprite"
	"github.com/programatta/spaceinvaders/internal/states"
)

type PlayState struct {
	spriteCreator *sprite.SpriteCreator
	textFace      *text.GoTextFace
	soundEffects  *sounds.SoundEffects
	nextState     states.StateId
}

func NewPlayState(spriteCreator *sprite.SpriteCreator, textFace *text.GoTextFace, soundEffects *sounds.SoundEffects) *PlayState {
	return &PlayState{
		spriteCreator: spriteCreator,
		textFace:      textFace,
		soundEffects:  soundEffects,
		nextState:     states.Play,
	}
}

func (ps *PlayState) ProcessEvents() {}

func (ps *PlayState) Update() {}

func (ps *PlayState) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x03, 0x04, 0x5e, 0xFF})
	ebitenutil.DebugPrint(screen, "Play")
}

func (ps *PlayState) NextState() states.StateId {
	return ps.nextState
}
