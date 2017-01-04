package main

import "github.com/forestgiant/eff"

const (
	defaultState   = 0
	xState         = 1
	fillState      = 2
	incorrectState = 3
)

type square struct {
	eff.Shape
	state     int
	point     eff.Point
	mouseOver bool
	leftDown  bool
	rightDown bool
	selected  bool
	onSelect  func(int, int)
	check     func(int, int) bool
}

func (s *square) setState(state int) {
	s.state = state
	s.Clear()
	if state == defaultState || state == incorrectState {
		color := eff.Color{R: 0xEE, G: 0xEE, B: 0xEE, A: 0xFF}
		if s.point.X%2 == 1 {
			color.Add(-5)
		}
		if s.point.Y%2 == 1 {
			color.Add(-5)
		}

		s.SetBackgroundColor(color)

		if state == incorrectState {
			strokeWidth := 2
			color := eff.Color{R: 0xFF, G: 0x00, B: 0x00, A: 0xFF}
			s.FillRect(eff.Rect{X: gridWidth, Y: gridWidth, W: s.Rect().W - gridWidth, H: strokeWidth}, color)
			s.FillRect(eff.Rect{X: s.Rect().W - strokeWidth, Y: strokeWidth, W: strokeWidth, H: s.Rect().H - (strokeWidth * 2)}, color)
			s.FillRect(eff.Rect{X: gridWidth, Y: s.Rect().H - strokeWidth, W: s.Rect().W - gridWidth, H: strokeWidth}, color)
			s.FillRect(eff.Rect{X: gridWidth, Y: strokeWidth, W: strokeWidth, H: s.Rect().H - (strokeWidth * 2)}, color)
		}

	} else if state == fillState {
		s.SetBackgroundColor(eff.Color{R: 0x33, B: 0x33, G: 0x33, A: 0xFF})
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

func (s *square) MouseDown(left bool, middle bool, right bool) {
	s.leftDown = left
	s.rightDown = right
}

func (s *square) MouseUp(left bool, middle bool, right bool) {
	if left && s.leftDown {
		s.leftDown = false
		if s.state == defaultState {
			if s.check != nil {
				if s.check(s.point.X, s.point.Y) {
					s.setState(fillState)
				} else {
					s.setState(incorrectState)
				}
			} else {
				s.setState(fillState)
			}
		}
	}

	if right && s.rightDown {
		s.rightDown = false
		if s.state == defaultState {
			s.setState(xState)
		} else if s.state == xState {
			s.setState(defaultState)
		}
	}
}

func (s *square) MouseOver() {
	s.mouseOver = true
	if s.onSelect != nil {
		s.onSelect(s.point.Y, s.point.X)
	}
}

func (s *square) MouseOut() {
	s.mouseOver = false
	s.leftDown = false
	s.rightDown = false
}
func (s *square) IsMouseOver() bool { return s.mouseOver }

func (s *square) SetSelected(selected bool) {
	if s.selected == selected {
		return
	}
	s.selected = selected
	color := s.BackgroundColor()
	if selected {
		color.Add(0x1E)
	} else {
		color.Add(-0x1E)
	}
	s.SetBackgroundColor(color)
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
