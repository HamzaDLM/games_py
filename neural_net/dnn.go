// FIXME might be conflicting dot product with multiplication of matrices.
// FIXME change float64 to float32 to reduce precision overhead
package main

import (
	"fmt"
	"math/rand"
	"time"
)

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"

type NeuralNetwork struct {
	inputNeuronsSize        int
	outputNeuronsSize       int
	hiddenLayers            int
	hiddenLayersNeuronsSize []int
	weights                 []Matrix
	biases                  []Matrix
	epochs                  int
	learningRate            float64
}

// Main training function
func nnLearn(nn *NeuralNetwork, x, y *Matrix) {

	// Check parameters
	if len(nn.hiddenLayersNeuronsSize) != nn.hiddenLayers {
		panic(fmt.Sprintf("The hidden layers neuron size array should contain %v elements.", nn.hiddenLayers))
	}

	// Create a random generator
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Initialize biases and weights arrays
	nn.weights = make([]Matrix, nn.hiddenLayers+1)
	nn.biases = make([]Matrix, nn.hiddenLayers+1)
	// Prealocate w, b space for l1
	nn.weights[0] = createMatrix(nn.hiddenLayersNeuronsSize[0], nn.inputNeuronsSize)
	nn.biases[0] = createMatrix(nn.hiddenLayersNeuronsSize[0], 1)
	// Prealocate w, b space for output
	nn.weights[nn.hiddenLayers] = createMatrix(nn.outputNeuronsSize, nn.hiddenLayersNeuronsSize[len(nn.hiddenLayersNeuronsSize)-1])
	nn.biases[nn.hiddenLayers] = createMatrix(nn.outputNeuronsSize, 1)

	// Preaclocate for every layer in between
	if nn.hiddenLayers > 1 {
		for i := 1; i < nn.hiddenLayers; i++ {
			nn.weights[i] = createMatrix(nn.hiddenLayersNeuronsSize[i], nn.hiddenLayersNeuronsSize[i-1])
			nn.biases[i] = createMatrix(nn.hiddenLayersNeuronsSize[i], 1)
		}
	}

	// Fill weights and biases with random floats
	for i := 0; i < len(nn.weights); i++ {
		for j := 0; j < len(nn.weights[i].data); j++ {
			nn.weights[i].data[j] = random.Float64()
		}
		for j := 0; j < len(nn.biases[i].data); j++ {
			nn.biases[i].data[j] = rand.Float64()
		}
	}

	// storing accuracy and loss
	var accuracyList []float64
	accuracyList = append(accuracyList, 0) // Initial accuracy

	// GD technique
	for i := 0; i < nn.epochs; i++ {
		// Feedforward
		// TODO make it dynamic (accepting any range of hidden layers)
		z1 := createMatrix(nn.weights[0].rowSize, x.colSize)
		z1.matrixDot(&nn.weights[0], x)
		z1.matrixAddArray(&z1, &nn.biases[0])
		a1 := applyToMatrix(relu, z1)

		z2 := createMatrix(nn.weights[1].rowSize, x.colSize)
		z2.matrixDot(&nn.weights[1], &a1)
		z2.matrixAddArray(&z2, &nn.biases[1])
		a2 := applyToMatrix(relu, z2)

		z3 := createMatrix(nn.weights[2].rowSize, x.colSize)
		z3.matrixDot(&nn.weights[2], &a2)
		z3.matrixAddArray(&z3, &nn.biases[2])
		z3.matrixMultScalar2(0.01) // scale down the values to avoid overflow when using exponent in softmax
		a3 := softmax(&z3)

		// Predict
		y_hat := getPrediction(&a3)

		// Accuracy
		accuracy := getAccuracy(*y, y_hat)
		accuracyList = append(accuracyList, accuracy)

		// Loss
		// TODO compute loss
		one_hot_y := one_hot(transpose(y))
		// loss := crossEntropy(one_hot(y), &a3)

		// Backpropagate
		// FIXME size of gradB
		var m float64 = float64(1.0 / float64(one_hot_y.colSize))

		// LAYER OUTPUT
		dZ3 := createMatrix(a3.rowSize, a3.colSize)
		dZ3.matrixSub(a3, one_hot_y)

		printMatrix(dZ3)
		a2Transposed := transpose(&a2)
		gradW3 := createMatrix(dZ3.rowSize, a2Transposed.colSize)
		gradW3.matrixDot(&dZ3, &a2Transposed)
		gradW3 = matrixMultScalar(&gradW3, m)

		gradB3 := m * matrixSum(&dZ3)

		// HIDDEN LAYER 2
		dZ2 := createMatrix(a2.rowSize, a2.colSize)
		w3Transpose := transpose(&nn.weights[1])
		fmt.Println("shapes", w3Transpose.shape(), dZ3.shape())
		dZ2.matrixDot(&w3Transpose, &dZ3)
		dZ2 = applyToMatrix(reluPrime, a2)

		a1Transposed := transpose(&a1)
		gradW2 := createMatrix(dZ2.rowSize, a1Transposed.colSize)
		gradW2.matrixDot(&dZ2, &a1Transposed)
		gradW2 = matrixMultScalar(&gradW2, m)

		gradB2 := matrixSum1Axis(&dZ2)
		gradB2.matrixMultScalar2(m)

		// HIDDEN LAYER 1
		dZ1 := createMatrix(a1.rowSize, a1.colSize)
		w2Transpose := transpose(&nn.weights[0])
		dZ1.matrixDot(&w2Transpose, &dZ2)
		zReluPrime := applyToMatrix(reluPrime, z1)
		dZ1 = matrixMult(&dZ1, &zReluPrime)

		a0Transposed := transpose(x)
		gradW1 := createMatrix(dZ1.rowSize, a0Transposed.colSize)
		gradW1.matrixDot(&dZ1, &a0Transposed)
		gradW1 = matrixMultScalar(&gradW1, m)

		gradB1 := matrixSum1Axis(&dZ1)
		gradB1.matrixMultScalar2(m)

		fmt.Println("Checking gradB1")
		fmt.Println(m)
		printMatrix(dZ1)
		printMatrix(gradB1)

		// fmt.Println("Checking weights w, b")
		// printMatrix(nn.weights[0])
		// printMatrix(nn.biases[0])
		// fmt.Println("Checking activations z1, a1")
		// printMatrix(z1)
		// printMatrix(a1)

		// Updating parameters
		nn.weights[0].matrixSub(nn.weights[0], matrixMultScalar(&gradW1, nn.learningRate))
		nn.biases[0].matrixSub(nn.biases[0], matrixMultScalar(&gradB1, nn.learningRate))

		nn.weights[1].matrixSub(nn.weights[1], matrixMultScalar(&gradW2, nn.learningRate))
		nn.biases[1].matrixSub(nn.biases[1], matrixMultScalar(&gradB2, nn.learningRate))

		nn.weights[2].matrixSub(nn.weights[2], matrixMultScalar(&gradW3, nn.learningRate))
		nn.biases[2] = matrixSubScalar(&nn.biases[2], gradB3*nn.learningRate)

		fmt.Println(Green, "Accuracy at Iteration:", i, "is", accuracy, ", Delta is:", accuracy-accuracyList[i], Reset)
		fmt.Println("Prediction Y hat")
		printMatrix(y_hat)
	}
}

