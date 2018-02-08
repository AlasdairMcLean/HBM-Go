package hbmutil

import (
	"fmt"
	"math/rand"
	"strings"
)

//Matrix is an overarching interface for the specific matrices to be derived from allowing overloading non-type specific methods like scalar multiplication
type Matrix interface {
	Scale()
	Dot()
	Transpose()
	Pretty()
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

//Scale performs simple scalar matrix multiplication for uint8 matrices
func (m1 Matrixi) Scale(s int) Matrixi {
	for i := 0; i < m1.Rows; i++ {
		for j := 0; j < m1.Cols; j++ {
			m1.Data[i][j] *= s
		}
	}
	return m1
}

//Scale performs simple scalar matrix multiplication for float32 matrices
func (m1 Matrixf) Scale(s float32) Matrixf {
	for i := 0; i < m1.Rows; i++ {
		for j := 0; j < m1.Cols; j++ {
			m1.Data[i][j] *= s
		}
	}
	return m1
}

//Scale performs simple scalar matrix multiplication for float64 matrices
func (m1 Matrixff) Scale(s float64) Matrixff {
	for i := 0; i < m1.Rows; i++ {
		for j := 0; j < m1.Cols; j++ {
			m1.Data[i][j] *= s
		}
	}
	return m1
}

//Dot performs simple dot product multiplication for two uint8 matrices
func (m1 Matrixi) Dot(m2 Matrixi) int {
	var sum int
	if m1.Rows == m2.Rows && m1.Cols == m2.Cols {
		for i := 0; i < m1.Rows; i++ {
			for j := 0; j < m1.Cols; j++ {
				sum += (m1.Data[i][j] * m2.Data[i][j])
			}
		}
	} else {
		pmsg := fmt.Sprintf("Matrices must be equal: Given matrix dimensions were: %v x %v, %v x %v.", m1.Rows, m1.Cols, m2.Rows, m2.Cols)
		panic(pmsg)
	}
	return sum
}

//Dot performs simple dot product multiplication for two float32 matrices
func (m1 Matrixf) Dot(m2 Matrixf) float32 {
	var sum float32
	if m1.Rows == m2.Rows && m1.Cols == m2.Cols {
		for i := 0; i < m1.Rows; i++ {
			for j := 0; j < m1.Cols; j++ {
				sum += (m1.Data[i][j] * m2.Data[i][j])
			}
		}
	} else {
		pmsg := fmt.Sprintf("Matrices must be equal: Given matrix dimensions were: %v x %v, %v x %v.", m1.Rows, m1.Cols, m2.Rows, m2.Cols)
		panic(pmsg)
	}
	return sum
}

//Dot performs simple dot product multiplication for two float64 matrices
func (m1 Matrixff) Dot(m2 Matrixff) float64 {
	var sum float64
	if m1.Rows == m2.Rows && m1.Cols == m2.Cols {
		for i := 0; i < m1.Rows; i++ {
			for j := 0; j < m1.Cols; j++ {
				sum += (m1.Data[i][j] * m2.Data[i][j])
			}
		}
	} else {
		pmsg := fmt.Sprintf("Matrices must be equal: Given matrix dimensions were: %v x %v, %v x %v.", m1.Rows, m1.Cols, m2.Rows, m2.Cols)
		panic(pmsg)
	}
	return sum
}

//Transpose returns the transpose of the input matrix
func (m1 Matrixi) Transpose() *Matrixi {
	tpose := NewMatrixi(m1.Cols, m1.Rows)
	for i := 0; i < m1.Rows; i++ {
		for j := 0; j < m1.Cols; j++ {
			tpose.Data[j][i] = m1.Data[i][j]
		}
	}
	return tpose
}

//Transpose returns the transpose of the input matrix
func (m1 Matrixf) Transpose() *Matrixf {
	tpose := NewMatrixf(m1.Cols, m1.Rows)
	for i := 0; i < m1.Rows; i++ {
		for j := 0; j < m1.Cols; j++ {
			tpose.Data[j][i] = m1.Data[i][j]
		}
	}
	return tpose
}

//Transpose returns the transpose of the input matrix
func (m1 Matrixff) Transpose() *Matrixff {
	tpose := NewMatrixff(m1.Cols, m1.Rows)
	for i := 0; i < m1.Rows; i++ {
		for j := 0; j < m1.Cols; j++ {
			tpose.Data[j][i] = m1.Data[i][j]
		}
	}
	return tpose
}

//Pretty prints out the values of the matrix in a visually readable way
func (m1 Matrixi) Pretty() {
	fmt.Println(strings.Repeat("-", m1.Cols*2+1))
	for j := 0; j < m1.Rows; j++ {
		fmt.Println(m1.Data[j][:])
	}
	fmt.Println(strings.Repeat("-", m1.Cols*2+1))
}

//Pretty prints out the values of the matrix in a visually readable way
func (m1 Matrixf) Pretty() {
	fmt.Println(strings.Repeat("-", m1.Cols*2+1))
	for j := 0; j < m1.Rows; j++ {
		fmt.Println(m1.Data[j][:])
	}
	fmt.Println(strings.Repeat("-", m1.Cols*2+1))
}

//Pretty prints out the values of the matrix in a visually readable way
func (m1 Matrixff) Pretty() {
	fmt.Println(strings.Repeat("-", m1.Cols*2+1))
	for j := 0; j < m1.Rows; j++ {
		fmt.Println(m1.Data[j][:])
	}
	fmt.Println(strings.Repeat("-", m1.Cols*2+1))
}

//Randmat returns a matrix of pseudo-random values
func Randmat(row, col, maxn int) *Matrixi {
	out := NewMatrixi(row, col)
	for i := 0; i < col; i++ {
		for j := 0; j < row; j++ {
			out.Data[j][i] = rand.Intn(maxn)
		}
	}
	return out
}
