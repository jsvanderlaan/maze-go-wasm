package utils

import (
	"bytes"
	"image"
	"image/color"
	"image/png"
	"log"
	"syscall/js"
	"time"

	"github.com/kelindar/bitmap"
)

func AddPath(img *image.NRGBA, mazeOptions MazeOptions, cellOptions CellOptions, path []uint32) {
	cellWidth := cellOptions.Width
	cellHeight := cellOptions.Height
	for x := 0; x < mazeOptions.Width; x++ {
		for y := 0; y < mazeOptions.Height; y++ {
			i := x + (y * mazeOptions.Width)
			xImg := x * cellWidth
			yImg := y * cellHeight
			for _, p := range path {
				if i == int(p) {
					Rect(img, color.RGBA{R: 0, G: 0, B: 255, A: 255}, xImg+(cellWidth/3), yImg+(cellHeight/3), xImg+cellWidth-(cellWidth/3), yImg+cellHeight-(cellHeight/3))
				}
			}
		}
	}
}

func AddStartEnd(img *image.NRGBA, mazeOptions MazeOptions, cellOptions CellOptions, start uint32, end uint32) {
	cellWidth := cellOptions.Width
	cellHeight := cellOptions.Height
	for x := 0; x < mazeOptions.Width; x++ {
		for y := 0; y < mazeOptions.Height; y++ {
			i := x + (y * mazeOptions.Width)
			xImg := x * cellWidth
			yImg := y * cellHeight

			if i == int(start) {
				Rect(img, color.RGBA{R: 0, G: 255, B: 0, A: 255}, xImg+(cellWidth/4), yImg+(cellHeight/4), xImg+cellWidth-(cellWidth/4), yImg+cellHeight-(cellHeight/4))
			}
			if i == int(end) {
				Rect(img, color.RGBA{R: 255, G: 0, B: 0, A: 255}, xImg+(cellWidth/4), yImg+(cellHeight/4), xImg+cellWidth-(cellWidth/4), yImg+cellHeight-(cellHeight/4))
			}
		}
	}
}

func MazeToImg(bm bitmap.Bitmap, mazeOptions MazeOptions, cellOptions CellOptions) image.NRGBA {
	mazeWidth := mazeOptions.Width
	mazeHeight := mazeOptions.Height

	cellWidth := cellOptions.Width
	cellHeight := cellOptions.Height
	cellShape := int(mazeOptions.Shape)
	borderCol := cellOptions.BorderColor

	bmCell := cellShape + 1

	imageWidth := mazeWidth * cellWidth
	imageHeight := mazeHeight * cellHeight

	img := image.NewNRGBA(image.Rect(0, 0, imageWidth+1, imageHeight+1))
	bmLength := mazeHeight * mazeWidth * bmCell

	for i := 0; i < bmLength; i += bmCell {

		bmI := uint32(i)
		if bm.Contains(bmI) {
			x := (i / bmCell % mazeWidth) * cellWidth
			y := (i / bmCell / mazeWidth) * cellHeight
			// log.Printf("x %d y %d before", x, y)
			// Rect(img, color.Black, x, y, x+cellWidth, y+cellHeight)

			if !bm.Contains(bmI + 1) {
				HLine(img, borderCol, y+cellHeight, x, x+cellWidth)
			}
			if !bm.Contains(bmI + 2) {
				VLine(img, borderCol, x, y, y+cellHeight)
			}
			if !bm.Contains(bmI + 3) {
				HLine(img, borderCol, y, x, x+cellWidth)
			}
			if !bm.Contains(bmI + 4) {
				VLine(img, borderCol, x+cellWidth, y, y+cellHeight)
			}
		}
	}

	return *img
}

func ImgToJs(img image.Image) js.Value {
	start := time.Now()
	buf := new(bytes.Buffer)
	log.Printf("create buffer: %v", time.Since(start))

	start = time.Now()
	enc := &png.Encoder{
		CompressionLevel: png.NoCompression,
	}
	error := enc.Encode(buf, img)
	log.Printf("encode: %v", time.Since(start))

	if error != nil {
		log.Fatal("Encoding Image failed", error)
	}
	bytes := buf.Bytes()
	dst := js.Global().Get("Uint8Array").New(len(bytes))
	start = time.Now()
	js.CopyBytesToJS(dst, bytes)
	log.Printf("CopyBytesToJS: %v", time.Since(start))
	return dst
}

func JsToImg(array js.Value) image.Image {
	inBuf := make([]uint8, array.Get("byteLength").Int())
	js.CopyBytesToGo(inBuf, array)

	reader := bytes.NewReader(inBuf)

	sourceImage, _, err := image.Decode(reader)

	if err != nil {
		log.Fatal("Failed to load image", err)
	}

	return sourceImage
}

func DetermineWidth(mazeHeight int, image image.Image) int {
	imageWidth := image.Bounds().Dx()
	imageHeight := image.Bounds().Dy()
	return int(float32(mazeHeight) * (float32(imageWidth) / float32(imageHeight)))
}
