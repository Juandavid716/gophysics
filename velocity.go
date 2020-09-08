package displacement

import "math"

// VelocityAcc calculates the Velocity from acceleration.
func VelocityAcc(v0, a float64) func(float64) float64 {
	fx := func(t float64) float64 {
		return v0 + a*t
	}
	return fx
}

//VelocityDisp calculates the Velocity from Displacement.
func VelocityDisp(v0, a, s0, sf float64) float64 {

	return math.Pow(v0, 2) + 2*a*(sf-s0)

}
