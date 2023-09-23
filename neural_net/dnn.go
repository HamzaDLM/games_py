// FIXME might be conflicting dot product with multiplication of matrices.
// FIXME change float64 to float32 to reduce precision overhead
package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/HamzaDLM/simulations_and_games/matrix"
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
	weights                 []matrix.Matrix
	biases                  []matrix.Matrix
	epochs                  int
	learningRate            float64
}

// Main training function
func nnLearn(nn *NeuralNetwork, x, y *matrix.Matrix) {

	// Check parameters
	if len(nn.hiddenLayersNeuronsSize) != nn.hiddenLayers {
		panic(fmt.Sprintf("The hidden layers neuron size array should contain %v elements.", nn.hiddenLayers))
	}

	// Create a random generator
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Initialize biases and weights arrays
	nn.weights = make([]matrix.Matrix, nn.hiddenLayers+1)
	nn.biases = make([]matrix.Matrix, nn.hiddenLayers+1)
	// Prealocate w, b space for l1
	nn.weights[0] = matrix.CreateMatrix(nn.hiddenLayersNeuronsSize[0], nn.inputNeuronsSize)
	nn.biases[0] = matrix.CreateMatrix(nn.hiddenLayersNeuronsSize[0], 1)
	// Prealocate w, b space for output
	nn.weights[nn.hiddenLayers] = matrix.CreateMatrix(nn.outputNeuronsSize, nn.hiddenLayersNeuronsSize[len(nn.hiddenLayersNeuronsSize)-1])
	nn.biases[nn.hiddenLayers] = matrix.CreateMatrix(nn.outputNeuronsSize, 1)

	// Prealocate for every layer in between
	if nn.hiddenLayers > 1 {
		for i := 1; i < nn.hiddenLayers; i++ {
			nn.weights[i] = matrix.CreateMatrix(nn.hiddenLayersNeuronsSize[i], nn.hiddenLayersNeuronsSize[i-1])
			nn.biases[i] = matrix.CreateMatrix(nn.hiddenLayersNeuronsSize[i], 1)
		}
	}
	
	// Fill weights and biases with random floats
	for i := 0; i < len(nn.weights); i++ {
		for j := 0; j < len(nn.weights[i].Data); j++ {
			nn.weights[i].Data[j] = random.Float64()
		}
		for j := 0; j < len(nn.biases[i].Data); j++ {
			nn.biases[i].Data[j] = rand.Float64()
		}
	}

	// storing accuracy and loss
	var accuracyList []float64
	accuracyList = append(accuracyList, 0) // Initial accuracy

	// GD technique
	for i := 0; i < nn.epochs; i++ {
		// Feedforward
		// TODO make it dynamic (accepting any range of hidden layers)
		z1 := matrix.CreateMatrix(nn.weights[0].RowSize, x.ColSize)
		z1.MatrixDot(&nn.weights[0], x)
		z1.MatrixAddArray(&z1, &nn.biases[0])
		a1 := matrix.ApplyToMatrix(relu, z1)

		z2 := matrix.CreateMatrix(nn.weights[1].RowSize, x.ColSize)
		z2.MatrixDot(&nn.weights[1], &a1)
		z2.MatrixAddArray(&z2, &nn.biases[1])
		a2 := matrix.ApplyToMatrix(relu, z2)

		z3 := matrix.CreateMatrix(nn.weights[2].RowSize, x.ColSize)
		z3.MatrixDot(&nn.weights[2], &a2)
		z3.MatrixAddArray(&z3, &nn.biases[2])
		z3.MatrixMultScalar2(0.01) // scale down the values to avoid overflow when using exponent in softmax
		a3 := softmax(&z3)

		// Predict
		y_hat := getPrediction(&a3)

		// Accuracy
		accuracy := getAccuracy(*y, y_hat)
		accuracyList = append(accuracyList, accuracy)

		// Loss
		// TODO compute loss
		one_hot_y := matrix.OneHot(matrix.Transpose(y))
		// loss := crossEntropy(one_hot(y), &a3)

		// Backpropagate
		// FIXME size of gradB
		var m float64 = float64(1.0 / float64(one_hot_y.ColSize))

		// LAYER OUTPUT
		dZ3 := matrix.CreateMatrix(a3.RowSize, a3.ColSize)
		dZ3.MatrixSub(a3, one_hot_y)

		matrix.PrintMatrix(&dZ3)
		a2Transposed := matrix.Transpose(&a2)
		gradW3 := matrix.CreateMatrix(dZ3.RowSize, a2Transposed.ColSize)
		gradW3.MatrixDot(&dZ3, &a2Transposed)
		gradW3 = matrix.MatrixMultScalar(&gradW3, m)

		gradB3 := m * matrix.MatrixSum(&dZ3)

		// HIDDEN LAYER 2
		dZ2 := matrix.CreateMatrix(a2.RowSize, a2.ColSize)
		w3Transpose := matrix.Transpose(&nn.weights[1])
		fmt.Println("shapes", w3Transpose.Shape(), dZ3.Shape())
		dZ2.MatrixDot(&w3Transpose, &dZ3)
		dZ2 = matrix.ApplyToMatrix(reluPrime, a2)

		a1Transposed := matrix.Transpose(&a1)
		gradW2 := matrix.CreateMatrix(dZ2.RowSize, a1Transposed.ColSize)
		gradW2.MatrixDot(&dZ2, &a1Transposed)
		gradW2 = matrix.MatrixMultScalar(&gradW2, m)

		gradB2 := matrix.MatrixSum1Axis(&dZ2)
		gradB2.MatrixMultScalar2(m)

		// HIDDEN LAYER 1
		dZ1 := matrix.CreateMatrix(a1.RowSize, a1.ColSize)
		w2Transpose := matrix.Transpose(&nn.weights[0])
		dZ1.MatrixDot(&w2Transpose, &dZ2)
		zReluPrime := matrix.ApplyToMatrix(reluPrime, z1)
		dZ1 = matrix.MatrixMult(&dZ1, &zReluPrime)

		a0Transposed := matrix.Transpose(x)
		gradW1 := matrix.CreateMatrix(dZ1.RowSize, a0Transposed.ColSize)
		gradW1.MatrixDot(&dZ1, &a0Transposed)
		gradW1 = matrix.MatrixMultScalar(&gradW1, m)

		gradB1 := matrix.MatrixSum1Axis(&dZ1)
		gradB1.MatrixMultScalar2(m)

		fmt.Println("Checking gradB1")
		fmt.Println(m)
		matrix.PrintMatrix(&dZ1)
		matrix.PrintMatrix(&gradB1)

		// fmt.Println("Checking weights w, b")
		// printMatrix(nn.weights[0])
		// printMatrix(nn.biases[0])
		// fmt.Println("Checking activations z1, a1")
		// printMatrix(z1)
		// printMatrix(a1)

		// Updating parameters
		nn.weights[0].MatrixSub(nn.weights[0], matrix.MatrixMultScalar(&gradW1, nn.learningRate))
		nn.biases[0].MatrixSub(nn.biases[0], matrix.MatrixMultScalar(&gradB1, nn.learningRate))

		nn.weights[1].MatrixSub(nn.weights[1], matrix.MatrixMultScalar(&gradW2, nn.learningRate))
		nn.biases[1].MatrixSub(nn.biases[1], matrix.MatrixMultScalar(&gradB2, nn.learningRate))

		nn.weights[2].MatrixSub(nn.weights[2], matrix.MatrixMultScalar(&gradW3, nn.learningRate))
		nn.biases[2] = matrix.MatrixSubScalar(&nn.biases[2], gradB3*nn.learningRate)

		fmt.Println(Green, "Accuracy at Iteration:", i, "is", accuracy, ", Delta is:", accuracy-accuracyList[i], Reset)
		fmt.Println("Prediction Y hat")
		matrix.PrintMatrix(&y_hat)
	}
}

