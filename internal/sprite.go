package internal

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	Image *ebiten.Image
	Color color.Color
	Data  [][]int
}
