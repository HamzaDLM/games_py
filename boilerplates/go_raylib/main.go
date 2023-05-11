package main

import (
	"fmt"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenW = 960
	screenH = 640
	N = 124
	iter = 
)

func renderFps() {
	fps := strconv.FormatInt(int64(rl.GetFPS()), 10)
	t := fmt.Sprintf("FPS: %s", fps)
	rl.DrawText(t, 20, 20, 20, rl.White)
}

func main() {
	fmt.Println("Starting Game")

	// Initialize window
	rl.InitWindow(screenW, screenH, "raylib [core] example - basic window")
	defer rl.CloseWindow()
	// FPS
	rl.SetTargetFPS(60)

	// Main game loop
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)
		rl.DrawText("Fluid simulation", screenH/2+40, 20, 30, rl.White)

		renderFps()

		rl.EndDrawing()
	}
}
