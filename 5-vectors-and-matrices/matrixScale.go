package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func main() {
	usd := mat.NewDense(3, 3, []float64{0.1, 0.05, 0.03, 0.06, 0.04, 0.02, 0.03, 0.02, 0.01})
	var cad mat.Dense
	cad.Scale(1.34, usd)
	result := mat.Formatted(&cad, mat.Prefix(""), mat.Squeeze())
	fmt.Println(result)
}
