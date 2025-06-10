package common

import "image/color"

type Notifier interface {
	OnChangeDirection(newDirection float32)
	OnCreateCannonBullet(posX, posY float32, color color.Color)
	OnCreateAlienBullet(posX, posy float32, color color.Color)
	OnResetUfo()
	OnResetCannon()
}
