package main

import "fmt"

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
	// row id * colsize + col id
	return row*size + col
}

// Returns Matrix shape
func (m Matrix) shape() string {
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

// Similar to np.sum with axis=1
func matrixSum1Axis(m *Matrix) Matrix {
	result := createMatrix(m.rowSize, 1)
	for r := 0; r < m.rowSize; r++ {
		for c := 0; c < m.colSize; c++ {
			result.data[r] += m.data[r*m.colSize+c]
		}
	}
	return result
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

// Perform Addition on two matrices of 1D form.
// Panics if dimensions are incorrect (dim A != dim B)
func matrixMult(A, B *Matrix) Matrix {
	if A.colSize != B.colSize || A.rowSize != B.rowSize {
		panic("The matrices can't be additioned.")
	}
	R := createMatrix(A.rowSize, A.colSize)

	for i := 0; i < A.rowSize; i++ {
		for j := 0; j < A.colSize; j++ {
			R.data[i*A.colSize+j] = A.data[i*A.colSize+j] * B.data[i*A.colSize+j]
		}
	}

	return R
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

// Perform substraction on matrix by a scalar
func ScalarSubMatrix(m *Matrix, s float64) Matrix {
	R := createMatrix(m.rowSize, m.colSize)
	for i := 0; i < R.rowSize; i++ {
		for j := 0; j < R.colSize; j++ {
			R.data[i*R.colSize+j] = s - R.data[i*R.colSize+j]
		}
	}
	return R
}

// Perform multiplication on matrix by a scalar
func matrixMultScalar(m *Matrix, s float64) Matrix {
	R := createMatrix(m.rowSize, m.colSize)
	for i := 0; i < R.rowSize; i++ {
		for j := 0; j < R.colSize; j++ {
			R.data[i*R.colSize+j] *= s
		}
	}
	return R
}

func (m *Matrix) matrixMultScalar2(s float64) {
	for i := 0; i < len(m.data); i++ {
		m.data[i] *= s
	}
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
