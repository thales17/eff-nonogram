package main

import (
	"math/rand"

	"github.com/forestgiant/eff"
)

type puzzleData struct {
	gridSize eff.Point
	squares  []eff.Point
}

func (pd *puzzleData) checkSquare(p eff.Point) bool {
	for _, s := range pd.squares {
		if s.X == p.X && s.Y == p.Y {
			return true
		}
	}

	return false
}

func randomPuzzleData(w int, h int) *puzzleData {
	pd := &puzzleData{}
	pd.gridSize = eff.Point{
		X: w,
		Y: h,
	}

	for i := 0; i < w*2+2*h; i++ {
		randPoint := func() eff.Point {
			return eff.Point{
				X: rand.Intn(w),
				Y: rand.Intn(h),
			}
		}

		p := randPoint()
		for pd.checkSquare(p) {
			p = randPoint()
		}

		pd.squares = append(pd.squares, p)
	}

	return pd
}
