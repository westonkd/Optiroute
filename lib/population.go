package geneticTSP

import (
	"errors"
	"math/rand"
	"time"
	"fmt"
	"github.com/kr/pretty"
)

// Represents a population of Chromosomes (tours)
type Population struct {
	Chromosomes    []Chromosome // Chromosomes in the population
	Recalculate    bool
	MutThreshold   float32        // Threshold for mutation
	CrossThreshold float32        // Threshold for crossover
	Matrix         *FitnessMatrix // Matrix of distances
}

// SetMatrix sets the cost matrix for the entire population
func (self *Population) SetMatrix(matrix *FitnessMatrix) {
	self.Matrix = matrix
}

// Size returns the population count
func (self *Population) Size() (size int) {
	size = len(self.Chromosomes)
	return
}

// Add a chromosome to the population
func (self *Population) Add(chromo *Chromosome) {
	self.Chromosomes = append(self.Chromosomes, *chromo)
}

// Get a chromosome from the population
func (self *Population) Get(index int) (*Chromosome, error) {
	if index < len(self.Chromosomes)-1 && index > -1 {
		return &self.Chromosomes[index], nil
	}

	return nil, errors.New("Index out of range.")
}

// GetFittest returns a pointer to the fittest
// chromosome in the population
func (self *Population) GetFittest() *Chromosome {
	fittest := self.Chromosomes[0]

	// Iterate over each chromosome to check fitness
	for _, val := range self.Chromosomes {
		if fittest.Fitness() <= val.Fitness() {
			fittest = val
		}
	}

	return &fittest
}

// TournamentSelect
func (self *Population) TournamentSelect(tournamentSize int) *Chromosome {

	// Make a new population of size tournamentSize
	tournamentPop := Population{}

	// Randomly add chromosomes to the new population
	for i := 0; i < tournamentSize; i++ {
		randomChromo, error := self.Get(rand.Intn(self.Size()))

		// If there was an error just try again later
		if error != nil {
			i--
			continue
		}

		tournamentPop.Add(randomChromo)
	}

	// Select the best chromosome
	return tournamentPop.GetFittest()
}

// Mutate iterates over all chromosomes and performs a random swap of locations
// if a randomly chosen number is bellow the mutation threshold level. Note the
// possibility of choosing the two swap values as the same number. This is less
// noticeable for larger chromosomes.
func (self *Population) Mutate() {
	// Seed the random number generator
	rand.Seed(time.Now().Unix())

	// Loop through each chromosome
	for i := range self.Chromosomes {
		// Generate a random number between 0 and 1
		mutVal := rand.Float32()

		// Mutate if the random number is bellow threshold
		if mutVal < self.MutThreshold {
			self.Chromosomes[i].RandSwap()
		}
	}

}

// Crossover
func (self *Population) Crossover(parentOne, parentTwo *Chromosome) (*Chromosome, error) {
	// Child locations
	initLocation, error := parentOne.Get(0)

	// If there was an error retrieving the first entry
	if error != nil {
		return nil, error
	}

	childLocations := []Location{
		initLocation,
	}

	 count := 0

	// While the child is not yet complete
	for (len(childLocations) < parentOne.Length()) {
		count++

		// Length of child
		childLen := len(childLocations)

		// Get the next valid location from each parent
		firstVal := parentOne.IndexOf(childLocations[childLen - 1].Id)
		secondVal := parentTwo.IndexOf(childLocations[childLen - 1].Id)

		if firstVal < 0 || secondVal < 0 {
			fmt.Println("+++++++++BANANA+++++++++++ ", childLen)
			pretty.Println(childLocations[childLen - 1]) // this is always empty so on the 5th iter there is an issue

			break
		}

		optionOne, eP1 := parentOne.Get((firstVal + 1) % childLen)
		optionTwo, eP2 := parentTwo.Get((secondVal + 1) % childLen)

		// Check for errors
		if eP1 != nil || !self.isValidLoc(optionOne.Id, childLocations)  {
			optionOne = self.nextValidLocation(childLocations, parentOne)
			//pretty.Println("Pulling from backup for option one: ", optionOne.Id)
		}

		if eP2 != nil || !self.isValidLoc(optionTwo.Id, childLocations){
			optionTwo = self.nextValidLocation(childLocations, parentTwo)
			//pretty.Println("Pulling from backup for option two: ", optionTwo.Id)
		}

		// Pick the shorter option
		distOne := parentOne.Matrix.GetDistance(childLocations[childLen - 1], optionOne)
		distTwo := parentOne.Matrix.GetDistance(childLocations[childLen - 1], optionTwo)

		//pretty.Println("Distances: ", distOne, " vs ", distTwo)

		if count == 5 {
			fmt.Println("aaaaaaaaaaaaaaaa")
			pretty.Println(optionOne)
			pretty.Println(optionTwo)
			fmt.Println("aaaaaaaaaaaaaaaa")
		}

		// Add the best location
		if distOne < distTwo {
			//pretty.Println("Chose 1")
			childLocations = append(childLocations, optionOne)
		} else {
			//pretty.Println("Chose 2")
			childLocations = append(childLocations, optionTwo)
		}
	}

	// Construct the chromosome
	childChromo := NewChromosome(childLocations, parentOne.Matrix)

	return childChromo, nil
}

// Checks if a given location id is valid for a set of locations
func (self *Population) isValidLoc(locId int, locations []Location) bool {
	for _, val := range locations {
		if val.Id == locId {
			return false
		}
	}

	return true
}

//TODO fix this method. It may just be that this method is not getting called when it supposed to
// (i.e. one of the parents really do have a valid location)
// Returns the next valid location in the set 2..n in the case
// that neither parent had a next valid location
func (self *Population) nextValidLocation(locations []Location, parent *Chromosome) Location {
	fmt.Println("Looking in")
	pretty.Println(locations)
	pretty.Println(parent)

	// For each number in the set 2..n
	for locId := 2; locId <= parent.Length(); locId++ {
		valid := true

		// Check if the locations already has this item
		for _, val := range locations {
			if locId == val.Id {
				valid = false
				break
			}
		}

		if valid && parent.IndexOf(locId) >= 0 && parent.IndexOf(locId) < parent.Length() {
			retVal, _ := parent.Get(parent.IndexOf(locId))
			return retVal
		} else {
			fmt.Println("Error: error in nextvalidlocation ", locId, valid)
		}
	}

	// There was an error
	return Location{Id:777}
}

