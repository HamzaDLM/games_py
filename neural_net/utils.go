package main

import "math"

// Relu function
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

func softmax(m *Matrix) Matrix {
	r := createMatrix(m.rowSize, m.colSize)
	mExp := applyToMatrix(math.Exp, *m)
	max := matrixSum(&mExp)
	for i := 0; i < len(m.data); i++ {
		r.data[i] = math.Exp(m.data[i]) / max
	}
	return r
}

func softmaxPrime(m *Matrix) Matrix {
	s := softmax(m)
	p := createMatrix(m.rowSize, m.colSize)

	for i := 0; i < len(p.data); i++ {
		p.data[i] = s.data[i] * (1 - s.data[i])
	}

	return p
}

// Sigmoid activation function
func sigmoid(z float64) float64 {
	return 1.0 / (1.0 + math.Exp(-z))
}

// Derivative of sigmoid activation function
func sigmoidPrime(z float64) float64 {
	return sigmoid(z) * (1.0 - sigmoid(z))
}
