package dbldbl

import "math"

// Neg negates n (exact).
func Neg(n Number) Number {
	return Number{y: -n.y, x: 0 - n.x}
}

// Abs returns the absolute value of n (exact).
func Abs(n Number) Number {
	if math.Signbit(n.y) {
		return Neg(n)
	}
	return n
}

// IsNaN reports whether n is a “not-a-number” value.
func IsNaN(n Number) bool {
	return math.IsNaN(n.y)
}

// IsInf reports whether f is an infinity, according to sign.
// If sign > 0, IsInf reports whether f is positive infinity.
// If sign < 0, IsInf reports whether f is negative infinity.
// If sign == 0, IsInf reports whether f is either infinity.
func IsInf(n Number, sign int) bool {
	return math.IsInf(n.y, sign)
}

// Trunc returns the integer value of n (exact).
func Trunc(n Number) Number {
	y := math.Trunc(n.y)
	switch {
	case y != n.y:
		return Number{y: y}
	case y < 0:
		return Number{y: y, x: math.Ceil(n.x)}
	default:
		return Number{y: y, x: math.Floor(n.x)}
	}
}

// Floor returns the greatest integer value less than or equal to n (exact).
func Floor(n Number) Number {
	y := math.Floor(n.y)
	if y != n.y {
		return Number{y: y}
	}
	return Number{y: y, x: math.Floor(n.x)}
}

// Ceil returns the least integer value greater than or equal to n (exact).
func Ceil(n Number) Number {
	y := math.Ceil(n.y)
	if y != n.y {
		return Number{y: y}
	}
	return Number{y: y, x: math.Ceil(n.x)}
}

// Shift returns the product of n by 2ⁱ (exact).
func Shift(n Number, i int8) Number {
	e := math.Float64frombits((1023 + uint64(i)) << 52)
	return Number{y: e * n.y, x: e * n.x}
}

// AddFloats returns the sum of a and b (exact).
func AddFloats(a, b float64) Number {
	s := twoSum(a, b)
	if !isFinite(s.y) {
		return Number{y: s.y}
	}
	return s
}

// AddFloat returns the sum of a and b (exactly rounded).
func AddFloat(a Number, b float64) Number {
	s := twoSum(a.y, b)
	if !isFinite(s.y) {
		return Number{y: s.y}
	}
	return twoSumQuick(s.y, s.x+a.x)
}

// Add returns the sum of a and b (exactly rounded).
func Add(a, b Number) Number {
	s := twoSum(a.y, b.y)
	if !isFinite(s.y) {
		return Number{y: s.y}
	}
	t := twoSum(a.x, b.x)
	s = twoSumQuick(s.y, s.x+t.y)
	s = twoSumQuick(s.y, s.x+t.x)
	return s
}

// SubFloats returns the difference of a and b (exact).
func SubFloats(a, b float64) Number {
	s := twoDiff(a, b)
	if !isFinite(s.y) {
		return Number{y: s.y}
	}
	return s
}

// SubFloat returns the difference of a and b (exactly rounded).
func SubFloat(a Number, b float64) Number {
	s := twoDiff(a.y, b)
	if !isFinite(s.y) {
		return Number{y: s.y}
	}
	return twoSumQuick(s.y, s.x+a.x)
}

// Sub returns the difference of a and b (exactly rounded).
func Sub(a, b Number) Number {
	s := twoDiff(a.y, b.y)
	if !isFinite(s.y) {
		return Number{y: s.y}
	}
	t := twoDiff(a.x, b.x)
	s = twoSumQuick(s.y, s.x+t.y)
	s = twoSumQuick(s.y, s.x+t.x)
	return s
}

// MulFloats returns the product of a and b (exact).
func MulFloats(a, b float64) Number {
	s := twoProd(a, b)
	if !isFinite(s.y) {
		return Number{y: s.y}
	}
	return s
}

// MulFloat returns the product of a and b (exactly rounded).
func MulFloat(a Number, b float64) Number {
	s := twoProd(a.y, b)
	if !isFinite(s.y) {
		return Number{y: s.y}
	}
	s.x = math.FMA(a.x, b, s.x)
	return twoSumQuick(s.y, s.x)
}

// Mul returns the product of a and b (approximate).
func Mul(a, b Number) Number {
	s := twoProd(a.y, b.y)
	if !isFinite(s.y) {
		return Number{y: s.y}
	}
	s.x = math.FMA(a.x, b.x, s.x)
	s.x = math.FMA(a.y, b.x, s.x)
	s.x = math.FMA(a.x, b.y, s.x)
	return twoSumQuick(s.y, s.x)
}

// Div returns the quotient of a and b (approximate).
func Div(a, b Number) Number {
	y := a.y / b.y
	if y == 0 || !isFinite(y) {
		return Number{y: y}
	}
	t := twoProd(y, b.y)
	x := (a.y - t.y - t.x + a.x - y*b.x) / b.y
	return twoSumQuick(y, x)
}

// Sqr returns the square of n (approximate).
func Sqr(n Number) Number {
	s := twoProd(n.y, n.y)
	if !isFinite(s.y) {
		return Number{y: s.y}
	}
	s.x = math.FMA(n.x, n.x, s.x)
	s.x = math.FMA(n.y, n.x+n.x, s.x)
	return twoSumQuick(s.y, s.x)
}

// Sqrt returns the square root of n (approximate).
func Sqrt(n Number) Number {
	y := math.Sqrt(n.y)
	if y == 0 || !isFinite(y) {
		return Number{y: y}
	}
	// Newton's method: y + (n-y²)/2y
	t := twoProd(y, y)
	x := (n.y - t.y - t.x + n.x) * 0.5 / y
	return twoSumQuick(y, x)
}

// Cbrt returns the cube root of n (approximate).
func Cbrt(n Number) Number {
	y := math.Cbrt(n.y)
	if y == 0 || !isFinite(y) {
		return Number{y: y}
	}
	// Newton's method: y + (n/y²-y)/3
	t := Div(n, twoProd(y, y))
	x := (t.y - y + t.x) / 3
	return twoSumQuick(y, x)
}

// FMAFloat returns a⋅b + c (exactly rounded).
func FMAFloat(a, b float64, c Number) Number {
	s := twoFMA(a, b, c)
	if !isFinite(s.y) {
		return Number{y: math.FMA(a, b, c.y)}
	}
	return s
}

// FMA returns a⋅b + c (approximate).
func FMA(a, b, c Number) Number {
	s := twoFMA(a.y, b.y, c)
	if !isFinite(s.y) {
		return Number{y: math.FMA(a.y, b.y, c.y)}
	}
	s = twoFMAQuick(a.x, b.x, s)
	s = twoFMAQuick(a.y, b.x, s)
	s = twoFMAQuick(a.x, b.y, s)
	return s
}

// Cmp compares x and y and returns:
//
//	-1 if x <  y (incl. NaN < !NaN)
//	 0 if x == y (incl. -0 == 0, -Inf == -Inf, +Inf == +Inf, NaN == NaN)
//	+1 if x >  y (incl. !NaN > NaN)
func Cmp(a, b Number) int {
	switch {
	case a.y < b.y:
		return -1
	case a.y > b.y:
		return +1
	case a.x < b.x:
		return -1
	case a.x > b.x:
		return +1
	case a.y == b.y:
		return 0
	case !IsNaN(b):
		return -1
	case !IsNaN(a):
		return +1
	default:
		return 0
	}
}

func isFinite(x float64) bool {
	return math.Float64bits(x)>>52&0x7ff != 0x7ff
}
