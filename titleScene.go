package main

import rl "github.com/gen2brain/raylib-go/raylib"

var (
	titleMeasureX int32
	titlePosX     int32

	versionPosX int32

	startPosX int32
)

// InitTitle inits logo
func InitTitle() {
	titleMeasureX = rl.MeasureText(titleText, 100)
	titlePosX = (width - titleMeasureX) / 2

	versionPosX = titlePosX + titleMeasureX + 10

	startPosX = (width - rl.MeasureText(startMessage, 30)) / 2
}

// DrawTitle draws the title scene
func DrawTitle() {
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

	rl.DrawText(titleText, titlePosX, 40, 100, rl.Black)
	rl.DrawText(versionText, versionPosX, 95, 30, rl.LightGray)
	rl.DrawText(startMessage, startPosX, height-70, 30, rl.LightGray)

}
