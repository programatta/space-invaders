package internal

import "github.com/hajimehoshi/ebiten/v2"

type ExplosionCannon struct {
	sprites       []Sprite
	currentSprite uint8
	posX          float32
	posY          float32
	notifier      Notifier
	time          float32
	repeatCount   uint8
	remove        bool
}

func NewExplosionCannon(posX, posY float32, sprite1, sprite2 Sprite, notifier Notifier) *ExplosionCannon {
	sprites := []Sprite{sprite1, sprite2}
	return &ExplosionCannon{sprites: sprites, currentSprite: 0, posX: posX, posY: posY, notifier: notifier, time: 0, repeatCount: 8, remove: false}
}

func (ec *ExplosionCannon) CanRemove() bool {
	return ec.remove
}

func (ec *ExplosionCannon) Update() {
	if ec.repeatCount > 0 {
		ec.time += dt
		if ec.time >= 0.35 {
			ec.currentSprite = (ec.currentSprite + 1) % 2
			ec.time = 0
			ec.repeatCount--
		}

		if ec.repeatCount == 0 {
			ec.remove = true
			ec.notifier.OnResetCannon()
		}
	}
}

func (ec *ExplosionCannon) Draw(screen *ebiten.Image) {
	opExplosionCannon := &ebiten.DrawImageOptions{}
	opExplosionCannon.GeoM.Translate(float64(ec.posX), float64(ec.posY))
	screen.DrawImage(ec.sprites[ec.currentSprite].Image, opExplosionCannon)
}
