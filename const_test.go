package dbldbl

import "testing"

func Test_constants(t *testing.T) {
	tests := []struct {
		name string
		want Number
		str  string
	}{
		{"E", E, "2.71828182845904523536028747135266249775724709369995957496"}, // https://oeis.org/A001113
		{"Pi", Pi, "3.141592653589793238462643383279502884197169399375105820"}, // https://oeis.org/A000796
		{"Phi", Phi, "1.6180339887498948482045868343656381177203091798057628"}, // https://oeis.org/A001622
		{"Sqrt2", Sqrt2, "1.414213562373095048801688724209698078569671875376"}, // https://oeis.org/A002193
		{"SqrtE", SqrtE, "1.648721270700128146848650787814163571653776100710"}, // https://oeis.org/A019774
		{"SqrtPi", SqrtPi, "1.7724538509055160272981674833411451827975494561"}, // https://oeis.org/A002161
		{"SqrtPhi", SqrtPhi, "1.27201964951406896425242246173749149171560804"}, // https://oeis.org/A139339
		{"Ln2", Ln2, "0.6931471805599453094172321214581765680755001343602552"}, // https://oeis.org/A002162
		{"Ln10", Ln10, "2.30258509299404568401799145468436420760110148862877"}, // https://oeis.org/A002392
		{"2/Pi", twoOfPi, "0.63661977236758134307553505349005744813783858296"}, // https://oeis.org/A060294
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parse(tt.str); !same(got, tt.want) {
				t.Errorf("%s() = %#v, want %#v", tt.name, got, tt.want)
			}
		})
	}
}
