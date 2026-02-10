package dbldbl

import "math"

// Sinh returns the hyperbolic sine of n (approximate).
func Sinh(n Number) Number {
	switch {
	case n.y == 0:
		return n
	case n.y < 0:
		return Neg(Sinh(Neg(n)))
	}
	if n = Expm1(n); !isFinite(n.y) {
		return n
	}
	t := scalb(n, -1)
	return Add(t, Div(t, AddFloat(n, 1)))
}

// Cosh returns the hyperbolic cosine of n (approximate).
func Cosh(n Number) Number {
	t := Exp(Abs(n))
	return Add(scalb(t, -1), scalb(Inv(t), -1))
}

// Tanh returns the hyperbolic tangent of n (approximate).
func Tanh(n Number) Number {
	switch {
	case n.y == 0:
		return n
	case math.Abs(n.y) > 100:
		return Number{y: math.Copysign(1, n.y)}
	}
	t := Expm1(scalb(n, 1))
	return Div(t, AddFloat(t, 2))
}

// Asinh returns the inverse hyperbolic sine of n (approximate).
func Asinh(n Number) Number {
	switch {
	case n.y == 0 || !isFinite(n.y):
		return n
	case n.y < 0:
		return Neg(Asinh(Neg(n)))
	case n.y > 100:
		return Add(Log(n), Ln2)
	}

	t := Sqr(n)
	return Log1p(Add(n, Div(t, AddFloat(Sqrt(AddFloat(t, 1)), 1))))
}

// Acosh returns the inverse hyperbolic cosine of n (approximate).
func Acosh(n Number) Number {
	switch {
	case n.y < 1:
		return NaN()
	case !isFinite(n.y):
		return n
	case n.y > 100:
		return Add(Log(n), Ln2)
	}

	t := AddFloat(n, -1)
	return Log1p(Add(t, Sqrt(Add(scalb(t, 1), Sqr(t)))))
}

// Atanh returns the inverse hyperbolic tangent of n (approximate).
func Atanh(n Number) Number {
	switch {
	case n.y == 0:
		return n // ±0
	case IsNaN(n):
		return n
	case n.y == 1 && n.x == 0:
		return Inf(1) // atanh(1) = +Inf
	case n.y == -1 && n.x == 0:
		return Inf(-1) // atanh(-1) = -Inf
	}

	// Check if |n| > 1
	absN := Abs(n)
	if absN.y > 1 || (absN.y == 1 && absN.x > 0) {
		return NaN() // atanh(n) = NaN if |n| > 1
	}

	// For better accuracy: atanh(n) = log1p(2n/(1-n)) / 2
	return scalb(Log1p(Div(scalb(n, 1), SubFloat(1, n))), -1)
}
