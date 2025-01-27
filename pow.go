package dbldbl

import "math"

// Pow returns bⁿ, the base-b exponential of n (approximate).
func Pow(b Number, n Number) Number {
	switch {
	case n.y == 0 || b == Float(1):
		return Float(1)
	case n == Float(1):
		return b
	case IsNaN(b) || IsNaN(n):
		return NaN()
	case b.y == 0:
		switch {
		case n.y < 0:
			if isOddInteger(n) {
				if math.Signbit(b.y) {
					return Inf(-1)
				}
				return Inf(1)
			}
			return Inf(1)
		case n.y > 0:
			if isOddInteger(n) {
				return b
			}
			return Number{}
		}
	case IsInf(n, 0):
		switch {
		case b == Float(-1):
			return Float(1)
		case (SubFloat(Abs(b), 1).y < 0) == IsInf(n, 1):
			return Number{}
		default:
			return Inf(1)
		}
	case IsInf(b, 0):
		if IsInf(b, -1) {
			return Pow(Float(1/b.y), Neg(n)) // Pow(-0, -n)
		}
		switch {
		case n.y < 0:
			return Number{}
		case n.y > 0:
			return Inf(1)
		}
	case n == Float(+0.5):
		return Sqrt(b)
	case n == Float(-0.5):
		return InvSqrt(b)
	case n == Floor(n):
		return pow(b, n)
	case b.y < 0:
		return NaN()
	}

	return Exp(Mul(Log(b), n))
}

func pow(b Number, n Number) Number {
	r := Float(1)
	i := Abs(n)
	for {
		if isOddInteger(i) {
			i = twoSumQuick(i.y, i.x-1)
			r = Mul(r, b)
		}
		if i.y == 0 {
			break
		}
		i = scalb(i, -1)
		b = Sqr(b)
	}
	if n.y < 0 {
		return Inv(r)
	}
	return r
}

func isOddInteger(n Number) bool {
	return isOddFloat(n.x) || isOddFloat(n.y)
}

func isOddFloat(x float64) bool {
	xi, xf := math.Modf(x)
	return xf == 0 && int64(xi)&1 == 1
}
