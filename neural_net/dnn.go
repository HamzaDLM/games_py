// FIXME might be conflicting dot product with multiplication of matrices.
package main

import (
	"fmt"
	"math"
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
	}
	for i := 0; i < len(nn.biases); i++ {
		for j := 0; j < len(nn.biases[i].data); j++ {
			nn.biases[i].data[j] = rand.Float64()
		}
	}

	// TODO remove after / Verifying dimensions
	for i := 0; i < 3; i++ {
		fmt.Println("Weights", i, nn.weights[i].dims())
		fmt.Println("Biases", i, nn.biases[i].dims())
		fmt.Println("-------------------")
	}

	// GD technique
	for i := 0; i < nn.epochs; i++ {
		// Feedforward
		// TODO make it dynamic (accepting any range of hidden layers)
		fmt.Println(Yellow, "FeedForward", Reset)
		fmt.Println("Input")
		fmt.Println(x.data[0:10])

		fmt.Println("Layer 1")
		// Z1
		layerL1 := createMatrix(nn.weights[0].rowSize, x.colSize)
		layerL1.matrixDot(&nn.weights[0], x)
		layerL1.matrixAddArray(&layerL1, &nn.biases[0])
		// A1
		activationL1 := applyToMatrix(sigmoid, layerL1)
		fmt.Println(activationL1.data[0:10])

		fmt.Println("Layer 2")
		// Z2
		layerL2 := createMatrix(nn.weights[1].rowSize, x.colSize)
		layerL2.matrixDot(&nn.weights[1], &layerL1)
		layerL2.matrixAddArray(&layerL2, &nn.biases[1])
		// A2
		activationL2 := applyToMatrix(sigmoid, layerL2)
		fmt.Println(activationL2.data[0:10])

		fmt.Println("Output")
		// Z3
		layerL3 := createMatrix(nn.weights[2].rowSize, x.colSize)
		layerL3.matrixDot(&nn.weights[2], &layerL2)
		layerL3.matrixAddArray(&layerL3, &nn.biases[2])
		// A3
		activationL3 := applyToMatrix(sigmoid, layerL3)
		fmt.Println(activationL3.data[0:10])

		// Predict
		fmt.Println(Yellow, "PREDICT", Reset)
		y_hat := getPredict(&activationL3)
		fmt.Println("correct", y.data[0:10])
		fmt.Println("predicted", y_hat.data[0:10])

		// Accuracy
		fmt.Println(Yellow, "ACCURACY", Reset)
		accuracy := getAccuracy(*y, y_hat)
		fmt.Println(accuracy)

		// Loss
		// TODO compute loss
		fmt.Println(Yellow, "LOSS / Y vs onehot(Y)", Reset)
		one_hot_y := one_hot(transpose(y))
		fmt.Println(y.dims(), one_hot_y.dims())
		// loss := crossEntropy(one_hot(y), &activationL3)
		// fmt.Println("Loss:", loss)

		// Backpropagate
		fmt.Println(Yellow, "BACKPROPAGATE", Reset)

		var m float64 = float64(1.0 / float64(one_hot_y.colSize))

		fmt.Println(Green, "Output", Reset)
		derivativeL3 := createMatrix(activationL3.rowSize, activationL3.colSize)
		derivativeL3.matrixSub(activationL3, one_hot_y)

		activationL2Transposed := transpose(&activationL2)
		gradW3 := createMatrix(derivativeL3.rowSize, activationL2Transposed.colSize)
		gradW3.matrixDot(&derivativeL3, &activationL2Transposed)
		gradW3 = matrixMultScalar(&gradW3, m)

		gradB3 := m * matrixSum(&derivativeL3)
		fmt.Println(gradB3)

		fmt.Println(Green, "Layer 2", Reset)
		derivativeL2 := createMatrix(activationL2.rowSize, activationL2.colSize)
		weights2Transpose := transpose(&nn.weights[2])
		derivativeL2.matrixDot(&weights2Transpose, &derivativeL3)
		derivativeL2 = applyToMatrix(sigmoidPrime, activationL2)

		activationL1Transposed := transpose(&activationL1)
		gradW2 := createMatrix(derivativeL2.rowSize, activationL1Transposed.colSize)
		gradW2.matrixDot(&derivativeL2, &activationL1Transposed)
		gradW2 = matrixMultScalar(&gradW2, m)

		gradB2 := m * matrixSum(&derivativeL2)
		fmt.Println(gradB2)

		fmt.Println(Green, "Layer 1", Reset)
		derivativeL1 := createMatrix(activationL1.rowSize, activationL1.colSize)
		weights1Transpose := transpose(&nn.weights[1])
		derivativeL1.matrixDot(&weights1Transpose, &derivativeL2)
		derivativeL1 = applyToMatrix(sigmoidPrime, activationL1)

		activationL0Transposed := transpose(x)
		gradW1 := createMatrix(derivativeL1.rowSize, activationL0Transposed.colSize)
		gradW1.matrixDot(&derivativeL1, &activationL0Transposed)
		gradW1 = matrixMultScalar(&gradW1, m)

		gradB1 := m * matrixSum(&derivativeL1)
		fmt.Println(gradB1)

		fmt.Println(Yellow, "ADJUST PARAMETERS", Reset)
		nn.weights[0].matrixSub(nn.weights[0], matrixMultScalar(&gradW1, nn.learningRate))
		nn.biases[0] = matrixSubScalar(&nn.biases[0], gradB1*nn.learningRate)

		nn.weights[1].matrixSub(nn.weights[1], matrixMultScalar(&gradW2, nn.learningRate))
		nn.biases[1] = matrixSubScalar(&nn.biases[1], gradB2*nn.learningRate)

		nn.weights[2].matrixSub(nn.weights[2], matrixMultScalar(&gradW3, nn.learningRate))
		nn.biases[2] = matrixSubScalar(&nn.biases[2], gradB3*nn.learningRate)

		fmt.Println(Green, "Accuracy at Iteration:", i, "is", accuracy, Reset)
	}
}

