package main

import rl "github.com/gen2brain/raylib-go/raylib"

var (
	logoPosX int32
	logoPosY int32
)

// InitLogo inits logo
func InitLogo() {
	logoPosX = (width - rl.MeasureText(logoText, 40)) / 2
	logoPosY = (height - 40) / 2
}

// DrawLogo draws the logo
func DrawLogo() {
	framesCounter++

	if framesCounter > 39 {
		screen = 1
		framesCounter = 0
	}

	rl.BeginDrawing()
	defer rl.EndDrawing()

	rl.ClearBackground(rl.RayWhite)

	alphaValue := framesCounter * 10
	if alphaValue > 254 {
		alphaValue = 255
	}

	rl.DrawText(logoText, logoPosX, logoPosY, 40, rl.Color{R: 0, G: 0, B: 0, A: uint8(alphaValue)})
}
