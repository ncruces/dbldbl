package dbldbl

import (
	"math"
	"math/big"
	"testing"
)

const prec = 104

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

func myBig() *big.Float {
	return new(big.Float).SetPrec(prec)
}

func toBig(n Number) *big.Float {
	return myBig().Add(big.NewFloat(n.y), big.NewFloat(n.x))
}
