package dbldbl

import "strconv"

func (n Number) String() string {
	return strconv.FormatFloat(n.y, 'g', 15, 64)
}

func (n Number) GoString() string {
	return "Number{" + strconv.FormatFloat(n.y, 'g', -1, 64) + ", " + strconv.FormatFloat(n.x, 'e', -1, 64) + "}"
}
