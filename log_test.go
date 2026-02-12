package dbldbl

import (
	"math"
	"testing"
)

func TestLog(t *testing.T) {
	tests := []struct {
		arg  Number
		want string
	}{
		{Pi, "1.144729885849400174143427351353058711647294812915311571513623"}, // https://oeis.org/A053510
		{Phi, "0.48121182505960344749775891342436842313518433438566051966101"}, // https://oeis.org/A002390
		{Float(2), "0.693147180559945309417232121458176568075500134360255254"}, // https://oeis.org/A002162
		{Float(3), "1.098612288668109691395245236922525704647490557822749451"}, // https://oeis.org/A002391
		{Float(10), "2.30258509299404568401799145468436420760110148862877297"}, // https://oeis.org/A002392
		{Float(0.5), "-0.693147180559945309417232121458176568075500134360255"},
		{Float(math.Nextafter(1, 2)), "2.2204460492503128343282295362059e-16"},
		{AddFloats(1, 0x1p-55), "2.77555756156289135105907917022705e-17"},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Log(tt.arg); !near(got, tt.want) {
				t.Errorf("Log() = %v, want %v", got, tt.want)
			}
		})
	}

	// Ensure no overflow.
	max := Number{math.MaxFloat64, math.MaxFloat64 * 0x1p-54}
	if got := Log(max); !isFinite(got.y) || !isFinite(got.x) {
		t.Errorf("Log() = %#v", got)
	}
	// Ensure no underflow.
	min := Number{math.SmallestNonzeroFloat64, 0}
	if got := Log(min); !isFinite(got.y) || !isFinite(got.x) {
		t.Errorf("Log() = %#v", got)
	}
}

func TestLog_specials(t *testing.T) {
	tests := []struct {
		arg  Number
		want Number
	}{
		{Float(1), Number{}},
		{Float(-1), Number{math.NaN(), 0}},
		{Number{}, Number{math.Inf(-1), 0}},
		{Inf(1), Number{math.Inf(1), 0}},
		{NaN(), Number{math.NaN(), 0}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Log(tt.arg); !same(got, tt.want) {
				t.Errorf("Log() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestExp(t *testing.T) {
	tests := []struct {
		arg  Number
		want string
	}{
		{Pi, "23.14069263277926900572908636794854738026610624260021199344504"}, // https://oeis.org/A039661
		{Phi, "5.04316564336002865131188218928542471032359017541384636030200"}, // https://oeis.org/A139341
		{Float(1), "2.718281828459045235360287471352662497757247093699959574"}, // https://oeis.org/A001113
		{Float(2), "7.389056098930650227230427460575007813180315570551847324"}, // https://oeis.org/A072334
		{Float(0x1p-55), "1.000000000000000027755575615628914"},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Exp(tt.arg); !near(got, tt.want) {
				t.Errorf("Exp() = %v, want %v", got, tt.want)
			}
		})
	}

	// Ensure no overflow.
	if got := Exp(Float(709)); !isFinite(got.y) || !isFinite(got.x) {
		t.Errorf("Exp() = %#v", got)
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
				t.Errorf("Exp() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestLog1p(t *testing.T) {
	// Ensure no overflow.
	max := Number{math.MaxFloat64, math.MaxFloat64 * 0x1p-54}
	if got := Log1p(max); !isFinite(got.y) || !isFinite(got.x) {
		t.Errorf("Log1p() = %#v", got)
	}
	// Ensure no underflow.
	min := Number{math.Nextafter(-1, 0), 0}
	if got := Log1p(min); !isFinite(got.y) || !isFinite(got.x) {
		t.Errorf("Log1p() = %#v", got)
	}
}

func TestLog1p_specials(t *testing.T) {
	tests := []struct {
		arg  Number
		want Number
	}{
		{Float(0), Number{}},
		{Float(-zero), Number{-zero, 0}},
		{AddFloats(-1, -0x1p-55), Number{math.NaN(), 0}},
		{Float(-1), Number{math.Inf(-1), 0}},
		{Inf(1), Number{math.Inf(1), 0}},
		{NaN(), Number{math.NaN(), 0}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Log1p(tt.arg); !same(got, tt.want) {
				t.Errorf("Log1p() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestExpm1(t *testing.T) {
	tests := []struct {
		arg  Number
		want string
	}{
		{Pi, "22.14069263277926900572908636794854738026610624260021199344504"}, // https://oeis.org/A039661
		{Phi, "4.04316564336002865131188218928542471032359017541384636030200"}, // https://oeis.org/A139341
		{Float(1), "1.718281828459045235360287471352662497757247093699959574"}, // https://oeis.org/A001113
		{Float(2), "6.389056098930650227230427460575007813180315570551847324"}, // https://oeis.org/A072334
		{Float(0x1p-55), "2.77555756156289135105907917022705e-17"},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Expm1(tt.arg); !near(got, tt.want) {
				t.Errorf("Expm1() = %v, want %v", got, tt.want)
			}
		})
	}

	// Ensure no overflow.
	if got := Expm1(Float(709)); !isFinite(got.y) || !isFinite(got.x) {
		t.Errorf("Expm1() = %#v", got)
	}
}

func TestExpm1_specials(t *testing.T) {
	tests := []struct {
		arg  Number
		want Number
	}{
		{Float(0), Number{}},
		{Float(-zero), Number{-zero, 0}},
		{Inf(+1), Number{math.Inf(1), 0}},
		{Inf(-1), Number{-1, 0}},
		{Float(-1024), Number{-1, 0}},
		{NaN(), Number{math.NaN(), 0}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Expm1(tt.arg); !same(got, tt.want) {
				t.Errorf("Expm1() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func Test_agm(t *testing.T) {
	// https://en.wikipedia.org/wiki/Arithmetic%E2%80%93geometric_mean#Example
	got := agm(Float(24), Float(6))
	want := "13.4581714817256154207668131569743992430538388544"
	if !near(got, want) {
		t.Errorf("agm = %v, want %v", got, want)
	}
}
