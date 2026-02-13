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
		if Signbit(n) {
			if isOddInteger(n) && Signbit(b) {
				return Inf(-1)
			}
			return Inf(0)
		} else {
			if isOddInteger(n) {
				return b
			}
			return Number{}
		}
	case IsInf(n, 0):
		switch {
		case b == Float(-1):
			return Float(1)
		case (SubFloat(1, Abs(b)).y > 0) == IsInf(n, 1):
			return Number{}
		default:
			return Inf(0)
		}
	case IsInf(b, 0):
		switch {
		case IsInf(b, -1):
			var zero float64
			return Pow(Float(-zero), Neg(n))
		case n.y < 0:
			return Number{}
		case n.y > 0:
			return Inf(0)
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

// Pow10 returns 10ⁱ, the base-10 exponential of i (approximate).
func Pow10(i int) Number {
	return pow(Float(10), Int(int64(i)))
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
		i = shift(i, -1)
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
