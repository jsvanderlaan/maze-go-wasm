package utils

import (
	"bytes"
	"image"
	"image/png"
	"log"
	"syscall/js"
	"time"
)

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
