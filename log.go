package dbldbl

func Log(n Number) Number {
	switch {
	case n.y < 0:
		return NaN()
	case n.y == 0:
		return Inf(-1)
	case !isFinite(n.y):
		return n
	case n == Number{1, 0}:
		return Number{}
	}

	negate := false
	if n.y < 1 {
		n = Div(Float(1), n)
		negate = true
	}

	var shift int8
	for n.y < 8 {
		n = Sqr(n)
		shift--
	}

	n = Shift(log(n), shift)
	if negate {
		n = Neg(n)
	}
	return n
}

func log(n Number) Number {
	// Sasaki and Kanada's formula, with q=q⁴
	// to avoid the ∜q̅ term in the θ₂(q) expansion:
	// https://maths-people.anu.edu.au/~brent/pd/RNC7t.pdf
	q := Div(Float(1), n)
	q4 := Sqr(Sqr(q))
	q8 := Sqr(q4)
	q16 := Sqr(q8)
	q24 := Mul(q16, q8)
	th2 := Shift(Mul(q, AddFloat(Add(q8, q24), 1)), 1)
	th3 := AddFloat(Shift(Add(q4, q16), 1), 1)
	return Div(Shift(Pi, -2), agm(Sqr(th2), Sqr(th3)))
}

func agm(a, g Number) Number {
	// https://en.wikipedia.org/wiki/Arithmetic%E2%80%93geometric_mean
	for {
		t := Shift(Add(a, g), -1)
		if a == t {
			return a
		}
		g = Sqrt(Mul(a, g))
		a = t
	}
}

func pi() Number {
	// https://en.wikipedia.org/wiki/Arithmetic%E2%80%93geometric_mean
	a := Float(1)
	g := Sqrt(Float(0.5))
	series := Float(1)

	for n := int8(1); n < 127; n++ {
		t := Shift(Add(a, g), -1)
		if a == t {
			break
		}
		g = Sqrt(Mul(a, g))
		a = t
		series = Sub(series, Shift(Sub(Sqr(a), Sqr(g)), n+1))
	}

	return Shift(Div(Sqr(a), series), 2)
}
