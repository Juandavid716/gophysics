package ellipse

import (
	"math"
)

// Init where center is 0,0
type Init struct {
	A, B float64
}

// Pythagorean theorem
func (e *Init) Pythagorean() float64 {
	return (math.Sqrt(math.Pow(e.A, 2) + math.Pow(e.B, 2)))
}

// GenDisplaceFn of one point
func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	fx := func(t float64) float64 {
		return 0.5*a*math.Pow(t, 2) + v0*t + s0
	}
	return fx
}
