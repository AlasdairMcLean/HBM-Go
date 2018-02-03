package hbmutil

//Matrix is an overarching interface for the specific matrices to be derived from allowing overloading non-type specific methods like scalar multiplication
type Matrix interface {
	STimes()
}

//Matrixi is an 8bit unsigned integer struct to streamline matrix operations
type Matrixi struct {
	Rows int
	Cols int
	Data [][]int
}

//Matrixf is a 32bit floating point struct to streamline matrix operations
type Matrixf struct {
	Rows int
	Cols int
	Data [][]float32
}

//Matrixff is a 64bit floating point struct to streamline matrix operations
type Matrixff struct {
	Rows int
	Cols int
	Data [][]float64
}

//
// Matrix preallocation implementations
//

//NewMatrixi constructs a integer containing Matrix structure with specified rows and columns, returning the pointer to the new matrix
func NewMatrixi(rows, cols int) *Matrixi {
	ndata := make([][]int, rows)
	for i := 0; i < rows; i++ {
		ndata[i] = make([]int, cols)
	}
	newmat := Matrixi{rows, cols, ndata}
	return &newmat
}

//NewMatrixf constructs a Matrix structure with specified rows and columns, returning the pointer to the new 32bit floating point matrix
func NewMatrixf(rows, cols int) *Matrixf {
	ndata := make([][]float32, rows)
	for i := 0; i < rows; i++ {
		ndata[i] = make([]float32, cols)
	}
	newmat := Matrixf{rows, cols, ndata}
	return &newmat
}

//NewMatrixff constructs a Matrix structure with specified rows and columns, returning the pointer to the new 64bit floating point matrix
func NewMatrixff(rows, cols int) *Matrixff {
	ndata := make([][]float64, rows)
	for i := 0; i < rows; i++ {
		ndata[i] = make([]float64, cols)
	}
	newmat := Matrixff{rows, cols, ndata}
	return &newmat
}

//
// Implementation of Matlab's zeros() function
//

//MatZerosi replicates matlab's zeros() function. Returns a new integer matrix full of 0's
func MatZerosi(rows, cols int) *Matrixi {
	ndata := make([][]int, rows)
	for i := 0; i < rows; i++ {
		ndata[i] = make([]int, cols)
	}
	for j := 0; j < rows; j++ {
		for i := 0; i < cols; i++ {
			ndata[j][i] = 0
		}
	}
	newmat := Matrixi{rows, cols, ndata}
	return &newmat
}

//MatZerosf replicates matlab's zeros() function. Returns a new 32bit floating point matrix full of 0's
func MatZerosf(rows, cols int) *Matrixf {
	ndata := make([][]float32, rows)
	for i := 0; i < rows; i++ {
		ndata[i] = make([]float32, cols)
	}
	for j := 0; j < rows; j++ {
		for i := 0; i < cols; i++ {
			ndata[j][i] = 0
		}
	}
	newmat := Matrixf{rows, cols, ndata}
	return &newmat
}

//MatZerosff replicates matlab's zeros() function. Returns a new 64bit floating point matrix full of 0's
func MatZerosff(rows, cols int) *Matrixff {
	ndata := make([][]float64, rows)
	for i := 0; i < rows; i++ {
		ndata[i] = make([]float64, cols)
	}
	for j := 0; j < rows; j++ {
		for i := 0; i < cols; i++ {
			ndata[j][i] = 0
		}
	}
	newmat := Matrixff{rows, cols, ndata}
	return &newmat
}

//
// Implementation of Matlab's ones() function:
//

//MatOnesi replicates matlab's ones() function. Returns a new integer matrix full of 1's
func MatOnesi(rows, cols int) *Matrixi {
	ndata := make([][]int, rows)
	for i := 0; i < rows; i++ {
		ndata[i] = make([]int, cols)
	}
	for j := 0; j < rows; j++ {
		for i := 0; i < cols; i++ {
			ndata[j][i] = 1
		}
	}
	newmat := Matrixi{rows, cols, ndata}
	return &newmat
}

//MatOnesf replicates matlab's ones() function. Returns a new 32bit floating point matrix full of 1's
func MatOnesf(rows, cols int) *Matrixf {
	ndata := make([][]float32, rows)
	for i := 0; i < rows; i++ {
		ndata[i] = make([]float32, cols)
	}
	for j := 0; j < rows; j++ {
		for i := 0; i < cols; i++ {
			ndata[j][i] = 1
		}
	}
	newmat := Matrixf{rows, cols, ndata}
	return &newmat
}

//MatOnesff replicates matlab's ones() function. Returns a new 64bit floating point matrix full of 1's
func MatOnesff(rows, cols int) *Matrixff {
	ndata := make([][]float64, rows)
	for i := 0; i < rows; i++ {
		ndata[i] = make([]float64, cols)
	}
	for j := 0; j < rows; j++ {
		for i := 0; i < cols; i++ {
			ndata[j][i] = 1
		}
	}
	newmat := Matrixff{rows, cols, ndata}
	return &newmat
}

//
// Scalar multiplication implementation
//

//STimes performs simple scalar matrix multiplication for uint8 matrices
func (m1 Matrixi) STimes(s int) Matrixi {
	for i := 0; i < m1.Rows; i++ {
		for j := 0; j < m1.Cols; j++ {
			m1.Data[i][j] *= s
		}
	}
	return m1
}

//STimes performs simple scalar matrix multiplication for float32 matrices
func (m1 Matrixf) STimes(s float32) Matrixf {
	for i := 0; i < m1.Rows; i++ {
		for j := 0; j < m1.Cols; j++ {
			m1.Data[i][j] *= s
		}
	}
	return m1
}

//STimes performs simple scalar matrix multiplication for float32 matrices
func (m1 Matrixff) STimes(s float64) Matrixff {
	for i := 0; i < m1.Rows; i++ {
		for j := 0; j < m1.Cols; j++ {
			m1.Data[i][j] *= s
		}
	}
	return m1
}
