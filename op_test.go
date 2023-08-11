package dbldbl

import (
	"math"
	"math/big"
	"reflect"
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
			if got := Neg(tt.arg); !reflect.DeepEqual(got, tt.want) {
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
		{Number{1, 0}, Number{1, 0}},
		{Number{-1, 0}, Number{1, -0}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Abs(tt.arg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Abs() = %v, want %v", got, tt.want)
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
		{Number{1.5, 0.5}, Number{1, 0}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Trunc(tt.arg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Trunc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShift(t *testing.T) {
	tests := []struct {
		arg  Number
		exp  int16
		want Number
	}{
		{Number{}, 1, Number{}},
		{Number{1, 0}, 1, Number{2, 0}},
		{Number{1, 0}, -1, Number{0.5, 0}},
		{Number{-1, 0.5}, 1, Number{-2, 1}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Shift(tt.arg, tt.exp); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Shift() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddFloats(t *testing.T) {
	want := Add(Float(math.E), Float(math.Pi))
	got := AddFloats(math.E, math.Pi)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("AddFloats() = %v, want %v", got, want)
	}
}

func TestSubFloats(t *testing.T) {
	want := Sub(Float(math.E), Float(math.Pi))
	got := SubFloats(math.E, math.Pi)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("SubFloats() = %v, want %v", got, want)
	}
}

func TestMulFloats(t *testing.T) {
	want := Mul(Float(math.E), Float(math.Pi))
	got := MulFloats(math.E, math.Pi)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("SubFloats() = %v, want %v", got, want)
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

func myBig() *big.Float {
	return new(big.Float).SetPrec(prec)
}

func toBig(n Number) *big.Float {
	return myBig().Add(big.NewFloat(n.y), big.NewFloat(n.x))
}
