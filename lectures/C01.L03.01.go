package lectures

import (
	"math"

	. "github.com/maseology/goClimate/climate"
)

// NakedPlanetTemp (bare rock) returns the temperature based on a steady-state energy balance model, the simplist climate model
// assumes temperature is uniform
func NakedPlanetTemp(L, alpha, epsilon float64) float64 {
	// aSphere, aShadow := 4.*math.Pi*math.Pow(EarthRadius, 2.), math.Pi*math.Pow(EarthRadius, 2.) // area of earth, area of shadow
	// out := StefanBoltzmannLaw(T, EarthEmissivity) * aSphere
	// in := SolarConstant * (1. - EarthAlbedo) * aShadow

	// L: solar constant; alpha: albedo; epsilon: emissivity
	return math.Pow(L*(1.-alpha)/4./epsilon/StefanBoltzmann, 0.25)
}
