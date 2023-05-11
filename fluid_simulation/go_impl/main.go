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

func renderDensity(f Fluid) {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			x := int32(i) * scale
			y := int32(j) * scale
			d := f.density[IX(i, j)]
			// d_scaled := constrain(d, 0, 255)
			// d_norm := uint8(d_scaled)
			// if d != 0 {
			// 	fmt.Printf("Density not 0 at: x: %d, y: %d, IX: %d, density: %v \n", x, y, IX(i, j), d)
			// }
			// fmt.Println(d_scaled, d_norm)
			rl.DrawRectangle(x, y, scale, scale, color.RGBA{255, 255, 255, uint8(int(d) % 255)})
		}
	}
}

func main() {
	// Initialize fluid
	f := newFluid(int64(N), 0.1, 0, 0)
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

		rl.DrawText("Fluid simulation", screenH/2+40, 20, 30, rl.White)

		fps := strconv.FormatInt(int64(rl.GetFPS()), 10)
		t := fmt.Sprintf("FPS: %s", fps)
		rl.DrawText(t, 20, 20, 20, rl.White)

		// fluid cycle
		if rl.IsMouseButtonDown(0) {
			pos := rl.GetMousePosition()
			max := float32(N * int(scale))
			if 0 < pos.X && pos.X < max && 0 < pos.Y && pos.Y < max {
				cx := int(pos.X / float32(scale))
				cy := int(pos.Y / float32(scale))
				addDensity(f, cx, cy, 100)
				addVelocity(f, cx, cy, 1, 1)
			}
		}
		step(f)
		renderDensity(f)
		// fadeDensity(f)

		rl.EndDrawing()
	}
}
