package internal

import (
	"image/color"
	"math/rand"
	"slices"

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

type Explosioner interface {
	Manageer
	Eraser
}

type Collider interface {
	Rect() (float32, float32, float32, float32)
	OnCollide()
}

type Game struct {
	spriteCreator     *SpriteCreator
	cannon            *Cannon
	cannonCount       uint8
	bullets           []*Bullet
	bunkers           []*Bunker
	ufo               *Ufo
	enemies           []*Alien
	enemiesCurrentDir float32
	newDirection      float32
	explosions        []Explosioner
	alienFireTime     float32
}

func NewGame() *Game {
	spriteCreator := NewSpriteCreator()

	game := &Game{}
	game.spriteCreator = spriteCreator

	spriteCannon, _ := spriteCreator.SpriteByName("cannon")
	game.cannon = NewCannon(float32(0), float32(DesignHeight-10), spriteCannon, game)

	bunkers := createBunkers(spriteCreator)
	game.bunkers = bunkers

	ufoSprite, _ := spriteCreator.SpriteByName("ufo")
	ufo := NewUfo(-20, 5, ufoSprite)
	game.ufo = ufo

	enemies := createEnemies(spriteCreator, game)
	game.enemies = enemies
	game.enemiesCurrentDir = 1
	game.newDirection = 1
	game.cannonCount = 3
	return game
}

// Implementación de la interface esperada por ebiten.
func (g *Game) Update() error {
	g.cannon.ProcessKeyEvents()

	g.cannon.Update()

	g.alienFireTime += dt
	if g.alienFireTime > 0.400 {
		if len(g.enemies) > 0 {
			pos := rand.Intn(len(g.enemies))
			g.enemies[pos].Fire()
		}
		g.alienFireTime = 0
	}

	for _, bullet := range g.bullets {
		bullet.Update()
	}

	g.ufo.Update()
	for _, enemy := range g.enemies {
		enemy.ChangeDirection(g.enemiesCurrentDir)
		enemy.Update()
	}

	for _, explosion := range g.explosions {
		explosion.Update()
	}

	//Colisiones.
	for _, bullet := range g.bullets {
		if bullet.dirY < 0 {
			//Bala de cañon
			for _, bunker := range g.bunkers {
				if g.checkCollision(bullet, bunker) {
					if bunker.DoDamage(bullet.posX, bullet.posY, -1) {
						bullet.OnCollide()
					}
				}
			}
			for _, enemy := range g.enemies {
				if g.checkCollision(bullet, enemy) {
					bullet.OnCollide()
					enemy.OnCollide()

					alienExplosionSprite, _ := g.spriteCreator.SpriteByName("alienExplosion")
					enemyX, enemyY := enemy.Position()
					explosion := NewExplosion(enemyX, enemyY, alienExplosionSprite, enemy.Color())
					g.explosions = append(g.explosions, explosion)
				}
			}
			if g.checkCollision(bullet, g.ufo) {
				bullet.OnCollide()
				g.ufo.OnCollide()

				ufoExplosionSprite, _ := g.spriteCreator.SpriteByName("ufoExplosion")
				ufoX, ufoY := g.ufo.Position()
				explosionUfo := NewExplosionUfo(ufoX, ufoY, ufoExplosionSprite, g)
				g.explosions = append(g.explosions, explosionUfo)
			}
		} else {
			//Bala de alien.
			for _, bunker := range g.bunkers {
				if g.checkCollision(bullet, bunker) {
					if bunker.DoDamage(bullet.posX, bullet.posY, 1) {
						bullet.OnCollide()
					}
				}
			}
			if g.checkCollision(bullet, g.cannon) {
				cannonExplosion1Sprite, _ := g.spriteCreator.SpriteByName("cannonExplosion1")
				cannonExplosion2Sprite, _ := g.spriteCreator.SpriteByName("cannonExplosion2")
				explosionCannon := NewExplosionCannon(g.cannon.posX, g.cannon.posY, cannonExplosion1Sprite, cannonExplosion2Sprite, g)
				g.explosions = append(g.explosions, explosionCannon)
				if g.cannonCount > 0 {
					g.cannonCount--
					g.cannon.OnCollide()
				}
				bullet.OnCollide()
			}
		}
	}

	//Colisines alien con bunker
	for _, enemy := range g.enemies {
		for _, bunker := range g.bunkers {
			if g.checkCollision(enemy, bunker) {
				bunker.OnCollide()
				break
			}
		}
	}

	if len(g.bullets) > 0 {
		g.bullets = slices.DeleteFunc(g.bullets, func(bullet *Bullet) bool {
			return bullet.CanRemove()
		})
	}

	if len(g.enemies) > 0 {
		g.enemies = slices.DeleteFunc(g.enemies, func(alien *Alien) bool {
			return alien.CanRemove()
		})
	} else {
		g.reset()
		g.cannon.Reset()
	}

	if len(g.explosions) > 0 {
		g.explosions = slices.DeleteFunc(g.explosions, func(explosion Explosioner) bool {
			return explosion.CanRemove()
		})
	}

	if len(g.bunkers) > 0 {
		g.bunkers = slices.DeleteFunc(g.bunkers, func(bunker *Bunker) bool {
			return bunker.CanRemove()
		})
	}

	if g.newDirection != g.enemiesCurrentDir {
		g.enemiesCurrentDir = g.newDirection
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x03, 0x04, 0x5e, 0xFF})

	g.ufo.Draw(screen)

	for _, enemy := range g.enemies {
		enemy.Draw(screen)
	}

	for _, explosion := range g.explosions {
		explosion.Draw(screen)
	}

	g.cannon.Draw(screen)
	for _, bullet := range g.bullets {
		bullet.Draw(screen)
	}

	for _, bunker := range g.bunkers {
		bunker.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return DesignWidth, DesignHeight
}

// Implementación de la interface Notifier
func (g *Game) OnCreateCannonBullet(posX, posY float32, color color.Color) {
	spriteBullet, _ := g.spriteCreator.SpriteByName("bullet")
	bullet := NewBullet(posX, posY, spriteBullet, color, -1)
	g.bullets = append(g.bullets, bullet)
}

func (g *Game) OnCreateAlienBullet(posX, posY float32, color color.Color) {
	spriteBullet, _ := g.spriteCreator.SpriteByName("bullet")
	bullet := NewBullet(posX, posY, spriteBullet, color, 1)
	g.bullets = append(g.bullets, bullet)
}

func (g *Game) OnChangeDirection(newDirection float32) {
	g.newDirection = newDirection
}

func (g *Game) OnResetUfo() {
	g.ufo.Reset()
}

func (g *Game) OnResetCannon() {
	if g.cannonCount == 0 {
		g.reset()
		g.cannonCount = 3
	}
	g.cannon.Reset()
}

func (g *Game) checkCollision(sourceObj, targetObj Collider) bool {
	sx0, sy0, sw, sh := sourceObj.Rect()
	tx0, ty0, tw, th := targetObj.Rect()

	hasCollision := sx0 < tx0+tw && sx0+sw > tx0 && sy0 < ty0+th && sh+sy0 > ty0

	return hasCollision
}

func (g *Game) reset() {
	bunkers := createBunkers(g.spriteCreator)
	enemies := createEnemies(g.spriteCreator, g)

	g.enemies = enemies
	g.bullets = []*Bullet{}
	g.bunkers = bunkers
	g.explosions = []Explosioner{}
	g.enemiesCurrentDir = 1
	g.newDirection = 1
	g.alienFireTime = 0
}

func createBunkers(spriteCreator *SpriteCreator) []*Bunker {
	bunkerSprite, _ := spriteCreator.SpriteByName("bunker")

	var posX float32 = 27
	bunkers := []*Bunker{}
	for range 4 {
		bunker := NewBunker(posX, float32(DesignHeight-40), bunkerSprite)
		bunkers = append(bunkers, bunker)
		posX += float32(bunkerSprite.Image.Bounds().Dx()) + 20
	}
	return bunkers
}

func createEnemies(spriteCreator *SpriteCreator, notifier Notifier) []*Alien {
	enemies := []*Alien{}

	squids := createSquids(11, 1, 11, 35, spriteCreator, notifier)
	enemies = append(enemies, squids...)

	crabs := createCrabs(11, 2, 10, 50, spriteCreator, notifier)
	enemies = append(enemies, crabs...)

	octopuses := createOctopuses(11, 2, 9, 80, spriteCreator, notifier)
	enemies = append(enemies, octopuses...)

	return enemies
}

func createCrabs(count, rows uint8, initX, initY float32, spriteCreator *SpriteCreator, notifier Notifier) []*Alien {
	sprite1, _ := spriteCreator.SpriteByName("crab1")
	sprite2, _ := spriteCreator.SpriteByName("crab2")
	crabs := []*Alien{}

	posX := initX
	posY := initY
	for i := range count * rows {
		crab := NewAlien(posX, posY, sprite1, sprite2, notifier)
		crabs = append(crabs, crab)
		posX += float32(sprite1.Image.Bounds().Dx() + 6)
		if i > 0 && (i+1)%count == 0 {
			posX = initX
			posY += float32(sprite1.Image.Bounds().Dy() + 5)
		}
	}
	return crabs
}

func createOctopuses(count, rows uint8, initX, initY float32, spriteCreator *SpriteCreator, notifier Notifier) []*Alien {
	sprite1, _ := spriteCreator.SpriteByName("octopus1")
	sprite2, _ := spriteCreator.SpriteByName("octopus2")
	octopuses := []*Alien{}

	posX := initX
	posY := initY
	for i := range count * rows {
		octopus := NewAlien(posX, posY, sprite1, sprite2, notifier)
		octopuses = append(octopuses, octopus)
		posX += float32(sprite1.Image.Bounds().Dx() + 5)
		if i > 0 && (i+1)%count == 0 {
			posX = initX
			posY += float32(sprite1.Image.Bounds().Dy() + 5)
		}
	}
	return octopuses
}

func createSquids(count, rows uint8, initX, initY float32, spriteCreator *SpriteCreator, notifier Notifier) []*Alien {
	sprite1, _ := spriteCreator.SpriteByName("squid1")
	sprite2, _ := spriteCreator.SpriteByName("squid2")
	squids := []*Alien{}

	posX := initX
	posY := initY
	for i := range count * rows {
		squid := NewAlien(posX, posY, sprite1, sprite2, notifier)
		squids = append(squids, squid)
		posX += float32(sprite1.Image.Bounds().Dx() + 9)
		if i > 0 && (i+1)%count == 0 {
			posX = initX
			posY += float32(sprite1.Image.Bounds().Dy())
		}
	}
	return squids
}

const dt float32 = float32(1.0 / 60)

const WindowWidth int = 642
const WindowHeight int = 642
const DesignWidth int = WindowWidth / 3
const DesignHeight int = WindowHeight / 3
