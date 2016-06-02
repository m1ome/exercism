package triangle

import "math"

const testVersion = 2

type Kind struct {
	name string
}

var NaT = Kind{"NaT"}
var Equ = Kind{"Equ"}
var Iso = Kind{"Iso"}
var Sca = Kind{"Sca"}

func KindFromSides(a, b, c float64) Kind {

	if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) || math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) {
		return NaT
	}

	if a == 0 && b == 0 && c == 0 {
		return NaT
	}

	if a+b < c || a+c < b || b+c < a {
		return NaT
	}

	if (a == b && a == c) || (b == c && b == a) || (c == a && c == b) {
		return Equ
	}

	if a == b || a == c || b == c {
		return Iso
	}

	return Sca
}
