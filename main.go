package main

import (
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
		newGame(randomPuzzleData(20, 20), canvas)
	})
}
