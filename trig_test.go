package dbldbl

import (
	"fmt"
	"math"
	"strings"
	"testing"
)

var neg0 = math.Copysign(0, -1)

func TestSin(t *testing.T) {
	tests := []struct {
		arg  Number
		want string
	}{
		{Div(Pi, Float(6)), "0.4999999999999999999999999999999"},                // sin(π/6) = 1/2
		{Div(Pi, Float(4)), "0.7071067811865475244008443621048490392848359376"}, // sin(π/4) = √2/2, https://oeis.org/A010503
		{Div(Pi, Float(3)), "0.8660254037844386467637231707529361834714026269"}, // sin(π/3) = √3/2, https://oeis.org/A010527
		{Float(1), "0.8414709848078965066525023216302989996225630607983710656"}, // https://oeis.org/A049469
		{Float(-1), "-0.84147098480789650665250232163029899962256306079837106"}, //
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Sin(tt.arg); !strings.HasPrefix(tt.want, fmt.Sprint(got)[:30]) {
				t.Errorf("Sin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSin_specials(t *testing.T) {
	tests := []struct {
		arg  Number
		want Number
	}{
		{Number{}, Number{}},
		{Float(neg0), Number{y: neg0}},
		{NaN(), Number{y: math.NaN()}},
		{Inf(+1), Number{y: math.NaN()}},
		{Inf(-1), Number{y: math.NaN()}},
		{Pi, Number{y: neg0}},  // sin(π) = ±0
		{twoPi, Number{}},      // sin(2π) = 0
		{halfPi, Number{y: 1}}, // sin(π/2) = 1
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Sin(tt.arg); !same(got, tt.want) {
				t.Errorf("Sin() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestCos(t *testing.T) {
	tests := []struct {
		arg  Number
		want string
	}{
		{Div(Pi, Float(3)), "0.4999999999999999999999999999999"},                // cos(π/3) = 1/2
		{Div(Pi, Float(4)), "0.7071067811865475244008443621048490392848359376"}, // cos(π/4) = √2/2, https://oeis.org/A010503
		{Div(Pi, Float(6)), "0.8660254037844386467637231707529361834714026269"}, // cos(π/6) = √3/2, https://oeis.org/A010527
		{Float(1), "0.5403023058681397174009366074429766037323104206179222276"}, // https://oeis.org/A049470
		{Float(-1), "0.540302305868139717400936607442976603732310420617922227"}, //
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Cos(tt.arg); !strings.HasPrefix(tt.want, fmt.Sprint(got)[:30]) {
				t.Errorf("Cos() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCos_specials(t *testing.T) {
	tests := []struct {
		arg  Number
		want Number
	}{
		{Number{}, Number{1, 0}},
		{NaN(), Number{y: math.NaN()}},
		{Inf(+1), Number{y: math.NaN()}},
		{Inf(-1), Number{y: math.NaN()}},
		{Pi, Number{y: -1}},       // cos(π) = -1
		{twoPi, Number{y: 1}},     // cos(2π) = 1
		{halfPi, Number{y: neg0}}, // cos(π/2) = ±0

	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Cos(tt.arg); !same(got, tt.want) {
				t.Errorf("Cos() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestTan(t *testing.T) {
	tests := []struct {
		arg  Number
		want string
	}{
		{Div(Pi, Float(4)), "0.99999999999999999999999999999999"},               // tan(π/4) = 1
		{Div(Pi, Float(3)), "1.7320508075688772935274463415058723669428052538"}, // tan(π/3) = √3, https://oeis.org/A002194
		{Float(1), "1.5574077246549022305069748074583601730872507723815200383"}, // https://oeis.org/A049471
		{Float(-1), "-1.55740772465490223050697480745836017308725077238152003"}, //
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Tan(tt.arg); !strings.HasPrefix(tt.want, fmt.Sprint(got)[:30]) {
				t.Errorf("Tan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTan_specials(t *testing.T) {
	tests := []struct {
		arg  Number
		want Number
	}{
		{Number{}, Number{}},
		{Float(neg0), Number{y: neg0}},
		{NaN(), Number{y: math.NaN()}},
		{Inf(+1), Number{y: math.NaN()}},
		{Inf(-1), Number{y: math.NaN()}},
		{Pi, Number{}},    // tan(π) = 0
		{twoPi, Number{}}, // tan(2π) = 0
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Tan(tt.arg); !same(got, tt.want) {
				t.Errorf("Tan() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestAsin(t *testing.T) {
	tests := []struct {
		arg  Number
		want string
	}{
		{Float(1), "1.570796326794896619231321691639751442098584699687552910"}, // π/2, https://oeis.org/A019669
		{Float(-1), "-1.5707963267948966192313216916397514420985846996875529"}, //
		{Float(0.5), "0.5235987755982988730771072305465838140328615665625176"}, // π/6, https://oeis.org/A019685
		{Float(-0.5), "-0.52359877559829887307710723054658381403286156656251"}, //
		{Sqrt(Float(0.5)), "0.7853981633974483096156608458198757210492923498"}, // π/4, https://oeis.org/A003881
		{Ldexp(Sqrt(Float(3)), -1), "1.0471975511965977461542144610931676280"}, // π/3, https://oeis.org/A019670
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Asin(tt.arg); !strings.HasPrefix(tt.want, fmt.Sprint(got)[:30]) {
				t.Errorf("Asin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAsin_specials(t *testing.T) {
	tests := []struct {
		arg  Number
		want Number
	}{
		{Number{}, Number{}},
		{Float(neg0), Number{y: neg0}},
		{NaN(), Number{y: math.NaN()}},
		{Float(+2), Number{y: math.NaN()}},
		{Float(-2), Number{y: math.NaN()}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Asin(tt.arg); !same(got, tt.want) {
				t.Errorf("Asin() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestAcos(t *testing.T) {
	tests := []struct {
		arg  Number
		want string
	}{
		{Float(-1), "3.14159265358979323846264338327950288419716939937510582"}, // π, https://oeis.org/A000796
		{Float(0), "1.570796326794896619231321691639751442098584699687552910"}, // π/2, https://oeis.org/A019669
		{Float(0.5), "1.0471975511965977461542144610931676280657231331250352"}, // π/3, https://oeis.org/A019670
		{Sqrt(Float(0.5)), "0.7853981633974483096156608458198757210492923498"}, // π/4, https://oeis.org/A003881
		{Ldexp(Sqrt(Float(3)), -1), "0.5235987755982988730771072305465838140"}, // π/6, https://oeis.org/A019685
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Acos(tt.arg); !strings.HasPrefix(tt.want, fmt.Sprint(got)[:30]) {
				t.Errorf("Acos() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAcos_specials(t *testing.T) {
	tests := []struct {
		arg  Number
		want Number
	}{
		{NaN(), Number{y: math.NaN()}},
		{Float(+2), Number{y: math.NaN()}},
		{Float(-2), Number{y: math.NaN()}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Acos(tt.arg); !same(got, tt.want) {
				t.Errorf("Acos() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestAtan(t *testing.T) {
	tests := []struct {
		arg  Number
		want string
	}{
		{Float(1), "0.785398163397448309615660845819875721049292349843776455"}, // π/4, https://oeis.org/A003881
		{Float(-1), "-0.7853981633974483096156608458198757210492923498437764"}, //
		{Inv(Sqrt(Float(3))), "0.5235987755982988730771072305465838140328615"}, // π/6, https://oeis.org/A019685
		{Sqrt(Float(3)), "1.047197551196597746154214461093167628065723133125"}, // π/3, https://oeis.org/A019670
		{Pi, "1.262627255678911683444322083605698343508947670424383596973809"}, // atan(π), https://oeis.org/A232273
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Atan(tt.arg); !strings.HasPrefix(tt.want, fmt.Sprint(got)[:30]) {
				t.Errorf("Atan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAtan_specials(t *testing.T) {
	tests := []struct {
		arg  Number
		want Number
	}{
		{Number{}, Number{}},
		{Float(neg0), Number{y: neg0}},
		{NaN(), Number{y: math.NaN()}},
		{Inf(1), Number{+Pi.y / 2, +Pi.x / 2}},
		{Inf(-1), Number{-Pi.y / 2, -Pi.x / 2}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Atan(tt.arg); !same(got, tt.want) {
				t.Errorf("Atan() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestAtan2(t *testing.T) {
	tests := []struct {
		y    Number
		x    Number
		want string
	}{
		{Inf(1), Inf(1), "0.785398163397448309615660845819875721049292349843"}, // π/4, https://oeis.org/A003881
		{Float(1), Float(1), "0.78539816339744830961566084581987572104929234"}, // π/4, https://oeis.org/A003881
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Atan2(tt.y, tt.x); !strings.HasPrefix(tt.want, fmt.Sprint(got)[:30]) {
				t.Errorf("Atan2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAtan2_specials(t *testing.T) {
	tests := []struct {
		y    Number
		x    Number
		want Number
	}{
		{Float(0), Float(0), Number{}},
		{NaN(), Float(1), Number{y: math.NaN()}},
		{Float(1), NaN(), Number{y: math.NaN()}},
		{Float(neg0), Float(0), Number{neg0, 0}},
		{Float(0), Float(neg0), Pi},
		{Float(neg0), Float(neg0), Number{-Pi.y, -Pi.x}},
		{Float(1), Inf(+1), Number{}},
		{Float(1), Inf(-1), Pi},
		{Inf(+1), Float(0), Number{+Pi.y / 2, +Pi.x / 2}},
		{Inf(-1), Float(0), Number{-Pi.y / 2, -Pi.x / 2}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Atan2(tt.y, tt.x); !same(got, tt.want) {
				t.Errorf("Atan2() = %#v, want %#v", got, tt.want)
			}
		})
	}
}
