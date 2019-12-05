package lectures

import (
	"fmt"
	"math"

	"github.com/maseology/goClimate/climate"
	"github.com/maseology/mmio"
)

const ( // see: week2 'Iterative Runaway Ice-Albedo Feedback Model.xlsx'
	ml    = 1.5
	bl    = -322.5
	ma    = -.01
	ba    = 2.8
	niter = 100
)

// IterativeRunawayIceAlbedoFeedbackModel is an Iterative Relaxation to Consistent T and Albedo Given L
func IterativeRunawayIceAlbedoFeedbackModel(epsilon float64) {

	ilist := make([]float64, niter)
	for i := 0; i < niter; i++ {
		ilist[i] = float64(i + 1)
	}
	xlist, tlist := []float64{}, []float64{}
	ttlist := [][]float64{}

	// cooling sweep
	albedo := 0.15
	for L := 1600.; L >= 1200.; L -= 10. { // W/m²
		t := 0.
		tlst := make([]float64, niter)
		for i := 0; i < niter; i++ {
			t = math.Pow(L*(1.-albedo)/4./climate.StefanBoltzmann/epsilon, .25)
			anew := math.Max(math.Min(ma*t+ba, .65), .15)
			// if math.Abs(anew-albedo) < 0.00001 {
			// 	break
			// }
			albedo = anew
			tlst[i] = t
		}
		xlist = append(xlist, L)
		tlist = append(tlist, t)
		ttlist = append(ttlist, tlst)
		// if albedo == .65 { // Code Trick: Hysteresis Into and Out Of the Snowball
		// 	fmt.Println("cool", L, albedo)
		// }
	}
	mmio.LinePoints("LvT_cooling.png", ilist, ttlist)

	// warming sweep
	ttlist = [][]float64{}                 // redim
	for L := 1200.; L <= 1600.; L += 10. { // W/m²
		t := 0.
		tlst := make([]float64, niter)
		for i := 0; i < niter; i++ {
			t = math.Pow(L*(1.-albedo)/4./climate.StefanBoltzmann/epsilon, .25)
			anew := math.Max(math.Min(ma*t+ba, .65), .15)
			// if math.Abs(anew-albedo) < 0.00001 {
			// 	break
			// }
			albedo = anew
			tlst[i] = t
		}
		xlist = append(xlist, L)
		tlist = append(tlist, t)
		ttlist = append(ttlist, tlst)
		// if albedo == .15 { // Code Trick: Hysteresis Into and Out Of the Snowball
		// 	fmt.Println("warm", L, albedo)
		// }
	}

	mmio.LinePoints1("LvT.png", xlist, tlist) // shows hysteresis
	mmio.LinePoints("LvT_warming.png", ilist, ttlist)

	// Code check
	inner(1280., .15, 1., 1)   // 263.17610875592277 0.16823891244077194
	inner(1280., .15, 1., 100) // 255.45242794389384 0.24547572056106137
	inner(1200., .65, 1., 100) // 207.4443257628261 0.65

}

func inner(L, albedo, epsilon float64, niter int) {
	t := 0.
	for i := 0; i < niter; i++ {
		t = math.Pow(L*(1.-albedo)/4./climate.StefanBoltzmann/epsilon, .25)
		anew := math.Max(math.Min(ma*t+ba, .65), .15)
		// if math.Abs(anew-albedo) < 0.00001 {
		// 	break
		// }
		albedo = anew
	}
	fmt.Println(t, albedo)
}
