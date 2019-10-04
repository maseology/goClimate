package climate

import "math"

// StefanBoltzmannLaw returns the black (grey when epsilon<1) body radiation [W/m2]
func StefanBoltzmannLaw(T, epsilon float64) float64 {
	return epsilon * StefanBoltzmann * math.Pow(T, 4.)
}
