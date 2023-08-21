package dbldbl

import (
	"fmt"
	"math/big"
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
				t.Errorf("Number.String() = %v, want %v", got, tt.want)
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
		{"E", E, "Number{2.718281828459045, 1.4456468917292502e-16}"},
		{"Pi", Pi, "Number{3.141592653589793, 1.2246467991473532e-16}"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.arg.GoString(); got != tt.want {
				t.Errorf("Number.GoString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func (n Number) Format(f fmt.State, verb rune) {
	if verb == 'v' && f.Flag('#') {
		fmt.Fprintf(f, "Number{%v, %v}", n.y, n.x)
	} else {
		n.toBig().Format(f, verb)
	}
}

func (n Number) toBig() *big.Float {
	var t big.Float
	t.SetFloat64(n.y)
	r := big.NewFloat(n.x).SetPrec(107)
	return r.Add(r, &t)
}
