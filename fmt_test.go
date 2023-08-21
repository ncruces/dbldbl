package dbldbl

import (
	"fmt"
	"math/big"
)

func (n Number) toBig() *big.Float {
	var t big.Float
	t.SetFloat64(n.y)
	r := big.NewFloat(n.x).SetPrec(107)
	return r.Add(r, &t)
}

func (n Number) String() string {
	return n.toBig().String()
}

func (n Number) GoString() string {
	return fmt.Sprintf("Number{%v, %v}", n.y, n.x)
}

func (n Number) Format(f fmt.State, verb rune) {
	if verb == 'v' && f.Flag('#') {
		fmt.Fprintf(f, "Number{%v, %v}", n.y, n.x)
	} else {
		n.toBig().Format(f, verb)
	}
}
