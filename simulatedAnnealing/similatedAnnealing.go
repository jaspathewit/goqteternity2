// similatedAnnealing.go
package simulatedAnnealing

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"bitbucket.com/jaspathewit/eternity2/model"
	"bitbucket.com/jaspathewit/eternity2/option"
	log "github.com/cihub/seelog"
)

//// Do the actual work
//func Process(pmodel *model.Model) error {
//	var err error
//
//	log.Info("Processing Started\n")
//
//	surface := model.NewSurface()
//	surface.FillFromTileSet()
//
//	txt := surface.String()
//	log.Infof("Best Surface found: \n%s", txt)
//	// bestSurface := surface.Clone()
//
//	// update the model with the current best surface
//	pmodel.SetSurface(surface)
//
//	log.Info("Processing Finished\n")
//
//	return err
//}

func Process(pmodel *model.Model) error {
	var err error

	log.Info("Processing Started\n")

	// initilise the random seed
	rand.Seed(time.Now().UTC().UnixNano())

	maxIterations := 5000000
	maxEnergy := 0

	// s ← s0;									// Initial state.
	surface := model.NewSurface()
	surface.FillFromTileSet()

	//  e ← E(s)								// Initial energy.
	energy := surface.Energy()
	// sbest ← s                              	// Initial "best" solution
	bestSurface := surface.Clone()

	// update the model with the current best surface
	pmodel.SetSurface(bestSurface)

	// ebest ← e                              	// Initial "best" energy
	bestEnergy := energy

	// k ← 0                                    // Energy evaluation count.
	// while k < kmax and e > emax              // While time left & not good enough:
	for k := 0; (k < maxIterations) && (energy > maxEnergy); k++ {
		//  T ← temperature(k/kmax)             // Temperature calculation.
		temp := temperature(k, maxIterations)

		//  snew ← neighbour(s)                             // Pick some neighbour.
		newSurface := surface.Clone()
		modify(newSurface)

		if option.Validate {
			valid, err := newSurface.IsValid()
			if !valid {
				txt := newSurface.String()
				fmt.Printf("Invalid Surface \n %s", txt)
				panic(fmt.Sprintf("Surface is no longer valid: %s", err))
			}
		}

		//  enew ← E(snew)                                  // Compute its energy.
		newEnergy := newSurface.Energy()

		//  if P(e, enew, T) > random() then                // Should we move to it?
		prob := acceptanceProbability(energy, newEnergy, temp)
		if prob > rand.Float64() {
			//    s ← snew; e ← enew                            // Yes, change state.
			surface = newSurface
			energy = newEnergy
		}

		//  if enew < ebest then                            // Is this a new best?
		if newEnergy < bestEnergy {
			//    sbest ← snew; ebest ← enew                    // Save 'new neighbour' to 'best found'.
			bestSurface = newSurface.Clone()

			// update the model with the current best surface
			pmodel.SetSurface(bestSurface)

			bestEnergy = newEnergy

			log.Infof("Current best energy Iteration: %d Temp: %f Energy: %d", k, temp, bestEnergy)

			//  k ← k + 1                                       // One more evaluation done
		}
	}

	//return sbest                                      // Return the best solution found.
	txt := bestSurface.String()
	log.Infof("Best Surface found: \n%s", txt)

	log.Info("Processing Finished\n")

	return err
}

// P(e, e', T) was defined as 1 if e' < e, and exp(-(e'-e)/T)
// Calculate the propability of moving to this higher energy state
func acceptanceProbability(energy int, newEnergy int, temp float64) float64 {
	result := 1.0

	if newEnergy > energy {
		delta := float64(newEnergy - energy)
		result = math.Exp(-(delta) / temp)
	}

	return result
}

// Calculate the temparature factor at this step
func temperature(k int, kmax int) float64 {
	t0 := 100.0
	lam := 0.9999

	result := t0 * math.Pow(lam, float64(k))

	//if k < kmax {
	//	result = float64(k) * math.Exp((-1)*lam*float64(k))
	//}

	// linear cooling
	/* t0 := 10000.0
	lam := t0 / float64(kmax)

	result := t0 - (lam * float64(k))

	if result < 0 {
		result = 0
	} */

	return result
}

// perform a modification to the surface to create the neighbour
func modify(surface *model.Surface) {
	// select the random tile
	srcTileNumber := randomTileNumber()

	// get the source tile
	location := surface.GetLocation(srcTileNumber)
	srcTile := surface.GetTile(location.Row, location.Column)
	typeOf := srcTile.TypeOf

	switch typeOf {
	case model.TYPEOF_CORNER:
		modifyCorner(surface, srcTileNumber)
		break
	case model.TYPEOF_EDGE:
		modifyEdge(surface, srcTileNumber)
		break
	case model.TYPEOF_MIDDLE:
		modifyMiddle(surface, srcTileNumber)
		break
	default:
		panic("Border tile!!!!")
		break
	}
}

// perform a modification on a middle tile
func modifyMiddle(surface *model.Surface, srcTileNumber byte) {
	// decide if we are going to switch or rotate
	prob := rand.Intn(2)
	if prob == 0 {
		if option.Verbose {
			srcNum := 1 + int(srcTileNumber)
			fmt.Printf("rotateMiddle src: %0.3d\n", srcNum)
		}
		surface.RotateTile(srcTileNumber)
	} else {
		// select a destination tile
		destTileNumber := randomMiddleTileNumber()
		if option.Verbose {
			srcNum := 1 + int(srcTileNumber)
			destNum := 1 + int(destTileNumber)
			fmt.Printf("modifyMiddle src: %0.3d, dest: %0.3d\n", srcNum, destNum)
		}
		surface.SwitchTiles(srcTileNumber, destTileNumber)
	}
}

// perform a modification on a side tile
func modifyEdge(surface *model.Surface, srcTileNumber byte) {
	// select a destination tile
	destTileNumber := randomEdgeTileNumber()
	if option.Verbose {
		srcNum := 1 + int(srcTileNumber)
		destNum := 1 + int(destTileNumber)
		fmt.Printf("modifyEdge   src: %0.3d, dest: %0.3d\n", srcNum, destNum)
	}
	surface.SwitchTiles(srcTileNumber, destTileNumber)
}

// perform a modification on a side tile
func modifyCorner(surface *model.Surface, srcTileNumber byte) {
	// select a destination tile
	destTileNumber := randomCornerTileNumber()
	if option.Verbose {
		srcNum := 1 + int(srcTileNumber)
		destNum := 1 + int(destTileNumber)
		fmt.Printf("modifyCorner src: %d, dest: %d\n", srcNum, destNum)
	}
	surface.SwitchTiles(srcTileNumber, destTileNumber)
}

// Returns a random tile number
func randomTileNumber() byte {
	number := rand.Intn(256)
	return byte(number)
}

// Returns a random tile number for middle tiles
func randomMiddleTileNumber() byte {
	// +60 for the corner tiles
	number := rand.Intn(196) + 60
	return byte(number)
}

// Returns a random tile number for side tiles
func randomEdgeTileNumber() byte {
	// +4 for the corner tiles
	number := rand.Intn(56) + 4
	return byte(number)
}

// Returns a random tile number for corner tiles
func randomCornerTileNumber() byte {
	number := rand.Intn(4)
	return byte(number)
}
