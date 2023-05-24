package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type NeuralNetwork struct {
	inputNeuronsSize        int
	outputNeuronsSize       int
	hiddenLayers            int
	hiddenLayersNeuronsSize []int
	weights                 [][]float64
	biases                  [][]float64
	epochs                  int
	learningRate            float64
}

// Main training function
func nnLearn(nn *NeuralNetwork) {

	// Check parameters
	if len(nn.hiddenLayersNeuronsSize) != nn.hiddenLayers {
		panic(fmt.Sprintf("The hidden layers neuron size array should contain %v elements.", nn.hiddenLayers))
	}

	// Create a random generator
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Initialize biases and weights arrays
	weights := make([][]float64, nn.hiddenLayers+1)
	biases := make([][]float64, nn.hiddenLayers+1)
	// Prealocate w, b space for l1
	weights[0] = make([]float64, nn.inputNeuronsSize*nn.hiddenLayersNeuronsSize[0])
	biases[0] = make([]float64, nn.hiddenLayersNeuronsSize[0])
	// Prealocate w, b space for output
	weights[nn.hiddenLayers] = make([]float64, nn.hiddenLayersNeuronsSize[len(nn.hiddenLayersNeuronsSize)-1]*nn.outputNeuronsSize)
	biases[nn.hiddenLayers] = make([]float64, nn.outputNeuronsSize)
	// Preaclocate for every layer in between
	if nn.hiddenLayers > 1 {
		for i := 1; i < nn.hiddenLayers; i++ {
			weights[i] = make([]float64, nn.hiddenLayersNeuronsSize[i-1]*nn.hiddenLayersNeuronsSize[i])
			biases[i] = make([]float64, nn.hiddenLayersNeuronsSize[i])
		}
	}

	for i := 0; i < len(weights); i++ {
		for j := 0; j < len(weights[i]); j++ {
			weights[i][j] = random.Float64()
		}
	}
	for i := 0; i < len(biases); i++ {
		for j := 0; j < len(biases[i]); j++ {
			biases[i][j] = rand.Float64()
		}
	}

	// Backpropagation technique
	for i := 0; i < nn.epochs; i++ {
		
	}
}

func feedForward() {}

func stochasticGradientDescent() {}

func updateMiniBatch() {}

func backPropagate() {}

func evaluate() {}

func costDerivative() {}

// ======= Helper functions

type Matrix struct {
	data    []float64
	rowSize float64
	colSize float64
}

// Prints a matrix in a 2D shape
func printMatrix(m Matrix) {
	fmt.Print("Matrix ------ rows ", m.rowSize, "x", m.colSize, " cols\n")
	for i := 0; i < int(m.rowSize); i++ {
		for j := 0; j < int(m.colSize); j++ {
			fmt.Print(m.data[i*int(m.colSize)+j], " ")
		}
		fmt.Print("\n")
	}
	fmt.Println("---------------------------")

}

// Perform multiplication on two matrices.
// Panics if n of cols in A doesn't equal n of rows in B.
func (R *Matrix) matrixMult(A, B *Matrix) {
	if A.colSize != B.rowSize {
		panic("The matrices aren't multipliable.")
	}

	for i := 0; i < int(A.rowSize); i++ {
		for j := 0; j < int(B.colSize); j++ {
			var sum float64 = 0
			for k := 0; k < int(A.colSize); k++ {
				sum += A.data[i*int(A.colSize)+k] * B.data[k*int(B.colSize)+j]
			}
			R.data[i*int(B.colSize)+j] = sum
		}
	}
}

// Perform Addition on two matrices.
// Panics if dimensions are incorrect (dim A != dim B)
func (R *Matrix) matrixAdd(A, B *Matrix) {
	if A.colSize != B.colSize || A.rowSize != B.rowSize {
		panic("The matrices can't be additioned.")
	}

	for i := 0; i < int(A.rowSize); i++ {
		for j := 0; j < int(A.colSize); j++ {
			R.data[i*int(A.colSize)+j] = A.data[i*int(A.colSize)+j] + B.data[i*int(A.colSize)+j]
		}
	}
}

// Activation function
func sigmoid(z float64) float64 {
	return 1.0 / (1.0 + math.Exp(-z))
}

// Derivative of sigmoid function
func sigmoidPrime(z float64) float64 {
	return sigmoid(z) * (1.0 - sigmoid(z))
}
