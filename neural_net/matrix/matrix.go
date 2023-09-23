package matrix

import "fmt"

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34"
)

type Matrix struct {
	Data    []float64
	RowSize int // How many rows
	ColSize int // How many cols
}

// Initialize a matrix and return it
func CreateMatrix(r int, c int) Matrix {
	return Matrix{
		Data:    make([]float64, r*c),
		RowSize: r,
		ColSize: c,
	}
}

// Get 1 dimensional index for a 2 dimensional array
func IX(rowN, colN, colSize int) int {
	// row id * ColSize + col id
	return rowN*colSize + colN
}

// Returns Matrix shape
func (m Matrix) Shape() string {
	return fmt.Sprintf("Matrix of size: %d X %d", m.RowSize, m.ColSize)
}

// Prints a matrix in a 2D shape
func PrintMatrix(m *Matrix) {
	limit := 10 //FIXME the fuk is this ?
	fmt.Print(Red, "Matrix ------ rows ", m.RowSize, "x", m.ColSize, " cols\n", Reset)
	for r := 0; r < int(m.RowSize); r++ {
		for c := 0; c < int(m.ColSize); c++ {
			if r < limit && c < limit {
				fmt.Printf("%10.4f", m.Data[r*int(m.ColSize)+c])
			}
		}
		if r < limit {
			fmt.Println()
		}
	}
	fmt.Println(Red, "---------------------------", Reset)

}

// Apply a function to matrix and return a copy of it
func ApplyToMatrix(f func(float64) float64, m Matrix) Matrix {
	r := CreateMatrix(m.RowSize, m.ColSize)
	for i := 0; i < len(r.Data); i++ {
		r.Data[i] = f(m.Data[i])
	}
	return r
}

// Sum elements of a matrix (np.Sum)
func MatrixSum(m *Matrix) float64 {
	var sum float64 = 0
	for i := 0; i < len(m.Data); i++ {
		sum += m.Data[i]
	}
	return sum
}

// Similar to np.sum with axis=1
func MatrixSum1Axis(m *Matrix) Matrix {
	result := CreateMatrix(m.RowSize, 1)
	for r := 0; r < m.RowSize; r++ {
		for c := 0; c < m.ColSize; c++ {
			result.Data[r] += m.Data[r*m.ColSize+c]
		}
	}
	return result
}

// Perform dot product on two matrices of 1D form (np.Dot).
// Panics if n of cols in A doesn't equal n of rows in B.
func (R *Matrix) MatrixDot(A, B *Matrix) {
	if A.ColSize != B.RowSize {
		panic("The matrices aren't multipliable.")
	}

	for i := 0; i < A.RowSize; i++ {
		for j := 0; j < B.ColSize; j++ {
			var sum float64 = 0
			for k := 0; k < A.ColSize; k++ {
				sum += A.Data[i*A.ColSize+k] * B.Data[k*B.ColSize+j]
			}
			index := i*B.ColSize + j
			R.Data[index] = sum
		}
	}
}

// Perform Addition on two matrices of 1D form.
// Panics if dimensions are incorrect (dim A != dim B)
func (R *Matrix) MatrixAdd(A, B *Matrix) {
	if A.ColSize != B.ColSize || A.RowSize != B.RowSize {
		panic("The matrices can't be additioned.")
	}

	for i := 0; i < A.RowSize; i++ {
		for j := 0; j < A.ColSize; j++ {
			R.Data[i*A.ColSize+j] = A.Data[i*A.ColSize+j] + B.Data[i*A.ColSize+j]
		}
	}
}

// Perform Addition on two matrices of 1D form.
// Panics if dimensions are incorrect (dim A != dim B)
func MatrixMult(A, B *Matrix) Matrix {
	if A.ColSize != B.ColSize || A.RowSize != B.RowSize {
		panic("The matrices can't be additioned.")
	}
	R := CreateMatrix(A.RowSize, A.ColSize)

	for i := 0; i < A.RowSize; i++ {
		for j := 0; j < A.ColSize; j++ {
			R.Data[i*A.ColSize+j] = A.Data[i*A.ColSize+j] * B.Data[i*A.ColSize+j]
		}
	}

	return R
}

