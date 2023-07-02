package utils

import (
	"image"
	"image/color"
)

func ImgToBlackWhite(img image.Image, threshold uint8) image.Image {
	return &BlackWhiteImg{img, threshold}
}

func colorToBlackWhite(c color.Color, threshold uint8) color.Color {
	if IsBlack(c, threshold) {
		return color.Black
	}

	return color.White
}

func IsBlack(c color.Color, threshold uint8) bool {
	r, g, b, a := c.RGBA()
	if a == 0 {
		return false
	}

	grayscale := (19595*r + 38470*g + 7471*b + 1<<15) >> 24

	if uint8(grayscale) < threshold {
		return true
	} else {
		return false
	}
}

type BlackWhiteImg struct {
	image.Image
	threshold uint8
}

func (m *BlackWhiteImg) At(x, y int) color.Color {
	col := m.Image.At(x, y)
	return colorToBlackWhite(col, m.threshold)
}
