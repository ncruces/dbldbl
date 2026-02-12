package dbldbl

import (
	"math"
	"math/big"
	"testing"
)

func TestInt(t *testing.T) {
	tests := []struct {
		arg  int64
		want Number
	}{
		{+1e03 + 1, Number{y: +1e03 + 1}},
		{-1e06 + 1, Number{y: -1e06 + 1}},
		{+1e09 + 1, Number{y: +1e09 + 1}},
		{-1e12 + 1, Number{y: -1e12 + 1}},
		{+1e15 + 1, Number{y: +1e15 + 1}},
		{-1e18 + 1, Number{y: -1e18, x: +1}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if tt.arg != tt.want.Int() {
				t.Errorf("arg = %#v, want %#v", tt.arg, tt.want)
			}
			if got := Int(tt.arg); !same(got, tt.want) {
				t.Errorf("Int() = %#v, want %#v", got, tt.want)
			}
			if f, ok := tt.want.Float(); ok != same(Float(f), tt.want) {
				t.Errorf("float = %#v, want %#v", f, tt.want)
			}
		})
	}
}

func TestUint(t *testing.T) {
	tests := []struct {
		arg  uint64
		want Number
	}{
		{1e03 + 1, Number{y: 1e03 + 1}},
		{1e06 + 1, Number{y: 1e06 + 1}},
		{1e09 + 1, Number{y: 1e09 + 1}},
		{1e12 + 1, Number{y: 1e12 + 1}},
		{1e15 + 1, Number{y: 1e15 + 1}},
		{1e18 + 1, Number{y: 1e18, x: +1}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if tt.arg != tt.want.Uint() {
				t.Errorf("arg = %#v, want %#v", tt.arg, tt.want)
			}
			if got := Uint(tt.arg); !same(got, tt.want) {
				t.Errorf("Int() = %#v, want %#v", got, tt.want)
			}
			if f, ok := tt.want.Float(); ok != same(Float(f), tt.want) {
				t.Errorf("float = %#v, want %#v", f, tt.want)
			}
		})
	}
}

var zero float64

func same(a, b Number) bool {
	return IsNaN(a) && IsNaN(b) || (true &&
		math.Float64bits(a.y) == math.Float64bits(b.y) &&
		math.Float64bits(a.x) == math.Float64bits(b.x))

}

func near(a Number, b string) bool {
	const prec = 101
	xa := a.toBig().SetPrec(prec)
	xp, _, _ := big.ParseFloat(b, 0, prec, big.ToPositiveInf)
	xn, _, _ := big.ParseFloat(b, 0, prec, big.ToNegativeInf)
	return xp.Cmp(xa) == 0 || xn.Cmp(xa) == 0
}
