package dbldbl

import "math"

// NaN returns a “not-a-number” value.
func NaN() Number {
	return Number{y: math.NaN()}
}

// Inf returns positive infinity if sign >= 0, negative infinity if sign < 0.
func Inf(sign int) Number {
	return Number{y: math.Inf(sign)}
}

var (
	E   = Number{2.718281828459045, +0x1.4d57ee2b1013ap-53} // https://oeis.org/A001113
	Pi  = Number{3.141592653589793, +0x1.1a62633145c07p-53} // https://oeis.org/A000796
	Phi = Number{1.618033988749895, -0x1.f506319fcfd18p-55} // https://oeis.org/A001622

	Sqrt2   = Number{1.4142135623730951, -0x1.bdd3413b26456p-54} // https://oeis.org/A002193
	SqrtE   = Number{1.6487212707001282, -0x1.b4690082a4906p-55} // https://oeis.org/A019774
	SqrtPi  = Number{1.772453850905516, -0x1.618f13eb7ca89p-54}  // https://oeis.org/A002161
	SqrtPhi = Number{1.272019649514069, -0x1.6daabafeede14p-56}  // https://oeis.org/A139339

	Ln2  = Number{0.6931471805599453, +0x1.abc9e3b39804p-56} // https://oeis.org/A002162
	Ln10 = Number{2.302585092994046, -0x1.f48ad494ea3e9p-53} // https://oeis.org/A002392

	twoPi   = Number{Pi.y * 2, Pi.x * 2}
	halfPi  = Number{Pi.y / 2, Pi.x / 2}
	twoOfPi = Number{0.6366197723675814, -0x1.6b01ec5417056p-55} // https://oeis.org/A060294
)
