package displacement

import "math"

// AccelerationVec calculates the acceleration from Velocity.
func AccelerationVec(v0, vf float64) func(float64) float64 {
	fx := func(t float64) float64 {
		return (vf - v0) / math.Abs(t)
	}
	return fx
}

//AccelerationDisp calculates the acceleration from Displacement.
func AccelerationDisp(v0, vf, s0, sf float64) float64 {

	return (math.Pow(vf, 2) - math.Pow(v0, 2)) / (2 * (sf - s0))

}


