package main

import (
	"fmt"
	"image/color"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenW = 1400
	screenH = 800
)

type Grid struct {
	size int
	r    int
	x    int
	y    int
}

func makeGrid(g Grid, l [][]uint8) {
	for i := 0; i < g.size; i++ {
		for j := 0; j < g.size; j++ {
			cx := int32(g.x + g.size + (i * g.size))
			cy := int32(g.y + g.size + (j * g.size))
			rl.DrawCircle(cx, cy, float32((g.r/2)-2), color.RGBA{255, 255, 255, l[i][j]})
			rl.DrawCircleLines(cx, cy, float32((g.r/2)-2), rl.White)
		}
	}
}

// // Input array containing values for color alpha
func initializeArray(size int) [][]uint8 {
	var inputArr [][]uint8
	for i := 0; i < size; i++ {
		inputArr = append(inputArr, make([]uint8, size))
	}
	return inputArr
}

func main() {
	fmt.Println("Neural Network")

	// Initialize window
	rl.InitWindow(screenW, screenH, "Primes")
	defer rl.CloseWindow()
	// set FPS
	rl.SetTargetFPS(40)

	gridSize := 20
	inputArray := initializeArray(gridSize)

	// Main game loop
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		rl.GetMouseX()

		rl.DrawText("Drag to write a number", 190, 170, 20, rl.White)
		makeGrid(Grid{size: gridSize, r: 20, x: 100, y: 200}, inputArray)

		// Show FPS
		fps := strconv.FormatInt(int64(rl.GetFPS()), 10)
		t := fmt.Sprintf("FPS: %s", fps)
		rl.DrawText(t, 20, 20, 20, rl.White)

		rl.EndDrawing()

	}
}
