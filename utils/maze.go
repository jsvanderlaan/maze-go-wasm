package utils

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"math/rand"
	"strings"

	"github.com/kelindar/bitmap"
)

type CellShape uint8

const (
	_ CellShape = iota
	_
	_
	Triangle
	Square
	_
	Hexagon
)

type ImageOptions struct {
	Threshold uint8
}

type MazeOptions struct {
	Width  int
	Height int
	Shape  CellShape
}

type CellOptions struct {
	Width       int
	Height      int
	BorderColor color.Color
}

func CreateMaze(image image.Image, imageOptions ImageOptions, mazeOptions MazeOptions, cellOptions CellOptions) *image.NRGBA {
	layout := calculateLayout(image, imageOptions, mazeOptions)
	// printBitmap(layout, mazeOptions.Width, mazeOptions.Height)

	shrinkedLayout, newWidth, newHeight := shrink(layout, mazeOptions)
	mazeOptions.Width = newWidth
	mazeOptions.Height = newHeight
	// printBitmap(shrinkedLayout, mazeOptions.Width, mazeOptions.Height)

	start, end := determineStartEnd(shrinkedLayout, mazeOptions)
	log.Printf("start %d end %d", start, end)

	maze, _ := dfsMaze(shrinkedLayout, mazeOptions, start, end)
	// printBitmap(maze, mazeOptions.Width*(int(mazeOptions.Shape)+1), mazeOptions.Height)
	// printPath(path)

	img := MazeToImg(maze, mazeOptions, cellOptions)
	// AddPath(&img, mazeOptions, cellOptions, path)
	AddStartEnd(&img, mazeOptions, cellOptions, start, end)

	// log.Println(printSquareMaze(maze, mazeOptions))

	return &img
}

func determineStartEnd(bm bitmap.Bitmap, mazeOptions MazeOptions) (uint32, uint32) {
	edges := []uint32{}
	max, _ := bm.Max()
	w := uint32(mazeOptions.Width)

	if max == 1 {
		return 1, 1
	}

	for i := uint32(0); i < max; i++ {
		y := i / w
		x := i % w
		if ((y == 0) ||
			(y == uint32(mazeOptions.Height)) ||
			(x == w-1) ||
			(x == 0)) && bm.Contains(i) {
			edges = append(edges, i)
		}
	}

	rndIndex := rand.Intn(len(edges))
	start := edges[rndIndex]
	end := start
	for end == start {
		rndIndex := rand.Intn(len(edges))
		end = edges[rndIndex]
	}

	return start, end
}

// returns new bitmap, new width, new height
func shrink(bm bitmap.Bitmap, mazeOptions MazeOptions) (bitmap.Bitmap, int, int) {
	originalWidth := uint32(mazeOptions.Width)
	min, _ := bm.Min()
	max, _ := bm.Max()
	minHeight := min / originalWidth
	maxHeight := max / originalWidth
	newHeight := maxHeight - minHeight
	log.Printf("minHeight %d", minHeight)
	log.Printf("maxHeight %d", maxHeight)
	log.Printf("newHeight %d", newHeight)

	minWidth := originalWidth
	maxWidth := uint32(0)

	bm.Range(func(x uint32) {
		w := x % originalWidth
		if w > uint32(maxWidth) {
			maxWidth = w
		}
		if w < uint32(minWidth) {
			minWidth = w
		}
	})

	newWidth := maxWidth - minWidth
	log.Printf("minWidth %d", minWidth)
	log.Printf("maxWidth %d", maxWidth)
	log.Printf("newWidth %d", newWidth)

	var shrunk bitmap.Bitmap
	shrunkI := uint32(0)
	for i := uint32(0); i <= max; i++ {
		y := i / originalWidth
		x := i % originalWidth
		// log.Printf("x %d y %d i %d sI %d", x, y, i, shrunkI)
		if (y < minHeight) ||
			(y > maxHeight) ||
			(x > maxWidth) ||
			(x < minWidth) {
			continue
		}

		if bm.Contains(i) {
			shrunk.Set(shrunkI)
		}

		shrunkI++
	}

	return shrunk, int(newWidth) + 1, int(newHeight) + 1
}

