package dbldbl

import "strconv"

// String implements [fmt.Stringer].
func (n Number) String() string {
	return strconv.FormatFloat(n.y, 'g', 15, 64)
}

// GoString implements [fmt.GoStringer].
func (n Number) GoString() string {
	return "Number{" + strconv.FormatFloat(n.y, 'g', -1, 64) + ", " + strconv.FormatFloat(n.x, 'e', -1, 64) + "}"
}
