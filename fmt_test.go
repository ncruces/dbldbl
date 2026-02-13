package dbldbl

import (
	"fmt"
	"testing"
)

func TestNumber_String(t *testing.T) {
	tests := []struct {
		name string
		arg  Number
		want string
	}{
		{"E", E, "2.71828182845905"},
		{"Pi", Pi, "3.14159265358979"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.arg.String(); got != tt.want {
				t.Errorf("Number.String() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestNumber_GoString(t *testing.T) {
	tests := []struct {
		name string
		arg  Number
		want string
	}{
		{"E", E, "Number{2.718281828459045, +0x1.4d57ee2b1013ap-53}"},
		{"Pi", Pi, "Number{3.141592653589793, +0x1.1a62633145c07p-53}"},
		{"Sqrt2", Sqrt2, "Number{1.4142135623730951, -0x1.bdd3413b26456p-54}"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.arg.GoString(); got != tt.want {
				t.Errorf("Number.GoString() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func (n Number) Format(f fmt.State, verb rune) {
	if verb == 'v' && f.Flag('#') {
		f.Write([]byte(n.GoString()))
	} else if IsNaN(n) {
		f.Write([]byte("NaN"))
	} else {
		n.toBig().Format(f, verb)
	}
}
