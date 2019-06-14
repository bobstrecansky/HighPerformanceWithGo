package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func main() {
	a := mat.NewDense(2, 3, []float64{0, 1, 2, 3, 4, 5})
	matrixPrint(a)
	matrixPrint(a.T())
}

func matrixPrint(m mat.Matrix) {
	formattedMatrix := mat.Formatted(m, mat.Prefix(""), mat.Squeeze())
	fmt.Printf("%v\n", formattedMatrix)
}
