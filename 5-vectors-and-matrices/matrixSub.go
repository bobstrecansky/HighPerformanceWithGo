package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func main() {
	a := mat.NewDense(4, 2, []float64{1345, 823, 346, 234, 843, 945, 442, 692})
	b := mat.NewDense(4, 2, []float64{920, 776, 498, 439, 902, 1023, 663, 843})
	var c mat.Dense
	c.Sub(b, a)
	result := mat.Formatted(&c, mat.Prefix(""), mat.Squeeze())
	fmt.Println(result)
}
