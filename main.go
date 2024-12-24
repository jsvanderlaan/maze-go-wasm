package main

import (
	"image"
	"log"
	"time"

	"maze.jurre.dev/utils"

	"syscall/js"

	"image/color"
	_ "image/jpeg"
)

func main() {
	c := make(chan struct{})
	js.Global().Set("processImage", js.FuncOf(processImage))
	js.Global().Set("processText", js.FuncOf(processText))
	<-c
}

func processText(this js.Value, args []js.Value) interface{} {
	gStart := time.Now()
	text := js.ValueOf(args[0]).String()
	outline := js.ValueOf(args[1]).Bool()
	threshold := js.ValueOf(args[2]).Float()
	height := js.ValueOf(args[3]).Int()

	start := time.Now()
	sourceImage, err := utils.RenderTextToJPG(text, outline)
	if err != nil {
		log.Printf("Err %v\n", err)
	}
	log.Printf("RenderTextToJPG: %v", time.Since(start))

	output := process(sourceImage, threshold, height)
	log.Printf("Total time: %v", time.Since(gStart))

	return output
}

func processImage(this js.Value, args []js.Value) interface{} {
	gStart := time.Now()
	imageArray := args[0]
	threshold := js.ValueOf(args[1]).Float()
	height := js.ValueOf(args[2]).Int()

	start := time.Now()
	sourceImage := utils.JsToImg(imageArray)
	log.Printf("JsToImg: %v", time.Since(start))

	output := process(sourceImage, threshold, height)
	log.Printf("Total time: %v", time.Since(gStart))

	return output
}

func process(img image.Image, t float64, height int) interface{} {
	threshold := uint8(float64(255) * t)
	mazeHeight := height
	mazeWidth := utils.DetermineWidth(mazeHeight, img)
	log.Printf("threshold: %d;  height: %d", threshold, mazeHeight)

	start := time.Now()
	maze := utils.CreateMaze(
		img,
		utils.ImageOptions{
			Threshold: threshold,
		},
		utils.MazeOptions{
			Width:  mazeWidth,
			Height: mazeHeight,
			Shape:  utils.Square,
		},
		utils.CellOptions{
			Width:       10,
			Height:      10,
			BorderColor: color.Black,
		})

	log.Printf("CreateMaze: %v", time.Since(start))

	// start = time.Now()
	// bwImage := utils.ImgToBlackWhite(img, threshold)
	// log.Printf("ImgToBlackWhite: %v", time.Since(start))

	// start = time.Now()
	// jsValue := utils.ImgToJs(bwImage)
	// log.Printf("ImgToJS: %v", time.Since(start))

	// jsValue := js.ValueOf(maze)
	jsValue := utils.ImgToJs(maze)

	return jsValue
}
