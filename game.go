package main

import (
	"log"
	"math"

	"fmt"

	"github.com/forestgiant/eff"
)

const (
	gridWidth = 1
	fiveWidth = 3
)

type game struct {
	eff.Shape
	pd         *puzzleData
	squares    []*square
	createMode bool
	path       string
}

func (g *game) init(c eff.Canvas) {
	g.SetRect(c.Rect())
	g.SetBackgroundColor(eff.White())

	// Squares
	cellW := c.Rect().W / (g.pd.gridSize.X + (g.pd.gridSize.X / 2))
	cellH := c.Rect().H / (g.pd.gridSize.Y + (g.pd.gridSize.Y / 2))
	squareSize := int(math.Min(float64(cellW), float64(cellH)))
	legendW := squareSize * (g.pd.gridSize.X / 2)
	legendH := squareSize * (g.pd.gridSize.Y / 2)
	boardLegendW := squareSize*(g.pd.gridSize.X) + legendW
	boardLegendH := squareSize*(g.pd.gridSize.Y) + legendH
	board := &eff.Shape{}
	board.SetRect(eff.Rect{
		X: (c.Rect().W-boardLegendW)/2 + legendW,
		Y: (c.Rect().H-boardLegendH)/2 + legendH,
		W: squareSize * g.pd.gridSize.X,
		H: squareSize * g.pd.gridSize.Y,
	})
	board.SetBackgroundColor(eff.Black())
	g.AddChild(board)
	c.AddChild(g)

	for i := 0; i < g.pd.gridSize.X*g.pd.gridSize.Y; i++ {
		s := newSquare(eff.Point{
			X: i % g.pd.gridSize.X,
			Y: i / g.pd.gridSize.X,
		}, squareSize)
		s.onSelect = func(r int, c int) {
			g.selectRowCol(r, c)
		}
		s.check = func(r int, c int) bool {
			if g.createMode {
				g.pd.squares = append(g.pd.squares, eff.Point{X: r, Y: c})
				g.pd.save(g.path)
				return true
			}

			for _, p := range g.pd.squares {
				if p.X == r && p.Y == c {
					return true
				}
			}

			return false
		}

		board.AddChild(s)
		g.squares = append(g.squares, s)
		c.AddClickable(s)
	}

	//Grid
	grid := &eff.Shape{}
	grid.SetRect(eff.Rect{
		X: 0,
		Y: 0,
		W: board.Rect().W,
		H: board.Rect().H,
	})
	board.AddChild(grid)
	gridBlue := eff.Color{R: 0x06, G: 0x5A, B: 0x82, A: 0xFF}
	for i := 0; i < g.pd.gridSize.X; i++ {
		color := eff.Black()
		if i%5 == 0 && i > 0 {
			color = gridBlue
		}

		grid.FillRect(eff.Rect{
			X: i * squareSize,
			Y: 0,
			W: gridWidth,
			H: grid.Rect().H,
		}, color)

		if i%5 == 0 && i > 0 {
			grid.FillRect(eff.Rect{
				X: i*squareSize - (fiveWidth / 2),
				Y: 0,
				W: fiveWidth,
				H: grid.Rect().H,
			}, color)
		}
	}
	grid.FillRect(eff.Rect{X: grid.Rect().W - gridWidth, Y: 0, W: gridWidth, H: grid.Rect().H}, eff.Black())

	for i := 0; i < g.pd.gridSize.Y; i++ {
		color := eff.Black()
		if i%5 == 0 && i > 0 {
			color = gridBlue
		}
		grid.FillRect(eff.Rect{X: 0, Y: i * squareSize, W: grid.Rect().W, H: gridWidth}, color)

		if i%5 == 0 && i > 0 {
			grid.FillRect(eff.Rect{
				X: 0,
				Y: i*squareSize - (fiveWidth / 2),
				W: grid.Rect().W,
				H: fiveWidth,
			}, color)
		}
	}

	grid.FillRect(eff.Rect{X: 0, Y: grid.Rect().H - gridWidth, W: grid.Rect().W, H: gridWidth}, eff.Black())

	//Legend
	font, err := c.OpenFont("assets/fonts/roboto/Roboto-Medium.ttf", squareSize/2)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < g.pd.gridSize.X; i++ {
		vals := g.pd.legendValuesForCol(i)
		squareRect := g.squares[i].Rect()
		legendH := len(vals) * squareSize
		for j, val := range vals {
			s := &eff.Shape{}
			s.SetRect(eff.Rect{
				X: board.Rect().X + squareRect.X,
				Y: board.Rect().Y + squareRect.Y - legendH + (j * squareSize),
				W: squareSize,
				H: squareSize,
			})
			s.SetBackgroundColor(eff.White())
			valStr := fmt.Sprintf("%d", val)
			textW, textH, err := c.Graphics().GetTextSize(font, valStr)
			if err != nil {
				log.Fatal(err)
			}
			textPoint := eff.Point{
				X: (squareSize - textW) / 2,
				Y: (squareSize - textH) / 2,
			}
			s.DrawText(font, valStr, eff.Black(), textPoint)
			g.AddChild(s)
		}
	}

	for i := 0; i < g.pd.gridSize.Y; i++ {
		vals := g.pd.legendValuesForRow(i)
		squareRect := g.squares[i*g.pd.gridSize.X].Rect()
		legendW := len(vals) * squareSize
		for j, val := range vals {
			s := &eff.Shape{}
			s.SetRect(eff.Rect{
				X: board.Rect().X + squareRect.X - legendW + (j * squareSize),
				Y: board.Rect().Y + squareRect.Y,
				W: squareSize,
				H: squareSize,
			})
			s.SetBackgroundColor(eff.White())
			valStr := fmt.Sprintf("%d", val)
			textW, textH, err := c.Graphics().GetTextSize(font, valStr)
			if err != nil {
				log.Fatal(err)
			}
			textPoint := eff.Point{
				X: (squareSize - textW) / 2,
				Y: (squareSize - textH) / 2,
			}
			s.DrawText(font, valStr, eff.Black(), textPoint)
			g.AddChild(s)
		}
	}
}

func (g *game) reveal() {
	for _, p := range g.pd.squares {
		index := p.X*g.pd.gridSize.X + p.Y
		g.squares[index].setState(fillState)
	}
}

func (g *game) selectRowCol(r int, c int) {
	if r > g.pd.gridSize.Y || c > g.pd.gridSize.X || r < 0 || c < 0 {
		fmt.Println("attempt to select invalid row col:", r, c)
		return
	}
	for _, s := range g.squares {
		s.SetSelected(false)
	}

	//Rows
	rStartIndex := r * g.pd.gridSize.X
	rEndIndex := rStartIndex + g.pd.gridSize.X
	rowSquares := g.squares[rStartIndex:rEndIndex]
	for _, row := range rowSquares {
		row.SetSelected(true)
	}

	//Cols
	for i := 0; i < g.pd.gridSize.Y; i++ {
		index := i*g.pd.gridSize.X + (c % g.pd.gridSize.X)
		g.squares[index].SetSelected(true)
	}

}

func newGame(pd *puzzleData, c eff.Canvas) *game {
	g := &game{}
	g.pd = pd
	g.init(c)
	return g
}
