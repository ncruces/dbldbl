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
		{Float(1), "1.17520119364380145688238185059560081515571798133410668596562850"}, // https://oeis.org/A073742
		{Float(-1), "-1.1752011936438014568823818505956008151557179813341066859656285"},
		{Float(0.5), "0.52109530549374736162242562641149155910592898261148052794609461"}, // https://oeis.org/A197868
		{Pi, "11.548739357257748377977334315388143695066417655546696754349045"},         // sinh(π), https://oeis.org/A073743
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Sinh(tt.arg); !strings.HasPrefix(tt.want, fmt.Sprint(got)[:30]) {
				t.Errorf("Sinh() = %v, want %v", got, tt.want)
			}
		})
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
		{Float(1), "1.54308063481524377847790562075706168260152911236336420171242837"},  // https://oeis.org/A073744
		{Float(-1), "1.5430806348152437784779056207570616826015291123633642017124283"}, //
		{Float(0.5), "1.1276259652063807852262251614026976547323565263257790509238666"}, // https://oeis.org/A197869
		{Pi, "11.591953275521520627751752052560051025883096072309132281892047"},        // cosh(π), https://oeis.org/A073745
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Cosh(tt.arg); !strings.HasPrefix(tt.want, fmt.Sprint(got)[:30]) {
				t.Errorf("Cosh() = %v, want %v", got, tt.want)
			}
		})
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
		{Float(1), "0.7615941559557648881194582826047935904127685972579365515968105"}, // https://oeis.org/A073746
		{Float(-1), "-0.761594155955764888119458282604793590412768597257936551596810"},
		{Float(0.5), "0.46211715726000975850231848364367254200461231635194031013072818"}, // https://oeis.org/A197870
		{Pi, "0.99627207622074994426469058001253671204139025728833936491954"},         // tanh(π), https://oeis.org/A073747
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
		{Float(1), "0.881373587019543025232609324979792309028160328261635410753295"},   // https://oeis.org/A068051
		{Float(-1), "-0.88137358701954302523260932497979230902816032826163541075329"}, //
		{Float(0.5), "0.48121182505960344749775891342436842313518433438566051966101"}, // asinh(0.5)
		{Pi, "1.86229574331084821988836132518264491594769850503039970476157"},        // asinh(π)
		{Float(2), "1.44363547517881034249327674027310420154256911586273682099405"},  // https://oeis.org/A102888
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
		{Float(2), "1.31695789692481670862504634730796844402698197146966603054147"},  // https://oeis.org/A068516
		{Float(1.5), "0.9624236501192068949955178268487368462704461920192687212"}, // https://oeis.org/A143526
		{Pi, "1.81152627246085310702185204930542051022064520104597705211413"},   // acosh(π)
		{E, "1.65745445415307727259382874228053473915829500605639201750514"},    // acosh(e)
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
		{Float(0.25), "0.25541281188299534160275704815183096743905539822288413508897"},  // atanh(0.25)
		{Inv(E), "0.38596841645265236253531957001759267189616767070472722560582"},         // atanh(1/e)
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
