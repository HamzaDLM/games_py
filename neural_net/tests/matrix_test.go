package main

import (
	"testing"
)

func TestMatrixSum1Axis(t *testing.T) {
	m := createMatrix(2, 3) // 2 rows x 3 cols
	m.data = []float64{1, 2, 3, 10, 20, 30}

	var r Matrix = matrixSum1Axis(&m)
	expectedResult := []float64{6, 60}

	if r.data[0] != expectedResult[0] || r.data[1] != expectedResult[1] {
		t.Errorf("Expected result not found!")
	}
}
