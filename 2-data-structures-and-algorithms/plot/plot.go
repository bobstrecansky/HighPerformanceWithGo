package main

import (
	"image/color"
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	p.Title.Text = "Big O Notation Plot"
	p.X.Label.Text = "Data Input"
	p.Y.Label.Text = "Execution Time"

	// Plot a Constant time: O(1) function
	constant := plotter.NewFunction(func(x float64) float64 { return 1.0 })
	constant.Width = vg.Points(4)

	// Plot a linear O(n) function
	linear := plotter.NewFunction(func(x float64) float64 { return x })
	linear.Width = vg.Points(2)
	linear.Dashes = []vg.Length{vg.Points(1), vg.Points(1)}

	// Plot a natural logarithm: O(log n) function
	log := plotter.NewFunction(func(x float64) float64 { return math.Log(x) })
	log.Dashes = []vg.Length{vg.Points(4), vg.Points(5)}
	log.Width = vg.Points(2)

	// Plot a linearithmic: O(n*log n) function
	linearithmic := plotter.NewFunction(func(x float64) float64 { return x * math.Log(x) })
	linearithmic.Dashes = []vg.Length{vg.Points(2), vg.Points(6)}
	linearithmic.Width = vg.Points(4)

	//	// Plot a quadratic: O(n^2) function
	quad := plotter.NewFunction(func(x float64) float64 { return x * x })
	quad.Dashes = []vg.Length{vg.Points(1), vg.Points(8)}
	quad.Width = vg.Points(4)

	//	// Plot an exponential: O(2^n) function
	exp := plotter.NewFunction(func(x float64) float64 { return math.Pow(2, x) })
	exp.Width = vg.Points(1)
	//
	// Add the functions and their legend entries.
	p.Add(constant, log, linear, linearithmic, quad, exp)
	p.Legend.Add("Constant: O(1)", constant)
	p.Legend.Add("Log Linear: O(log n)", log)
	p.Legend.Add("Linear: O(n)", linear)
	p.Legend.Add("Linearithmic: O(n*log n)", linearithmic)
	p.Legend.Add("Quadratic: O(n^2)", quad)
	p.Legend.Add("Exponential: O(2^n)", exp)
	p.Legend.YOffs = 100
	p.Legend.Color = color.Black
	p.Legend.ThumbnailWidth = 0.5 * vg.Inch

	// Set the axis ranges.  Unlike other data sets,
	// functions don't set the axis ranges automatically
	// since functions don't necessarily have a
	// finite range of x and y values.
	p.X.Min = 1
	p.X.Max = 50
	p.Y.Min = 1
	p.Y.Max = 50

	// Save the plot to a PNG file.
	if err := p.Save(10*vg.Inch, 10*vg.Inch, "functions.png"); err != nil {
		panic(err)
	}
}
