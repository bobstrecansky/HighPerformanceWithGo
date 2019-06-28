package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func main() {
	a := mat.NewDense(2, 2, []float64{1, 2, 3, 4})
	b := mat.NewDense(2, 3, []float64{1, 2, 3, 4, 5, 6})
	var c mat.Dense
	c.Mul(a, b)
	result := mat.Formatted(&c, mat.Prefix(""), mat.Squeeze())
	fmt.Println(result)
}
