package main

import rl "github.com/gen2brain/raylib-go/raylib"

// InitEnd inits the playing scene
func InitEnd() {

}

// DrawEnd draws the playing scene
func DrawEnd() {

	if rl.IsKeyPressed(32) {
		screen = 2
	}

	framesCounter++

	if framesCounter > 59 {
		framesCounter = 0
	}

	rl.BeginDrawing()
	defer rl.EndDrawing()

	rl.ClearBackground(rl.RayWhite)

	rl.DrawText("ENDING SCREEN", 20, 20, 40, rl.DarkBlue)

}
