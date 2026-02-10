package dbldbl

import (
	"fmt"
	"math"
	"strings"
	"testing"
)

func TestSinh(t *testing.T) {
	tests := []struct {
		arg  Number
		want string
	}{
		{Float(1), "1.175201193643801456882381850595600815155717981334095870"}, // https://oeis.org/A073742
		{Float(-1), "-1.1752011936438014568823818505956008151557179813340958"},
		{Pi, "11.54873935725774837797733431538840968449518906639478945523216"}, // https://oeis.org/A334401
		{E, "7.5441371028169758263418200425165327402949857443016716663691364"}, // https://oeis.org/A334399
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Sinh(tt.arg); !strings.HasPrefix(tt.want, fmt.Sprint(got)[:30]) {
				t.Errorf("Sinh() = %v, want %v", got, tt.want)
			}
		})
	}

	// Ensure no overflow.
	if got := Sinh(Float(709)); !isFinite(got.y) || !isFinite(got.x) {
		t.Errorf("Sinh() = %#v", got)
	}
}

func TestSinh_specials(t *testing.T) {
	tests := []struct {
		arg  Number
		want Number
	}{
		{Number{}, Number{}},
		{Float(neg0), Number{y: neg0}},
		{NaN(), Number{y: math.NaN()}},
		{Inf(+1), Number{y: math.Inf(1)}},
		{Inf(-1), Number{y: math.Inf(-1)}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Sinh(tt.arg); !same(got, tt.want) {
				t.Errorf("Sinh() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestCosh(t *testing.T) {
	tests := []struct {
		arg  Number
		want string
	}{
		{Float(1), "1.543080634815243778477905620757061682601529112365863704"}, // https://oeis.org/A073743
		{Float(-1), "1.54308063481524377847790562075706168260152911236586370"},
		{Pi, "11.59195327552152062775175205256013769577091717620542253821288"}, // https://oeis.org/A334402
		{E, "7.6101251386622883634186102301133791652335627925544681027716099"}, // https://oeis.org/A334400
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Cosh(tt.arg); !strings.HasPrefix(tt.want, fmt.Sprint(got)[:30]) {
				t.Errorf("Cosh() = %v, want %v", got, tt.want)
			}
		})
	}

	// Ensure no overflow.
	if got := Cosh(Float(709)); !isFinite(got.y) || !isFinite(got.x) {
		t.Errorf("Cosh() = %#v", got)
	}
}

func TestCosh_specials(t *testing.T) {
	tests := []struct {
		arg  Number
		want Number
	}{
		{Number{}, Number{1, 0}},
		{Float(neg0), Number{1, 0}},
		{NaN(), Number{y: math.NaN()}},
		{Inf(+1), Number{y: math.Inf(1)}},
		{Inf(-1), Number{y: math.Inf(1)}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Cosh(tt.arg); !same(got, tt.want) {
				t.Errorf("Cosh() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestTanh(t *testing.T) {
	tests := []struct {
		arg  Number
		want string
	}{
		{Float(1), "0.761594155955764888119458282604793590412768597257936551"}, // https://oeis.org/A073744
		{Float(-1), "-0.7615941559557648881194582826047935904127685972579365"},
		{Pi, "0.996272076220749944264690580012536711896899190804587614362612"}, // https://oeis.org/A344505
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Tanh(tt.arg); !strings.HasPrefix(tt.want, fmt.Sprint(got)[:30]) {
				t.Errorf("Tanh() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTanh_specials(t *testing.T) {
	tests := []struct {
		arg  Number
		want Number
	}{
		{Number{}, Number{}},
		{Float(neg0), Number{y: neg0}},
		{NaN(), Number{y: math.NaN()}},
		{Inf(+1), Number{1, 0}},
		{Inf(-1), Number{-1, 0}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Tanh(tt.arg); !same(got, tt.want) {
				t.Errorf("Tanh() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestAsinh(t *testing.T) {
	tests := []struct {
		arg  Number
		want string
	}{
		{E, "1.7253825588523150939450979704048887562745572746729386688142115"}, // https://oeis.org/A366599
		{Pi, "1.862295743310848219888361325182620574902674184961554765612879"}, // https://oeis.org/A360938
		{Float(0.25), "0.247466461547263452944781549788359289253766903098567"}, // https://oeis.org/A129200
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Asinh(tt.arg); !strings.HasPrefix(tt.want, fmt.Sprint(got)[:30]) {
				t.Errorf("Asinh() = %v, want %v", got, tt.want)
			}
		})
	}

	// Ensure no overflow.
	max := Number{math.MaxFloat64, math.MaxFloat64 * 0x1p-54}
	if got := Asinh(max); !isFinite(got.y) || !isFinite(got.x) {
		t.Errorf("Asinh() = %#v", got)
	}
}

func TestAsinh_specials(t *testing.T) {
	tests := []struct {
		arg  Number
		want Number
	}{
		{Number{}, Number{}},
		{Float(neg0), Number{y: neg0}},
		{NaN(), Number{y: math.NaN()}},
		{Inf(+1), Number{y: math.Inf(1)}},
		{Inf(-1), Number{y: math.Inf(-1)}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Asinh(tt.arg); !same(got, tt.want) {
				t.Errorf("Asinh() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestAcosh(t *testing.T) {
	tests := []struct {
		arg  Number
		want string
	}{
		{E, "1.6574544541530772725938287422805347391583927620336768258485822"}, // https://oeis.org/A365927
		{Pi, "1.811526272460853107021852049305420510220702081057922474861595"}, // https://oeis.org/A359540
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Acosh(tt.arg); !strings.HasPrefix(tt.want, fmt.Sprint(got)[:30]) {
				t.Errorf("Acosh() = %v, want %v", got, tt.want)
			}
		})
	}

	// Ensure no overflow.
	max := Number{math.MaxFloat64, math.MaxFloat64 * 0x1p-54}
	if got := Acosh(max); !isFinite(got.y) || !isFinite(got.x) {
		t.Errorf("Acosh() = %#v", got)
	}
}

func TestAcosh_specials(t *testing.T) {
	tests := []struct {
		arg  Number
		want Number
	}{
		{Float(1), Number{}},
		{Float(0), Number{y: math.NaN()}},
		{Inf(+1), Number{y: math.Inf(1)}},
		{NaN(), Number{y: math.NaN()}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Acosh(tt.arg); !same(got, tt.want) {
				t.Errorf("Acosh() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestAtanh(t *testing.T) {
	tests := []struct {
		arg  Number
		want string
	}{
		{AddFloat(Phi, -1), "0.721817737589405171246638370136552634702776501"}, // https://oeis.org/A377813
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Atanh(tt.arg); !strings.HasPrefix(tt.want, fmt.Sprint(got)[:30]) {
				t.Errorf("Atanh() = %v, want %v", got, tt.want)
			}
		})
	}

	// Ensure no overflow.
	max := AddFloats(1, -math.Sqrt(math.SmallestNonzeroFloat64))
	if got := Atanh(max); !isFinite(got.y) || !isFinite(got.x) {
		t.Errorf("Atanh() = %#v", got)
	}
}

func TestAtanh_specials(t *testing.T) {
	tests := []struct {
		arg  Number
		want Number
	}{
		{Number{}, Number{}},
		{Float(neg0), Number{y: neg0}},
		{Float(+1), Number{y: math.Inf(+1)}},
		{Float(-1), Number{y: math.Inf(-1)}},
		{Float(2), Number{y: math.NaN()}},
		{NaN(), Number{y: math.NaN()}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Atanh(tt.arg); !same(got, tt.want) {
				t.Errorf("Atanh() = %#v, want %#v", got, tt.want)
			}
		})
	}
}
