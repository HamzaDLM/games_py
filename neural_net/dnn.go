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
	weights                 []uint8
	biases                  []uint8
	epochs                  int
	learningRate            float64
}

// Main training function
func nnLearn(nn *NeuralNetwork) {

	// Check parameters
	if len(nn.hiddenLayersNeuronsSize) != nn.hiddenLayers {
		panic(fmt.Sprintf("The hidden layers neuron size array should contain %v elements.", nn.hiddenLayers))
	}

	// Initialize biases and weights
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	biases := make([]uint8, size)
	weights := make([]uint8, size)
	for i := 0; i < int(size); i++ {
		biases[i] = uint8(random.Float64())
		weights[i] = uint8(random.Float64())
	}

	for i := 0; i < nn.epochs; i++ {

	}
}

func fastForward() {}

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

	// result := Matrix{
	// 	data:    make([]float64, int(A.rowSize*B.colSize)),
	// 	rowSize: A.rowSize,
	// 	colSize: B.colSize,
	// }

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

	// result := Matrix{
	// 	data:    make([]float64, int(A.colSize*A.colSize)),
	// 	rowSize: A.rowSize,
	// 	colSize: A.colSize,
	// }

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
