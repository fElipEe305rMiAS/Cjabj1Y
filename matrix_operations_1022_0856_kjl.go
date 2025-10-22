// 代码生成时间: 2025-10-22 08:56:49
package matrix

import (
	"errors"
	"fmt"
)

// Matrix represents a 2D matrix
type Matrix struct {
	Rows [][]float64
}

// NewMatrix creates a new matrix from a 2D slice
func NewMatrix(rows [][]float64) *Matrix {
	return &Matrix{Rows: rows}
}

// Add adds two matrices together
func (m *Matrix) Add(other *Matrix) (*Matrix, error) {
	if len(m.Rows) != len(other.Rows) || len(m.Rows[0]) != len(other.Rows[0]) {
		return nil, errors.New("matrices must be the same size")
	}
	
	result := make([][]float64, len(m.Rows))
	for i := range m.Rows {
		result[i] = make([]float64, len(m.Rows[i]))
		for j := range m.Rows[i] {
			result[i][j] = m.Rows[i][j] + other.Rows[i][j]
		}
	}
	return NewMatrix(result), nil
}

// Subtract subtracts one matrix from another
func (m *Matrix) Subtract(other *Matrix) (*Matrix, error) {
	if len(m.Rows) != len(other.Rows) || len(m.Rows[0]) != len(other.Rows[0]) {
		return nil, errors.New("matrices must be the same size")
	}
	
	result := make([][]float64, len(m.Rows))
	for i := range m.Rows {
		result[i] = make([]float64, len(m.Rows[i]))
		for j := range m.Rows[i] {
			result[i][j] = m.Rows[i][j] - other.Rows[i][j]
		}
	}
	return NewMatrix(result), nil
}

// Multiply multiplies two matrices
func (m *Matrix) Multiply(other *Matrix) (*Matrix, error) {
	if len(m.Rows[0]) != len(other.Rows) {
		return nil, errors.New("number of columns in the first matrix must be equal to the number of rows in the second matrix")
	}
	
	result := make([][]float64, len(m.Rows))
	for i := range m.Rows {
		result[i] = make([]float64, len(other.Rows[0]))
		for j := range other.Rows[0] {
			sum := 0.0
			for k := range m.Rows[i] {
				sum += m.Rows[i][k] * other.Rows[k][j]
			}
			result[i][j] = sum
		}
	}
	return NewMatrix(result), nil
}

// Print prints the matrix
func (m *Matrix) Print() {
	for _, row := range m.Rows {
		fmt.Println(row)
	}
}

// Example usage
func Example() {
	matrix1 := [][]float64{
		{1, 2},
		{3, 4},
	}
	matrix2 := [][]float64{
		{5, 6},
		{7, 8},
	}

	mat1 := NewMatrix(matrix1)
	mat2 := NewMatrix(matrix2)

	addResult, err := mat1.Add(mat2)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Addition Result: ")
		addResult.Print()
	}

	subResult, err := mat1.Subtract(mat2)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Subtraction Result: ")
		subResult.Print()
	}

	mulResult, err := mat1.Multiply(mat2)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Multiplication Result: ")
		mulResult.Print()
	}
}

func main() {
	Example()
}