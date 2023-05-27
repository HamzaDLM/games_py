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
	size    = 28
)

type Grid struct {
	size int
	r    int
	x    int
	y    int
}

func makeGrid(g Grid, l []uint8) {
	for i := 0; i < g.size; i++ {
		for j := 0; j < g.size; j++ {
			cx := int32(g.x + 15 + (i * 15))
			cy := int32(g.y + 15 + (j * 15))
			rl.DrawCircle(cx, cy, float32(g.r), color.RGBA{255, 255, 255, l[IX(i, j)]})
			rl.DrawCircleLines(cx, cy, float32(g.r), color.RGBA{255, 255, 255, 100})
		}
	}
}

func drawNeuralNetwork() {
	rl.DrawCircle(0, 0, 10, color.RGBA{255, 255, 255, 255})
}

// Train csv file contains following format: Label, Pixel1, ..., PixelN
func parseTrain(filename string) (Matrix, Matrix) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	dataInputs := Matrix{
		data:    make([]float64, 0),
		rowSize: 784,
		colSize: -1,
	}
	dataLabels := Matrix{
		data:    make([]float64, 0),
		rowSize: 0,
		colSize: 1,
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		s := strings.Split(row, ",")
		dataInputs.colSize += 1
		for i := 0; i < len(s); i++ {
			d, err := strconv.ParseFloat(s[i], 64)
			if err == nil {
				if i == 0 {
					dataLabels.data = append(dataLabels.data, d)
					dataLabels.rowSize += 1
				} else {
					dataInputs.data = append(dataInputs.data, d/255)
				}

			}
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return dataInputs, dataLabels
}

func main() {
	// Create the neural net
	nn := NeuralNetwork{
		inputNeuronsSize:        size * size,
		outputNeuronsSize:       10,
		hiddenLayers:            2,
		hiddenLayersNeuronsSize: []int{16, 16}, // for each hidden layer
		weights:                 make([]Matrix, size),
		biases:                  make([]Matrix, size),
		epochs:                  1000,
		learningRate:            0.1,
	}

	// Import training set
	trainInputs, trainLabels := parseTrain("data/train.csv")
	// fmt.Println(trainInputs.dims())
	// fmt.Println(trainLabels.dims())
	// Start the training
	fmt.Println("Starting the learning process")
	go nnLearn(&nn, &trainInputs, &trainLabels)

	// m1 := createMatrix(3, 4)
	// m1.data = []float64{1, 2, 3, 4, 0, 0, 0, 0, 0, 0, 0, 0}
	// m1.matrixMultScalar(2)
	// printMatrix(m1)

	// Initialize window
	rl.InitWindow(screenW, screenH, "Neural Network Visualization")
	defer rl.CloseWindow()
	// set FPS
	rl.SetTargetFPS(60)

	gridInfo := Grid{size: size, r: 5, x: 100, y: 200}
	gridArray := make([]uint8, size*size)

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
						gridArray[IX(i, j)] = 255
					}
				}
			}
		}

		drawNeuralNetwork()

		rl.DrawText("Drag to write a number", 190, 170, 20, rl.White)
		makeGrid(gridInfo, gridArray)
		// Grid clear button
		// TODO make the values into variables
		if rl.IsMouseButtonDown(0) && pos.X > 225 && pos.X < 375 && pos.Y > 650 && pos.Y < 680 {
			rl.DrawRectangleLines(225, 650, 150, 30, rl.White)
			rl.DrawText("Clear Grid", 255, 655, 18, rl.White)
			gridArray = make([]uint8, size*size)
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
