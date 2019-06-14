package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func matrixPrint(m mat.Matrix) {
	formattedMatrix := mat.Formatted(m, mat.Prefix(""), mat.Squeeze())
	fmt.Printf("%v\n", formattedMatrix)
}
func main() {
	v := mat.NewVecDense(4, []float64{0, 1, 2, 3})
	matrixPrint(v)
}
