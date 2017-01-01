package main

import (
	"flag"
	"math/rand"
	"time"

	"github.com/forestgiant/eff"
	"github.com/forestgiant/eff/sdl"
)

const (
	windowW = 1024
	windowH = 768
)

func main() {
	canvas := sdl.NewCanvas("Eff Nonogram", windowW, windowH, eff.Color{R: 0xFF, B: 0xFF, G: 0xFF, A: 0xFF}, 60, true)
	canvas.Run(func() {
		rand.Seed(time.Now().UnixNano())
		createPtr := flag.Bool("create", false, "creates a new puzzle file, requires the path")
		pathPtr := flag.String("path", "", "specifies the path to the puzzle, if none provide a random puzzle will be generated")
		gridSizeXPtr := flag.Uint("gridSizeX", 10, "specify the width of the puzzle grid, this is used when creating a new puzzle or playing a random one")
		gridSizeYPtr := flag.Uint("gridSizeY", 10, "specify the height of the puzzle grid, this is used when creating a new puzzle or playing a random one")
		// revealPtr := flag.Bool("reveal", true, "reveals the puzzle")
		flag.Parse()
		if *createPtr != false && len(*pathPtr) == 0 {
			// log.Fatal(flag.Usage())
		}

		if *createPtr == false && len(*pathPtr) == 0 {
			/*g := */ newGame(randomPuzzleData(int(*gridSizeXPtr), int(*gridSizeYPtr)), canvas)
			// g.reveal()
		}

	})
}
