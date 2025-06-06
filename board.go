// Copyright (c) 2025, Mads Moi-Aune <mads@moiaune.dev>
//
// SPDX-License-Identifier: BSD-3-Clause

package main

import rl "github.com/gen2brain/raylib-go/raylib"

type board struct {
	w, h  float32
	x, y  float32
	scale float32
}

func NewBoard(w, h, x, y, scale float32) *board {
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
