package dbldbl

import "math"

// Neg negates a Number (exact).
func Neg(n Number) Number {
	return Number{y: -n.y, x: -n.x}
}

// Abs returns the absolute value of n (exact).
func Abs(n Number) Number {
	if n.y >= 0 {
		return n
	}
	return Neg(n)
}

// Trunc returns the integer value of n (exact).
func Trunc(n Number) (rv Number) {
	return Number{y: math.Trunc(n.y), x: math.Trunc(n.x)}
}

// Shift returns the product of n by 2ⁱ, for −1023 < i ≤ +1023 (exact).
func Shift(n Number, i int16) Number {
	e := math.Float64frombits(uint64(1023+i) << 52)
	return Number{y: e * n.y, x: e * n.x}
}

// AddFloats returns the sum of a and b (exact).
func AddFloats(a, b float64) Number {
	return twoSum(a, b)
}

// AddFloat returns the sum of a and b (approximate).
func AddFloat(a Number, b float64) Number {
	s := twoSum(a.y, b)
	return twoSumQuick(s.y, s.x+a.x)
}

// Add returns the sum of a and b (approximate).
func Add(a, b Number) Number {
	s := twoSum(a.y, b.y)
	t := twoSum(a.x, b.x)
	s = twoSumQuick(s.y, s.x+t.y)
	s = twoSumQuick(s.y, s.x+t.x)
	return s
}

// SubFloats returns the difference of a and b (exact).
func SubFloats(a, b float64) Number {
	return twoDiff(a, b)
}

// SubFloat returns the difference of a and b (approximate).
func SubFloat(a Number, b float64) Number {
	s := twoDiff(a.y, b)
	return twoSumQuick(s.y, s.x+a.x)
}

// Sub returns the difference of a and b (approximate).
func Sub(a, b Number) Number {
	s := twoDiff(a.y, b.y)
	t := twoDiff(a.x, b.x)
	s = twoSumQuick(s.y, s.x+t.y)
	s = twoSumQuick(s.y, s.x+t.x)
	return s
}

// MulFloats returns the product of a and b (exact).
func MulFloats(a, b float64) Number {
	return twoProd(a, b)
}

// MulFloat returns the product of a and b (approximate).
func MulFloat(a Number, b float64) Number {
	t := twoProd(a.y, b)
	t.x = math.FMA(a.x, b, t.x)
	return twoSumQuick(t.y, t.x)
}

// Mul returns the product of a and b (approximate).
func Mul(a, b Number) Number {
	t := twoProd(a.y, b.y)
	t.x = math.FMA(a.x, b.x, t.x)
	t.x = math.FMA(a.y, b.x, t.x)
	t.x = math.FMA(a.x, b.y, t.x)
	return twoSumQuick(t.y, t.x)
}

// Div returns the quotient of a and b (approximate).
func Div(a, b Number) Number {
	y := a.y / b.y
	t := twoProd(y, b.y)
	x := (a.y - t.y - t.x + a.x - y*b.x) / b.y
	return twoSumQuick(y, x)
}

// Sqr returns the square of n (approximate).
func Sqr(n Number) Number {
	return Mul(n, n)
}

// Sqrt returns the square root of n (approximate).
func Sqrt(n Number) Number {
	y := math.Sqrt(n.y)
	t := twoProd(y, y)
	x := (n.y - t.y - t.x + n.x) * 0.5 / y
	return twoSumQuick(y, x)
}
