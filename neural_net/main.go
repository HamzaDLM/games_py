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

// Draw a neural network architecture
// x, y are the drawing starting position
// config list contains number of nodes in each layer
func drawNeuralNetwork(x, y, w, h, radius int, config []int, nn *NeuralNetwork) {
	var offset int32 = 40 // How far the nodes are from each other vertically
	// Connect nodes
	for layer := 0; layer < len(config)-1; layer++ {
		currentLayerNeurons := config[layer]
		// limit number of neurons
		if currentLayerNeurons > 20 {
			currentLayerNeurons = 10
		}
		nextLayerNeurons := config[layer+1]

		for currentNeuron := 0; currentNeuron < currentLayerNeurons; currentNeuron++ {
			for nextNeuron := 0; nextNeuron < nextLayerNeurons; nextNeuron++ {
				currentX := int32((w/(len(config)+1))*(layer+1)) + int32(x) + int32(radius)
				nextX := int32((w/(len(config)+1))*(layer+2)) + int32(x) - int32(radius)

				currentY := int32((h-(int(currentLayerNeurons)*50))/2) + int32(y)
				nextY := int32((h-(int(nextLayerNeurons)*50))/2) + int32(y)
				rl.DrawLineEx(
					rl.Vector2{X: float32(currentX), Y: float32(currentY + (int32(currentNeuron) * offset))},
					rl.Vector2{X: float32(nextX), Y: float32(nextY + (int32(nextNeuron) * offset))},
					1,
					rl.Gray)
			}
		}
	}
	// Display nodes
	for layer := 0; layer < len(config); layer++ {
		layerNeurons := config[layer]
		// limit number of neurons
		if layerNeurons > 20 {
			layerNeurons = 10
		}

		y := int32((h-(int(layerNeurons)*50))/2) + int32(y)

		for neuron := 0; neuron < layerNeurons; neuron++ {
			x := int32((w/(len(config)+1))*(layer+1)) + int32(x)

			rl.DrawCircle(x, y+(int32(neuron)*offset), float32(radius), rl.Black)
			rl.DrawCircleLines(x, y+(int32(neuron)*offset), float32(radius), rl.White)
			// Draw the numbers corresponding to the output
			if layer == len(config)-1 {
				rl.DrawText(strconv.Itoa(neuron), x+40, y+(int32(neuron)*offset)-8, 20, rl.Gray)
			}
		}
	}
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

func clearGridButton(bounds rl.Rectangle, text string) bool {
	pos := rl.GetMousePosition()
	state := true

	if rl.CheckCollisionPointRec(pos, bounds) {
		rl.DrawRectangleLines(bounds.ToInt32().X, bounds.ToInt32().Y, bounds.ToInt32().Width, bounds.ToInt32().Height, rl.Red)
		rl.DrawText("text", 255, 655, 18, rl.White)
	} else {
		rl.DrawRectangleLines(bounds.ToInt32().X, bounds.ToInt32().Y, bounds.ToInt32().Width, bounds.ToInt32().Height, rl.Yellow)
		rl.DrawText("text", 255, 655, 18, rl.White)
	}

	return state
}

func main() {
	// Create the neural net
	nn := NeuralNetwork{
		inputNeuronsSize:        size * size,
		outputNeuronsSize:       10,
		hiddenLayers:            2,
		hiddenLayersNeuronsSize: []int{16, 16}, // for each hidden layer
		weights:                 make([]Matrix, 0),
		biases:                  make([]Matrix, 0),
		epochs:                  1000,
		learningRate:            0.1,
	}

	// // Import training set
	// trainInputs, trainLabels := parseTrain("data/train.csv")
	// // Start the training
	// fmt.Println("Starting the learning process")
	// nnLearn(&nn, &trainInputs, &trainLabels)

	// Initialize window
	rl.InitWindow(screenW, screenH, "Neural Network Visualization")
	defer rl.CloseWindow()

	froboto := rl.LoadFontEx("static/Roboto-Black.ttf", 24, nil)
	defer rl.UnloadFont(froboto)

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
			for i := 0; i < gridInfo.size; i++ {
				for j := 0; j < gridInfo.size; j++ {
					cx := int32(gridInfo.x + 15 + (i * 15))
					cy := int32(gridInfo.y + 15 + (j * 15))
					if cx-int32(gridInfo.r)-4 < int32(pos.X) && int32(pos.X) <= cx+int32(gridInfo.r)+4 &&
						cy-int32(gridInfo.r)-4 < int32(pos.Y) && int32(pos.Y) <= cy+int32(gridInfo.r)+4 {
						gridArray[IX(i, j)] = 255
					}
				}
			}
		}

		drawNeuralNetwork(400, 300, 1200, 500, 10, []int{700, 16, 16, 10}, &nn)

		rl.DrawText("Drag to write a number", 190, 170, 20, rl.White)
		makeGrid(gridInfo, gridArray)
		// Grid clear button
		// TODO make the values into variables

		if rl.IsMouseButtonDown(0) && pos.X > 225 && pos.X < 375 && pos.Y > 650 && pos.Y < 680 {
			rl.DrawRectangleLines(225, 650, 150, 30, rl.White)
			rl.DrawText("Clear Grid", 255, 655, 18, rl.White)
			gridArray = make([]uint8, size*size)
		} else {
			rl.DrawRectangle(240, 650, 150, 30, rl.Gray)
			rl.DrawText("Clear Grid", 270, 655, 18, rl.Black)

		}
		rl.DrawText("28 x 28", 460, 635, 16, rl.Gray)

		// rl.DrawTextEx(froboto, "Basic Neural Network", rl.Vector2{X: screenW/2 - 200, Y: 20}, 48, 1, rl.White)
		rl.DrawText("Deep Neural Network", screenW/2-200, 20, 36, rl.White)
		a := clearGridButton(rl.Rectangle{X: 100, Y: 100, Width: 100, Height: 100}, "Button")
		fmt.Println("clicked:", a)
		// Show FPS
		fps := strconv.FormatInt(int64(rl.GetFPS()), 10)
		t := fmt.Sprintf("FPS: %s", fps)
		rl.DrawText(t, 20, 20, 20, rl.White)

		rl.EndDrawing()

	}
}
