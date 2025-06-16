package common

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Notifier interface {
	OnChangeDirection(newDirection float32)
	OnCreateCannonBullet(posX, posY float32, color color.Color)
	OnCreateAlienBullet(posX, posy float32, color color.Color)
	OnResetUfo()
	OnResetCannon()
}

type Manageer interface {
	Update()
	Draw(screen *ebiten.Image)
}

type Eraser interface {
	CanRemove() bool
}
