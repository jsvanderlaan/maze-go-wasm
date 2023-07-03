package utils

import (
	"image"
	"log"
	"math/rand"
	"strings"
	"time"

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

type Maze struct {
	width  uint8
	height uint8
	shape  CellShape
	// a cell consists of one bit to indicate if the cell is part of the maze
	// followed by a number of cells equal to the number of neighbours (depends on the shape)
	// indicating of the neighbour is connected.
	// Starting south and going clockwise.
	// for example for shape Square a cell with a connection to the upper and rightmost neighbours looks like
	// 10011
	cells bitmap.Bitmap
}

func CreateMaze(image image.Image, height int, cellShape CellShape, threshold uint8) string {
	rand.Seed(time.Now().Unix())

	imageWidth := image.Bounds().Dx()
	imageHeight := image.Bounds().Dy()
	width := int(float32(height) * (float32(imageWidth) / float32(imageHeight)))

	layout := calculateLayout(image, height, width, threshold)
	printBitmap(layout, width, height)

	maze := dfsMaze(layout, uint32(height), uint32(width), cellShape)
	return printSquareMaze(maze, width, height)

	// return Maze{
	// 	uint8(width),
	// 	uint8(height),
	// 	cellShape,
	// 	layout,
	// }
}

func dfsMaze(layout bitmap.Bitmap, height uint32, width uint32, cellShape CellShape) bitmap.Bitmap {
	var maze bitmap.Bitmap
	var visited bitmap.Bitmap
	stack := []uint32{}

	start, _ := layout.Min()
	log.Printf("start %d", start)
	stack = append(stack, start)

	for len(stack) > 0 {
		// printSquareMaze(maze, int(width), int(height))
		current := stack[len(stack)-1]
		// log.Printf("current %d", current)
		visited.Set(current)
		index := current * (uint32(cellShape) + 1)
		// log.Printf("index %d", index)
		maze.Set(index)

		// Set previous connection
		if len(stack) > 1 {
			prev := stack[len(stack)-2]
			// log.Printf("prev %d", prev)
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
		if layout.Contains(east) && !visited.Contains(east) && !maze.Contains(index+4) {
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

	return maze
}

func calculateLayout(img image.Image, height int, width int, threshold uint8) bitmap.Bitmap {
	imageWidth := img.Bounds().Dx()
	imageHeight := img.Bounds().Dy()

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
					if IsBlack(img.At(int(float32(x)*pixelsPerCellWidthF)+dx, int(float32(y)*pixelsPerCellHeightF)+dy), threshold) {
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

func printSquareMaze(bm bitmap.Bitmap, width int, height int) string {
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
