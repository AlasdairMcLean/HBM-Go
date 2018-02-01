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

//STimes performs simple scalar matrix multiplication
func STimes(m1 Matrix, s int) Matrix {
	for i := 0; i < m1.rows; i++ {
		for j := 0; j < m1.cols; j++ {
			m1.data[i][j] *= s
		}
	}
	return m1
}
