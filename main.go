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
	rl.SetTargetFPS(targetFps)

	// Init all scenes
	InitLogo()
	InitTitle()

	for isRunning && !rl.WindowShouldClose() {
		switch screen {
		case 0:
			DrawLogo()
		case 1:
			DrawTitle()
		}
		/*
			// Game logic
			doGameLogic()

			// Draw
			drawGame()
		*/
	}

	rl.CloseWindow()
}

func drawGame() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)

	switch screen {
	case 0:
		// Logo Screen
		alphaValue := framesCounter * 10
		if alphaValue > 254 {
			alphaValue = 255
		}
		rl.DrawText(logoText, logoPosX, logoPosY, 40, rl.Color{R: 0, G: 0, B: 0, A: uint8(alphaValue)})
	case 1:
		// Title Screen
		rl.DrawText(titleText, titlePosX, 40, 100, rl.Black)
		rl.DrawText(versionText, versionPosX, 95, 30, rl.LightGray)
		rl.DrawText(startMessage, startPosX, height-70, 30, rl.LightGray)
	case 2:
		// Update gameplay screen here
		rl.DrawText("GAMEPLAY SCREEN", 20, 20, 40, rl.Maroon)
	case 3:
		// Update END screen here
		rl.DrawText("ENDING SCREEN", 20, 20, 40, rl.DarkBlue)
	}

	rl.EndDrawing()
}

func doGameLogic() {
	switch screen {
	case 0:
		// Draw logo here for 180 frames
		framesCounter++

		if framesCounter > 39 {
			screen = 1
			framesCounter = 0
		}

	case 1:
		// Title screen here
		if rl.IsKeyPressed(32) {
			screen = 2
			return
		}

		framesCounter++

		if framesCounter > 59 {
			framesCounter = 0
		}
	case 2:
		// Update gameplay screen here
	case 3:
		// Update END screen here
	}
}

func mapRanges(value int, start1 int, stop1 int, start2 int, stop2 int) int {
	return start2 + (stop2-start2)*((value-start1)/(stop1-start1))
}
