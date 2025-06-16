package explosion

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/programatta/spaceinvaders/internal/config"
	"github.com/programatta/spaceinvaders/internal/sprite"
	"github.com/programatta/spaceinvaders/internal/states/play/common"
)

type ExplosionUfo struct {
	sprite   sprite.Sprite
	posX     float32
	posY     float32
	face     *text.GoTextFace
	score    uint16
	notifier common.Notifier
	time     float32
	remove   bool
}

func NewExplosionUfo(posX, posY float32, sprite sprite.Sprite, face *text.GoTextFace, score uint16, notifier common.Notifier) *ExplosionUfo {
	return &ExplosionUfo{sprite: sprite, posX: posX, posY: posY, face: face, score: score, notifier: notifier}
}

// Implementación de la interface Explosioner
func (eu *ExplosionUfo) CanRemove() bool {
	return eu.remove
}

func (eu *ExplosionUfo) Update() {
	eu.time += config.Dt
	if eu.time >= 0.35 {
		eu.time = 0
		eu.remove = true
		eu.notifier.OnResetUfo()
	}
}

func (eu *ExplosionUfo) Draw(screen *ebiten.Image) {
	opExplosionUfo := &ebiten.DrawImageOptions{}
	opExplosionUfo.GeoM.Translate(float64(eu.posX), float64(eu.posY))
	screen.DrawImage(eu.sprite.Image, opExplosionUfo)

	uiScoreMisteryText := fmt.Sprintf("+%03d", eu.score)
	widthText, _ := text.Measure(uiScoreMisteryText, eu.face, 0)
	textX := eu.posX + float32(eu.sprite.Image.Bounds().Dx())
	textY := eu.posY + float32(eu.sprite.Image.Bounds().Dy())/2 - 4

	if textX+float32(widthText)+4 > float32(config.DesignWidth) {
		textX = eu.posX - float32(widthText) - 4
	}
	op := &text.DrawOptions{}
	op.GeoM.Translate(float64(textX), float64(textY))
	op.ColorScale.ScaleWithColor(eu.sprite.Color)
	text.Draw(screen, uiScoreMisteryText, eu.face, op)
}
