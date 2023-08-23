package dbldbl

import "math"

// Log returns the natural logarithm of n (approximate).
func Log(n Number) Number {
	switch {
	case n.y < 0:
		return NaN()
	case n.y == 0:
		return Inf(-1)
	case n.y == 1:
		return Number{y: math.Log1p(n.x)}
	case !isFinite(n.y):
		return n
	}

	negate := true
	if n.y > 1 {
		n = Div(Float(1), n)
		negate = false
	}

	var shift int8
	for n.y > 0x1p-54 {
		n = Sqr(n)
		shift--
	}

	// log(1/n) = π / 2⋅AGM(1, 4⋅n)
	n = Div(Pi, agm(Float(1), Shift(n, 2)))
	n = Shift(n, shift-1)
	if negate {
		n = Neg(n)
	}
	return n
}

// Exp returns eⁿ, the base-e exponential of n (approximate).
func Exp(n Number) Number {
	y := math.Exp(n.y)
	if y == 0 || !isFinite(y) {
		return Number{y: y}
	}
	// Newton's method: y + y⋅(n-log(y))
	t := Sub(n, Log(Float(y)))
	return twoSumQuick(y, y*t.y)
}

// Log1p returns the natural logarithm of 1 plus n (approximate).
// It is more accurate than Log(AddFloat(n, 1)) when n is near zero.
func Log1p(n Number) Number {
	// https://www.johndcook.com/blog/2012/07/25/trick-for-computing-log1x/
	u := AddFloat(n, 1)
	switch {
	case u.y < 0:
		return NaN()
	case u == Float(1):
		return n
	case !isFinite(u.y):
		return n
	}
	// log(1+n) = n⋅log(u)/(u-1), u=1+n
	return Div(Mul(n, Log(u)), SubFloat(u, 1))
}

// Expm1 returns eⁿ-1, the base-e exponential of n minus 1 (approximate).
// It is more accurate than SubFloat(Exp(n), 1) when n is near zero.
func Expm1(n Number) Number {
	y := math.Expm1(n.y)
	switch {
	case y == 0 || !isFinite(y):
		return Number{y: y}
	case y == -1:
		return SubFloat(Exp(n), 1)
	}
	// Newton's method: y + (y+1)⋅(n-log1p(y))
	t := Sub(n, Log1p(Float(y)))
	return twoSumQuick(y, (y+1)*t.y)
}

func agm(a, g Number) Number {
	// https://en.wikipedia.org/wiki/Arithmetic%E2%80%93geometric_mean
	for {
		t := Shift(Add(a, g), -1)
		if a == t {
			return a
		}
		g = Sqrt(Mul(a, g))
		a = t
		if a == g {
			return a
		}
	}
}
