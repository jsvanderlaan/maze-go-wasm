package utils

import (
	"image"
	"image/color"
	"image/draw"
	"log"

	"golang.org/x/image/font"

	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/fixed"

	_ "embed"
)

//go:embed fonts/Montserrat-Bold.ttf
var montserratBold []byte

//go:embed fonts/Pacifico-Regular.ttf
var pacificoRegular []byte

func getTextWidth(text string, fontFace font.Face) int {
	drawer := &font.Drawer{
		Face: fontFace,
	}
	return drawer.MeasureString(text).Ceil()
}

func getTextHeight(fontFace font.Face) (int, int) {
	metrics := fontFace.Metrics()
	ascent := metrics.Ascent.Round()
	descent := metrics.Descent.Round()
	log.Printf("ascent %v descent %v", ascent, descent)

	return ascent, descent
}

func RenderTextToJPG(text string, outline bool) (image.Image, error) {
	var fontBytes []byte
	if outline {
		fontBytes = pacificoRegular
	} else {
		fontBytes = montserratBold
	}

	ttfFont, err := opentype.Parse(fontBytes)
	if err != nil {
		return nil, err
	}

	const maxChars = 200
	if len(text) > maxChars {
		text = text[:maxChars]
	}

	fontSize := 100.0

	faceOptions := &opentype.FaceOptions{
		Size:    fontSize,
		DPI:     72,
		Hinting: font.HintingFull,
	}
	face, err := opentype.NewFace(ttfFont, faceOptions)
	if err != nil {
		return nil, err
	}

	textWidth := getTextWidth(text, face)
	textHeightAbove, textHeightBelow := getTextHeight(face)
	textHeight := textHeightAbove + textHeightBelow
	log.Printf("textHeight %v textWidth %v", textHeightAbove+textHeightBelow, textWidth)

	pX := 30

	var textColor color.Color
	var bgColor color.Color

	if outline {
		textColor = color.White
		bgColor = color.Black
	} else {
		textColor = color.Black
		bgColor = color.White
	}

	rect := image.Rect(0, 0, textWidth+(pX*2), textHeight)
	img := image.NewRGBA(rect)
	draw.Draw(img, img.Bounds(), &image.Uniform{C: textColor}, image.Point{}, draw.Src)

	drawer := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(bgColor),
		Face: face,
		Dot:  fixed.P(pX, textHeightAbove),
	}
	drawer.DrawString(text)

	return img.SubImage(rect), nil
}
