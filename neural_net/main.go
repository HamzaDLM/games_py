package main

import (
	"bufio"
	"fmt"
	"image/color"
	"os"
	"strconv"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenW = 1600
	screenH = 900
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
			cx := int32(g.x + 15 + (i * 15))
			cy := int32(g.y + 15 + (j * 15))
			rl.DrawCircle(cx, cy, float32(g.r), color.RGBA{255, 255, 255, l[i][j]})
			rl.DrawCircleLines(cx, cy, float32(g.r), color.RGBA{255, 255, 255, 100})
		}
	}
}

// // Input array containing values for color alpha
func initializeMatrix(size int) [][]uint8 {
	var inputArr [][]uint8
	for i := 0; i < size; i++ {
		inputArr = append(inputArr, make([]uint8, size))
	}
	return inputArr
}

// Train csv file contains following format: Label, Pixel1, ..., PixelN
func parseTrain(filename string) map[uint8][]uint8 {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data := make(map[uint8][]uint8, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		s := strings.Split(row, ",")
		// sublist holding pixel values
		var lints []uint8
		// Parse string to uint and add to sublist
		for i := 1; i < len(s); i++ {
			d, err := strconv.ParseUint(s[i], 10, 64)
			if err == nil {
				lints = append(lints, uint8(d))
			}
		}
		key, err := strconv.ParseUint(s[0], 10, 64)
		if err == nil {
			data[uint8(key)] = lints
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return data
}

// Turn a 1 dimension array to a 2 dimensions array
func map1Dto2D(l1 []uint8, l2 [][]uint8) [][]uint8 {
	for i := 0; i < len(l2); i++ {
		for j := 0; j < len(l2[i]); j++ {
			index := j*len(l2) + i
			l2[i][j] = l1[index]
		}
	}
	return l2
}

func matrixDot() {

}

func main() {
	fmt.Println("Neural Network")

	// Import training set
	// trainList := parseTrain("train.csv")

	// Initialize window
	rl.InitWindow(screenW, screenH, "Neural Network Visualization")
	defer rl.CloseWindow()
	// set FPS
	rl.SetTargetFPS(60)

	gridInfo := Grid{size: 28, r: 5, x: 100, y: 200}
	gridArray := initializeMatrix(gridInfo.size)
	// gridArray = map1Dto2D(trainList[6], gridArray)
	// Main game loop
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		pos := rl.GetMousePosition()

		// User to draw a number on the grid
		if rl.IsMouseButtonDown(0) {
			// Add the side fades of each click (near cells of the clicked cell should be greyish)
			for i := 0; i < gridInfo.size; i++ {
				for j := 0; j < gridInfo.size; j++ {
					cx := int32(gridInfo.x + 15 + (i * 15))
					cy := int32(gridInfo.y + 15 + (j * 15))
					if int32(pos.X) > cx-int32(gridInfo.r)-4 && int32(pos.X) <= cx+int32(gridInfo.r)-4 &&
						int32(pos.Y) > cy-int32(gridInfo.r)+4 && int32(pos.Y) <= cy+int32(gridInfo.r)+4 {
						gridArray[i][j] = 255
					}
				}
			}
		}

		rl.DrawText("Drag to write a number", 190, 170, 20, rl.White)
		makeGrid(gridInfo, gridArray)
		// Grid clear button
		// TODO make the values into variables
		if rl.IsMouseButtonDown(0) && pos.X > 225 && pos.X < 375 && pos.Y > 650 && pos.Y < 680 {
			rl.DrawRectangleLines(225, 650, 150, 30, rl.White)
			rl.DrawText("Clear Grid", 255, 655, 18, rl.White)
			gridArray = initializeMatrix(gridInfo.size)
		} else {
			rl.DrawRectangle(225, 650, 150, 30, rl.Gray)
			rl.DrawText("Clear Grid", 255, 655, 18, rl.Black)
		}
		rl.DrawText("28 x 28", 460, 635, 16, rl.Gray)

		rl.DrawText("Basic Neural Network", screenW/2-200, 20, 36, rl.Blue)
		// Show FPS
		fps := strconv.FormatInt(int64(rl.GetFPS()), 10)
		t := fmt.Sprintf("FPS: %s", fps)
		rl.DrawText(t, 20, 20, 20, rl.White)

		rl.EndDrawing()

	}
}
