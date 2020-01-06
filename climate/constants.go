package climate

const (
	SolarConstant = 1360.8 // (L) 1360.8 ± 0.5 W/m²

	StefanBoltzmann = 5.670374419e-8 // (sigma) [W m-2 K-4]

	EarthTemperature  = 284.42  // [K]
	EarthEmissivity   = 1.      // assumed close to black body (as is for most solids and liquids)
	EarthAlbedo       = .3      // (alpha)
	EarthRadius       = 6371.e3 // [m]
	EarthCO2Degassing = 7.5e12  // [mol K yr-1]

	AtmosLapseRate = 6.   // [K/km] (10. for dry air)
	AtmosHeight    = 15.  // stratospheric height [km]
	AtmospCO2      = 400. // concentration of CO2 [ppm]
	AtmosCH4       = 1.7  // concentration of methane [ppm]
	AtmosRH        = .8   // relative humidity [-]

	EqualibriumpCO2 = 280. // natural/baseline CO2 concentration [ppm]

	CloudDropRadius = 1.e-6 // [m] seeded droplets = ~1-5μm

	μm = 3. // nothing, just seeing if μ can be a variable name
)
