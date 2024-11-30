package utils

import (
	"image"
	"image/color"
)

func ToGrayScale(img image.Image) image.Image {
	return &GrayScaleImg{img}
}

type GrayScaleImg struct {
	image.Image
}

func (m *GrayScaleImg) At(x, y int) color.Color {
	col := m.Image.At(x, y)
	return color.GrayModel.Convert(col)
}
