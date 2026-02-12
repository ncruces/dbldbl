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
	E   = Number{2.718281828459045, 1.4456468917292502e-16} // https://oeis.org/A001113
	Pi  = Number{3.141592653589793, 1.2246467991473532e-16} // https://oeis.org/A000796
	Phi = Number{1.618033988749895, -5.432115203682506e-17} // https://oeis.org/A001622

	Sqrt2   = Number{1.4142135623730951, -9.667293313452913e-17} // https://oeis.org/A002193
	SqrtE   = Number{1.6487212707001282, -4.731568479435833e-17} // https://oeis.org/A019774
	SqrtPi  = Number{1.772453850905516, -7.666586499825799e-17}  // https://oeis.org/A002161
	SqrtPhi = Number{1.272019649514069, -1.982284322743343e-17}  // https://oeis.org/A139339

	Ln2  = Number{0.6931471805599453, 2.3190468138462996e-17} // https://oeis.org/A002162
	Ln10 = Number{2.302585092994046, -2.1707562233822494e-16} // https://oeis.org/A002392

	twoPi   = Number{Pi.y * 2, Pi.x * 2}
	halfPi  = Number{Pi.y / 2, Pi.x / 2}
	twoOfPi = Number{0.6366197723675814, -3.935735335036497e-17} // https://oeis.org/A060294
)
