package main

import (
	"math"
	"math/rand"
	"strconv"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	playerX int32 = 10
	playerY int32 = 10

	enemyX int32 = 10
	enemyY int32 = 50

	playerScore = 0
	enemyScore  = 0

	puckX int32
	puckY int32

	puckAngle float64

	puckXSpeed float32
	puckYSpeed float32
)

var (
	scoreSound rl.Sound
	wallSound  rl.Sound
	barSound   rl.Sound
)

// InitPlaying inits the playing scene
func InitPlaying() {
	scoreSound = rl.LoadSound("res/score.mp3")
	wallSound = rl.LoadSound("res/wall.mp3")
	rl.SetSoundPitch(wallSound, 2)
	barSound = rl.LoadSound("res/bar.mp3")

	resetPuck()
}

func resetPuck() {
	var direction float32 = 1

	playerX = width - 20

	puckX = width / 2
	puckY = height / 2

	puckAngle = degToRad(float64(getRandom(-45, 45)))

	//puckAngle = degToRad(100)
	if getRandom(0, 1) == 1 {
		direction = -1
	}

	puckXSpeed = float32(puckSpeed*math.Cos(puckAngle)) * direction
	puckYSpeed = float32(puckSpeed*math.Sin(puckAngle)) * direction
}

// DrawPlaying draws the playing scene
func DrawPlaying() {
	movePlayer()

	movePuck()
	checkPuckCollide()

	rl.BeginDrawing()
	defer rl.EndDrawing()

	rl.ClearBackground(rl.RayWhite)

	// ScoreBoard
	scoreString := strconv.Itoa(enemyScore) + " | " + strconv.Itoa(playerScore)
	scoreX := (width - rl.MeasureText(scoreString, 40)) / 2
	rl.DrawText(scoreString, scoreX, 20, 40, rl.Maroon)

	// Player
	rl.DrawRectangle(playerX, playerY, 10, barHeight, rl.Black)

	// Enemy
	rl.DrawRectangle(enemyX, enemyY, 10, barHeight, rl.Black)

	// Puck
	rl.DrawRectangle(puckX, puckY, puckSize, puckSize, rl.Black)
}

func checkPuckCollide() {
	// Puck touches top or bottom of screen
	if puckY <= 0 || puckY+puckSize >= height {
		puckYSpeed *= -1
		rl.PlaySound(wallSound)
	}

	// Puck collided with enemy
	enemyRect := rl.Rectangle{X: float32(enemyX), Y: float32(enemyY), Width: 10.0, Height: float32(barHeight)}
	playerRect := rl.Rectangle{X: float32(playerX), Y: float32(playerY), Width: 10.0, Height: float32(barHeight)}
	puckRect := rl.Rectangle{X: float32(puckX), Y: float32(puckY), Width: float32(puckSize), Height: float32(puckSize)}

	collidedWithEnemy := rl.CheckCollisionRecs(puckRect, enemyRect)
	collidedWithPlayer := rl.CheckCollisionRecs(puckRect, playerRect)
	if collidedWithEnemy || collidedWithPlayer {
		if puckXSpeed > 0 {
			puckXSpeed += speedIncrement
		} else {
			puckXSpeed -= speedIncrement
		}

		if puckYSpeed > 0 {
			puckYSpeed += speedIncrement
		} else {
			puckYSpeed -= speedIncrement
		}

		puckXSpeed *= -1

		// Prevent it from getting stuck
		if collidedWithEnemy {
			puckX += 10
		} else {
			puckX -= 10
		}

		rl.PlaySound(barSound)
	}

	// Puck reaches left side of the screen
	if puckX <= -puckSize {
		resetPuck()
		playerScore++
		rl.PlaySound(scoreSound)
	}

	// Puck reaches right side of the screen
	if puckX >= width {
		resetPuck()
		enemyScore++
		rl.PlaySound(scoreSound)
	}
}

func movePuck() {
	puckX += int32(puckXSpeed)
	puckY += int32(puckYSpeed)
}

func movePlayer() {
	// 265 is KeyUP, 87 is KeyW
	if rl.IsKeyDown(265) {
		if playerY > 10 {
			playerY -= movementSpeed
		}
	}
	if rl.IsKeyDown(87) {
		if enemyY > 10 {
			enemyY -= movementSpeed
		}
	}

	// 264 is KeyDOWN, 83 is KeyS
	if rl.IsKeyDown(264) {
		if playerY < (height - barHeight - 15) {
			playerY += movementSpeed
		}
	}
	if rl.IsKeyDown(83) {
		if enemyY < (height - barHeight - 15) {
			enemyY += movementSpeed
		}
	}
}

func getRandom(a int, b int) int {
	rand.Seed(time.Now().UnixNano())
	return a + rand.Intn(b-a+1)
}

func radToDeg(rad float64) float64 {
	return (180 * rad) / math.Pi
}

func degToRad(deg float64) float64 {
	return (deg * math.Pi) / 180
}
