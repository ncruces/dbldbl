package dbldbl

import "math"

// Sinh returns the hyperbolic sine of n (approximate).
func Sinh(n Number) Number {
	t := Expm1(n)
	if t.y == 0 || !isFinite(t.y) {
		return Number{y: math.Copysign(t.y, n.y)}
	}
	t2 := scalb(t, -1)
	return Add(t2, Div(t2, AddFloat(t, 1)))
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
	case n.y == 0:
		return n // ±0
	case IsNaN(n):
		return n
	case IsInf(n, 0):
		return n // ±Inf
	}

	// For negative n: asinh(-n) = -asinh(n)
	if n.y < 0 {
		return Neg(Asinh(Neg(n)))
	}

	// For better accuracy: asinh(n) = log1p(n + n²/(√(n²+1)+1))
	n2 := Sqr(n)
	return Log1p(Add(n, Div(n2, AddFloat(Sqrt(AddFloat(n2, 1)), 1))))
}

// Acosh returns the inverse hyperbolic cosine of n (approximate).
func Acosh(n Number) Number {
	switch {
	case n.y < 1:
		return NaN() // acosh(n) = NaN if n < 1
	case n.y == 1 && n.x == 0:
		return Number{} // acosh(1) = 0
	case IsInf(n, 1):
		return n // acosh(+Inf) = +Inf
	case IsNaN(n):
		return n
	}

	// For better accuracy: acosh(n) = log1p((n-1) + √((n-1)·(n+1)))
	t := AddFloat(n, -1)
	return Log1p(Add(t, Sqrt(Mul(t, AddFloat(n, 1)))))
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