func crossEntropy(y_one_hot matrix.Matrix, y_hat *matrix.Matrix) float64 {

	return 0.0
}

func getAccuracy(correct matrix.Matrix, predicted matrix.Matrix) float64 {
	if len(correct.Data) != len(predicted.Data) {
		panic("Can't get accuracy, sizes are different.")
	}
	var count float64 = 0
	for i := 0; i < len(correct.Data); i++ {
		if correct.Data[i] == predicted.Data[i] {
			count += 1
		}
	}
	return count / float64(len(correct.Data))
}

// Predict result with size 1 x m
func getPrediction(m *matrix.Matrix) matrix.Matrix {
	r := matrix.CreateMatrix(1, m.ColSize)
	for c := 0; c < m.ColSize; c++ {
		max := 0 // max is the first row value in col i
		for r := 0; r < m.RowSize; r++ {
			if m.Data[matrix.IX(r, c, size)] > m.Data[matrix.IX(max, c, size)] {
				max = r
			}
		}
		r.Data[c] = float64(max)
	}
	return r
}

func relu(z float64) float64 {
	if z < 0 {
		return 0
	}
	return z
}

func reluPrime(z float64) float64 {
	if z < 0 {
		return 0
	}
	return 1
}

func softmax(m *matrix.Matrix) matrix.Matrix {
	r := matrix.CreateMatrix(m.RowSize, m.ColSize)
	mExp := matrix.ApplyToMatrix(math.Exp, *m)
	max := matrix.MatrixSum(&mExp)
	for i := 0; i < len(m.Data); i++ {
		r.Data[i] = math.Exp(m.Data[i]) / max
	}
	return r
}

func softmaxPrime(m *matrix.Matrix) matrix.Matrix {
	s := softmax(m)
	p := matrix.CreateMatrix(m.RowSize, m.ColSize)

	for i := 0; i < len(p.Data); i++ {
		p.Data[i] = s.Data[i] * (1 - s.Data[i])
	}

	return p
}

func sigmoid(z float64) float64 {
	return 1.0 / (1.0 + math.Exp(-z))
}

// Derivative of sigmoid activation function
func sigmoidPrime(z float64) float64 {
	return sigmoid(z) * (1.0 - sigmoid(z))
}
