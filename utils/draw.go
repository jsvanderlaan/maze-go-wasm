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

func Rect(img *image.NRGBA, col color.Color, x1, y1, x2, y2 int) {
	for x := x1; x <= x2; x++ {
		for y := y1; y <= y2; y++ {
			img.Set(x, y, col)
		}
	}
}
