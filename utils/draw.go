package utils

import (
	"image"
	"image/color"
)

func HLine(img *image.NRGBA, col color.Color, y, x1, x2 int) {
	for ; x1 <= x2; x1++ {
		img.Set(x1, y, col)
	}
}

func VLine(img *image.NRGBA, col color.Color, x, y1, y2 int) {
	for ; y1 <= y2; y1++ {
		img.Set(x, y1, col)
	}
}
