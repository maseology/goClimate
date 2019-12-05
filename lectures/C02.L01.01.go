package lectures

import (
	"github.com/maseology/goClimate/climate"
)

// TransientNakedEarthModel
func TransientNakedEarthModel(L, alpha, epsilon, oceanDepth, T0, ts float64) []float64 {
	T := T0                                    // initial temperature [K]
	hcap := 4.184 * 1000. * 1000. * oceanDepth // heat capacity of water of given depth [J/K/m²]
	tf := 86400. * 365.24 * float64(ts)        // temporal conversion factor [s/ts]
	heatin := func() float64 {
		return L * (1. - alpha) / 4.
	}()
	heatout := func(T float64) float64 {
		return climate.StefanBoltzmannLaw(T, epsilon)
	}
	hcon := heatout(T) * tf // initial heat content (J/m²)
	fout := []float64{}
	for t := 0.; t <= 2000.; t += ts {
		fout = append(fout, T)
		// if t >= 100 {
		// 	fmt.Println(t, T, hcon, heatin, heatout(T))
		// 	break
		// }
		hcon += (heatin - heatout(T)) * tf
		T = hcon / hcap
	}
	return fout
}
