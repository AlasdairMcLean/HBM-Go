//Package hbmutil adds various functions relating to matrix operations.
package hbmutil

import (
	"fmt"
	"math/rand"
	"strings"
)

//Matrix is an overarching interface for the specific matrices to be derived from allowing overloading non-type specific methods like scalar multiplication
type Matrix interface {
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

//Matrixb defines a 2-dimensional binary bit map for increased efficiency and performance in matrix operations
type Matrixb struct {
	Data [][]bool
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

//NewMatrixb constructs a binary Matrix with specified rows and columns, returning the pointer to the new binary matrix
func NewMatrixb(rows, cols int) *Matrixb {
	ndata := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		ndata[i] = make([]bool, cols)
	}
	newmat := Matrixb{ndata}
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

//MatOnesb returns a binary matrix full of 'true' values instead of the default 'false' map
func MatOnesb(rows, cols int) *Matrixb {
	ndata := make([][]bool, rows)
	for i := 0; i < cols; i++ {
		ndata[i] = make([]bool, cols)
	}
	for j := 0; j < rows; j++ {
		for i := 0; i < cols; i++ {
			ndata[j][i] = true
		}
	}
	newmat := Matrixb{ndata}
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

//Det takes the determinant of the integer matrix
func (m1 Matrixi) Det() int {
	var sum int
	if m1.Rows != m1.Cols {
		pmsg := fmt.Sprintf("Matrix must be an N x N; Given matrix dimensions were: %v x %v.", m1.Rows, m1.Cols)
		panic(pmsg)
	}
	for j := 0; j < m1.Rows; j++ {
		line := m1.Data[j][0]
		//fmt.Printf("New Line: %v\n", line)
		for i, r, c := 1, j, 0; i < m1.Cols; i, r, c = i+1, r+1, c+1 {
			currow := r + 1
			curcol := c + 1
			if currow > m1.Rows-1 {
				currow -= m1.Rows
			}
			if curcol > m1.Cols-1 {
				curcol -= m1.Cols
			}
			//fmt.Printf("Row:%v, Col:%v, times:%v, line:%v\n", currow, curcol, m1.Data[currow][curcol], (line * m1.Data[currow][curcol]))
			line *= m1.Data[currow][curcol]
		}
		//fmt.Printf("Line done: %v\n", line)
		sum += line
	}
	for j := 0; j < m1.Rows; j++ {
		line := m1.Data[j][m1.Cols-1]
		//fmt.Printf("New Line: %v\n", line)
		for i, r, c := 1, j, m1.Cols-1; i < m1.Cols; i, r, c = i+1, r+1, c-1 {
			currow := r + 1
			curcol := c - 1
			if currow > m1.Rows-1 {
				currow -= m1.Rows
			}
			if curcol > m1.Cols-1 {
				curcol -= m1.Cols
			}
			//fmt.Printf("Row:%v, Col:%v, times:%v, line:%v\n", currow, curcol, m1.Data[currow][curcol], (line * m1.Data[currow][curcol]))
			line *= m1.Data[currow][curcol]
		}
		//fmt.Printf("Line done: %v\n", line)
		sum -= line
	}
	return sum
}

//Det takes the determinant of the integer matrix
func (m1 Matrixf) Det() float32 {
	var sum float32
	if m1.Rows != m1.Cols {
		pmsg := fmt.Sprintf("Matrix must be an N x N; Given matrix dimensions were: %v x %v.", m1.Rows, m1.Cols)
		panic(pmsg)
	}
	for j := 0; j < m1.Rows; j++ {
		line := m1.Data[j][0]
		//fmt.Printf("New Line: %v\n", line)
		for i, r, c := 1, j, 0; i < m1.Cols; i, r, c = i+1, r+1, c+1 {
			currow := r + 1
			curcol := c + 1
			if currow > m1.Rows-1 {
				currow -= m1.Rows
			}
			if curcol > m1.Cols-1 {
				curcol -= m1.Cols
			}
			//fmt.Printf("Row:%v, Col:%v, times:%v, line:%v\n", currow, curcol, m1.Data[currow][curcol], (line * m1.Data[currow][curcol]))
			line *= m1.Data[currow][curcol]
		}
		//fmt.Printf("Line done: %v\n", line)
		sum += line
	}
	for j := 0; j < m1.Rows; j++ {
		line := m1.Data[j][m1.Cols-1]
		//fmt.Printf("New Line: %v\n", line)
		for i, r, c := 1, j, m1.Cols-1; i < m1.Cols; i, r, c = i+1, r+1, c-1 {
			currow := r + 1
			curcol := c - 1
			if currow > m1.Rows-1 {
				currow -= m1.Rows
			}
			if curcol > m1.Cols-1 {
				curcol -= m1.Cols
			}
			//fmt.Printf("Row:%v, Col:%v, times:%v, line:%v\n", currow, curcol, m1.Data[currow][curcol], (line * m1.Data[currow][curcol]))
			line *= m1.Data[currow][curcol]
		}
		//fmt.Printf("Line done: %v\n", line)
		sum -= line
	}
	return sum
}

//Det takes the determinant of the integer matrix
func (m1 Matrixff) Det() float64 {
	var sum float64
	if m1.Rows != m1.Cols {
		pmsg := fmt.Sprintf("Matrix must be an N x N; Given matrix dimensions were: %v x %v.", m1.Rows, m1.Cols)
		panic(pmsg)
	}
	for j := 0; j < m1.Rows; j++ {
		line := m1.Data[j][0]
		//fmt.Printf("New Line: %v\n", line)
		for i, r, c := 1, j, 0; i < m1.Cols; i, r, c = i+1, r+1, c+1 {
			currow := r + 1
			curcol := c + 1
			if currow > m1.Rows-1 {
				currow -= m1.Rows
			}
			if curcol > m1.Cols-1 {
				curcol -= m1.Cols
			}
			//fmt.Printf("Row:%v, Col:%v, times:%v, line:%v\n", currow, curcol, m1.Data[currow][curcol], (line * m1.Data[currow][curcol]))
			line *= m1.Data[currow][curcol]
		}
		//fmt.Printf("Line done: %v\n", line)
		sum += line
	}
	for j := 0; j < m1.Rows; j++ {
		line := m1.Data[j][m1.Cols-1]
		//fmt.Printf("New Line: %v\n", line)
		for i, r, c := 1, j, m1.Cols-1; i < m1.Cols; i, r, c = i+1, r+1, c-1 {
			currow := r + 1
			curcol := c - 1
			if currow > m1.Rows-1 {
				currow -= m1.Rows
			}
			if curcol > m1.Cols-1 {
				curcol -= m1.Cols
			}
			//fmt.Printf("Row:%v, Col:%v, times:%v, line:%v\n", currow, curcol, m1.Data[currow][curcol], (line * m1.Data[currow][curcol]))
			line *= m1.Data[currow][curcol]
		}
		//fmt.Printf("Line done: %v\n", line)
		sum -= line
	}
	return sum
}

//Times returns the matrix multiplication of the input matrices
func (m1 Matrixi) Times(m2 Matrixi) *Matrixi {
	if m1.Cols != m2.Rows {
		pmsg := fmt.Sprintf("The number of columns in m1 must equal the number of rows in m2. Instead, matrices were %v x %v and %v x %v respectively.", m1.Rows, m1.Cols, m2.Rows, m2.Cols)
		panic(pmsg)
	}
	out := NewMatrixi(m1.Rows, m2.Cols)
	for k := 0; k < m2.Cols; k++ {
		for j := 0; j < m1.Rows; j++ {
			for i := 0; i < m1.Cols; i++ {
				out.Data[j][k] += m1.Data[j][i] * m2.Data[i][k] //remember that m1.Cols==m2.Rows so i is also the row index for m2.Data
			}
		}
	}
	return out
}

//Times returns the matrix multiplication of the input matrices
func (m1 Matrixf) Times(m2 Matrixf) *Matrixf {
	if m1.Cols != m2.Rows {
		pmsg := fmt.Sprintf("The number of columns in m1 must equal the number of rows in m2. Instead, matrices were %v x %v and %v x %v respectively.", m1.Rows, m1.Cols, m2.Rows, m2.Cols)
		panic(pmsg)
	}
	out := NewMatrixf(m1.Rows, m2.Cols)
	for k := 0; k < m2.Cols; k++ {
		for j := 0; j < m1.Rows; j++ {
			for i := 0; i < m1.Cols; i++ {
				out.Data[j][k] += m1.Data[j][i] * m2.Data[i][k] //remember that m1.Cols==m2.Rows so i is also the row index for m2.Data
			}
		}
	}
	return out
}

//Times returns the matrix multiplication of the input matrices
func (m1 Matrixff) Times(m2 Matrixff) *Matrixff {
	if m1.Cols != m2.Rows {
		pmsg := fmt.Sprintf("The number of columns in m1 must equal the number of rows in m2. Instead, matrices were %v x %v and %v x %v respectively.", m1.Rows, m1.Cols, m2.Rows, m2.Cols)
		panic(pmsg)
	}
	out := NewMatrixff(m1.Rows, m2.Cols)
	for k := 0; k < m2.Cols; k++ {
		for j := 0; j < m1.Rows; j++ {
			for i := 0; i < m1.Cols; i++ {
				out.Data[j][k] += m1.Data[j][i] * m2.Data[i][k] //remember that m1.Cols==m2.Rows so i is also the row index for m2.Data
			}
		}
	}
	return out
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

//Pretty prints out the values of the matrix in a visually readable way
func (m1 Matrixb) Pretty() {
	fmt.Println(strings.Repeat("-", (len(m1.Data[0])*6 + 1)))
	for j := 0; j < len(m1.Data); j++ {
		fmt.Println(m1.Data[j][:])
	}
	fmt.Println(strings.Repeat("-", len(m1.Data)*6+1))
}

//Unpackr unpacks all values into a given row
func (m1 *Matrixi) Unpackr(row int, vals ...int) {
	if row > m1.Rows-1 {
		pmsg := fmt.Sprintf("Attempt to unpack exceeded rows dimension of input matrix. Matrix has %v rows, but attempted to add vals to row %v (index %v)", m1.Rows, row+1, row)
		panic(pmsg)
	}
	if m1.Rows < len(vals) {
		pmsg := fmt.Sprintf("Attempt to unpack exceeded columns dimension of input matrix. Matrix has %v columns, but attempted to add %v vals", m1.Cols, len(vals))
		panic(pmsg)
	}
	for i, val := range vals {
		m1.Data[row][i] = val
	}
}

//Unpackr unpacks all values into a given row
func (m1 *Matrixf) Unpackr(row int, vals ...float32) {
	if row > m1.Rows-1 {
		pmsg := fmt.Sprintf("Attempt to unpack exceeded rows dimension of input matrix. Matrix has %v rows, but attempted to add vals to row %v (index %v)", m1.Rows, row+1, row)
		panic(pmsg)
	}
	if m1.Cols < len(vals) {
		pmsg := fmt.Sprintf("Attempt to unpack exceeded columns dimension of input matrix. Matrix has %v columns, but attempted to add %v vals", m1.Cols, len(vals))
		panic(pmsg)
	}
	for i, val := range vals {
		m1.Data[row][i] = val
	}
}

//Unpackr unpacks all values into a given row
func (m1 *Matrixff) Unpackr(row int, vals ...float64) {
	if row > m1.Rows-1 {
		pmsg := fmt.Sprintf("Attempt to unpack exceeded rows dimension of input matrix. Matrix has %v rows, but attempted to add vals to row %v (index %v)", m1.Rows, row+1, row)
		panic(pmsg)
	}
	if m1.Cols < len(vals) {
		pmsg := fmt.Sprintf("Attempt to unpack exceeded columns dimension of input matrix. Matrix has %v columns, but attempted to add %v vals", m1.Cols, len(vals))
		panic(pmsg)
	}
	for i, val := range vals {
		m1.Data[row][i] = val
	}
}

//Unpackr unpacks all values into a given row
func (m1 *Matrixb) Unpackr(row int, vals ...bool) {
	if row > len(m1.Data)-1 {
		pmsg := fmt.Sprintf("Attempt to unpack exceeded rows dimension of input matrix. Matrix has %v rows, but attempted to add vals to row %v (index %v)", len(m1.Data), row+1, row)
		panic(pmsg)
	}
	if len(m1.Data) < len(vals) {
		pmsg := fmt.Sprintf("Attempt to unpack exceeded columns dimension of input matrix. Matrix has %v columns, but attempted to add %v vals", len(m1.Data[0]), len(vals))
		panic(pmsg)
	}
	for i, val := range vals {
		m1.Data[row][i] = val
	}
}

//Unpackc unpacks all values into a given column
func (m1 *Matrixi) Unpackc(col int, vals ...int) {
	if col > m1.Cols-1 {
		pmsg := fmt.Sprintf("Attempt to unpack exceeded columns dimension of input matrix. Matrix has %v columns, but attempted to add vals to col %v (index %v)", m1.Cols, col+1, col)
		panic(pmsg)
	}
	if m1.Rows < len(vals) {
		pmsg := fmt.Sprintf("Attempt to unpack exceeded rows dimension of input matrix. Matrix has %v rows, but attempted to add %v vals", m1.Rows, len(vals))
		panic(pmsg)
	}
	for j, val := range vals {
		m1.Data[j][col] = val
	}
}

//Unpackc unpacks all values into a given column
func (m1 *Matrixf) Unpackc(col int, vals ...float32) {
	if col > m1.Cols-1 {
		pmsg := fmt.Sprintf("Attempt to unpack exceeded columns dimension of input matrix. Matrix has %v columns, but attempted to add vals to column %v (index %v)", m1.Cols, col+1, col)
		panic(pmsg)
	}
	if m1.Rows < len(vals) {
		pmsg := fmt.Sprintf("Attempt to unpack exceeded rows dimension of input matrix. Matrix has %v rows, but attempted to add %v vals", m1.Cols, len(vals))
		panic(pmsg)
	}
	for j, val := range vals {
		m1.Data[j][col] = val
	}
}

//Unpackc unpacks all values into a given column
func (m1 *Matrixff) Unpackc(col int, vals ...float64) {
	if col > m1.Cols-1 {
		pmsg := fmt.Sprintf("Attempt to unpack exceeded cols dimension of input matrix. Matrix has %v cols, but attempted to add vals to col %v (index %v)", m1.Cols, col+1, col)
		panic(pmsg)
	}
	if m1.Rows < len(vals) {
		pmsg := fmt.Sprintf("Attempt to unpack exceeded columns dimension of input matrix. Matrix has %v columns, but attempted to add %v vals", m1.Cols, len(vals))
		panic(pmsg)
	}
	for j, val := range vals {
		m1.Data[j][col] = val
	}
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

//ToMati will panic if you try to convert an integer matrix into an integer matrix
func (m1 *Matrixi) ToMati() {
	panic("Matrix is already an integer matrix")
}

//ToMati will reassign the values of a float32 matrix into an integer matrix
func (m1 *Matrixf) ToMati() *Matrixi {
	out := NewMatrixi(m1.Rows, m1.Cols)
	for i := 0; i < m1.Cols; i++ {
		for j := 0; j < m1.Rows; j++ {
			out.Data[j][i] = int(m1.Data[j][i])
		}
	}
	return out
}

//ToMati will reassign the values of a float64 matrix into an integer matrix
func (m1 *Matrixff) ToMati() *Matrixi {
	out := NewMatrixi(m1.Rows, m1.Cols)
	for i := 0; i < m1.Cols; i++ {
		for j := 0; j < m1.Rows; j++ {
			out.Data[j][i] = int(m1.Data[j][i])
		}
	}
	return out
}

//ToMatf will reassign the values of an integer matrix into a float32 matrix
func (m1 *Matrixi) ToMatf() *Matrixf {
	out := NewMatrixf(m1.Rows, m1.Cols)
	for i := 0; i < m1.Cols; i++ {
		for j := 0; j < m1.Rows; j++ {
			out.Data[j][i] = float32(m1.Data[j][i])
		}
	}
	return out
}

//ToMatf will panic if you try to convert a float32 matrix into a float32 matrix
func (m1 *Matrixf) ToMatf() {
	panic("Matrix is already a float32 matrix")
}

//ToMatf will reassign the values of a float64 matrix into a float32 matrix
func (m1 *Matrixff) ToMatf() *Matrixf {
	out := NewMatrixf(m1.Rows, m1.Cols)
	for i := 0; i < m1.Cols; i++ {
		for j := 0; j < m1.Rows; j++ {
			out.Data[j][i] = float32(m1.Data[j][i])
		}
	}
	return out
}

//ToMatff will reassign the values of a float64 matrix into a float32 matrix
func (m1 *Matrixi) ToMatff() *Matrixff {
	out := NewMatrixff(m1.Rows, m1.Cols)
	for i := 0; i < m1.Cols; i++ {
		for j := 0; j < m1.Rows; j++ {
			out.Data[j][i] = float64(m1.Data[j][i])
		}
	}
	return out
}

//ToMatff will reassign the values of a float64 matrix into a float32 matrix
func (m1 *Matrixf) ToMatff() *Matrixff {
	out := NewMatrixff(m1.Rows, m1.Cols)
	for i := 0; i < m1.Cols; i++ {
		for j := 0; j < m1.Rows; j++ {
			out.Data[j][i] = float64(m1.Data[j][i])
		}
	}
	return out
}

//ToMatff will panic if you try to convert a float64 matrix into a float64 matrix
func (m1 *Matrixff) ToMatff() {
	panic("Matrix is already a float64 matrix")
}

//Maxr will find the maximum in each row of a matrix
func (m1 *Matrixi) Maxr() []int {
	maxes := make([]int, m1.Rows)
	for j := 0; j < m1.Rows; j++ {
		rowmax := m1.Data[j][0]
		for i := 1; i < m1.Cols; i++ {
			if m1.Data[j][i] > rowmax {
				rowmax = m1.Data[j][i]
			}
		}
		maxes[j] = rowmax
	}
	return maxes
}

//Maxr will find the maximum in each row of a matrix
func (m1 *Matrixf) Maxr() []float32 {
	maxes := make([]float32, m1.Rows)
	for j := 0; j < m1.Rows; j++ {
		rowmax := m1.Data[j][0]
		for i := 1; i < m1.Cols; i++ {
			if m1.Data[j][i] > rowmax {
				rowmax = m1.Data[j][i]
			}
		}
		maxes[j] = rowmax
	}
	return maxes
}

//Maxr will find the maximum in each row of a matrix
func (m1 *Matrixff) Maxr() []float64 {
	maxes := make([]float64, m1.Rows)
	for j := 0; j < m1.Rows; j++ {
		rowmax := m1.Data[j][0]
		for i := 1; i < m1.Cols; i++ {
			if m1.Data[j][i] > rowmax {
				rowmax = m1.Data[j][i]
			}
		}
		maxes[j] = rowmax
	}
	return maxes
}

//Maxc will find the maximum in each column of a matrix
func (m1 *Matrixi) Maxc() []int {
	maxes := make([]int, m1.Cols)
	for i := 0; i < m1.Cols; i++ {
		colmax := m1.Data[0][i]
		for j := 1; j < m1.Rows; j++ {
			if m1.Data[j][i] > colmax {
				colmax = m1.Data[j][i]
			}
		}
		maxes[i] = colmax
	}
	return maxes
}

//Maxc will find the maximum in each column of a matrix
func (m1 *Matrixf) Maxc() []float32 {
	maxes := make([]float32, m1.Cols)
	for i := 0; i < m1.Cols; i++ {
		colmax := m1.Data[0][i]
		for j := 1; j < m1.Rows; j++ {
			if m1.Data[j][i] > colmax {
				colmax = m1.Data[j][i]
			}
		}
		maxes[i] = colmax
	}
	return maxes
}

//Maxc will find the maximum in each column of a matrix
func (m1 *Matrixff) Maxc() []float64 {
	maxes := make([]float64, m1.Cols)
	for i := 0; i < m1.Cols; i++ {
		colmax := m1.Data[0][i]
		for j := 1; j < m1.Rows; j++ {
			if m1.Data[j][i] > colmax {
				colmax = m1.Data[j][i]
			}
		}
		maxes[i] = colmax
	}
	return maxes
}

//Maxa will find the total maximum of a matrix
func (m1 *Matrixi) Maxa() int {
	max := m1.Data[0][0]
	for i := 0; i < m1.Cols; i++ {
		for j := 0; j < m1.Rows; j++ {
			if m1.Data[j][i] > max {
				max = m1.Data[j][i]
			}
		}
	}
	return max
}

//Maxa will find the total maximum of a matrix
func (m1 *Matrixf) Maxa() float32 {
	max := m1.Data[0][0]
	for i := 0; i < m1.Cols; i++ {
		for j := 0; j < m1.Rows; j++ {
			if m1.Data[j][i] > max {
				max = m1.Data[j][i]
			}
		}
	}
	return max
}

//Maxa will find the total maximum of a matrix
func (m1 *Matrixff) Maxa() float64 {
	max := m1.Data[0][0]
	for i := 0; i < m1.Cols; i++ {
		for j := 0; j < m1.Rows; j++ {
			if m1.Data[j][i] > max {
				max = m1.Data[j][i]
			}
		}
	}
	return max
}

//Getcol will retrieve all values from a given column and return them as a slice
func (m1 *Matrixi) Getcol(n int) []int {
	out := make([]int, m1.Rows)
	for j := 0; j < m1.Rows; j++ {
		out[j] = m1.Data[j][n-1]
	}
	return out
}

//Getcol will retrieve all values from a given column and return them as a slice
func (m1 *Matrixf) Getcol(n int) []float32 {
	out := make([]float32, m1.Rows)
	for j := 0; j < m1.Rows; j++ {
		out[j] = m1.Data[j][n-1]
	}
	return out
}

//Getcol will retrieve all values from a given column and return them as a slice
func (m1 *Matrixff) Getcol(n int) []float64 {
	out := make([]float64, m1.Rows)
	for j := 0; j < m1.Rows; j++ {
		out[j] = m1.Data[j][n-1]
	}
	return out
}

//Getcol will retrieve all values from a given column and return them as a slice
func (m1 *Matrixb) Getcol(n int) []bool {
	out := make([]bool, len(m1.Data))
	for j := 0; j < len(m1.Data); j++ {
		out[j] = m1.Data[j][n-1]
	}
	return out
}

//MatEqual determines if multiple matrices are completely equivalent
func (m1 Matrixi) MatEqual(m ...Matrixi) bool {
	rows := m1.Rows
	cols := m1.Cols
	for i := range m {
		if m[i].Rows != rows || m[i].Cols != cols {
			panic("Error: mismatching indices")
		}
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			val := m1.Data[j][i]
			for k := 0; k < len(m); k++ {
				if m[k].Data[j][i] != val {
					return false
				}
			}
		}
	}
	return true
}

//MatEqual determines if multiple matrices are completely equivalent
func (m1 Matrixf) MatEqual(m ...Matrixf) bool {
	rows := m1.Rows
	cols := m1.Cols
	for i := range m {
		if m[i].Rows != rows || m[i].Cols != cols {
			panic("Error: mismatching indices")
		}
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			val := m1.Data[j][i]
			for k := 0; k < len(m); k++ {
				if m[k].Data[j][i] != val {
					return false
				}
			}
		}
	}
	return true
}

//MatEqual determines if multiple matrices are completely equivalent
func (m1 Matrixff) MatEqual(m ...Matrixff) bool {
	rows := m1.Rows
	cols := m1.Cols
	for i := range m {
		if m[i].Rows != rows || m[i].Cols != cols {
			panic("Error: mismatching indices")
		}
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			val := m1.Data[j][i]
			for k := 0; k < len(m); k++ {
				if m[k].Data[j][i] != val {
					return false
				}
			}
		}
	}
	return true
}

//MatEqual determines if multiple matrices are completely equivalent
func (m1 Matrixb) MatEqual(m ...Matrixb) bool {
	rows := len(m1.Data)
	cols := len(m1.Data[0])
	for i := range m {
		if len(m[i].Data) != rows || len(m[i].Data[0]) != cols {
			panic("Error: mismatching indices")
		}
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			val := m1.Data[j][i]
			for k := 0; k < len(m); k++ {
				if m[k].Data[j][i] != val {
					return false
				}
			}
		}
	}
	return true
}

//Combinelr combines multiple matrices into one larger matrix, appending to the m1 matrix from left to right. If one matrix has more rows than another, the extraneous rows will be zero-valued. Example: if [1,2;3,4] and [5] are appended, the output will be [1,2,5;3,4,0].
func (m1 Matrixi) Combinelr(m ...Matrixi) *Matrixi {
	cols := m1.Cols
	for ind := 0; ind < len(m); ind++ {
		cols += m[ind].Cols
	}
	rowsl := make([]int, len(m)+1)
	rowsl[0] = m1.Rows
	for i := 0; i < len(m); i++ {
		rowsl[i+1] = m[i].Rows
	}
	rows := Maxil(rowsl) // find the maximum number of columns in all matrices
	out := NewMatrixi(rows, cols)
	for i := 0; i < m1.Cols; i++ {
		for j := 0; j < m1.Rows; j++ {
			out.Data[j][i] = m1.Data[j][i]
		}
	}
	curcol := m1.Cols
	for mat := 0; mat < len(m); mat++ {
		for i := curcol; i < curcol+m[mat].Cols; i++ {
			for j := 0; j < m[mat].Rows; j++ {
				out.Data[j][i] = m[mat].Data[j][i-curcol]
			}
		}
		curcol += m[mat].Cols
	}
	return out
}

//Combinelr combines multiple matrices into one larger matrix, appending to the m1 matrix from left to right. If one matrix has more rows than another, the extraneous rows will be zero-valued. Example: if [1,2;3,4] and [5] are appended, the output will be [1,2,5;3,4,0].
func (m1 Matrixf) Combinelr(m ...Matrixf) *Matrixf {
	cols := m1.Cols
	for ind := 0; ind < len(m); ind++ {
		cols += m[ind].Cols
	}
	rowsl := make([]int, len(m)+1)
	rowsl[0] = m1.Rows
	for i := 0; i < len(m); i++ {
		rowsl[i+1] = m[i].Rows
	}
	rows := Maxil(rowsl) // find the maximum number of columns in all matrices
	out := NewMatrixf(rows, cols)
	for i := 0; i < m1.Cols; i++ {
		for j := 0; j < m1.Rows; j++ {
			out.Data[j][i] = m1.Data[j][i]
		}
	}
	curcol := m1.Cols
	for mat := 0; mat < len(m); mat++ {
		for i := curcol; i < curcol+m[mat].Cols; i++ {
			for j := 0; j < m[mat].Rows; j++ {
				out.Data[j][i] = m[mat].Data[j][i-curcol]
			}
		}
		curcol += m[mat].Cols
	}
	return out
}

//Combinelr combines multiple matrices into one larger matrix, appending to the m1 matrix from left to right. If one matrix has more rows than another, the extraneous rows will be zero-valued. Example: if [1,2;3,4] and [5] are appended, the output will be [1,2,5;3,4,0].
func (m1 Matrixff) Combinelr(m ...Matrixff) *Matrixff {
	cols := m1.Cols
	for ind := 0; ind < len(m); ind++ {
		cols += m[ind].Cols
	}
	rowsl := make([]int, len(m)+1)
	rowsl[0] = m1.Rows
	for i := 0; i < len(m); i++ {
		rowsl[i+1] = m[i].Rows
	}
	rows := Maxil(rowsl) // find the maximum number of columns in all matrices
	out := NewMatrixff(rows, cols)
	for i := 0; i < m1.Cols; i++ {
		for j := 0; j < m1.Rows; j++ {
			out.Data[j][i] = m1.Data[j][i]
		}
	}
	curcol := m1.Cols
	for mat := 0; mat < len(m); mat++ {
		for i := curcol; i < curcol+m[mat].Cols; i++ {
			for j := 0; j < m[mat].Rows; j++ {
				out.Data[j][i] = m[mat].Data[j][i-curcol]
			}
		}
		curcol += m[mat].Cols
	}
	return out
}

//Linspacei returns an array of equally spaced numbers between start and end (inclusive!) just like matlab's linspace() function
func Linspacei(start, end, vals int) []int {
	if vals < 2 {
		panic("You must have at least 2 values in your array")
	}
	out := make([]int, vals)            //create the n x 3 colormap rgb matrix to be returned and dereference it.
	level := (end - start) / (vals - 1) //determine the increments necessary to create a clean transition between inclusive 0 and 1
	for i := 0; i < vals; i++ {         //iterate through each value of the colormap
		out[i] = level*(i) + start // since this colormap is greyscale, all values will be the same per row.
	}
	return out //return the pointer to the new matrix
}

//Linspacef returns an array of equally spaced numbers between start and end (inclusive!) just like matlab's linspace() function
func Linspacef(start, end float32, vals int) []float32 {
	if vals < 2 {
		panic("You must have at least 2 values in your array")
	}
	out := make([]float32, vals)                 //create the n x 3 colormap rgb matrix to be returned and dereference it.
	level := (end - start) / (float32(vals) - 1) //determine the increments necessary to create a clean transition between inclusive 0 and 1
	for i := 0; i < vals; i++ {                  //iterate through each value of the colormap
		out[i] = level*float32(i) + start // since this colormap is greyscale, all values will be the same per row.
	}
	return out //return the pointer to the new matrix
}

//Linspaceff returns an array of equally spaced numbers between start and end (inclusive!) just like matlab's linspace() function
func Linspaceff(start, end float64, vals int) []float64 {
	if vals < 2 {
		panic("You must have at least 2 values in your array")
	}
	out := make([]float64, vals)                 //create the n x 3 colormap rgb matrix to be returned and dereference it.
	level := (end - start) / (float64(vals) - 1) //determine the increments necessary to create a clean transition between inclusive 0 and 1
	for i := 0; i < vals; i++ {                  //iterate through each value of the colormap
		out[i] = level*float64(i) + start // since this colormap is greyscale, all values will be the same per row.
	}
	return out //return the pointer to the new matrix
}
