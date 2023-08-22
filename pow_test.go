package dbldbl

import (
	"fmt"
	"math"
	"strings"
	"testing"
)

var negZero = Float(math.Copysign(0, -1))

func TestPow(t *testing.T) {
	tests := []struct {
		arg0 Number
		arg1 Number
		want string
	}{
		{Pi, E, "22.4591577183610454734271522045437350275893151339966922492030025"}, // https://oeis.org/A059850
		{Pi, Phi, "6.373902142303394651671647908775251983155144621316827413455420"}, // https://oeis.org/A182549
		{Float(2), Float(+0.5), "1.4142135623730950488016887242096980785696718753"}, // https://oeis.org/A182549
		{Float(2), Float(-0.5), "0.7071067811865475244008443621048490392848359376"}, // https://oeis.org/A182549
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Pow(tt.arg0, tt.arg1); !strings.HasPrefix(tt.want, fmt.Sprint(got)[:30]) {
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
		{Inf(-1), Float(-3), Number{}},
		{Inf(-1), negZero, Float(1)},
		{Inf(-1), Number{}, Float(1)},
		{Inf(-1), Float(1), Inf(-1)},
		{Inf(-1), Float(3), Inf(-1)},
		{Inf(-1), Pi, Inf(1)},
		{Inf(-1), Float(0.5), Inf(1)},
		{Inf(-1), NaN(), NaN()},
		{Neg(Pi), Inf(-1), Number{}},
		{Neg(Pi), Neg(Pi), NaN()},
		{Neg(Pi), negZero, Float(1)},
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
		{negZero, Inf(-1), Inf(1)},
		{negZero, Neg(Pi), Inf(1)},
		{negZero, Float(-0.5), Inf(1)},
		{negZero, Float(-3), Inf(-1)},
		{negZero, Float(3), negZero},
		{negZero, Pi, Number{}},
		{negZero, Float(0.5), Number{}},
		{negZero, Inf(1), Number{}},
		{Number{}, Inf(-1), Inf(1)},
		{Number{}, Neg(Pi), Inf(1)},
		{Number{}, Float(-3), Inf(1)},
		{Number{}, negZero, Float(1)},
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
		{Pi, negZero, Float(1)},
		{Pi, Number{}, Float(1)},
		{Pi, Float(1), Pi},
		{Pi, Inf(1), Inf(1)},
		{Pi, NaN(), NaN()},
		{Inf(1), Neg(Pi), Number{}},
		{Inf(1), negZero, Float(1)},
		{Inf(1), Number{}, Float(1)},
		{Inf(1), Float(1), Inf(1)},
		{Inf(1), Pi, Inf(1)},
		{Inf(1), NaN(), NaN()},
		{NaN(), Neg(Pi), NaN()},
		{NaN(), negZero, Float(1)},
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
				t.Errorf("Pow() = %v, want %v", got, tt.want)
			}
		})
	}
}
