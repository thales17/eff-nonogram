package main

import "github.com/forestgiant/eff"

const (
	defaultState = 0
	xState       = 1
	fillState    = 2
)

type square struct {
	eff.Shape
	state int
	point eff.Point
}

func (s *square) setState(state int) {
	s.state = state
	s.Clear()
	if state == defaultState {
		evenEvenColor := eff.Color{R: 0xEE, G: 0xEE, B: 0xEE, A: 0xFF}
		evenOddColor := eff.Color{R: 0xE9, G: 0xE9, B: 0xE9, A: 0xFF}
		oddEvenColor := eff.Color{R: 0xE5, G: 0xE5, B: 0xE5, A: 0xFF}
		oddOddColor := eff.Color{R: 0xE1, G: 0xE1, B: 0xE1, A: 0xFF}
		color := evenEvenColor
		if s.point.X%2 == 1 && s.point.Y%2 == 0 {
			color = evenOddColor
		} else if s.point.X%2 == 0 && s.point.Y%2 == 1 {
			color = oddEvenColor
		} else if s.point.X%2 == 1 && s.point.Y%2 == 1 {
			color = oddOddColor
		}

		s.SetBackgroundColor(color)
	} else if state == fillState {
		s.SetBackgroundColor(eff.Black())
	} else if state == xState {
		s.DrawLine(
			eff.Point{X: 0, Y: 0},
			eff.Point{X: s.Rect().W, Y: s.Rect().H},
			eff.Black(),
		)

		s.DrawLine(
			eff.Point{X: s.Rect().W, Y: 0},
			eff.Point{X: 0, Y: s.Rect().H},
			eff.Black(),
		)
	}
}

func newSquare(point eff.Point, size int) *square {
	s := &square{}
	s.SetRect(eff.Rect{
		X: point.X * size,
		Y: point.Y * size,
		W: size,
		H: size,
	})
	s.point = point
	s.setState(defaultState)

	return s

}
