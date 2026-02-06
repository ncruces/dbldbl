package dbldbl

import "testing"

func TestSincos(t *testing.T) {
	t.Log(Sincos(Div(Pi, Float(6))))

	t.Log(Sincos(Ldexp(Pi, -2)))
	t.Log(Inv(Sqrt2))
}
