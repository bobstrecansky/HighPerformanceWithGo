package main

import (
	"fmt"

	"github.com/james-bowman/sparse"
	"gonum.org/v1/gonum/mat"
)

func main() {
	sparseMatrix := sparse.NewDOK(3, 3)
	sparseMatrix.Set(0, 0, 5)
	sparseMatrix.Set(1, 1, 1)
	sparseMatrix.Set(2, 1, -3)
	fmt.Println(mat.Formatted(sparseMatrix))
	csrMatrix := sparseMatrix.ToCSR()
	fmt.Println(mat.Formatted(csrMatrix))
	cscMatrix := sparseMatrix.ToCSC()
	fmt.Println(mat.Formatted(cscMatrix))
}
