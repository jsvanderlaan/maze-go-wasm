package main

import (
	"log"
	"time"

	"maze.jurre.dev/utils"

	// https://github.com/golang/go/tree/master/src/syscall/js
	"syscall/js"

	"image/color"
	_ "image/jpeg"
)

func main() {
	c := make(chan struct{})
	js.Global().Set("processImage", js.FuncOf(processImage))
	<-c
}

func processImage(this js.Value, args []js.Value) interface{} {
	imageArray := args[0]

	start := time.Now()
	sourceImage := utils.JsToImg(imageArray)
	log.Printf("JsToImg: %v", time.Since(start))

	threshold := uint8(float64(255) * js.ValueOf(args[1]).Float())
	mazeHeight := int(js.ValueOf(args[2]).Int())
	mazeWidth := utils.DetermineWidth(mazeHeight, sourceImage)
	log.Printf("threshold: %d;  height: %d", threshold, mazeHeight)

	start = time.Now()
	maze := utils.CreateMaze(
		sourceImage,
		utils.ImageOptions{
			Threshold: threshold,
		},
		utils.MazeOptions{
			Width:  mazeWidth,
			Height: mazeHeight,
			Shape:  utils.Square,
		},
		utils.CellOptions{
			Width:       8,
			Height:      8,
			BorderColor: color.Black,
		})

	log.Printf("CreateMaze: %v", time.Since(start))

	// start = time.Now()
	// bwImage := utils.ImgToBlackWhite(sourceImage, threshold)
	// log.Printf("ImgToBlackWhite: %v", time.Since(start))

	// start = time.Now()
	// jsValue := utils.ImgToJs(bwImage)
	// log.Printf("ImgToJS: %v", time.Since(start))

	// jsValue := js.ValueOf(maze)
	jsValue := utils.ImgToJs(maze)

	return jsValue
}
