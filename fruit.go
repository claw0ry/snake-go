package main

import (
	"fmt"
	"math/rand"
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type fruit struct {
	board *board
	pos   point
	color color.RGBA
}

func NewFruit(b *board) *fruit {
	f := &fruit{
		board: b,
		color: color.RGBA{0, 255, 0, 255},
	}

	f.Update()

	return f
}

func (f *fruit) Update() {
	f.pos = point{
		x: int32(rand.Intn((f.board.w+f.board.x)/f.board.scale) + 1),
		y: int32(rand.Intn((f.board.h+f.board.y-20)/f.board.scale) + 2),
	}

	fmt.Printf("Fruit pos: %v\n", f.pos)
}

func (f *fruit) Draw() {
	rl.DrawRectangle(
		f.pos.x * int32(f.board.scale),
		f.pos.y * int32(f.board.scale),
		int32(f.board.scale),
		int32(f.board.scale),
		color.RGBA{
			f.color.R,
			f.color.G,
			f.color.B,
			f.color.A,
		},
	);
}
