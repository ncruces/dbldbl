package dbldbl

// Number is a double-double precision number.
type Number struct {
	y, x float64
}

// Float creates a Number from a float64.
func Float(a float64) Number {
	return Number{y: a}
}

// Int creates a Number from an int64.
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

// Int converts this Number to an int64.
func (n Number) Int() int64 {
	return int64(n.y) + int64(n.x)
}

// Uint converts this Number to a uint64.
func (n Number) Uint() uint64 {
	return uint64(n.y) + uint64(n.x)
}
