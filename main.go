package main

import (
	"fmt"

	// . "github.com/maseology/goClimate/climate"
	"github.com/maseology/goClimate/lectures"
	"github.com/maseology/mmio"
)

func main() {
	tt := mmio.NewTimer()

	if false {
		func() {
			fmt.Printf("\nCOURSE 1, LECTURE 3 =============\n")
			// Lecture 3, model 1
			fmt.Println("\nBare-rock temperatures [K]:")
			fmt.Printf(" Venus = %.1f (actual %d)\n", lectures.NakedPlanetTemp(2600, .7, 1.), 700)
			fmt.Printf(" Earth = %.1f (actual %d)\n", lectures.NakedPlanetTemp(1350, .3, 1.), 295)
			fmt.Printf(" Mars  = %.1f (actual %d)\n", lectures.NakedPlanetTemp(600, .15, 1.), 259)
			// fmt.Printf(" Earth = %.1f\n", lectures.NakedPlanetTemp(SolarConstant, EarthAlbedo, EarthEmissivity))

			// Lecture 3, quiz 1
			fmt.Printf("\n Moon (at equator, sun directly above) = %.1f\n", lectures.NakedPlanetTemp(4*1350, .33, 1.))
			fmt.Printf(" Moon (at night) = %.1f\n", lectures.NakedPlanetTemp(0, .33, 1.))

			// Lecture 3, model 2: greenhouse effect
			fmt.Println("\nground temperatures (with atmosphere) [K]:")
			fmt.Printf(" Venus = %.1f\n", lectures.GreenHouseEffectTemp(2600, .7, 1.))
			fmt.Printf(" Earth = %.1f\n", lectures.GreenHouseEffectTemp(1350, .3, 1.))
			fmt.Printf(" Mars  = %.1f\n", lectures.GreenHouseEffectTemp(600, .15, 1.))

			// Lecture 3, model 2: greenhouse effect (2-layer system)
			fmt.Println("\nground temperatures (with 2-Layered atmosphere) [K]:")
			fmt.Printf(" Venus = %.1f\n", lectures.TwoLayeredAtmosphereTemp(2600, .7, 1.))
			fmt.Printf(" Earth = %.1f\n", lectures.TwoLayeredAtmosphereTemp(1350, .3, 1.))
			fmt.Printf(" Mars  = %.1f\n", lectures.TwoLayeredAtmosphereTemp(600, .15, 1.))
		}()

		func() {
			fmt.Printf("\nCOURSE 2, LECTURE 1 =============\n")
			fmt.Println("\nBare-rock temperatures (transient model) [K]:")
			for _, v := range lectures.TransientNakedEarthModel(1350., 0.3, 1., 1000, 0., 20) {
				fmt.Printf("  %.3f\n", v)
			}
			// fmt.Println("\nQuiz: Code Tricks: Heat Capacity, Time Steps, and Equilibration Time")
			// for _, v := range lectures.TransientNakedEarthModel(1350., 0.3, 1., 4000., 400., 10.) {
			// 	fmt.Printf("  %.3f\n", v)
			// }
		}()

		func() {
			fmt.Printf("\nCOURSE 2, LECTURE 2 =============\n")
			fmt.Println("\nIterative Runaway Ice-Albedo Feedback Model:")
			lectures.IterativeRunawayIceAlbedoFeedbackModel(1.)
		}()
	}

	tt.Print("\ngoClimate complete\n")
}
