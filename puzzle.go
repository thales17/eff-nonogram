package main

import (
	"math/rand"
	"sort"

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

func (pd *puzzleData) legendValuesForRow(row int) []int {
	vals := []int{}
	rowVals := []int{}

	for _, s := range pd.squares {
		if s.Y != row {
			continue
		}
		rowVals = append(rowVals, s.X)
	}
	sort.Ints(rowVals)

	val := 0
	last := 0
	for _, v := range rowVals {
		if v-last > 1 && val > 0 {
			vals = append(vals, val)
			val = 0
		}
		val++
		last = v
	}

	vals = append(vals, val)

	return vals

}

func (pd *puzzleData) legendValuesForCol(col int) []int {
	vals := []int{}
	colVals := []int{}
	for _, s := range pd.squares {
		if s.X != col {
			continue
		}
		colVals = append(colVals, s.Y)
	}
	sort.Ints(colVals)

	val := 0
	last := 0
	for _, v := range colVals {
		if v-last > 1 && val > 0 {
			vals = append(vals, val)
			val = 0
		}
		val++
		last = v
	}

	vals = append(vals, val)
	return vals
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
