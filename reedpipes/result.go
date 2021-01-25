package reedpipes

import (
	"fmt"
	"math"
)

var (
	r0 float64
	r5 float64
	r10 float64
	r15 float64
	r20 float64
	n float64
	vector = []float64{0, 0, 0, 0, 0}
)

func spline() []float64 {
	var res []float64

	ordinate := []float64{r0, r5, r10, r15, r20}
	abscissa := []float64{0, 5, 10, 15, 20}
	A := 6 * (r10 - 2 * r5 + r0) / 50
	B := 6 * (r15 - 2 * r10 + r5) / 50
	C := 6 * (r20 - 2 * r15 + r10) / 50
	vector[2] = (B - (A + C) / 4) * 4 / 7
	vector[1] = A / 2 - 0.25 * vector[2]
	vector[3] = C / 2 - 0.25 * vector[2]

	for d := 0.0; d < n; d++ {
		x := 20 / (n - 1) * d
		i := (int)((x - 0.01) / 5) + 1

		result := - vector[i - 1] / 30.0 * math.Pow(x - abscissa[i], 3.0) +
			      	vector[i] / 30.0 * math.Pow(x - abscissa[i - 1], 3.0) -
			      	(ordinate[i - 1] / 5.0 - 5.0 / 6.0 * vector[i - 1]) *
			      	(x - abscissa[i]) +
			      	(ordinate[i] / 5.0 - 5.0 / 6.0 * vector[i]) *
			      	(x - abscissa[i - 1])
		res = append(res, result)
	}
	return res
}

func displayResult(res []float64) {
	fmt.Printf("vector result: [")
	for i := 0; i < len(vector); i++ {
		if (vector[i] > 0 && vector[i] < 0.1) || (vector[i] > -0.1 && vector[i] < 0) {
			fmt.Printf("0.0")
		} else {
			fmt.Printf("%.1f", vector[i])
		}
		if i + 1 < len(vector) {
			fmt.Printf(", ")
		}
	}
	fmt.Printf("]\n");
	for i := 0.0; i < n; i++ {
		fmt.Printf("abscissa: %.1f cm\tradius: %.1f cm\n", 20 / (n - 1) * i, res[int(i)])
	}
}

// Main - 308reedpipes main
func Main() error {
	res := spline()
	displayResult(res)
	return nil
}