// Transpose a matrix of n x m dimensions to m x n
func transpose(m *Matrix) Matrix {
	t := createMatrix(m.colSize, m.rowSize)
	t.data = m.data
	return t
}

// Convert Matrix of shape n x 1 to one-hot-binary
func one_hot(m Matrix) Matrix {
	one_hot := createMatrix(m.rowSize*10, m.colSize)
	for c := 0; c < 10; c++ {
		for r := 0; r < 10; r++ {
			if int(m.data[c]) == r {
				one_hot.data[r*one_hot.colSize+c] = 1
			}
		}
	}
	return one_hot
}

func crossEntropy(y_one_hot Matrix, y_hat *Matrix) float64 {

	return 0.0
}

func getAccuracy(correct Matrix, predicted Matrix) float64 {
	if len(correct.data) != len(predicted.data) {
		panic("Can't get accuracy, sizes are different.")
	}
	var count float64 = 0
	for i := 0; i < len(correct.data); i++ {
		if correct.data[i] == predicted.data[i] {
			count += 1
		}
	}
	return count / float64(len(correct.data))
}

// Predict result with size 1 x m
func getPrediction(m *Matrix) Matrix {
	r := createMatrix(1, m.colSize)
	for c := 0; c < m.colSize; c++ {
		max := 0 // max is the first row value in col i
		for r := 0; r < m.rowSize; r++ {
			if m.data[IX(r, c)] > m.data[IX(max, c)] {
				max = r
			}
		}
		r.data[c] = float64(max)
	}
	return r
}
