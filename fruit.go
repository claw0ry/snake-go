// Copyright (c) 2025, Mads Moi-Aune <mads@moiaune.dev>
//
// SPDX-License-Identifier: BSD-3-Clause

package main

import (
	"fmt"
	"math/rand"
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type fruit struct {
	board *board
	pos   rl.Vector2
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
	f.pos = rl.Vector2{
		X: float32(rand.Intn(int((f.board.w+f.board.x)/f.board.scale) + 1)),
		Y: float32(rand.Intn(int((f.board.h+f.board.y-20)/f.board.scale) + 2)),
	}

	fmt.Printf("Fruit pos: %v\n", f.pos)
}

func (f *fruit) Draw() {
	rl.DrawRectangle(
		int32(f.pos.X * f.board.scale),
		int32(f.pos.Y * f.board.scale),
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
