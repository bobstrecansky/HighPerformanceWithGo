package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func main() {
	a := mat.NewDense(1, 3, []float64{0.10, 0.42, 0.37})
	b := mat.NewDense(3, 2, []float64{5, 8, 10, 6, 2, 3})
	var c mat.Dense
	c.Mul(a, b)
	result := mat.Formatted(&c, mat.Prefix(""), mat.Squeeze())
	fmt.Println(result)
}
