package main

import "github.com/forestgiant/eff"
import "math"

type game struct {
	eff.Shape
	pd      *puzzleData
	squares []*square
}

func (g *game) init(c eff.Canvas) {
	g.SetRect(c.Rect())
	g.SetBackgroundColor(eff.White())
	cellW := c.Rect().W / g.pd.gridSize.X
	cellH := c.Rect().H / g.pd.gridSize.Y
	squareSize := int(math.Min(float64(cellW), float64(cellH)))
	board := &eff.Shape{}
	board.SetRect(eff.Rect{
		X: (c.Rect().W - (squareSize * g.pd.gridSize.X)) / 2,
		Y: (c.Rect().H - (squareSize * g.pd.gridSize.Y)) / 2,
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
		board.AddChild(s)
		g.squares = append(g.squares, s)
		c.AddClickable(s)
	}

	grid := &eff.Shape{}
	grid.SetRect(eff.Rect{
		X: 0,
		Y: 0,
		W: board.Rect().W,
		H: board.Rect().H,
	})
	board.AddChild(grid)
	for i := 0; i < g.pd.gridSize.X; i++ {
		grid.DrawLine(
			eff.Point{
				X: i * squareSize,
				Y: 0,
			},
			eff.Point{
				X: i * squareSize,
				Y: grid.Rect().H,
			},
			eff.Black(),
		)
	}
	grid.DrawLine(
		eff.Point{
			X: grid.Rect().W - 1,
			Y: 0,
		},
		eff.Point{
			X: grid.Rect().W - 1,
			Y: grid.Rect().H,
		},
		eff.Black(),
	)

	for i := 0; i < g.pd.gridSize.Y; i++ {
		grid.DrawLine(
			eff.Point{
				X: 0,
				Y: i * squareSize,
			},
			eff.Point{
				X: grid.Rect().W,
				Y: i * squareSize,
			},
			eff.Black(),
		)
	}

	grid.DrawLine(
		eff.Point{
			X: 0,
			Y: grid.Rect().H - 1,
		},
		eff.Point{
			X: grid.Rect().W,
			Y: grid.Rect().H - 1,
		},
		eff.Black(),
	)

	for _, p := range g.pd.squares {
		index := p.Y*g.pd.gridSize.X + p.X
		g.squares[index].setState(fillState)
	}
}

func newGame(pd *puzzleData, c eff.Canvas) *game {
	g := &game{}
	g.pd = pd
	g.init(c)
	return g
}
