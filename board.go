package main

import rl "github.com/gen2brain/raylib-go/raylib"

type board struct {
	w, h  int
	x, y  int
	scale int
}

func NewBoard(w, h, x, y, scale int) *board {
	return &board{
		w:     w,
		h:     h,
		x:     x,
		y:     y,
		scale: scale,
	}
}

func (b *board) Draw() {
	rl.DrawRectangle(
		int32(b.x),
		int32(b.y),
		int32(b.w),
		int32(b.h),
		rl.RayWhite,
	)
}
