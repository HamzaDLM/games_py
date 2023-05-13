package main

import (
	"fmt"
	"image/color"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	scale   int32 = 5
	screenW int32 = int32(N) * scale
	screenH int32 = int32(N) * scale
)

func lerpColor(a float32) {
	return
}

func renderDensity(f Fluid) {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			x := int32(i) * scale
			y := int32(j) * scale
			d := f.density[IX(i, j)]
			rl.DrawRectangle(x, y, scale, scale, color.RGBA{255, 255, 255, uint8(int(d) % 255)})
		}
	}
}

func main() {
	previousPos := rl.GetMousePosition()
	// Initialize fluid
	f := newFluid(int64(N), 0.2, 0, 0.0000001)
	fmt.Println(len(f.density))

	// Initialize window
	rl.InitWindow(screenW, screenH, "Fluid Simulation")
	defer rl.CloseWindow()
	// FPS
	rl.SetTargetFPS(40)

	// Main game loop
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		if rl.IsMouseButtonDown(0) {
			pos := rl.GetMousePosition()
			max := float32(N * int(scale))
			if 0 < pos.X && pos.X < max && 0 < pos.Y && pos.Y < max {
				cx := int(pos.X / float32(scale))
				cy := int(pos.Y / float32(scale))
				addDensity(f, cx, cy, 200)
				amtX := pos.X - previousPos.X
				amtY := pos.Y - previousPos.Y
				addVelocity(f, cx, cy, float64(amtX)+1, float64(amtY)+1)
				previousPos = pos
			}
		}
		step(f)
		renderDensity(f)
		fadeDensity(f, 0.99)

		rl.DrawText("Fluid simulation", screenH/2+40, 20, 30, rl.White)

		fps := strconv.FormatInt(int64(rl.GetFPS()), 10)
		t := fmt.Sprintf("FPS: %s (capped at 40)", fps)
		rl.DrawText(t, 20, 20, 20, rl.White)
		t1 := fmt.Sprintf("Size: %d", N)
		rl.DrawText(t1, 20, 40, 20, rl.White)
		t2 := fmt.Sprintf("Iter: %d", iter)
		rl.DrawText(t2, 20, 60, 20, rl.White)

		rl.EndDrawing()
	}
}
