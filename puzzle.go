package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"sort"
	"strings"

	"strconv"

	"github.com/forestgiant/eff"
	"github.com/pkg/errors"
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

func (pd *puzzleData) save(path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(fmt.Sprintf("%d,%d\n", pd.gridSize.X, pd.gridSize.Y))
	if err != nil {
		return err
	}
	f.Sync()
	for _, s := range pd.squares {
		_, err = f.WriteString(fmt.Sprintf("%d,%d\n", s.X, s.Y))
		if err != nil {
			return err
		}
		f.Sync()
	}
	return nil
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

func load(path string) (*puzzleData, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	rows := strings.Split(string(data), "\n")
	if len(rows) < 2 {
		return nil, errors.New("invalid file")
	}

	parseRow := func(row string) (int, int, error) {
		rowVals := strings.Split(rows[0], ",")
		if len(rowVals) != 2 {
			return 0, 0, errors.New("invalid file")
		}
		x, err := strconv.Atoi(rowVals[0])
		if err != nil {
			return 0, 0, err
		}
		y, err := strconv.Atoi(rowVals[1])
		if err != nil {
			return 0, 0, err
		}

		return x, y, nil
	}

	pd := &puzzleData{}
	x, y, err := parseRow(rows[0])
	if err != nil {
		return nil, err
	}
	pd.gridSize.X = x
	pd.gridSize.Y = y

	for _, row := range rows[1:] {
		x, y, err := parseRow(row)
		if err != nil {
			return nil, err
		}

		pd.squares = append(pd.squares, eff.Point{X: x, Y: y})
	}
	return pd, nil
}
