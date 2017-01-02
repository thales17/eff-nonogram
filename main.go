package main

import (
	"flag"
	"log"
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
		revealPtr := flag.Bool("reveal", false, "reveals the puzzle")
		flag.Parse()
		if *createPtr && len(*pathPtr) == 0 {
			log.Fatal("Create flag passed with out a path")
		}

		var pd *puzzleData
		var err error
		if !*createPtr && len(*pathPtr) == 0 {
			pd = randomPuzzleData(int(*gridSizeXPtr), int(*gridSizeYPtr))
		} else if !*createPtr && len(*pathPtr) > 0 {
			pd, err = load(*pathPtr)
			if err != nil {
				log.Fatal(err)
			}
		} else if *createPtr && len(*pathPtr) > 0 {
			pd = &puzzleData{}
			pd.gridSize.X = int(*gridSizeXPtr)
			pd.gridSize.Y = int(*gridSizeYPtr)
		}

		g := newGame(pd, canvas)
		if *createPtr {
			g.createMode = true
			g.path = *pathPtr
		}
		if *revealPtr {
			g.reveal()
		}
	})
}
