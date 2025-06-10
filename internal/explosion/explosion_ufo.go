package explosion

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/programatta/spaceinvaders/internal/common"
	"github.com/programatta/spaceinvaders/internal/config"
	"github.com/programatta/spaceinvaders/internal/sprite"
)

type ExplosionUfo struct {
	sprite   sprite.Sprite
	posX     float32
	posY     float32
	notifier common.Notifier
	time     float32
	remove   bool
}

func NewExplosionUfo(posX, posY float32, sprite sprite.Sprite, notifier common.Notifier) *ExplosionUfo {
	return &ExplosionUfo{sprite: sprite, posX: posX, posY: posY, notifier: notifier}
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
}
