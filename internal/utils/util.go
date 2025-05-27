package utils

import (
	"bytes"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"

	"github.com/programatta/spaceinvaders/internal/assets/fonts"
)

func SpriteFromArray(data [][]int, pixelSize int, colorOn color.Color) *ebiten.Image {
	h := len(data)
	w := len(data[0])
	img := ebiten.NewImage(w*pixelSize, h*pixelSize)

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if data[y][x] == 1 {
				rect := ebiten.NewImage(pixelSize, pixelSize)
				rect.Fill(colorOn)

				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(x*pixelSize), float64(y*pixelSize))
				img.DrawImage(rect, op)
			}
		}
	}

	return img
}

func LoadEmbeddedFont(size float64) *text.GoTextFace {
	faceSource, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.FontFiles))
	if err != nil {
		log.Fatal(err)
	}

	return &text.GoTextFace{
		Source: faceSource,
		Size:   size,
	}
}
