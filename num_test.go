package dbldbl

import (
	"reflect"
	"testing"
)

func TestInt(t *testing.T) {
	tests := []struct {
		args int64
		want Number
	}{
		{+1e03 + 1, Number{y: +1e03 + 1}},
		{-1e06 + 1, Number{y: -1e06 + 1}},
		{+1e09 + 1, Number{y: +1e09 + 1}},
		{-1e12 + 1, Number{y: -1e12 + 1}},
		{+1e15 + 1, Number{y: +1e15 + 1}},
		{-1e18 + 1, Number{y: -1e18, x: +1}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Int(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint(t *testing.T) {
	tests := []struct {
		args uint64
		want Number
	}{
		{1e03 + 1, Number{y: 1e03 + 1}},
		{1e06 + 1, Number{y: 1e06 + 1}},
		{1e09 + 1, Number{y: 1e09 + 1}},
		{1e12 + 1, Number{y: 1e12 + 1}},
		{1e15 + 1, Number{y: 1e15 + 1}},
		{1e18 + 1, Number{y: 1e18, x: +1}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Uint(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int() = %v, want %v", got, tt.want)
			}
		})
	}
}
