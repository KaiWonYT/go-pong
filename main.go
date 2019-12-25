package main

import rl "github.com/gen2brain/raylib-go/raylib"

// Game variables
var (
	screen        = 0
	framesCounter = 0
	isRunning     = true
)

func main() {
	rl.InitWindow(width, height, title)
	defer rl.CloseWindow()

	rl.InitAudioDevice()
	defer rl.CloseAudioDevice()

	rl.SetTargetFPS(targetFps)

	// Init all scenes
	InitLogo()
	InitTitle()
	InitPlaying()
	InitEnd()

	for isRunning && !rl.WindowShouldClose() {
		switch screen {
		case 0:
			DrawLogo()
		case 1:
			DrawTitle()
		case 2:
			DrawPlaying()
		case 3:
			DrawEnd()
		}
	}
}
