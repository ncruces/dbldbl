package dbldbl

import (
	"math"
	"testing"
)

func TestNeg(t *testing.T) {
	tests := []struct {
		arg  Number
		want Number
	}{
		{Number{}, Number{-zero, 0}},
		{Float(1), Number{-1, 0}},
		{Number{1, 0.5}, Number{-1, -0.5}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Neg(tt.arg); !same(got, tt.want) {
				t.Errorf("Neg() = %#v, want %#v", got, tt.want)
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
		{Float(+1), Number{1, 0}},
		{Float(-1), Number{1, 0}},
		{Number{+1, 0.5}, Number{1, +0.5}},
		{Number{-1, 0.5}, Number{1, -0.5}},
		{Inf(-1), Number{math.Inf(1), 0}},
		{NaN(), Number{math.NaN(), 0}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Abs(tt.arg); !same(got, tt.want) {
				t.Errorf("Abs() = %#v, want %#v", got, tt.want)
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
		{Float(+1), false},
		{Float(-1), false},
		{Inf(+1), true},
		{Inf(-1), true},
		{NaN(), false},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := IsInf(tt.arg, 0); got != tt.want {
				t.Errorf("IsInf() = %#v, want %#v", got, tt.want)
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
		{Float(1), Number{1, 0}},
		{Float(0.5), Number{0, 0}},
		{Float(1.5), Number{1, 0}},
		{Number{10, -1.5}, Number{10, -2}},
		{Number{-10, 1.5}, Number{-10, 2}},
		{Inf(1), Number{math.Inf(1), 0}},
		{NaN(), Number{math.NaN(), 0}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Trunc(tt.arg); !same(got, tt.want) {
				t.Errorf("Trunc() = %#v, want %#v", got, tt.want)
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
		{Float(1), Number{1, 0}},
		{Float(0.5), Number{0, 0}},
		{Float(1.5), Number{1, 0}},
		{Number{10, -1.5}, Number{10, -2}},
		{Number{-10, 1.5}, Number{-10, 1}},
		{Inf(1), Number{math.Inf(1), 0}},
		{NaN(), Number{math.NaN(), 0}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Floor(tt.arg); !same(got, tt.want) {
				t.Errorf("Floor() = %#v, want %#v", got, tt.want)
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
		{Float(1), Number{1, 0}},
		{Float(0.5), Number{1, 0}},
		{Float(1.5), Number{2, 0}},
		{Number{10, -1.5}, Number{10, -1}},
		{Number{-10, 1.5}, Number{-10, 2}},
		{Inf(1), Number{math.Inf(1), 0}},
		{NaN(), Number{math.NaN(), 0}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Ceil(tt.arg); !same(got, tt.want) {
				t.Errorf("Ceil() = %#v, want %#v", got, tt.want)
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
		{Float(1), +1, Number{2, 0}},
		{Float(1), -1, Number{0.5, 0}},
		{Number{-1, 0.5}, 1, Number{-2, 1}},
		{Float(1), +127, Number{0x1p+127, 0}},
		{Float(1), -128, Number{0x1p-128, 0}},
		{Inf(1), +1, Number{math.Inf(1), 0}},
		{Inf(1), -1, Number{math.Inf(1), 0}},
		{NaN(), +1, Number{math.NaN(), 0}},
		{NaN(), -1, Number{math.NaN(), 0}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Shift(tt.arg, tt.exp); !same(got, tt.want) {
				t.Errorf("Shift() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestAdd(t *testing.T) {
	want := Add(Float(math.E), Float(math.Pi))
	got := AddFloats(math.E, math.Pi)
	if !same(got, want) {
		t.Errorf("AddFloats() = %#v, want %#v", got, want)
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
		{math.MaxFloat64, math.MaxFloat64, math.Inf(+1)},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := AddFloats(tt.arg1, tt.arg2); !same(got, Float(tt.want)) {
				t.Errorf("AddFloats() = %#v, want %#v", got, tt.want)
			}
			if got := AddFloat(Float(tt.arg1), tt.arg2); !same(got, Float(tt.want)) {
				t.Errorf("AddFloat() = %#v, want %#v", got, tt.want)
			}
			if got := Add(Float(tt.arg1), Float(tt.arg2)); !same(got, Float(tt.want)) {
				t.Errorf("Add() = %#v, want %#v", got, tt.want)
			}
		})
	}

	// Ensure no overflow.
	max := Number{math.MaxFloat64 / 2, math.MaxFloat64 / 2 * 0x1p-54}
	if got := Add(max, max); !isFinite(got.y) || !isFinite(got.x) {
		t.Errorf("Add() = %#v", got)
	}
}

func TestSub(t *testing.T) {
	want := Sub(Float(math.E), Float(math.Pi))
	got := SubFloats(math.E, math.Pi)
	if !same(got, want) {
		t.Errorf("SubFloats() = %#v, want %#v", got, want)
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
		{-math.MaxFloat64, math.MaxFloat64, math.Inf(-1)},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := SubFloats(tt.arg1, tt.arg2); !same(got, Float(tt.want)) {
				t.Errorf("SubFloats() = %#v, want %#v", got, tt.want)
			}
			if got := SubFloat(Float(tt.arg1), tt.arg2); !same(got, Float(tt.want)) {
				t.Errorf("SubFloat() = %#v, want %#v", got, tt.want)
			}
			if got := Sub(Float(tt.arg1), Float(tt.arg2)); !same(got, Float(tt.want)) {
				t.Errorf("Sub() = %#v, want %#v", got, tt.want)
			}
		})
	}

	// Ensure no underflow.
	max := Number{math.MaxFloat64 / 2, math.MaxFloat64 / 2 * 0x1p-54}
	if got := Sub(Neg(max), max); !isFinite(got.y) || !isFinite(got.x) {
		t.Errorf("Sub() = %#v", got)
	}
}

func TestMul(t *testing.T) {
	want := Mul(Float(math.E), Float(math.Pi))
	got := MulFloats(math.E, math.Pi)
	if !same(got, want) {
		t.Errorf("MulFloats() = %#v, want %#v", got, want)
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
		{math.MaxFloat64, -math.MaxFloat64, math.Inf(-1)},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := MulFloats(tt.arg1, tt.arg2); !same(got, Float(tt.want)) {
				t.Errorf("MulFloats() = %#v, want %#v", got, tt.want)
			}
			if got := MulFloat(Float(tt.arg1), tt.arg2); !same(got, Float(tt.want)) {
				t.Errorf("MulFloat() = %#v, want %#v", got, tt.want)
			}
			if got := Mul(Float(tt.arg1), Float(tt.arg2)); !same(got, Float(tt.want)) {
				t.Errorf("Mul() = %#v, want %#v", got, tt.want)
			}
		})
	}

	// Ensure no overflow.
	max := Number{math.Sqrt(math.MaxFloat64), math.Sqrt(math.MaxFloat64) * 0x1p-54}
	if got := Mul(max, max); !isFinite(got.y) || !isFinite(got.x) {
		t.Errorf("Mul() = %#v", got)
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
		{0, math.Inf(-1), -zero},
		{1, -math.SmallestNonzeroFloat64, math.Inf(-1)},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Div(Float(tt.arg1), Float(tt.arg2)); !same(got, Float(tt.want)) {
				t.Errorf("Div() = %#v, want %#v", got, tt.want)
			}
		})
	}

	// Ensure no overflow.
	max := Number{math.Sqrt(math.MaxFloat64), math.Sqrt(math.MaxFloat64) * 0x1p-54}
	if got := Div(max, Div(Float(1), max)); !isFinite(got.y) || !isFinite(got.x) {
		t.Errorf("Div() = %#v", got)
	}
}

func TestSqr(t *testing.T) {
	tests := []struct {
		arg  float64
		want float64
	}{
		{0, 0},
		{1, 1},
		{-1, 1},
		{math.Inf(1), math.Inf(1)},
		{math.Inf(-1), math.Inf(+1)},
		{math.NaN(), math.NaN()},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Sqr(Float(tt.arg)); !same(got, Float(tt.want)) {
				t.Errorf("Sqr() = %#v, want %#v", got, tt.want)
			}
		})
	}

	// Ensure no overflow.
	max := Number{math.Sqrt(math.MaxFloat64), math.Sqrt(math.MaxFloat64) * 0x1p-54}
	if got := Sqr(max); !isFinite(got.y) || !isFinite(got.x) {
		t.Errorf("Sqr() = %#v", got)
	}
}

func TestSqrt(t *testing.T) {
	got := Sqrt(Float(2))
	if got.y != math.Sqrt2 {
		t.Fatalf("Sqrt() = %#v, want %#v", got.y, math.Sqrt2)
	}

	if res := Sqr(got); res.y != 2 || res.x != 0 {
		t.Fatalf("Sqrt() = %#v, want %#v", res, 2)
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
		{math.NaN(), math.NaN()},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Sqrt(Float(tt.arg)); !same(got, Float(tt.want)) {
				t.Errorf("Sqrt() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestCbrt(t *testing.T) {
	got := Cbrt(Float(-3))
	if got.y != math.Cbrt(-3) {
		t.Fatalf("Cbrt() = %#v, want %#v", got.y, math.Cbrt(-3))
	}

	if res := Mul(got, Sqr(got)); res.y != -3 || math.Abs(res.x) > 0x1p-104 {
		t.Fatalf("Cbrt() = %#v, want %#v", res, -3)
	}

	tests := []struct {
		arg  float64
		want float64
	}{
		{0, 0},
		{1, 1},
		{-1, -1},
		{math.Inf(1), math.Inf(1)},
		{math.Inf(-1), math.Inf(-1)},
		{math.NaN(), math.NaN()},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Cbrt(Float(tt.arg)); !same(got, Float(tt.want)) {
				t.Errorf("Cbrt() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestFMA(t *testing.T) {
	a := AddFloats(1, math.E/0x1p100)
	b := Sub(Float(1), Shift(a, 1))
	if got := FMA(a, a, b); !same(got, Float(a.x*a.x)) {
		t.Fatalf("FMA() = %#v, want %#v", got.y, a.x*a.x)
	}

	tests := []struct {
		arg0 float64
		arg1 float64
		arg2 Number
		want Number
	}{
		{0, 0, Number{}, Number{}},
		{0, 0, Shift(Pi, -128), Shift(Pi, -128)},
		{-math.MaxFloat32, -math.MaxFloat64, Float(+math.MaxFloat64), Inf(+1)},
		{+math.MaxFloat64, -math.MaxFloat32, Float(-math.MaxFloat32), Inf(-1)},
		{+math.MaxFloat64, -math.MaxFloat64, Float(-math.MaxFloat32), Inf(-1)},
		{+math.MaxFloat64, -math.MaxFloat64, Inf(+1), Inf(+1)},
		{-math.MaxFloat64, -math.MaxFloat64, Inf(-1), Inf(-1)},
		{-math.MaxFloat32, math.NaN(), Inf(0), NaN()},
		{-math.MaxFloat32, math.NaN(), Float(-math.MaxFloat32), NaN()},
		{-math.MaxFloat32, +math.MaxFloat32, NaN(), NaN()},
		{math.Inf(-1), math.MaxFloat32, Float(-math.MaxFloat32), Inf(-1)},
		{math.Inf(+1), math.MaxFloat32, Float(-math.MaxFloat32), Inf(+1)},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := FMAFloat(tt.arg0, tt.arg1, tt.arg2); !same(got, tt.want) {
				t.Errorf("FMA() = %#v, want %#v", got, tt.want)
			}
			if got := FMA(Float(tt.arg0), Float(tt.arg1), tt.arg2); !same(got, tt.want) {
				t.Errorf("FMA() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestSqrDivSqrt(t *testing.T) {
	got := Sqr(Div(Float(1), Sqrt(Float(2))))
	if got != Float(0.5) { // (1/√2)² = 0.5
		t.Fatalf("got %#v", got)
	}
}

func TestCmp(t *testing.T) {
	tests := []struct {
		arg1 Number
		arg2 Number
		want int
	}{
		{Number{}, Number{}, 0},
		{Float(+1), Number{}, +1},
		{Float(-1), Number{}, -1},
		{Number{1, +1}, Number{1, 0}, +1},
		{Number{1, -1}, Number{1, 0}, -1},
		{Number{}, Number{math.NaN(), 0}, +1},
		{Inf(1), Number{math.Inf(1), 0}, 0},
		{NaN(), Number{math.NaN(), 0}, 0},
		{NaN(), Number{}, -1},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Cmp(tt.arg1, tt.arg2); got != tt.want {
				t.Errorf("Cmp() = %#v, want %#v", got, tt.want)
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
