package main

import "github.com/forestgiant/eff"

const (
	defaultState = 0
	xState       = 1
	fillState    = 2
)

type square struct {
	eff.Shape
	state     int
	point     eff.Point
	mouseOver bool
}

func (s *square) setState(state int) {
	s.state = state
	s.Clear()
	if state == defaultState {
		color := eff.Color{R: 0xEE, G: 0xEE, B: 0xEE, A: 0xFF}
		if s.point.X%2 == 1 {
			color.Add(-5)
		}
		if s.point.Y%2 == 1 {
			color.Add(-5)
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

func (s *square) Hitbox() eff.Rect {
	return s.ParentOffsetRect()
}

func (s *square) MouseDown(left bool, middle bool, right bool) {}
func (s *square) MouseUp(left bool, middle bool, right bool)   {}
func (s *square) MouseOver() {
	s.mouseOver = true
	color := s.BackgroundColor()
	color.Add(0x1E)
	s.SetBackgroundColor(color)
}
func (s *square) MouseOut() {
	s.mouseOver = false
	color := s.BackgroundColor()
	color.Add(-0x1E)
	s.SetBackgroundColor(color)
}
func (s *square) IsMouseOver() bool { return s.mouseOver }

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
