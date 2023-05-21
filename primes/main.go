// idea: objects made using fibo sequence represent prime values in spiral
// one of them suppleis an argument to the other.
package main

import (
	"fmt"
	"image/color"
	"math"

	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenW = 1600
	screenH = 1000
)

var white = color.RGBA{255, 255, 255, 255}
var whitefaded = color.RGBA{255, 255, 255, 80}

type polar struct {
	Radius float64
	Angle  float64
}

type cartesian struct {
	X int32
	Y int32
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

func initializeArray(a []bool, n int) {
	for i := 0; i < n; i++ {
		a[i] = true
	}
}

// Prime number is a natural number that is divisable by 1 and themselves.
// Calculate prime numbers using Sieve of Eratosthenes algorithm
func findPrimes(n int) []int {
	a := make([]bool, n+1)
	initializeArray(a, n+1)
	for i := 2; i*i <= n; i++ {
		if a[i] == true {
			j := int(math.Pow(float64(i), 2))
			for j = 0; j < n; j++ {
				a[j] = false
				j = j + i
			}
		}
	}
	var lastPrimes []int
	for i := 2; i <= n; i++ {
		if a[i] {
			lastPrimes = append(lastPrimes, i)
		}
	}
	return lastPrimes

}

// Polar coordinates to Cartesian:
// x = r • cos(O), y = r • sin(O)
func polarToCartesian(p polar) cartesian {
	var x int32 = int32(p.Radius*math.Cos(p.Angle) + screenW/2)
	var y int32 = int32(p.Radius*math.Sin(p.Angle) + screenH/2)

	return cartesian{X: x, Y: y}
}

// Draw a polar coordinate system
// splits determines the minimum angle between lines to display
func drawPolarPlane(minRadius int) {
	rl.DrawCircle(screenW/2, screenH/2, 5, whitefaded)
	rl.DrawCapsule(rl.Vector3{X: screenW / 2, Y: screenH / 2, Z: 0}, rl.Vector3{X: screenW, Y: screenH, Z: 100}, 50, 10, 10, whitefaded)
	rl.DrawLine(screenW/2, screenH/2, screenW/2, 0, whitefaded)
	rl.DrawLine(screenW/2, screenH/2, 0, screenH/2, whitefaded)
	rl.DrawLine(screenW/2, screenH/2, screenW, screenH/2, whitefaded)
	rl.DrawLine(screenW/2, screenH/2, screenW/2, screenH, whitefaded)
}

func main() {
	fmt.Println("Starting Game")

	// Initialize window
	rl.InitWindow(screenW, screenH, "Primes")
	defer rl.CloseWindow()
	// FPS
	rl.SetTargetFPS(60)

	// fmt.Println(findLastPrimes(20))
	// Main game loop
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		// rl.DrawText("Primes Spiral", screenW-200, 20, 30, rl.White)

		// Draw the spiral plane
		drawPolarPlane(10)
		// Draw the prime points
		n := 400
		listOfPrimes := findPrimes(n)
		for i := 0; i < len(listOfPrimes); i++ {
			c := polarToCartesian(polar{Radius: float64(listOfPrimes[i]), Angle: float64(listOfPrimes[i])})
			t := (float64(listOfPrimes[i]) - 2) / (float64(n) - 2)
			rl.DrawCircle(c.X, c.Y, 5,
				lerp_rbga_triple(
					color.RGBA{255, 0, 0, 255},
					color.RGBA{255, 255, 0, 255},
					color.RGBA{0, 255, 255, 255}, t))
		}

		// Show FPS
		fps := strconv.FormatInt(int64(rl.GetFPS()), 10)
		t := fmt.Sprintf("FPS: %s", fps)
		rl.DrawText(t, 20, 20, 20, rl.White)

		rl.EndDrawing()
	}
}
