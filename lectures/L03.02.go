package lectures

import "math"

// GreenHouseEffectTemp builds on naked earth model but with an atmosphere,
// assumes emissivity of the ground=emissivity of the atmosphere
func GreenHouseEffectTemp(L, alpha, epsilon float64) float64 {
	Ta := NakedPlanetTemp(L, alpha, epsilon) // temperature of the atmosphere
	return math.Pow(2., .25) * Ta            // temperature of the ground
}

// TwoLayeredAtmosphereTemp builds on greenhouse model but with a 2-layer atmospheric system,
// assumes emissivity of the ground=emissivity of the atmosphere
func TwoLayeredAtmosphereTemp(L, alpha, epsilon float64) float64 {
	Ta := NakedPlanetTemp(L, alpha, epsilon) // temperature of the upper atmosphere
	return math.Pow(3., .25) * Ta            // temperature of the ground
}

// NuclearWinterTemp builds on greenhouse model where atmospheric system absorbs all sunlight,
// assumes emissivity of the ground=emissivity of the atmosphere
func NuclearWinterTemp(L, alpha, epsilon float64) float64 {
	return NakedPlanetTemp(L, alpha, epsilon) // in this case, ground temperature equlibrates with the temperature of the atmosphere.
}
