// Copyright (c) 2025, Mads Moi-Aune <mads@moiaune.dev>
//
// SPDX-License-Identifier: BSD-3-Clause

package main

import (
	"fmt"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	WIN_W, WIN_H int32 = 420, 240
	BOARD_SCALE  int32 = 10
)

var (
	MAX_FPS int32 = 15
)

type point struct {
	x int32
	y int32
}

func run() (err error) {
	rl.InitWindow(WIN_W, WIN_H, "Snake-GO")
	defer rl.CloseWindow()

	rl.SetTargetFPS(MAX_FPS)

	// TODO: Need label for points
	// TODO: Need label for highscore

	board := NewBoard(float32(WIN_W-20), float32(WIN_H-30), 10, 20, float32(BOARD_SCALE))
	snake := NewSnake(board)
	fruit := NewFruit(board)
	pause := false

	// Game Loop
	for !rl.WindowShouldClose() {
		// 1. Process input
		if rl.IsKeyPressed(rl.KeyW) {
			snake.ChangeDirection(DirectionUp)
		}

		if rl.IsKeyPressed(rl.KeyD) {
			snake.ChangeDirection(DirectionRight)
		}

		if rl.IsKeyPressed(rl.KeyS) {
			snake.ChangeDirection(DirectionDown)
		}

		if rl.IsKeyPressed(rl.KeyA) {
			snake.ChangeDirection(DirectionLeft)
		}

		if rl.IsKeyPressed(rl.KeyP) {
			pause = !pause
		}

		if rl.IsKeyPressed(rl.KeySpace) {
			snake.Reset()
			fruit.Update()
		}

		// 2. Update objects
		if !pause {
			update(snake, fruit)
		}

		// 3. Draw objects
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		board.Draw()
		snake.Draw()
		fruit.Draw()

		rl.DrawText(fmt.Sprintf("%d", snake.GetPoints()), 10, 10, 12, rl.Black)

		if pause {
			rl.DrawText("PAUSED", WIN_W/2, WIN_H/2, 30, rl.Gray)
		}

		rl.EndDrawing()
	}


	return
}

func update(s *snake, f *fruit) {
	if s.points == 5 {
		// speed = 20
		MAX_FPS = 20
	} else if s.points == 10 {
		// speed = 25
		MAX_FPS = 25
	} else if s.points == 15 {
		// speed = 30
		MAX_FPS = 30
	} else if s.points == 20 {
		// speed = 35
		MAX_FPS = 35
	} else if s.points < 5 {
		// speed = 15
		MAX_FPS = 15
	}

	rl.SetTargetFPS(MAX_FPS)

	if s.DetectCollision() {
		s.Reset()
	} else {
		s.Eat(f)
		s.Update()
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "ERR: %+v\n", err)
		os.Exit(1)
	}
}