// Perform substraction on matrix by a scalar
func MatrixSubScalar(m *Matrix, s float64) Matrix {
	R := CreateMatrix(m.RowSize, m.ColSize)
	for i := 0; i < R.RowSize; i++ {
		for j := 0; j < R.ColSize; j++ {
			R.Data[i*R.ColSize+j] = R.Data[i*R.ColSize+j] - s
		}
	}
	return R
}

// Perform substraction on matrix by a scalar
func ScalarSubMatrix(m *Matrix, s float64) Matrix {
	R := CreateMatrix(m.RowSize, m.ColSize)
	for i := 0; i < R.RowSize; i++ {
		for j := 0; j < R.ColSize; j++ {
			R.Data[i*R.ColSize+j] = s - R.Data[i*R.ColSize+j]
		}
	}
	return R
}

// Perform multiplication on matrix by a scalar
func MatrixMultScalar(m *Matrix, s float64) Matrix {
	resultMatrix := CreateMatrix(m.RowSize, m.ColSize)
	for i := 0; i < resultMatrix.RowSize; i++ {
		for j := 0; j < resultMatrix.ColSize; j++ {
			resultMatrix.Data[IX(i, j, m.ColSize)] = m.Data[IX(i, j, m.ColSize)] * s
		}
	}
	return resultMatrix
}

func (m *Matrix) MatrixMultScalar2(s float64) {
	for i := 0; i < len(m.Data); i++ {
		m.Data[i] *= s
	}
}

// Perform Substraction on two matrices of 1D form.
// Panics if dimensions are incorrect (dim A != dim B)
func (R *Matrix) MatrixSub(A, B Matrix) {
	if A.ColSize != B.ColSize || A.RowSize != B.RowSize {
		panic("The matrices can't be substracted.")
	}

	for i := 0; i < A.RowSize; i++ {
		for j := 0; j < A.ColSize; j++ {
			R.Data[i*A.ColSize+j] = A.Data[i*A.ColSize+j] - B.Data[i*A.ColSize+j]
		}
	}
}

func (R *Matrix) MatrixAddArray(A, B *Matrix) {
	if A.RowSize != B.RowSize {
		panic("Can't add the array to this matrix.")
	}

	for i := 0; i < A.RowSize; i++ {
		for j := 0; j < A.RowSize; j++ {
			R.Data[i*A.ColSize+j] = A.Data[i*A.ColSize+j] + B.Data[i]
		}
	}
}

// Compares two matricies for if they similar
func CompareMatricies(A, B *Matrix) bool {
	if A.RowSize != B.RowSize || A.ColSize != B.ColSize {
		return false
	}
	for i := 0; i < A.ColSize * A.RowSize; i++ {
		if A.Data[i] != B.Data[i] {
			return false
		}
	}

	return true
}

// Transpose a matrix of n x m dimensions to m x n
func Transpose(m *Matrix) Matrix {
	transposed := CreateMatrix(m.ColSize, m.RowSize)
	transposed.Data = []float64{}	
	for c := 0; c < m.ColSize; c++ {
		for r := 0; r < m.RowSize; r++ {
			transposed.Data = append(transposed.Data, m.Data[IX(r, c, m.ColSize)])
		}
	}
	return transposed
}

// Convert Matrix of shape n x 1 to one-hot-binary
func OneHot(m Matrix) Matrix {
	// Get the max value of the matrix to know how many columns we should have
	maxVal := m.Data[0]
	for i := 0; i < m.RowSize; i++ {
		if m.Data[i] > maxVal {
			maxVal = m.Data[i]
		}
	}
	colValue := m.ColSize*(int(maxVal) + 1)
	
	oneHotMatrix := CreateMatrix(m.RowSize, colValue) // + 1 because 0 is represented by 0..01
	
	for i := 0; i < m.RowSize; i++ {
		for j := 0; j < colValue; j++ {
			if int(m.Data[i]) == j {
				oneHotMatrix.Data[IX(i, j, colValue)] = 1
			}
		}
	}
	
	return oneHotMatrix
}
