package dbldbl

import "math"

// Sincos returns Sin(n), Cos(n) (approximate).
func Sincos(n Number) (sin, cos Number) {
	switch {
	case n.y == 0:
		return n, Float(1) // return ±0.0, 1.0
	case IsNaN(n) || IsInf(n, 0):
		return NaN(), NaN()
	}

	// Range reduction modulo π/2.
	k := math.Round(n.y / (math.Pi * 0.5))
	r := Sub(n, MulFloat(scalb(Pi, -1), k))

	// Halve the angle until it is less than 2⁻⁵³.
	var halvings int8
	if _, e := math.Frexp(r.y); e > -53 {
		halvings = int8(53 + e)
	}

	// These approximations are good enough for 107-bits of precision.
	sin = scalb(r, -halvings) // sin(r) ≈ r
	cos = Float(1)            // cos(r) ≈ 1

	// Angle doubling.
	// The cosine formula offers the best numeric stability.
	for range halvings {
		s, c := sin, cos
		sin = scalb(Mul(s, c), 1)                // sin(2⋅t) = 2⋅sin(t)⋅cos(t)
		cos = Neg(SubFloat(scalb(Sqr(s), 1), 1)) // cos(2⋅t) = 1 - 2⋅sin²(t)
	}

	switch int64(k) & 3 {
	default:
		return sin, cos
	case 1:
		return cos, Neg(sin)
	case 2:
		return Neg(sin), Neg(cos)
	case 3:
		return Neg(cos), sin
	}
}

// Sin returns the sine of the radian argument n.
func Sin(n Number) Number {
	sin, _ := Sincos(n)
	return sin
}

// Cos returns the cosine of the radian argument n.
func Cos(n Number) Number {
	_, cos := Sincos(n)
	return cos
}

// Tan returns the tangent of the radian argument n.
func Tan(n Number) Number {
	sin, cos := Sincos(n)
	return Div(sin, cos)
}
