package dbldbl

import "testing"

func TestPow(t *testing.T) {
	tests := []struct {
		arg0 Number
		arg1 Number
		want string
	}{
		{Pi, E, "22.45915771836104547342715220454373502758931513399669224920"}, // https://oeis.org/A059850
		{Pi, Phi, "6.3739021423033946516716479087752519831551446213168274134"}, // https://oeis.org/A182549
		{Float(2), Float(+0.5), "1.41421356237309504880168872420969807856967"}, // https://oeis.org/A002193
		{Float(2), Float(-0.5), "0.70710678118654752440084436210484903928483"}, // https://oeis.org/A010503
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Pow(tt.arg0, tt.arg1); !near(got, tt.want) {
				t.Errorf("Pow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPow_specials(t *testing.T) {
	tests := []struct {
		arg0 Number
		arg1 Number
		want Number
	}{
		{Inf(-1), Neg(Pi), Number{}},
		{Inf(-1), Float(-3), Float(-zero)},
		{Inf(-1), Float(-zero), Float(1)},
		{Inf(-1), Number{}, Float(1)},
		{Inf(-1), Float(1), Inf(-1)},
		{Inf(-1), Float(3), Inf(-1)},
		{Inf(-1), Pi, Inf(1)},
		{Inf(-1), Float(0.5), Inf(1)},
		{Inf(-1), NaN(), NaN()},
		{Neg(Pi), Inf(-1), Number{}},
		{Neg(Pi), Neg(Pi), NaN()},
		{Neg(Pi), Float(-zero), Float(1)},
		{Neg(Pi), Number{}, Float(1)},
		{Neg(Pi), Float(1), Neg(Pi)},
		{Neg(Pi), Pi, NaN()},
		{Neg(Pi), Inf(1), Inf(1)},
		{Neg(Pi), NaN(), NaN()},
		{Float(-1), Inf(-1), Float(1)},
		{Float(-1), Inf(1), Float(1)},
		{Float(-1), NaN(), NaN()},
		{Float(-0.5), Inf(-1), Inf(1)},
		{Float(-0.5), Inf(1), Number{}},
		{Float(-zero), Inf(-1), Inf(1)},
		{Float(-zero), Neg(Pi), Inf(1)},
		{Float(-zero), Float(-0.5), Inf(1)},
		{Float(-zero), Float(-3), Inf(-1)},
		{Float(-zero), Float(3), Float(-zero)},
		{Float(-zero), Pi, Number{}},
		{Float(-zero), Float(0.5), Number{}},
		{Float(-zero), Inf(1), Number{}},
		{Number{}, Inf(-1), Inf(1)},
		{Number{}, Neg(Pi), Inf(1)},
		{Number{}, Float(-3), Inf(1)},
		{Number{}, Float(-zero), Float(1)},
		{Number{}, Number{}, Float(1)},
		{Number{}, Float(3), Number{}},
		{Number{}, Pi, Number{}},
		{Number{}, Inf(1), Number{}},
		{Number{}, NaN(), NaN()},
		{Float(0.5), Inf(-1), Inf(1)},
		{Float(0.5), Inf(1), Number{}},
		{Float(1), Inf(-1), Float(1)},
		{Float(1), Inf(1), Float(1)},
		{Float(1), NaN(), Float(1)},
		{Pi, Inf(-1), Number{}},
		{Pi, Float(-zero), Float(1)},
		{Pi, Number{}, Float(1)},
		{Pi, Float(1), Pi},
		{Pi, Inf(1), Inf(1)},
		{Pi, NaN(), NaN()},
		{Inf(1), Neg(Pi), Number{}},
		{Inf(1), Float(-zero), Float(1)},
		{Inf(1), Number{}, Float(1)},
		{Inf(1), Float(1), Inf(1)},
		{Inf(1), Pi, Inf(1)},
		{Inf(1), NaN(), NaN()},
		{NaN(), Neg(Pi), NaN()},
		{NaN(), Float(-zero), Float(1)},
		{NaN(), Number{}, Float(1)},
		{NaN(), Float(1), NaN()},
		{NaN(), Pi, NaN()},
		{NaN(), NaN(), NaN()},

		{Float(2), Float(+0x1p32), Inf(1)},
		{Float(2), Float(-0x1p32), Number{}},
		{Float(0.5), Float(+0x1p45), Number{}},
		{Float(0.5), Float(-0x1p45), Inf(1)},
		{Float(-2), AddFloats(0x1p45, 1), Inf(-1)},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Pow(tt.arg0, tt.arg1); !same(got, tt.want) {
				t.Errorf("Pow() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestPow10_specials(t *testing.T) {
	tests := []struct {
		arg  int
		want Number
	}{
		{309, Inf(1)},
		{-309, Number{}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Pow10(tt.arg); !same(got, tt.want) {
				t.Errorf("Pow10() = %#v, want %#v", got, tt.want)
			}
		})
	}
}
