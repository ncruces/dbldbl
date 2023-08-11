package dbldbl

// Number is a double-double precision number.
type Number struct {
	y, x float64
}

// Float creates a Number from a float64.
func Float(a float64) Number {
	return Number{y: a}
}

// Int creates a Number from a int64.
func Int(a int64) Number {
	if y := a >> 48 << 48; y == 0 {
		return Number{y: float64(a)}
	} else {
		return twoSumQuick(float64(y), float64(a-y))
	}
}

// Uint creates a Number from a uint64.
func Uint(a uint64) Number {
	if y := a >> 48 << 48; y == 0 {
		return Number{y: float64(a)}
	} else {
		return twoSumQuick(float64(y), float64(a-y))
	}
}

// Float converts this Number to a float64.
func (n Number) Float() (_ float64, exact bool) {
	return n.y, n.x == 0
}

// Accum adds a float64 to this Number (approximate).
func (n *Number) Accum(a float64) {
	s := twoSum(n.y, a)
	*n = twoSumQuick(s.y, s.x+n.x)
}
