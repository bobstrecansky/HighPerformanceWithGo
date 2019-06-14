package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func main() {
	a := mat.NewDense(2, 2, []float64{0, 1, 2, 3})
	b := mat.NewDense(2, 2, []float64{4, 5, 6, 7})
	c := mat.NewDense(2, 2, nil)
	c.Sub(b, a)
	fmt.Println(c)
}
