package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// Matrix represents a two dimensional array
type Matrix [][]int

// New is a Matrix constructor
func New(input string) (Matrix, error) {
	rows := strings.Split(input, "\n")
	m := make([][]int, len(rows))
	var ncol int
	for i, r := range rows {
		col := strings.Fields(r)
		if 0 == len(col) {
			return Matrix{}, errors.New("empty row")
		}
		if i > 0 && ncol != len(col) {
			return Matrix{}, errors.New("uneven rows")
		}
		ncol = len(col)
		m[i] = make([]int, ncol)
		for j, v := range col {
			n, err := strconv.Atoi(v)
			if err != nil {
				return Matrix{}, errors.New("string to integer conversion failed")
			}
			m[i][j] = n
		}
	}
	return m, nil
}

// Rows returns the rows of the matrix
func (m Matrix) Rows() [][]int {
	rows := make([][]int, len(m))
	for i, r := range m {
		rows[i] = make([]int, len(r))
		copy(rows[i], r)
	}
	return rows
}

// Cols returns the columns of the matrix
func (m Matrix) Cols() [][]int {
	cols := make([][]int, len(m[0]))
	for j := 0; j < len(m[0]); j++ {
		cols[j] = make([]int, len(m))
		for i := 0; i < len(m); i++ {
			cols[j][i] = m[i][j]
		}
	}
	return cols
}

// Set sets a matrix cell of given row and column to a given val
func (m Matrix) Set(row, column, val int) bool {
	if row < 0 || row >= len(m) || column < 0 || column >= len(m[0]) {
		return false
	}
	m[row][column] = val
	return true
}
