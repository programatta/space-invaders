package player

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/programatta/spaceinvaders/internal/common"
	"github.com/programatta/spaceinvaders/internal/config"
	"github.com/programatta/spaceinvaders/internal/sprite"
)

type Cannon struct {
	sprite       sprite.Sprite
	posX         float32
	posY         float32
	originalPosX float32
	dirX         float32
	speed        float32
	notifier     common.Notifier
	canFired     bool
	time         float32
	active       bool
}

func NewCannon(posX, posY float32, sprite sprite.Sprite, notifier common.Notifier) *Cannon {
	return &Cannon{
		sprite:       sprite,
		posX:         posX,
		posY:         posY,
		originalPosX: posX,
		notifier:     notifier,
		speed:        60,
		canFired:     true,
		active:       true,
	}
}

func (c *Cannon) Position() (float32, float32) {
	return c.posX, c.posY
}

func (c *Cannon) ProcessKeyEvents() {
	c.dirX = 0
	if c.active {
		if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
			c.dirX = 1
		} else if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
			c.dirX = -1
		}

		if c.canFired && ebiten.IsKeyPressed(ebiten.KeySpace) {
			c.canFired = false
			c.notifier.OnCreateCannonBullet(c.posX+6, c.posY, c.sprite.Color)
		}
	}
}

func (c *Cannon) Update() error {
	if c.active {
		if !c.canFired {
			c.time += config.Dt
			if c.time >= 0.35 {
				c.canFired = true
				c.time = 0
			}
		}

		c.posX += c.speed * config.Dt * c.dirX
		if c.posX <= 0 {
			c.posX = 0
		} else if c.posX+float32(c.sprite.Image.Bounds().Dx()) >= float32(config.DesignWidth) {
			c.posX = float32(config.DesignWidth) - float32(c.sprite.Image.Bounds().Dx())
		}
	}
	return nil
}

func (c *Cannon) Draw(screen *ebiten.Image) {
	if c.active {
		opCannon := &ebiten.DrawImageOptions{}
		opCannon.GeoM.Translate(float64(c.posX), float64(c.posY))
		screen.DrawImage(c.sprite.Image, opCannon)
	}
}

func (c *Cannon) IsActive() bool {
	return c.active
}

// Implementación de la interface Collider.
func (c *Cannon) Rect() (float32, float32, float32, float32) {
	width := float32(c.sprite.Image.Bounds().Dx())
	height := float32(c.sprite.Image.Bounds().Dy())
	return c.posX, c.posY, width, height
}

func (c *Cannon) OnCollide() {
	c.active = false
}

func (c *Cannon) Reset() {
	c.active = true
	c.posX = c.originalPosX
}
