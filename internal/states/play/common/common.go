package common

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Notifier interface {
	OnChangeDirection(newDirection float32)
	OnCreateCannonBullet(posX, posY float32, color color.Color)
	OnCreateAlienBullet(posX, posY float32, color color.Color)
	OnResetCannon()
	OnResetUfo()
}

type Manageer interface {
	Update()
	Draw(screen *ebiten.Image)
}

type Eraser interface {
	CanRemove() bool
}
