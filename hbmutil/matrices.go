package hbmutil

//Matrix is a struct to streamline matrix operations
type Matrix struct {
	rows int
	cols int
	data [][]int
}

//NewMatrix constructs a Matrix structure with specified rows and columns, returning the pointer to the new matrix
func NewMatrix(rows, cols int) *Matrix {
	ndata := make([][]int, rows)
	for i := 0; i < rows; i++ {
		ndata[i] = make([]int, cols)
	}
	newmat := Matrix{rows, cols, ndata}
	return &newmat
}

//MatZeros replicates matlab's zeros() function. Returns a new matrix full of 0's
func MatZeros(rows, cols int) *Matrix {
	ndata := make([][]int, rows)
	for i := 0; i < rows; i++ {
		ndata[i] = make([]int, cols)
	}
	for j := 0; j < rows; j++ {
		for i := 0; i < cols; i++ {
			ndata[j][i] = 0
		}
	}
	newmat := Matrix{rows, cols, ndata}
	return &newmat
}

//MatOnes replicates matlab's ones() function. Returns a new matrix full of 1's
func MatOnes(rows, cols int) *Matrix {
	ndata := make([][]int, rows)
	for i := 0; i < rows; i++ {
		ndata[i] = make([]int, cols)
	}
	for j := 0; j < rows; j++ {
		for i := 0; i < cols; i++ {
			ndata[j][i] = 1
		}
	}
	newmat := Matrix{rows, cols, ndata}
	return &newmat
}

//STimes performs simple scalar matrix multiplication
func STimes(m1 Matrix, s int) Matrix {
	for i := 0; i < m1.rows; i++ {
		for j := 0; j < m1.cols; j++ {
			m1.data[i][j] *= s
		}
	}
	return m1
}
