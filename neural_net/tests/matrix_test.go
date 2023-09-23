package main

import (
	_ "fmt"
	"github.com/HamzaDLM/simulations_and_games/matrix"
	"testing"
)

/* Helper functions */

func initTestMatrix() matrix.Matrix {

	m := matrix.CreateMatrix(2, 3) // 2 rows x 3 cols

	m.Data = []float64{
		1, 2, 3,
		10, 20, 30,
	}

	return m
}

func multNumberByThree(n float64) float64 { return n * 3 }

/* Test functions */

func TestIX(t *testing.T) {

	m := initTestMatrix()

	// Get IX position of row 1 col 2 of size 10
	position_ix := matrix.IX(1, 2, 3)

	if position_ix != 5 {
		t.Errorf("Expected 5 but found %d", position_ix)
	}

	// Test if it gets the right item from matrix
	a := m.Data[position_ix]

	if a != 30 {
		t.Errorf("Expected 30 but found %f", a)
	}
}

func TestApplyToMatrix(t *testing.T) {

	m := initTestMatrix()

	mSquared := matrix.ApplyToMatrix(multNumberByThree, m)

	if mSquared.Data[2] != 9 {
		t.Errorf("Expected 9 but found %f", mSquared.Data[2])
	}
}

func TestMatrixSum(t *testing.T) {

	m := initTestMatrix()

	if matrix.MatrixSum(&m) != 66 {
		t.Errorf("Expected 66 but found %f", matrix.MatrixSum(&m))
	}
}

func TestMatrixSum1Axis(t *testing.T) {

	m := initTestMatrix()

	resultMatrix := matrix.MatrixSum1Axis(&m)
	expectedResult := []float64{6, 60}

	for i, v := range resultMatrix.Data {
		if v != expectedResult[i] {
			t.Errorf("Expected %f but found %f", expectedResult[i], v)
		}
	}
}

func TestOneHot(t *testing.T) {

	m := matrix.CreateMatrix(3, 1) // 2 rows x 3 cols

	m.Data = []float64{
		1,
		2,
		3,
	}

	// Get a n x 1 matrix by summing to the first axis
	summedMatrix := matrix.MatrixSum1Axis(&m)

	oneHotMatrix := matrix.OneHot(summedMatrix)

	for i, v := range oneHotMatrix.Data {
		if (i == 1 || i == 6 || i == 11) && int(v) != 1 {
			t.Errorf("Expected 1 but found %d", int(v))
		} else if (i != 1 && i != 6 && i != 11) && int(v) != 0 {
			t.Errorf("Expected 0 but found %d", int(v))
		}
	}
}

func TestTranspose(t *testing.T) {

	m := initTestMatrix()
	transposedMatrix := matrix.Transpose(&m)
	
	expectedMatrix := matrix.CreateMatrix(3, 2)
	expectedMatrix.Data = []float64{1, 10, 2, 20, 3, 30}
	
	if !matrix.CompareMatricies(&transposedMatrix, &expectedMatrix) {
		t.Error("Expected matrix is not found")
	}
}

func TestMatrixSub(t *testing.T) {
	
	m1 := initTestMatrix()
	m2 := initTestMatrix()
	
	m3 := matrix.CreateMatrix(m1.RowSize, m1.ColSize)

	m3.MatrixSub(m1, m2)

	for _, v := range m3.Data {
		if int(v) != 0 {
			t.Errorf("Expected 0 but found %d", int(v))
		}
	}
}

func TestMatrixMultScalar(t *testing.T) {
	
	m := initTestMatrix()

	mResult := matrix.MatrixMultScalar(&m, 2)

	expectedMatrix := matrix.CreateMatrix(2, 3)
	expectedMatrix.Data = []float64{2, 4, 6, 20, 40, 60}

	if !matrix.CompareMatricies(&mResult, &expectedMatrix) {
		t.Error("Expected matrix is not found")
	}
}


