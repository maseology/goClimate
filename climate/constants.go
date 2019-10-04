package climate

const (
	SolarConstant = 1360.8 // (L) 1360.8 ± 0.5  W/m²

	StefanBoltzmann = 5.670374419e-8 // (sigma) W m-2 K-4

	EarthTemperature = 284.42  // K
	EarthEmissivity  = 1.      // assumed close to black body (as is for most solids and liquids)
	EarthAlbedo      = .3      // (alpha)
	EarthRadius      = 6371.e3 // m

	AtmosLapseRate = 6.   // K/km (10. for dry air)
	AtmosHeight    = 15.  // stratospheric height km
	AtmosCO2       = 400. // concentration of CO2 ppm
	AtmosCH4       = 1.7  // concentration of methane ppm
	AtmosRH        = .8   // relative humidity
)
