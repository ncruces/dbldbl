package dbldbl

import (
	"math"
	"math/big"
	"testing"
)

const prec = 104

func TestNeg(t *testing.T) {
	tests := []struct {
		arg  Number
		want Number
	}{
		{Number{}, Number{}},
		{Number{1, 0}, Number{-1, -0}},
		{Number{0, 1}, Number{-0, -1}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Neg(tt.arg); !same(got, tt.want) {
				t.Errorf("Neg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAbs(t *testing.T) {
	tests := []struct {
		arg  Number
		want Number
	}{
		{Number{}, Number{}},
		{Number{+1, 0}, Number{1, 0}},
		{Number{-1, 0}, Number{1, 0}},
		{Number{+1, 0.5}, Number{1, +0.5}},
		{Number{-1, 0.5}, Number{1, -0.5}},
		{Number{math.Inf(+1), 0}, Number{math.Inf(1), 0}},
		{Number{math.Inf(-1), 0}, Number{math.Inf(1), 0}},
		{Number{math.NaN(), 0}, Number{math.NaN(), 0}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Abs(tt.arg); !same(got, tt.want) {
				t.Errorf("Abs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsInf(t *testing.T) {
	tests := []struct {
		arg  Number
		want bool
	}{
		{Number{}, false},
		{Number{+1, 0}, false},
		{Number{-1, 0}, false},
		{Number{math.Inf(+1), 0}, true},
		{Number{math.Inf(-1), 0}, true},
		{Number{math.NaN(), 0}, false},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := IsInf(tt.arg, 0); got != tt.want {
				t.Errorf("IsInf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrunc(t *testing.T) {
	tests := []struct {
		arg  Number
		want Number
	}{
		{Number{}, Number{}},
		{Number{1, 0}, Number{1, 0}},
		{Number{0.5, 0}, Number{0, 0}},
		{Number{1.5, 0}, Number{1, 0}},
		{Number{10, -1.5}, Number{10, -2}},
		{Number{-10, 1.5}, Number{-10, 2}},
		{Number{math.Inf(1), 0}, Number{math.Inf(1), 0}},
		{Number{math.NaN(), 0}, Number{math.NaN(), 0}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Trunc(tt.arg); !same(got, tt.want) {
				t.Errorf("Trunc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloor(t *testing.T) {
	tests := []struct {
		arg  Number
		want Number
	}{
		{Number{}, Number{}},
		{Number{1, 0}, Number{1, 0}},
		{Number{0.5, 0}, Number{0, 0}},
		{Number{1.5, 0}, Number{1, 0}},
		{Number{10, -1.5}, Number{10, -2}},
		{Number{-10, 1.5}, Number{-10, 1}},
		{Number{math.Inf(1), 0}, Number{math.Inf(1), 0}},
		{Number{math.NaN(), 0}, Number{math.NaN(), 0}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Floor(tt.arg); !same(got, tt.want) {
				t.Errorf("Floor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCeil(t *testing.T) {
	tests := []struct {
		arg  Number
		want Number
	}{
		{Number{}, Number{}},
		{Number{1, 0}, Number{1, 0}},
		{Number{0.5, 0}, Number{1, 0}},
		{Number{1.5, 0}, Number{2, 0}},
		{Number{10, -1.5}, Number{10, -1}},
		{Number{-10, 1.5}, Number{-10, 2}},
		{Number{math.Inf(1), 0}, Number{math.Inf(1), 0}},
		{Number{math.NaN(), 0}, Number{math.NaN(), 0}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Ceil(tt.arg); !same(got, tt.want) {
				t.Errorf("Ceil() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShift(t *testing.T) {
	tests := []struct {
		arg  Number
		exp  int8
		want Number
	}{
		{Number{}, +1, Number{}},
		{Number{}, -1, Number{}},
		{Number{1, 0}, +1, Number{2, 0}},
		{Number{1, 0}, -1, Number{0.5, 0}},
		{Number{-1, 0.5}, 1, Number{-2, 1}},
		{Number{1, 0}, +127, Number{0x1p+127, 0}},
		{Number{1, 0}, -128, Number{0x1p-128, 0}},
		{Number{math.Inf(1), 0}, +1, Number{math.Inf(1), 0}},
		{Number{math.Inf(1), 0}, -1, Number{math.Inf(1), 0}},
		{Number{math.NaN(), 0}, +1, Number{math.NaN(), 0}},
		{Number{math.NaN(), 0}, -1, Number{math.NaN(), 0}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Shift(tt.arg, tt.exp); !same(got, tt.want) {
				t.Errorf("Shift() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdd(t *testing.T) {
	want := Add(Float(math.E), Float(math.Pi))
	got := AddFloats(math.E, math.Pi)
	if !same(got, want) {
		t.Errorf("AddFloats() = %v, want %v", got, want)
	}

	tests := []struct {
		arg1 float64
		arg2 float64
		want float64
	}{
		{math.Inf(+1), -1, math.Inf(+1)},
		{math.Inf(-1), +1, math.Inf(-1)},
		{+1, math.Inf(+1), math.Inf(+1)},
		{-1, math.Inf(-1), math.Inf(-1)},
		{math.Inf(+1), math.Inf(+1), math.Inf(+1)},
		{math.Inf(-1), math.Inf(-1), math.Inf(-1)},
		{math.Inf(+1), math.Inf(-1), math.NaN()},
		{math.Inf(-1), math.Inf(+1), math.NaN()},
		{math.NaN(), 0, math.NaN()},
		{0, math.NaN(), math.NaN()},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := AddFloats(tt.arg1, tt.arg2); !same(got, Float(tt.want)) {
				t.Errorf("AddFloats() = %v, want %v", got, tt.want)
			}
			if got := AddFloat(Float(tt.arg1), tt.arg2); !same(got, Float(tt.want)) {
				t.Errorf("AddFloat() = %v, want %v", got, tt.want)
			}
			if got := Add(Float(tt.arg1), Float(tt.arg2)); !same(got, Float(tt.want)) {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSub(t *testing.T) {
	want := Sub(Float(math.E), Float(math.Pi))
	got := SubFloats(math.E, math.Pi)
	if !same(got, want) {
		t.Errorf("SubFloats() = %v, want %v", got, want)
	}

	tests := []struct {
		arg1 float64
		arg2 float64
		want float64
	}{
		{math.Inf(+1), -1, math.Inf(+1)},
		{math.Inf(-1), +1, math.Inf(-1)},
		{+1, math.Inf(+1), math.Inf(-1)},
		{-1, math.Inf(-1), math.Inf(+1)},
		{math.Inf(+1), math.Inf(+1), math.NaN()},
		{math.Inf(-1), math.Inf(-1), math.NaN()},
		{math.Inf(+1), math.Inf(-1), math.Inf(+1)},
		{math.Inf(-1), math.Inf(+1), math.Inf(-1)},
		{math.NaN(), 0, math.NaN()},
		{0, math.NaN(), math.NaN()},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := SubFloats(tt.arg1, tt.arg2); !same(got, Float(tt.want)) {
				t.Errorf("SubFloats() = %v, want %v", got, tt.want)
			}
			if got := SubFloat(Float(tt.arg1), tt.arg2); !same(got, Float(tt.want)) {
				t.Errorf("SubFloat() = %v, want %v", got, tt.want)
			}
			if got := Sub(Float(tt.arg1), Float(tt.arg2)); !same(got, Float(tt.want)) {
				t.Errorf("Sub() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMul(t *testing.T) {
	want := Mul(Float(math.E), Float(math.Pi))
	got := MulFloats(math.E, math.Pi)
	if !same(got, want) {
		t.Errorf("SubFloats() = %v, want %v", got, want)
	}

	tests := []struct {
		arg1 float64
		arg2 float64
		want float64
	}{
		{math.Inf(+1), -1, math.Inf(-1)},
		{math.Inf(-1), +1, math.Inf(-1)},
		{+1, math.Inf(+1), math.Inf(+1)},
		{-1, math.Inf(-1), math.Inf(+1)},
		{math.Inf(+1), math.Inf(+1), math.Inf(+1)},
		{math.Inf(-1), math.Inf(-1), math.Inf(+1)},
		{math.Inf(+1), math.Inf(-1), math.Inf(-1)},
		{math.Inf(-1), math.Inf(+1), math.Inf(-1)},
		{math.Inf(+1), 0, math.NaN()},
		{math.Inf(-1), 0, math.NaN()},
		{0, math.Inf(+1), math.NaN()},
		{0, math.Inf(-1), math.NaN()},
		{math.NaN(), 0, math.NaN()},
		{0, math.NaN(), math.NaN()},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := MulFloats(tt.arg1, tt.arg2); !same(got, Float(tt.want)) {
				t.Errorf("MulFloats() = %v, want %v", got, tt.want)
			}
			if got := MulFloat(Float(tt.arg1), tt.arg2); !same(got, Float(tt.want)) {
				t.Errorf("MulFloat() = %v, want %v", got, tt.want)
			}
			if got := Mul(Float(tt.arg1), Float(tt.arg2)); !same(got, Float(tt.want)) {
				t.Errorf("Mul() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDiv(t *testing.T) {
	tests := []struct {
		arg1 float64
		arg2 float64
		want float64
	}{
		{math.Inf(+1), -1, math.Inf(-1)},
		{math.Inf(-1), +1, math.Inf(-1)},
		{math.Inf(+1), math.Inf(+1), math.NaN()},
		{math.Inf(-1), math.Inf(-1), math.NaN()},
		{math.Inf(+1), math.Inf(-1), math.NaN()},
		{math.Inf(-1), math.Inf(+1), math.NaN()},
		{math.Inf(+1), 0, math.Inf(+1)},
		{math.Inf(-1), 0, math.Inf(-1)},
		{math.NaN(), 0, math.NaN()},
		{0, math.NaN(), math.NaN()},
		{0, 0, math.NaN()},
		{0, math.Inf(+1), 0},
		{0, math.Inf(-1), 0},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Div(Float(tt.arg1), Float(tt.arg2)); !same(got, Float(tt.want)) {
				t.Errorf("Div() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSqrt(t *testing.T) {
	got := Sqrt(Float(2))
	if got.y != math.Sqrt2 {
		t.Fatalf("got %v want %v", got.y, math.Sqrt2)
	}

	a := toBig(got)
	b := myBig().Sqrt(big.NewFloat(2))

	if a.Cmp(b) != 0 {
		t.Fatalf("got %v want %v", a, b)
	}

	tests := []struct {
		arg  float64
		want float64
	}{
		{0, 0},
		{1, 1},
		{-1, math.NaN()},
		{math.Inf(1), math.Inf(1)},
		{math.Inf(-1), math.NaN()},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Sqrt(Float(tt.arg)); !same(got, Float(tt.want)) {
				t.Errorf("Div() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSqrDivSqrt(t *testing.T) {
	got := Sqr(Div(Float(1), Sqrt(Float(2))))
	if got.y != 0.5 {
		t.Fatalf("got %v want %v", got.y, math.Sqrt2)
	}

	a := toBig(got)
	b := big.NewFloat(0.5)

	if a.Cmp(b) != 0 {
		t.Fatalf("got %v want %v", a, b)
	}
}

func TestCmp(t *testing.T) {
	tests := []struct {
		arg1 Number
		arg2 Number
		want int
	}{
		{Number{}, Number{}, 0},
		{Number{+1, 0}, Number{}, +1},
		{Number{-1, 0}, Number{}, -1},
		{Number{1, +1}, Number{1, 0}, +1},
		{Number{1, -1}, Number{1, 0}, -1},
		{Number{math.Inf(1), 0}, Number{math.Inf(1), 0}, 0},
		{Number{math.NaN(), 0}, Number{math.NaN(), 0}, 0},
		{Number{}, Number{math.NaN(), 0}, +1},
		{Number{math.NaN(), 0}, Number{}, -1},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Cmp(tt.arg1, tt.arg2); got != tt.want {
				t.Errorf("Cmp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPolynomial(t *testing.T) {
	poly := [...]float64{1, -78, 2717, -55770, 749463, -6926634,
		44990231, -206070150, 657206836, -1414014888, 1931559552,
		-1486442880, 479001600}

	var got Number
	const x = 12.001
	for _, c := range poly {
		got = MulFloat(got, x)
		got = AddFloat(got, c)
	}

	const want = 40037.49486
	if d := math.Abs(SubFloat(got, want).y); d > 0.5e-5 {
		t.Fatalf("got %.5f want %.5f", got.y, want)
	}
}

func same(a, b Number) bool {
	return a == b || IsNaN(a) && IsNaN(b)
}

func myBig() *big.Float {
	return new(big.Float).SetPrec(prec)
}

func toBig(n Number) *big.Float {
	return myBig().Add(big.NewFloat(n.y), big.NewFloat(n.x))
}