func backPropagate(activations Matrix) {}

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
func getPredict(m *Matrix) Matrix {
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

func feedForward() {}

func stochasticGradientDescent() {}

func updateMiniBatch() {}

func evaluate() {}

func costDerivative() {}

// ======= Helper functions

type Matrix struct {
	data    []float64
	rowSize int // How many rows
	colSize int // How many cols
}

// Initialize a matrix and return it
func createMatrix(r int, c int) Matrix {
	return Matrix{
		data:    make([]float64, r*c),
		rowSize: r,
		colSize: c,
	}

}

// Get 1 dimensional index for a 2 dimensional array
func IX(row, col int) int {
	return row*size + col
}

// Returns Matrix dimensions in better formatted way
func (m Matrix) dims() string {
	return fmt.Sprintf("Matrix of size: %d X %d", m.rowSize, m.colSize)
}

// Prints a matrix in a 2D shape
func printMatrix(m Matrix) {
	limit := 10
	fmt.Print(Red, "Matrix ------ rows ", m.rowSize, "x", m.colSize, " cols\n", Reset)
	for r := 0; r < int(m.rowSize); r++ {
		for c := 0; c < int(m.colSize); c++ {
			if r < limit && c < limit {
				fmt.Print(m.data[r*int(m.colSize)+c], " ")
			}
		}
		if r < limit {
			fmt.Print("\n")
		}
	}
	fmt.Println(Red, "---------------------------", Reset)

}

// Apply a function to matrix and return a copy of it
func applyToMatrix(f func(float64) float64, m Matrix) Matrix {
	r := createMatrix(m.rowSize, m.colSize)
	for i := 0; i < len(r.data); i++ {
		r.data[i] = f(m.data[i])
	}
	return r
}

// Sum elements of a matrix (np.Sum)
func matrixSum(m *Matrix) float64 {
	var sum float64 = 0
	for i := 0; i < len(m.data); i++ {
		sum += m.data[i]
	}
	return sum
}

// Perform dot product on two matrices of 1D form (np.Dot).
// Panics if n of cols in A doesn't equal n of rows in B.
func (R *Matrix) matrixDot(A, B *Matrix) {
	if A.colSize != B.rowSize {
		panic("The matrices aren't multipliable.")
	}

	for i := 0; i < A.rowSize; i++ {
		for j := 0; j < B.colSize; j++ {
			var sum float64 = 0
			for k := 0; k < A.colSize; k++ {
				sum += A.data[i*A.colSize+k] * B.data[k*B.colSize+j]
			}
			index := i*B.colSize + j
			R.data[index] = sum
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

// Perform substraction on matrix by a scalar
func matrixSubScalar(m *Matrix, s float64) Matrix {
	R := createMatrix(m.rowSize, m.colSize)
	for i := 0; i < R.rowSize; i++ {
		for j := 0; j < R.colSize; j++ {
			R.data[i*R.colSize+j] = R.data[i*R.colSize+j] - s
		}
	}
	return R
}

// Perform multiplication on matrix by a scalar
func matrixMultScalar(m *Matrix, s float64) Matrix {
	R := createMatrix(m.rowSize, m.colSize)
	for i := 0; i < R.rowSize; i++ {
		for j := 0; j < R.colSize; j++ {
			R.data[i*R.colSize+j] = R.data[i*R.colSize+j] * s
		}
	}
	return R
}

// Perform Substraction on two matrices of 1D form.
// Panics if dimensions are incorrect (dim A != dim B)
func (R *Matrix) matrixSub(A, B Matrix) {
	if A.colSize != B.colSize || A.rowSize != B.rowSize {
		panic("The matrices can't be substracted.")
	}

	for i := 0; i < A.rowSize; i++ {
		for j := 0; j < A.colSize; j++ {
			R.data[i*A.colSize+j] = A.data[i*A.colSize+j] - B.data[i*A.colSize+j]
		}
	}
}

func (R *Matrix) matrixAddArray(A, B *Matrix) {
	if A.rowSize != B.rowSize {
		panic("Can't add the array to this matrix.")
	}

	for i := 0; i < A.rowSize; i++ {
		for j := 0; j < A.rowSize; j++ {
			R.data[i*A.colSize+j] = A.data[i*A.colSize+j] + B.data[i]
		}
	}
}

// Sigmoid activation function
func sigmoid(z float64) float64 {
	return 1.0 / (1.0 + math.Exp(-z))
}

// Derivative of sigmoid activation function
func sigmoidPrime(z float64) float64 {
	return sigmoid(z) * (1.0 - sigmoid(z))
}
