package dbldbl

import (
	"fmt"
	"math"
	"strings"
	"testing"
)

func TestLog(t *testing.T) {
	tests := []struct {
		arg  Number
		want string
	}{
		{Pi, "1.14472988584940017414342735135305871164729481291531157151362307"},        // https://oeis.org/A053510
		{Phi, "0.481211825059603447497758913424368423135184334385660519661018168"},      // https://oeis.org/A002390
		{Float(2), "0.693147180559945309417232121458176568075500134360255254120680009"}, // https://oeis.org/A002162
		{Float(3), "1.09861228866810969139524523692252570464749055782274945173469433"},  // https://oeis.org/A002391
		{Float(10), "2.30258509299404568401799145468436420760110148862877297603332790"}, // https://oeis.org/A002392
		{Float(0.5), "-0.693147180559945309417232121458176568075500134360255254120680009"},
		{AddFloats(1, 0x1p-55), "2.7755575615628913510590791702e-17"},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Log(tt.arg); !strings.HasPrefix(tt.want, fmt.Sprint(got)[:30]) {
				t.Errorf("Log() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLog_specials(t *testing.T) {
	tests := []struct {
		arg  Number
		want Number
	}{
		{Number{1, 0}, Number{}},
		{Number{}, Number{math.Inf(-1), 0}},
		{Number{-1, 0}, Number{math.NaN(), 0}},
		{Inf(1), Number{math.Inf(1), 0}},
		{NaN(), Number{math.NaN(), 0}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Log(tt.arg); !same(got, tt.want) {
				t.Errorf("Log() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExp(t *testing.T) {
	tests := []struct {
		arg  Number
		want string
	}{
		{Pi, "23.1406926327792690057290863679485473802661062426002119934450464"},       // https://oeis.org/A039661
		{Phi, "5.04316564336002865131188218928542471032359017541384636030200019"},      // https://oeis.org/A139341
		{Float(1), "2.71828182845904523536028747135266249775724709369995957496696762"}, // https://oeis.org/A072334
		{Float(2), "7.38905609893065022723042746057500781318031557055184732408712782"}, // https://oeis.org/A001113
		{Float(2.7755575615628913510590791702e-17), "1.000000000000000027755575615628914"},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Exp(tt.arg); !strings.HasPrefix(tt.want, fmt.Sprint(got)[:30]) {
				t.Errorf("Exp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExp_specials(t *testing.T) {
	tests := []struct {
		arg  Number
		want Number
	}{
		{Number{}, Number{1, 0}},
		{NaN(), Number{math.NaN(), 0}},
		{Inf(-1), Number{}},
		{Inf(+1), Number{math.Inf(1), 0}},
		{Float(math.SmallestNonzeroFloat64), Number{1, math.SmallestNonzeroFloat64}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Exp(tt.arg); !same(got, tt.want) {
				t.Errorf("Exp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_agm(t *testing.T) {
	got := agm(Float(24), Float(6))
	want := "13.4581714817256154207668131569743992430538388544"
	if !strings.HasPrefix(want, fmt.Sprint(got)[:30]) {
		t.Errorf("agm = %v, want %v", got, want)
	}
}

func Test_pi(t *testing.T) {
	got := pi()
	want := "3.14159265358979323846264338327950288419716939937510582097494459"
	if !strings.HasPrefix(want, fmt.Sprint(got)[:30]) {
		t.Errorf("agm = %v, want %v", got, want)
	}
}
