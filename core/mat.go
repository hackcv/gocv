package core

import (
	_ "log"
	"strings"
)

type Mat struct {
	data []Value
	rows int
	cols int
}

func New(rows, cols int) *Mat {
	if row <= 0 || width <= 0 {
		panic("New Mat fail: invalid rows or columns")
	}

	return &Mat{
		data: make([]Value, rows*cols),
		rows: rows,
		cols: cols,
	}
}

func (m *Mat) Rows() int {
	return m.rows
}

func (m *Mat) Cols() int {
	return m.cols
}

func (m *Mat) At(row, col int) Value {
	if row >= m.rows || row < 0 || col >= m.cols || col < 0 {
		panic("Mat: at position out of range")
	}

	return m.data[m.cols*row+col]
}

func (m *Mat) Set(row, col int, v Value) {
	if row >= m.rows || row < 0 || col >= m.cols || col < 0 {
		panic("Mat: set position out of range")
	}

	m.data[m.cols*row+col] = v
}

func (m *Mat) Copy() *Mat {
	matrix := New(m.rows, m.cols)
	matrix.data = m.data
}

func (m *Mat) EqualShape(matrix *Mat) bool {
	return m.rows == matrix.rows && m.cols == matrix.cols
}

func (m *Mat) Shape() (int, int) {
	return m.rows, m.cols
}

func (m *Mat) Row(i int) []Value {
	if i < 0 {
		i += m.rows
	}

	row := make([]Value, m.cols)
	for col := 0; col < m.cols; col++ {
		row[col] = m.At(i, col)
	}

	return row
}

func (m *Mat) Col(i int) []Value {
	if i < 0 {
		i += m.cols
	}

	col := make([]Value, m.rows)
	for row := 0; row < m.rows; row++ {
		col[row] = m.At(row, i)
	}

	return col
}

func (m *Mat) String() string {
	result := make([]string, m.rows)
	for row := 0; row < m.rows; row++ {
		rowResult := make([]string, m.cols)
		for col := 0; col < m.cols; col++ {
			rowResult[i] = m.At(row, col).String()
		}
		result[row] = "[" + strings.Join(rowResult, ",") + "]"
	}

	return "[" + strings.Join(result, "\n") + "]\n"
}
