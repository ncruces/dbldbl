package dbldbl

import "math"

func twoSumQuick(x, y float64) Number {
	// log₂|x| ≥ log₂|y|
	r := float64(x + y)
	e := y - float64(r-x)
	return Number{r, e}
}

func twoSum(x, y float64) Number {
	r := float64(x + y)
	t := float64(r - x)
	e := float64(x-float64(r-t)) + float64(y-t)
	return Number{r, e}
}

func twoDiff(x, y float64) Number {
	r := float64(x - y)
	t := float64(r - x)
	e := float64(x-float64(r-t)) - float64(y+t)
	return Number{r, e}
}

func twoProd(x, y float64) Number {
	r := float64(x * y)
	e := math.FMA(x, y, -r)
	return Number{r, e}
}

func twoFMA(x, y float64, n Number) Number {
	p := twoProd(x, y)
	s := twoSum(n.y, p.y)
	t := twoSum(n.x, p.x)
	s = twoSumQuick(s.y, s.x+t.y)
	s = twoSumQuick(s.y, s.x+t.x)
	return s
}

func twoFMAQuick(x, y float64, n Number) Number {
	p := twoProd(x, y)
	s := twoSum(n.y, p.y)
	return twoSumQuick(s.y, s.x+(n.x+p.x)) // approximation
}
