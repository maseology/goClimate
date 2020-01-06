package climate

import (
	"math"

	"github.com/maseology/mmio"
)

const (
	dt = 5. // timestep [yr]

	delT2xCO2     = 3. // climate sensitivity: degrees (°C) per doubling of CO2
	TResponseTime = 20.

	initC02       = 290.
	pA            = .0225 // (CO2_exp) growth rate parameter [p/yr]. A value of 0.0225/yr, gives an atmospheric rise rate of 2.5 ppm when the pCO2 value is 400 ppm; a reasonable fit to the current pCO2 value and rate of rise.
	CO2RampExp    = .01   // (CO2RampExp)
	aerosolWm2now = -.75
)

// Nearterm simulates the near-term future of Earth’s climate as a function of the three main uncertainties: future CO2 emission, the climate sensitivity, and the current cooling impact of industrial aerosols and short-lived greenhouse gases (the masking effect).
// This is a time-stepping simulation with two stages: one with people releasing CO2 and generating aerosols, which we'll call Business-as-Usual, and a second, diverging from the first, in which we suddenly stop both of those things, at which point the CO2 largely persists, while the aerosols go away immediately.
// from: Lecture 5 of Coursera course: Global Warming II create your own models [2019]
func Nearterm() {

	pCO2 := initC02
	climatesensitivityWm2 := delT2xCO2 / 4. // climate sensitivity to forcing equivalent [W/m²]. doubling CO2 gives ~4 W/m²/2xCO2 of forcing
	// T := 0.

	radiativeForcing := func(pCO2 float64) float64 {
		return 4. * math.Log(pCO2/EqualibriumpCO2) / math.Ln2 // radiative forcing from CO2 (where the last factor, ln(pCO2/280) / ln(2), gives the number of doublings of the CO2 concentration over an initial value of 280.) (doubling CO2 gives ~4 W/m²/2xCO2 of forcing.)
	}

	// Business as usual
	yr, i := 1900., 0
	xr := make(map[float64]int)
	pCO2s, incCO2, rfCO2 := []float64{pCO2}, []float64{0.}, []float64{radiativeForcing(pCO2)}
	for {
		pCO2 = EqualibriumpCO2 + (pCO2-EqualibriumpCO2)*(1+pA*dt) // update pCO2
		pCO2s = append(pCO2s, pCO2)
		incCO2 = append(incCO2, (pCO2-pCO2s[len(pCO2s)-2])/dt)
		rfCO2 = append(rfCO2, radiativeForcing(pCO2))
		xr[yr] = i
		i++
		yr += dt
		if yr > 2100. {
			break
		}
	}

	i2015, ok := xr[2015.]
	if !ok {
		panic("0001")
	}
	aerosolCoeff := aerosolWm2now / ((pCO2s[i2015] - pCO2s[i2015-1]) / dt) // -.3

	n := len(pCO2s)
	yr, i = 1900., 1
	yrs := make([]float64, n)
	yrs[0] = yr
	rfmask, rftot, T := make([]float64, n), make([]float64, n), make([]float64, n)
	for {
		rfmask[i] = math.Max(incCO2[i]*aerosolCoeff, aerosolWm2now) // radiative impact of short-lived industrial stuff, primarily the sum of cooling from sulfate aerosols and warming from short-lived greenhouse gases like methane. [W/m²/yr]
		// The model looks more reasonable in the future if we assume that this function only works for the past, but that further increases in CO2 emissions don't lead to still further increases in aerosols. It is easy for coal plants to clean up their sulfur, which leads to the aerosols (and lots of air quality problems), without cutting the carbon emissions.
		rftot[i] = rfCO2[i] + rfmask[i]
		Teq := rftot[i] * climatesensitivityWm2
		T[i] = T[i-1] + (Teq-T[i-1])*dt/TResponseTime
		yr += dt
		yrs[i] = yr
		i++
		if yr > 2100. {
			break
		}
	}

	// Rampdown
	pCO2sRD, rfCO2rd := make([]float64, n), make([]float64, n)
	rfmaskRD, rftotRD, Trd := make([]float64, n), make([]float64, n), make([]float64, n)
	copy(pCO2sRD, pCO2s)
	copy(rfCO2rd, rfCO2)
	copy(rfmaskRD, rfmask)
	copy(rftotRD, rftot)
	copy(Trd, T)
	for i := i2015; i < n; i++ {
		pCO2sRD[i] = pCO2sRD[i-1] + (1.2*EqualibriumpCO2-pCO2sRD[i-1])*CO2RampExp*dt // where the CO2 concentration is now relaxing toward a higher concentration than the initial, due to the long-term change in ocean chemistry (acidification). The time scale for the CO2 invasion into the ocean is 100 years (1 / 0.01), which is a composite of fast equilibration times with the surface ocean and the land biosphere, and slower with the deep ocean.
		rfCO2rd[i] = radiativeForcing(pCO2sRD[i])
		rftotRD[i] = rfCO2rd[i]
		rfmaskRD[i] = 0.
		Teq := rftotRD[i] * climatesensitivityWm2
		Trd[i] = Trd[i-1] + (Teq-Trd[i-1])*dt/TResponseTime
	}

	mmio.Line("out1.png", yrs, map[string][]float64{"pCO2": pCO2s, "pCO2 (ramp down)": pCO2sRD}, 12.)
	mmio.Line("out2.png", yrs, map[string][]float64{"RF tot": rftot, "RF tot (ramp down)": rftotRD, "RF mask (ramp down)": rfmaskRD}, 12.)
	mmio.Line("out3.png", yrs, map[string][]float64{"delTemp": T, "delTemp (ramp down)": Trd}, 12.)

	// old
	// armageddon := false // if true, there no antrhopogenic forcing
	// for {               // annual timesteps

	// 	rfmask := func() float64 {
	// 		const (
	// 			rfmask2015 = -.75  // [W/m²]
	// 			pCO22015   = 380.  // [ppm]
	// 			pCO22014   = 377.5 // [ppm]
	// 		)
	// 		pB := rfmask2015 / (pCO22015 - pCO22014) // -.3
	// 		rfscaled := pB * (pCO22015 - pCO22014)
	// 		return math.Max(rfmask2015, rfscaled)
	// 	}

	// 	rftot := rf + rfmask() // Compute the total radiative forcing

	// 	// T += (Teq - T) / Tresponse20 * xyrs / dt

	// 	if armageddon { // 8. Beginning in the year when pCO2 hits 400, make a new column or list with pCO2 values that decrease rather than increase, using the relationship
	// 		pCO2 += (340. - pCO2) * .01 * dt
	// 		rf = radiativeForcing(pCO2)

	// 	}

	// 	fmt.Println(rftot)
	// 	yr += dt
	// 	if pCO2 > 400. {
	// 		armageddon = true
	// 	}
	// 	if yr > 2100. {
	// 		break
	// 	}
	// }
}
