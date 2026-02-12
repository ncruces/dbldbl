package dbldbl

import "math"

// Sincos returns Sin(n), Cos(n) (approximate).
func Sincos(n Number) (sin, cos Number) {
	switch {
	case n.y == 0:
		return n, Float(1)
	case !isFinite(n.y):
		return NaN(), NaN()
	}

	// Range reduction modulo π/2.
	k := Round(Mul(n, twoOfPi))
	t := Sub(n, Mul(k, halfPi))

	// Halve the angle until it is less than 2⁻⁵³.
	var halvings int8
	if _, e := math.Frexp(t.y); t.y != 0 && e > -53 {
		halvings = int8(53 + e)
	}

	// For |θ|<2⁻⁵³ these are accurate to 107 bits.
	sin = scalb(t, -halvings) // sin(θ) ≈ θ
	cos = Float(1)            // cos(θ) ≈ 1

	// Double-angle formulae.
	for range halvings {
		s, c := sin, cos
		sin = scalb(Mul(s, c), 1)           // sin(2⋅t) = 2⋅sin(θ)⋅cos(θ)
		cos = SubFloat(1, scalb(Sqr(s), 1)) // cos(2⋅t) = 1 - 2⋅sin²(θ)
	}

	yi, _ := math.Modf(k.y)
	xi, _ := math.Modf(k.x)
	switch (int64(xi) | int64(yi)) & 3 {
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

// Sin returns the sine of the radian argument n (approximate).
func Sin(n Number) Number {
	sin, _ := Sincos(n)
	return sin
}

// Cos returns the cosine of the radian argument n (approximate).
func Cos(n Number) Number {
	_, cos := Sincos(n)
	return cos
}

// Tan returns the tangent of the radian argument n (approximate).
func Tan(n Number) Number {
	sin, cos := Sincos(n)
	return Div(sin, cos)
}

// Asin returns the arcsine, in radians, of n (approximate).
func Asin(n Number) Number {
	// asin(θ) = atan2(θ, √(1-θ²))
	return Atan2(n, Sqrt(SubFloat(1, Sqr(n))))
}

// Acos returns the arccosine, in radians, of n (approximate).
func Acos(n Number) Number {
	// acos(θ) = atan2(√(1-θ²), θ)
	return Atan2(Sqrt(SubFloat(1, Sqr(n))), n)
}

// Atan returns the arctangent, in radians, of n (approximate).
func Atan(n Number) Number {
	switch {
	case n.y == 0:
		return n
	case n.y < 0:
		// atan(θ) = -atan(-θ)
		return Neg(Atan(Neg(n)))
	case n.y > 1:
		// atan(θ) = π/2 - atan(1/θ), if θ>0
		return Sub(halfPi, Atan(Inv(n)))
	case IsNaN(n):
		return NaN()
	}

	// Reduce the argument until it is less than 2⁻⁵³.
	var doublings int8
	for math.Abs(n.y) > 0x1p-53 {
		// atan(θ) = 2·atan(θ/(1+√(1+θ²)))
		n = Div(n, AddFloat(Sqrt(AddFloat(Sqr(n), 1)), 1))
		doublings++
	}

	// For |θ|<2⁻⁵³ this is accurate to 107 bits
	return scalb(n, doublings) // atan(θ) ≈ θ
}

// Atan2 returns the arc tangent of y/x, using the signs of the two
// to determine the quadrant of the return value (approximate).
func Atan2(y, x Number) Number {
	switch {
	case y.y == 0 && x.y == 0:
		switch {
		case !Signbit(x):
			return y
		case !Signbit(y):
			return Pi
		default:
			return Neg(Pi)
		}
	case IsInf(y, 0) && IsInf(x, 0):
		y = Number{y: math.Copysign(1, y.y)}
		x = Number{y: math.Copysign(1, x.y)}
	}

	z := Atan(Div(y, x))

	switch {
	case !Signbit(x):
		return z
	case !Signbit(y):
		return Add(z, Pi)
	default:
		return Sub(z, Pi)
	}
}
