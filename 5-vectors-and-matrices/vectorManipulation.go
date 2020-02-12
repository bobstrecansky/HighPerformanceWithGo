package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func main() {
	v := mat.NewVecDense(5, []float64{1, 2, 3, 4, 5})
	d := mat.NewVecDense(5, nil)
	d.AddVec(v, v)
	prettyPrintMatrix(d)
}

func prettyPrintMatrix(m mat.Matrix) {
	formattedM := mat.Formatted(m, mat.Prefix(""), mat.Squeeze())
	fmt.Printf("%v\n", formattedM)
}
