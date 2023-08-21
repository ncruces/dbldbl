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
		return Number{y: n.x}
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

	// log(1/x) = π / 2⋅AGM(1, 4⋅x)
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
	// Newton's method: y + (n-log(y))⋅y
	t := Sub(n, Log(Float(y)))
	return twoSumQuick(y, t.y*y)
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
	}
}

func pi() Number {
	// https://en.wikipedia.org/wiki/Arithmetic%E2%80%93geometric_mean
	a := Float(1)
	g := Sqrt(Float(0.5))
	series := Float(1)

	for n := int8(1); n < 127; n++ {
		t := Shift(Add(a, g), -1)
		if a == t {
			break
		}
		g = Sqrt(Mul(a, g))
		a = t
		series = Sub(series, Shift(Sub(Sqr(a), Sqr(g)), n+1))
	}

	return Shift(Div(Sqr(a), series), 2)
}
