package geneticTSP

import (
	"errors"
	"math/rand"
	"github.com/kr/pretty"
)

// Represents a population of Chromosomes (tours)
type Population struct {
	Chromosomes    []Chromosome // Chromosomes in the population
	Recalculate    bool
	MutThreshold   float32        // Threshold for mutation
	CrossThreshold float32        // Threshold for crossover
	Matrix         *FitnessMatrix // Matrix of distances
	IDCounter int
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

func (self *Population) IndexOf(id int) int {
	for i, val := range self.Chromosomes {
		if val.Id == id {
			return i
		}
	}

	return -1
}

// TournamentSelect
func (self *Population) TournamentSelect(tournamentSize int, prevChosen_opt ...int) *Chromosome {

	// Set the dault value if non was specified
	prevChosen := -999

	if len(prevChosen_opt) > 0 {
		prevChosen = prevChosen_opt[0]
	}

	// Make an empty list of chromosomes
	chromosomes := make([]Chromosome, 0)

	for i := 0; i < tournamentSize; i++ {
		randomIndex := rand.Intn(len(self.Chromosomes) - 1)

		randChromo, err := self.Get(randomIndex)

		if err != nil {
			pretty.Println(err)
		}

		if prevChosen == randChromo.Distance() {
			i--
			continue
		}

		chromosomes = append(chromosomes, *randChromo)
	}

	// create the tournament population
	tournPop := Population{
		Chromosomes: chromosomes,
		Matrix: self.Matrix,
	}

	return tournPop.GetFittest()
}

// Mutate iterates over all chromosomes and performs a random swap of locations
// if a randomly chosen number is bellow the mutation threshold level. Note the
// possibility of choosing the two swap values as the same number. This is less
// noticeable for larger chromosomes.
func (self *Population) Mutate() {
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

// RSMutate randomly reverses a sub slice in the slice
// of locations.
func (self *Population) RSMutate() {
	// Loop through each chromosome
	for i := range self.Chromosomes {
		// Generate a random number between 0 and 1
		mutVal := rand.Float32()

		// Mutate if the random number is bellow threshold
		if mutVal < self.MutThreshold {
			self.Chromosomes[i].RandInvert()
		}
	}
}

// Simple crossover picks a random subsection of each parent and combines them
// without duplicating locations in the route.
func (self *Population) SimpleCrossover(parentOne, parentTwo *Chromosome) (*Chromosome, error) {
	startPos := rand.Intn(parentOne.Length() - 1)
	endPos := rand.Intn(parentOne.Length() - 1)

	// Make the child locations list
	childLocations := make([]Location, parentOne.Length())

	if startPos > endPos {
		//Swap
		temp := endPos
		endPos = startPos
		startPos = temp
	}

	// Add the subset to the child
	for i := startPos; i <= endPos; i++ {
		copy := parentOne.Locations[i]
		childLocations[i] = copy
	}

	childIndex := 0
	// Add the missing values from parent 2
	for _, val := range parentTwo.Locations {
		for childLocations[childIndex].Id != 0 {
			childIndex++
		}

		if self.isValidId(val.Id, childLocations) {
			copy := val
			childLocations[childIndex] = copy
		}
	}

	// Create the resulting chromosome
	child := &Chromosome{
		Locations: childLocations,
		Matrix: parentOne.Matrix,
		Id: self.IDCounter,
	}

	self.IDCounter++

	return child, nil
}

// Implementation of SCX crossover. This is not the optimal crossover operator for the
// traveling salesman problem.
func (self *Population) Crossover(parentOne, parentTwo *Chromosome) (*Chromosome, error) {
	// Locations for the child
	childLocations := make([]Location, 0)

	// Add the starting location
	childLocations = append(childLocations, parentOne.Locations[0])

	// While the child is not long enough
	for len(childLocations) < parentOne.Length() {
		canOne, canTwo := Location{}, Location{}

		// Get the first candidate location
		if parentOne.IndexOf(childLocations[len(childLocations) - 1].Id) > -1 {
			// Get the value
			canOne, _ = parentOne.Get((parentOne.IndexOf(childLocations[len(childLocations) - 1].Id) + 1) % 7)
		} else {
			return &Chromosome{}, errors.New("There was a problem indexing parent 1")
		}

		// Get the second candidate
		if parentTwo.IndexOf(childLocations[len(childLocations) - 1].Id) > -1 {
			// Get the value
			canTwo, _ = parentTwo.Get((parentTwo.IndexOf(childLocations[len(childLocations) - 1].Id) + 1) % 7)
		} else {
			return &Chromosome{}, errors.New("There was a problem indexing parent 2")
		}

		// Check if each candidate is valid
		if !self.isValidId(canOne.Id, childLocations) {
			// pick a new one
			newId, err := self.nextValidId(childLocations, parentOne.Length())

			// check for an error
			if err != nil {
				return parentOne, err
			}

			// Assign the new location
			canOne, err = parentOne.Get(parentOne.IndexOf(newId))

			// Check for errors
			if err != nil {
				return parentOne, err
			}
		}

		if !self.isValidId(canTwo.Id, childLocations) {
			// pick a new one
			newId, err := self.nextValidId(childLocations, parentOne.Length())

			// check for an error
			if err != nil {
				return parentTwo, err
			}

			// Assign the new location
			canTwo, err = parentTwo.Get(parentTwo.IndexOf(newId))

			// Check for errors
			if err != nil {
				return parentTwo, err
			}
		}

		// Append the choice with the smallest distance
		prevLoc := childLocations[len(childLocations) - 1]

		if parentOne.Matrix.GetDistance(prevLoc, canOne) < parentOne.Matrix.GetDistance(prevLoc, canTwo) {
			childLocations = append(childLocations, canOne)
		} else {
			childLocations = append(childLocations, canTwo)
		}

	}

	// Create the child chromosome
	child := &Chromosome{
		Locations: childLocations,
		Matrix: parentOne.Matrix,
		Id: self.IDCounter,
	}

	self.IDCounter++

	return child, nil
}

// Helper function for SCX Crossover operator. Gets the next valid id in the case
// the normal flow does not find one.
func (self *Population) nextValidId(locations []Location, size int) (int, error) {
	for i := 2; i <= size; i++ {
		if self.isValidId(i, locations) {
			return i, nil
		}
	}

	pretty.Println(locations)
	return -1, errors.New("Could not find a valid ID in the set 2..n")
}

// isValidId is a helper function for the SCX Crossover operator. It checks
// to see if a 
func (self *Population) isValidId(id int, locations []Location) bool {
	// Loop through each item in the locations slice
	for _, val := range locations {
		if val.Id == id {
			return false
		}
	}

	// The location was not included yet
	return true
}







