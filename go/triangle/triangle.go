package triangle

import "math"

type Kind int

const (
	Equ Kind = iota
	Iso
	Sca
	NaT
)

func KindFromSides(a, b, c float64) Kind {
	// I create an array with the sides so that I can easily loop on them without using multiple if
	triangle := []float64{a, b, c}
	for _, side := range triangle {
		if side < 0 || math.IsNaN(side) {
			return NaT
		}
	}

	switch {
	case a+b <= c || a+c <= b || b+c <= a:
		return NaT
	case a == b && b == c:
		return Equ
	case a == b || a == c || b == c:
		return Iso
	default:
		return Sca
	}
}