func dfsMaze(layout bitmap.Bitmap, mazeOptions MazeOptions, start uint32, end uint32) (bitmap.Bitmap, []uint32) {
	width := uint32(mazeOptions.Width)
	cellShape := mazeOptions.Shape

	var maze bitmap.Bitmap
	var visited bitmap.Bitmap
	stack := []uint32{}
	path := []uint32{}

	stack = append(stack, start)

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		visited.Set(current)
		index := current * (uint32(cellShape) + 1)
		if len(path) == 0 && current == end {
			path = append(path, stack...)
		}
		maze.Set(index)

		// Set previous connection
		if len(stack) > 1 {
			prev := stack[len(stack)-2]
			if prev == current+width { // south
				maze.Set(index + 1)
			} else if prev == current-1 { // west
				maze.Set(index + 2)
			} else if prev == current-width { // north
				maze.Set(index + 3)
			} else if prev == current+1 { // east
				maze.Set(index + 4)
			}
		}

		allowedNeighbours := []uint32{}

		// Check neighbours
		south := current + width
		if layout.Contains(south) && !visited.Contains(south) && !maze.Contains(index+1) {
			allowedNeighbours = append(allowedNeighbours, south)
		}
		west := int(current) - 1
		uwest := uint32(west)
		if west >= 0 && uwest%(width) != width-1 && layout.Contains(uwest) && !visited.Contains(uwest) && !maze.Contains(index+2) {
			allowedNeighbours = append(allowedNeighbours, uwest)
		}
		north := int(current) - int(width)
		unorth := uint32(north)
		if north >= 0 && layout.Contains(unorth) && !visited.Contains(unorth) && !maze.Contains(index+3) {
			allowedNeighbours = append(allowedNeighbours, unorth)
		}
		east := current + 1
		if layout.Contains(east) && east%(width) != 0 && !visited.Contains(east) && !maze.Contains(index+4) {
			allowedNeighbours = append(allowedNeighbours, east)
		}

		if len(allowedNeighbours) > 0 {
			rndIndex := rand.Intn(len(allowedNeighbours))
			next := allowedNeighbours[rndIndex]
			if next == south {
				maze.Set(index + 1)
			} else if next == uwest {
				maze.Set(index + 2)
			} else if next == unorth {
				maze.Set(index + 3)
			} else if next == east {
				maze.Set(index + 4)
			}
			stack = append(stack, next)
			continue
		}

		stack = stack[:len(stack)-1]
	}

	return maze, path
}

func calculateLayout(img image.Image, imageOptions ImageOptions, mazeOptions MazeOptions) bitmap.Bitmap {
	imageWidth := img.Bounds().Dx()
	imageHeight := img.Bounds().Dy()
	width := mazeOptions.Width
	height := mazeOptions.Height

	pixelsPerCellWidthF := float32(imageWidth) / float32(width)
	pixelsPerCellHeightF := float32(imageHeight) / float32(height)
	pixelsPerCellWidth := int(pixelsPerCellWidthF)
	pixelsPerCellHeight := int(pixelsPerCellHeightF)

	max := int(int(pixelsPerCellHeight*pixelsPerCellWidth) >> 1)

	if pixelsPerCellWidth <= 0 {
		pixelsPerCellWidth = 1
	}
	if pixelsPerCellHeight <= 0 {
		pixelsPerCellHeight = 1
	}
	if max <= 0 {
		max = 1
	}

	log.Printf("pixelsPerCellWidthF: %f", pixelsPerCellWidthF)
	log.Printf("pixelsPerCellHeightF: %f", pixelsPerCellHeightF)
	log.Printf("pixelsPerCellWidth: %d", pixelsPerCellWidth)
	log.Printf("pixelsPerCellHeight: %d", pixelsPerCellHeight)
	log.Printf("width: %d", width)
	log.Printf("height: %d", height)
	log.Printf("max: %d", max)

	var layout bitmap.Bitmap
	for y := int(height) - 1; y >= 0; y-- {
		for x := int(width) - 1; x >= 0; x-- {
			sum := 0
			for dy := 0; dy < pixelsPerCellWidth; dy++ {
				for dx := 0; dx < pixelsPerCellHeight; dx++ {
					if IsBlack(img.At(int(float32(x)*pixelsPerCellWidthF)+dx, int(float32(y)*pixelsPerCellHeightF)+dy), imageOptions.Threshold) {
						sum++
					}
				}
			}
			if sum >= max {
				layout.Set(uint32(y*width + x))
			}
		}
	}

	return layout
}

func printPath(path []uint32) {
	var sb strings.Builder
	for _, p := range path {
		sb.WriteString(fmt.Sprintf("%d, ", p))
	}
	log.Println(sb.String())
}

func printBitmap(bm bitmap.Bitmap, width int, height int) {
	var sb strings.Builder
	sb.WriteString("\n")

	for i := uint32(0); i < uint32(height*width); i++ {
		if bm.Contains(i) {
			sb.WriteString("1")
		} else {
			sb.WriteString("0")
		}

		if i%uint32(width) == uint32(width)-1 {
			sb.WriteString("\n")
		}
	}

	log.Println(sb.String())
}

func printSquareMaze(bm bitmap.Bitmap, mazeOptions MazeOptions) string {
	width := mazeOptions.Width
	height := mazeOptions.Height
	var sb strings.Builder
	sb.WriteString("\n")

	for i := uint32(0); i < uint32(height*width); i++ {
		index := i * (uint32(Square) + 1)

		if !bm.Contains(index) {
			sb.WriteString(" ")
		} else {
			south := bm.Contains(index + 1)
			west := bm.Contains(index + 2)
			north := bm.Contains(index + 3)
			east := bm.Contains(index + 4)

			if south && west && north && east {
				sb.WriteString("╬")
			} else if south && west && north {
				sb.WriteString("╣")
			} else if south && west && east {
				sb.WriteString("╦")
			} else if south && north && east {
				sb.WriteString("╠")
			} else if west && north && east {
				sb.WriteString("╩")
			} else if south && west {
				sb.WriteString("╗")
			} else if south && north {
				sb.WriteString("║")
			} else if south && east {
				sb.WriteString("╔")
			} else if west && east {
				sb.WriteString("═")
			} else if west && north {
				sb.WriteString("╝")
			} else if north && east {
				sb.WriteString("╚")
			} else {
				sb.WriteString("●")
			}
		}
		if i%uint32(width) == uint32(width-1) {
			sb.WriteString("\n")
		}
	}

	return sb.String()
}
