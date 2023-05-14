package main

import (
	"fmt"
	"image/color"
	"strconv"

	rlgui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	scale   int32 = 5
	screenW int32 = int32(N) * scale
	screenH int32 = int32(N) * scale
)

var flame = []color.RGBA{
	{0, 0, 0, 255},
	{255, 0, 0, 255},
	{226, 88, 34, 255},
}

var checked = false

var customColor = []rl.Color{{0, 0, 0, 255}, {125, 125, 125, 255}, {0, 0, 255, 255}}

var density float32 = 200
var velocity float32 = 5

func lerp_rbga(color1 color.RGBA, color2 color.RGBA, t float64) color.RGBA {

	var rgba color.RGBA

	rgba.R = uint8(float64(color1.R)*(1-t) + float64(color2.R)*t)
	rgba.G = uint8(float64(color1.G)*(1-t) + float64(color2.G)*t)
	rgba.B = uint8(float64(color1.B)*(1-t) + float64(color2.B)*t)
	rgba.A = 255

	return rgba

}

func lerp_rbga_triple(color1 color.RGBA, color2 color.RGBA, color3 color.RGBA, t float64) color.RGBA {

	var rgba_a color.RGBA
	var rgba_b color.RGBA
	var rgba_c color.RGBA

	rgba_a.R = uint8(float64(color1.R)*(1-t) + float64(color2.R)*t)
	rgba_a.G = uint8(float64(color1.G)*(1-t) + float64(color2.G)*t)
	rgba_a.B = uint8(float64(color1.B)*(1-t) + float64(color2.B)*t)

	rgba_b.R = uint8(float64(color2.R)*(1-t) + float64(color3.R)*t)
	rgba_b.G = uint8(float64(color2.G)*(1-t) + float64(color3.G)*t)
	rgba_b.B = uint8(float64(color2.B)*(1-t) + float64(color3.B)*t)

	rgba_c.R = uint8(float64(rgba_a.R)*(1-t) + float64(rgba_b.R)*t)
	rgba_c.G = uint8(float64(rgba_a.G)*(1-t) + float64(rgba_b.G)*t)
	rgba_c.B = uint8(float64(rgba_a.B)*(1-t) + float64(rgba_b.B)*t)

	rgba_c.A = 255

	return rgba_c
}

func normalize(a float64, min float64, max float64) float64 {
	return a - min/max - min
}

func renderDensity(f Fluid) {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			x := int32(i) * scale
			y := int32(j) * scale
			d := f.density[IX(i, j)]
			n := normalize(d, 0, 10000)
			lerp_color := lerp_rbga(customColor[0], customColor[2], n)
			rl.DrawRectangle(x, y, scale, scale, lerp_color)
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
				addDensity(f, cx, cy, float64(density))
				amtX := pos.X - previousPos.X
				amtY := pos.Y - previousPos.Y
				addVelocity(f, cx, cy, float64(amtX)+float64(velocity), float64(amtY)+float64(velocity))
				previousPos = pos
			}
		}

		// Start fluid simulation
		step(f)
		renderDensity(f)
		fadeDensity(f, 0.90)

		rl.DrawText("Fluid simulation", screenH/2+40, 20, 30, rl.White)

		// alternating variables via GUI

		// Choosing colors to interpolate
		rec := rl.Rectangle{X: 20, Y: float32(screenH) - 70, Width: 200, Height: 50}
		customColor[2] = rlgui.ColorPicker(rec, "text", customColor[2])
		// Choose density
		rec1 := rl.Rectangle{X: 270, Y: float32(screenH) - 44, Width: 200, Height: 20}
		density = rlgui.SliderBar(rec1, "", "", density, 0, 1000)
		dentxt := fmt.Sprintf("Density: %v", density)
		rl.DrawText(dentxt, 480, screenH-44, 16, rl.White)
		// Choose velocity
		rec2 := rl.Rectangle{X: 270, Y: float32(screenH) - 68, Width: 200, Height: 20}
		velocity = rlgui.SliderBar(rec2, "", "", velocity, 0, 10)
		veltxt := fmt.Sprintf("Velocity: %v", velocity)
		rl.DrawText(veltxt, 480, screenH-68, 16, rl.White)

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
