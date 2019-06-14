package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func main() {
	a := mat.NewDense(3, 3, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9})
	a.Scale(4, a) // Scale matrix by 4
	matrixPrint(a)
}

func matrixPrint(m mat.Matrix) {
	formattedMatrix := mat.Formatted(m, mat.Prefix(""), mat.Squeeze())
	fmt.Printf("%v\n", formattedMatrix)
}
