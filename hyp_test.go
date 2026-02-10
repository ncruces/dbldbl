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
		{Float(1), "0.881373587019543025232609324979792309028160328261635410753295"},  // https://oeis.org/A068051
		{Float(-1), "-0.88137358701954302523260932497979230902816032826163541075329"}, //
		{Float(0.5), "0.48121182505960344749775891342436842313518433438566051966101"}, // asinh(0.5)
		{Pi, "1.86229574331084821988836132518264491594769850503039970476157"},         // asinh(π)
		{Float(2), "1.44363547517881034249327674027310420154256911586273682099405"},   // https://oeis.org/A102888
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Asinh(tt.arg); !strings.HasPrefix(tt.want, fmt.Sprint(got)[:30]) {
				t.Errorf("Asinh() = %v, want %v", got, tt.want)
			}
		})
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
		{Float(2), "1.31695789692481670862504634730796844402698197146966603054147"}, // https://oeis.org/A068516
		{Float(1.5), "0.9624236501192068949955178268487368462704461920192687212"},   // https://oeis.org/A143526
		{Pi, "1.81152627246085310702185204930542051022064520104597705211413"},       // acosh(π)
		{E, "1.65745445415307727259382874228053473915829500605639201750514"},        // acosh(e)
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Acosh(tt.arg); !strings.HasPrefix(tt.want, fmt.Sprint(got)[:30]) {
				t.Errorf("Acosh() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAcosh_specials(t *testing.T) {
	tests := []struct {
		arg  Number
		want Number
	}{
		{Float(1), Number{}},
		{Float(0.5), Number{y: math.NaN()}},
		{Float(-1), Number{y: math.NaN()}},
		{NaN(), Number{y: math.NaN()}},
		{Inf(+1), Number{y: math.Inf(1)}},
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
		{Float(0.5), "0.54930614433405484569762261846126285232374527891137472586735"}, // https://oeis.org/A068631
		{Float(-0.5), "-0.549306144334054845697622618461262852323745278911374725867"},
		{Float(0.25), "0.25541281188299534160275704815183096743905539822288413508897"}, // atanh(0.25)
		{Inv(E), "0.38596841645265236253531957001759267189616767070472722560582"},      // atanh(1/e)
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Atanh(tt.arg); !strings.HasPrefix(tt.want, fmt.Sprint(got)[:30]) {
				t.Errorf("Atanh() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAtanh_specials(t *testing.T) {
	tests := []struct {
		arg  Number
		want Number
	}{
		{Number{}, Number{}},
		{Float(neg0), Number{y: neg0}},
		{Float(1), Number{y: math.Inf(1)}},
		{Float(-1), Number{y: math.Inf(-1)}},
		{Float(2), Number{y: math.NaN()}},
		{Float(-2), Number{y: math.NaN()}},
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
