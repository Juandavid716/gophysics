package displacement

import (
	"math"
)

// DisplaceAcc calculates the displacement from acceleration.
func DisplaceAcc(v0, a, s0 float64) func(float64) float64 {
	fx := func(t float64) float64 {
		return 0.5*a*math.Pow(t, 2) + v0*t + s0
	}
	return fx
}

//DisplaceVec calculates the displacement from Velocity.
func DisplaceVec(v0, vf, s0 float64) func(float64) float64 {
	fx := func(t float64) float64 {
		return 0.5*(v0+vf)*t + s0
	}
	return fx
}
