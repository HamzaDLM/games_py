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
func nnLearn(nn *NeuralNetwork, x, y []float64) {

	// Check parameters
	if len(nn.hiddenLayersNeuronsSize) != nn.hiddenLayers {
		panic(fmt.Sprintf("The hidden layers neuron size array should contain %v elements.", nn.hiddenLayers))
	}

	// Create a random generator
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Initialize biases and weights arrays
	weights := make([]Matrix, nn.hiddenLayers+1)
	biases := make([]Matrix, nn.hiddenLayers+1)
	// Prealocate w, b space for l1
	weights[0] = Matrix{
		data:    make([]float64, nn.inputNeuronsSize*nn.hiddenLayersNeuronsSize[0]),
		rowSize: nn.hiddenLayersNeuronsSize[0],
		colSize: nn.inputNeuronsSize,
	}
	biases[0] = Matrix{
		data:    make([]float64, nn.hiddenLayersNeuronsSize[0]),
		rowSize: nn.hiddenLayersNeuronsSize[0],
		colSize: 1,
	}
	// Prealocate w, b space for output
	weights[nn.hiddenLayers] = Matrix{
		data:    make([]float64, nn.hiddenLayersNeuronsSize[len(nn.hiddenLayersNeuronsSize)-1]*nn.outputNeuronsSize),
		rowSize: nn.outputNeuronsSize,
		colSize: nn.hiddenLayersNeuronsSize[len(nn.hiddenLayersNeuronsSize)-1],
	}
	biases[nn.hiddenLayers] = Matrix{
		data:    make([]float64, nn.outputNeuronsSize),
		rowSize: nn.outputNeuronsSize,
		colSize: 1,
	}
	// Preaclocate for every layer in between
	if nn.hiddenLayers > 1 {
		for i := 1; i < nn.hiddenLayers; i++ {
			weights[i] = Matrix{
				data:    make([]float64, nn.hiddenLayersNeuronsSize[i-1]*nn.hiddenLayersNeuronsSize[i]),
				rowSize: nn.hiddenLayersNeuronsSize[i],
				colSize: nn.hiddenLayersNeuronsSize[i-1],
			}
			biases[i] = Matrix{
				data:    make([]float64, nn.hiddenLayersNeuronsSize[i]),
				rowSize: nn.hiddenLayersNeuronsSize[i],
				colSize: 1,
			}
		}
	}
	// Fill weights and biases with random floats
	for i := 0; i < len(weights); i++ {
		for j := 0; j < len(weights[i].data); j++ {
			weights[i].data[j] = random.Float64()
		}
	}
	for i := 0; i < len(biases); i++ {
		for j := 0; j < len(biases[i].data); j++ {
			biases[i].data[j] = rand.Float64()
		}
	}

	// TODO: remove after / Verifying dimensions
	for i := 0; i < 3; i++ {
		fmt.Println("Weights", i, weights[i].dims())
		fmt.Println("Biases", i, biases[i].dims())
		fmt.Println("-------------------")
	}

	// // Backpropagation technique
	// for i := 0; i < nn.epochs; i++ {
	// 	// feedforward

	// 	// activationL1 := Matrix{
	// 	// 	data: make([]float64, )
	// 	// }

	// 	// backpropagate

	// }
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
	rowSize int // How many rows
	colSize int // How many cols
}

// Returns Matrix dimensions in better formatted way
func (m Matrix) dims() string {
	return fmt.Sprintf("Matrix of size: %d X %d", m.rowSize, m.colSize)
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

// Perform multiplication on two matrices of 1D form.
// Panics if n of cols in A doesn't equal n of rows in B.
func (R *Matrix) matrixMult(A, B *Matrix) {
	if A.colSize != B.rowSize {
		panic("The matrices aren't multipliable.")
	}

	for i := 0; i < A.rowSize; i++ {
		for j := 0; j < B.colSize; j++ {
			var sum float64 = 0
			for k := 0; k < A.colSize; k++ {
				sum += A.data[i*A.colSize+k] * B.data[k*B.colSize+j]
			}
			R.data[i*B.colSize+j] = sum
		}
	}
}

// Perform Addition on two matrices of 1D form.
// Panics if dimensions are incorrect (dim A != dim B)
func (R *Matrix) matrixAdd(A, B *Matrix) {
	if A.colSize != B.colSize || A.rowSize != B.rowSize {
		panic("The matrices can't be additioned.")
	}

	for i := 0; i < A.rowSize; i++ {
		for j := 0; j < A.colSize; j++ {
			R.data[i*A.colSize+j] = A.data[i*A.colSize+j] + B.data[i*A.colSize+j]
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
