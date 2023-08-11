package dbldbl

import "math"

func twoSumQuick(x, y float64) Number {
	// |x| > |y|
	r := x + y
	e := y - (r - x)
	return Number{r, e}
}

func twoSum(x, y float64) Number {
	r := x + y
	t := r - x
	e := (x - (r - t)) + (y - t)
	return Number{r, e}
}

func twoDiff(x, y float64) Number {
	r := x - y
	t := r - x
	e := (x - (r - t)) - (y + t)
	return Number{r, e}
}

func twoProd(x, y float64) Number {
	r := x * y
	e := math.FMA(x, y, -r)
	return Number{r, e}
}
