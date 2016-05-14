//Package genetic TSP provides structures to approximate a solution
//to the traveling salesman problem using a genetic algorithm.
package geneticTSP

import (
	"errors"
	"strconv"
)

// Chromosome represents a set of genes (locations).
type Chromosome struct {
	Locations []Location
	Matrix *FitnessMatrix
}

// NewChromosome is the constructor for Chromosome
func NewChromosome(locations []Location, matrix *FitnessMatrix) *Chromosome {
	c := Chromosome{
		Locations: locations,
		Matrix: matrix,
	}

	return &c
}


// Get retrieves element i from the chromosome. Returns an error
// if i is out of bounds.
func (self *Chromosome) Get(i int) (Location, error)  {
	if i < len(self.Locations) {
		return self.Locations[i], nil
	}

	return Location{0,0,"error",0}, errors.New("Index out of boounds")
}

// Add adds n locations to the chromosome.
func (self *Chromosome) Add(locations ...Location) {
	//Add each location
	for i := range locations {
		self.Locations = append(self.Locations, locations[i])
	}
}

// Remove removes element i from the chromosome.
// Returns an error if i is out of bounds.
func (self *Chromosome) Remove(i int) error {
	//if the index is in bounds
	if i < len(self.Locations) {
		//remove the specified element
		self.Locations = append(self.Locations[:i], self.Locations[i + 1:]...)
		return nil
	}
	return errors.New("Index ouf of bounds.")
}

// Swap swaps element first and element second. If either first
// or second is out of bounds, an error is returned.
func (self *Chromosome) Swap(first, second int) error {
	if first < len(self.Locations) && second < len(self.Locations) {
		//To a simple swap
		temp := self.Locations[first]
		self.Locations[first] = self.Locations[second]
		self.Locations[second] = temp

		return nil
	}

	return errors.New("Swap indeces out of bounds.")
}

// Length returns a count of the number of elements in
// the chromosome.
func (self *Chromosome) Length() (location int) {
	location = len(self.Locations)
	return
}


// Fitness returns the fitness of the chromosome
func (self *Chromosome) Fitness() (fitness int) {
	fitness = 0

	for i, location := range self.Locations {
		if i + 1 < self.Length() {
			fitness += self.Matrix.DistanceMap[strconv.Itoa(location.Id) + strconv.Itoa(self.Locations[i + 1].Id)]
		} else {
			fitness += self.Matrix.DistanceMap[strconv.Itoa(location.Id) + strconv.Itoa(self.Locations[0].Id)]
		}
	}

	return
}