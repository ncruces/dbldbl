package dbldbl

import "strconv"

// String implements [fmt.Stringer].
func (n Number) String() string {
	return strconv.FormatFloat(n.y, 'g', 15, 64)
}

// GoString implements [fmt.GoStringer].
func (n Number) GoString() string {
	y := strconv.FormatFloat(n.y, 'g', -1, 64)
	x := strconv.FormatFloat(n.x, 'x', -1, 64)
	sep := ", +"
	if x[0] == '-' {
		sep = sep[:2]
	}
	return "Number{" + y + sep + x + "}"
}
