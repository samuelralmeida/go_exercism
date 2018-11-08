package triangle

import (
	"math"
	"sort"
)

type Kind int

// Pick values for the following identifiers used by the test program.
const (
	NaT Kind = iota // not a triangle
	Equ             // equilateral
	Iso             // isosceles
	Sca             // scalene
)

func KindFromSides(a, b, c float64) Kind {
	sides := []float64{a, b, c}
	sort.Float64s(sides)
	if sides[0] <= 0 || sides[0]+sides[1] < sides[2] || math.IsNaN(sides[0]) || math.IsInf(sides[len(sides)-1], 0) {
		return NaT
	} else if sides[0] == sides[1] && sides[1] == sides[2] {
		return Equ
	} else if sides[0] == sides[1] || sides[1] == sides[2] {
		return Iso
	}
	return Sca
}
