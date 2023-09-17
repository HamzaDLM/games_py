package main

import (
	"testing"

	"github.com/HamzaDLM/simulations_and_games/matrix"
)

func TestMatrixSum1Axis(t *testing.T) {
	m := matrix.CreateMatrix(2, 3) // 2 rows x 3 cols
	m.Data = []float64{1, 2, 3, 10, 20, 30}

	matrix.PrintMatrix(m)
	// var r Matrix = MatrixSum1Axis(&m)
	// expectedResult := []float64{6, 60}
	//
	// if r.data[0] != expectedResult[0] || r.data[1] != expectedResult[1] {
	// 	t.Errorf("Expected result not found!")
	// }
}
