package dbldbl

import "math"

// Sinh returns the hyperbolic sine of n (approximate).
func Sinh(n Number) Number {
	switch {
	case n.y == 0:
		return n // ±0
	case IsNaN(n):
		return n
	case IsInf(n, 0):
		return n // ±Inf
	}

	// For small |n| use Expm1 for better accuracy.
	// sinh(n) = (expm1(n) + expm1(n)/(expm1(n)+1)) / 2
	if math.Abs(n.y) < 0.5 {
		t := Expm1(n)
		return scalb(Add(t, Div(t, AddFloat(t, 1))), -1)
	}

	// For large |n| use (e^n - e^(-n)) / 2
	ep := Exp(n)
	en := Exp(Neg(n))
	return scalb(Sub(ep, en), -1)
}

// Cosh returns the hyperbolic cosine of n (approximate).
func Cosh(n Number) Number {
	switch {
	case n.y == 0:
		return Float(1) // cosh(±0) = 1
	case IsNaN(n):
		return n
	case IsInf(n, 0):
		return Inf(1) // cosh(±Inf) = +Inf
	}

	// cosh(n) = (e^n + e^(-n)) / 2
	ep := Exp(n)
	en := Exp(Neg(n))
	return scalb(Add(ep, en), -1)
}

// Tanh returns the hyperbolic tangent of n (approximate).
func Tanh(n Number) Number {
	switch {
	case n.y == 0:
		return n // ±0
	case IsNaN(n):
		return n
	case IsInf(n, 1):
		return Float(1) // tanh(+Inf) = +1
	case IsInf(n, -1):
		return Float(-1) // tanh(-Inf) = -1
	}

	// For better accuracy: tanh(n) = expm1(2n) / (expm1(2n) + 2)
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
