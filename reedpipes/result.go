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
)

func getP(i float64, iD float64, a float64, b float64, c float64, x []float64, y []float64, f []float64) float64 {
	//first := -f[int(i) - 1] / 30 * math.Pow(iD - x[int(i)], 3)
	//second := f[int(i)] / 30 * math.Pow(iD - x[int(i) - 1], 3)
	//third := (y[int(i)-1]/5 - 5/6*f[int(i)-1]) * (iD-x[int(i)])
	//fourth := (y[int(i)]/5 - 5/6*f[int(i)]) * (iD-x[int(i)-1])
	//res := first + second - third + fourth
	//return res

	res := (-f[int(i) - 1] / 30 * math.Pow(iD - x[int(i)], 3) + f[int(i)] / 30 * math.Pow(iD - x[int(i) - 1], 3) - (y[int(i)-1] / 5 - 5/ 6 * f[int(i)-1]) * (iD-x[int(i)]) + (y[int(i)]/5 - 5 / 6 * f[int(i)]) * (iD-x[int(i)-1]))
	return res
}

func spline() {
	X := []float64{0, 5, 10, 15, 20}
	Y := []float64{r0, r5, r10, r15, r20}
	A := 6 * (r10 - 2 * r5 + r0) / 50
	B := 6 * (r15 - 2 * r10 + r5) / 50
	C := 6 * (r20 - 2 * r15 + r10) / 50
	f := []float64{0, 0, (B -(A + C) / 4) * 4 / 7, 0, 0}
	f[3] = C / 2 - 0.25 * f[2]
	f[1] = A / 2 - 0.25 * f[2]

	fmt.Printf("vector result: [%.1f, %.1f, %.1f, %.1f, %.1f]\n", f[0], f[1], f[2], f[3], f[4])

	for i := 0; i < int(n - 1); i++ {
		t := 20 / (n - 1) * float64(i)
		radius := getP(t/5+1, t, A, B, C, X, Y, f)
		fmt.Printf("abscissa: %.1f cm\tradius: %.1f cm\n", t, radius)
	}
	radius := getP(4, 20, A, B, C, X, Y, f)
	fmt.Printf("abscissa: 20.0 cm\tradius: %.1f cm\n", radius)
}

func displayResult() {

}

// Main - 308reedpipes main
func Main() error {
	spline()
	displayResult()
	return nil
}