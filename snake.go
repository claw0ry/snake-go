// Copyright (c) 2025, Mads Moi-Aune <mads@moiaune.dev>
//
// SPDX-License-Identifier: BSD-3-Clause

package main

import (
	"fmt"
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Direction int
var DirectionUp Direction = 1
var DirectionRight Direction = 2
var DirectionDown Direction = 3
var DirectionLeft Direction = 4

type snake struct {
	board     *board
	body      []rl.Vector2
	direction rl.Vector2
	points    int
	color     color.RGBA
}

func NewSnake(b *board) *snake {
	return &snake{
		board: b,
		body: []rl.Vector2{
			{X: 10, Y: 10},
			{X: 9, Y: 10},
			{X: 8, Y: 10},
		},
		direction: rl.Vector2{X: 0, Y: 0},
		points:    0,
		color:     color.RGBA{255, 0, 0, 255},
	}
}

func (s *snake) GetPoints() int {
	return s.points
}

func (s *snake) Update() {
	if len(s.body) > 1 {
		if s.direction.X != 0 || s.direction.Y != 0 {
			// Remove last element of s.Body and assign to tail. (pop)
			var tail rl.Vector2
			tail, s.body = s.body[len(s.body)-1], s.body[:len(s.body)-1]

			// Set coordinates of tail based on current direction
			if s.direction.X == 1 {
				tail.X = s.body[0].X + 1
				tail.Y = s.body[0].Y
			} else if s.direction.Y == 1 {
				tail.X = s.body[0].X
				tail.Y = s.body[0].Y + 1
			} else if s.direction.X == -1 {
				tail.X = s.body[0].X - 1
				tail.Y = s.body[0].Y
			} else if s.direction.Y == -1 {
				tail.X = s.body[0].X
				tail.Y = s.body[0].Y - 1
			}

			// Insert tail as the first element in s.Body (unshift)
			s.body = append([]rl.Vector2{tail}, s.body...)

			fmt.Printf("Body pos: %v\n", s.body[0])
		}
	} else {
		s.body[0] = rl.Vector2{
			X: s.direction.X,
			Y: s.direction.Y,
		}
	}
}

func (s *snake) getDirection() Direction {
	if s.direction.X == 0 && s.direction.Y == -1 { // up
		return DirectionUp
	} else if s.direction.X == 1 && s.direction.Y == 0 { // right
		return DirectionRight
	} else if s.direction.X == 0 && s.direction.Y == 1 { // down
		return DirectionDown
	} else if s.direction.X == -1 && s.direction.Y == 0 { // left
		return DirectionLeft
	} else { // stop
		return 0
	}
}

func (s *snake) ChangeDirection(dir Direction) {
	if dir == DirectionUp {
		if s.getDirection() != DirectionDown {
			s.direction = rl.Vector2{X: 0, Y: -1}
		}
	} else if dir == DirectionRight {
		if s.getDirection() != DirectionLeft {
			s.direction = rl.Vector2{X: 1, Y: 0}
		}
	} else if dir == DirectionDown {
		if s.getDirection() != DirectionUp {
			s.direction = rl.Vector2{X: 0, Y: 1}
		}
	} else if dir == DirectionLeft {
		if s.getDirection() != DirectionRight {
			s.direction = rl.Vector2{X: -1, Y: 0}
		}
	}
}

func (s *snake) Reset() {
	s.body = []rl.Vector2{
		{X: 10, Y: 10},
		{X: 9, Y: 10},
		{X: 8, Y: 10},
	}
	s.direction = rl.Vector2{X: 0, Y: 0}
	s.points = 0
	s.color = color.RGBA{255, 0, 0, 255}
}

func (s *snake) Eat(f *fruit) {

	// If snake head is on the same position as fruit
	if s.body[0].X == f.pos.X && s.body[0].Y == f.pos.Y {

		fmt.Println("Fruit match Snake pos")

		var new_x float32
		var new_y float32

		// Calculate where to add the new bodypart
		if s.direction.X == 1 {
			new_x = s.body[len(s.body)-1].X + 1
			new_y = s.body[len(s.body)-1].Y
		} else if s.direction.X == -1 {
			new_x = s.body[len(s.body)-1].X - 1
			new_y = s.body[len(s.body)-1].Y
		} else if s.direction.Y == 1 {
			new_x = s.body[len(s.body)-1].X
			new_y = s.body[len(s.body)-1].Y + 1
		} else if s.direction.Y == -1 {
			new_x = s.body[len(s.body)-1].X
			new_y = s.body[len(s.body)-1].Y - 1
		}

		// Append new bodypart to the end of snake
		s.body = append(s.body, rl.Vector2{X: new_x, Y: new_y})

		// Increase score
		s.points += 1

		// Place fruit on another random spot
		f.Update()
	}
}

func (s *snake) DetectCollision() bool {

	min_x := s.board.x / s.board.scale
	max_x := (s.board.w + s.board.x) / s.board.scale
	min_y := s.board.y / s.board.scale
	max_y := (s.board.h + s.board.y) / s.board.scale

	if s.body[0].X >= float32(max_x) ||
		s.body[0].X < float32(min_x) ||
		s.body[0].Y >= float32(max_y) ||
		s.body[0].Y < float32(min_y) {
		return true
	}

	return false
}

func (s *snake) Draw() {
	for _, v := range s.body {
		rl.DrawRectangle(
			int32(v.X * float32(s.board.scale)),
			int32(v.Y * float32(s.board.scale)),
			int32(s.board.scale),
			int32(s.board.scale),
			color.RGBA{
				s.color.R,
				s.color.G,
				s.color.B,
				s.color.A,
			},
		)
	}
}
