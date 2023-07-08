package utils

import (
	"bytes"
	"image"
	"image/png"
	"log"
	"syscall/js"
	"time"

	"github.com/kelindar/bitmap"
)

func MazeToImg(bm bitmap.Bitmap, mazeOptions MazeOptions, cellOptions CellOptions) image.Image {
	mazeWidth := mazeOptions.Width
	mazeHeight := mazeOptions.Height

	cellWidth := cellOptions.Width
	cellHeight := cellOptions.Height
	cellShape := int(mazeOptions.Shape)
	borderCol := cellOptions.BorderColor

	bmCell := cellShape + 1

	imageWidth := mazeWidth * cellWidth
	imageHeight := mazeHeight * cellHeight

	img := image.NewNRGBA(image.Rect(0, 0, imageWidth, imageHeight))
	bmLength := mazeHeight * mazeWidth * bmCell

	for i := 0; i < bmLength; i += bmCell {
		bmI := uint32(i)
		if bm.Contains(bmI) {
			x := (i / bmCell % mazeWidth) * cellWidth
			y := (i / bmCell / mazeWidth) * cellHeight

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

	return img
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
