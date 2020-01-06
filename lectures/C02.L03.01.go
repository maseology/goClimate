package lectures

import "fmt"

const (
	ts = 100  // years
	nc = 10   // n cells
	dx = 1.e5 // cell size [m]
	ny = 50000
	k  = 1.e4 // m/yr
	c  = 1.
)

// IceSheetModel is a simple 1D ice sheet model.
// snowfall in m/yr
func IceSheetModel(snowfall float64) {
	hs, qs := make([]float64, nc+2), make([]float64, nc+1)
	hs[0], hs[nc+1] = 0., 0.       // edge boundary conditions (ice is confined to a land mass)
	for t := 0; t <= ny; t += ts { // 100 yr timestep
		for f := 0; f <= nc; f++ {
			qs[f] = k * (hs[f] - hs[f+1]) / dx   // flux
			qs[f] *= (hs[f] + hs[f+1]) / 2. / dx // corrects for the aspect ratio of the grid cell
		}
		for i := 1; i <= nc; i++ {
			hs[i] += ts * (snowfall + qs[i-1] - qs[i])
		}
		fmt.Printf("%10d %10.5f\n", t, qs[nc]-qs[0])
	}
	fmt.Println(hs)
}
