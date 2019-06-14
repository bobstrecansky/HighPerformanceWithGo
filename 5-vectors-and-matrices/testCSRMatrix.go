package main

import (
	"fmt"

	"github.com/james-bowman/sparse"
	"gonum.org/v1/gonum/mat"
)

func main() {
	sparseMatrix := sparse.NewDOK(4, 4)
	sparseMatrix.Set(1, 0, 5)
	sparseMatrix.Set(1, 1, 8)
	sparseMatrix.Set(2, 2, 3)
	sparseMatrix.Set(3, 1, 6)
	fmt.Print("DOK Matrix:\n", mat.Formatted(sparseMatrix), "\n\n") // Dictionary of Keys
	fmt.Print("CSR Matrix:\n", sparseMatrix.ToCSR(), "\n\n")        // Print CSR version of the matrix
}
