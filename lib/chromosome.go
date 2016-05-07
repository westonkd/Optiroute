//Package genetic TSP provides structures to approximate a solution
//to the traveling salesman problem using a genetic algorithm.
package geneticTSP

import (
	"errors"
)

// Chromosome represents a set of genes (locations).
type Chromosome struct {
	locations []Location
}

// Get retrieves element i from the chromosome. Returns an error
// if i is out of bounds.
func (self *Chromosome) Get(i int) (Location, error)  {
	if i < len(self.locations) {
		return self.locations[i], nil
	}

	return Location{0,0,"error",0}, errors.New("Index out of boounds")
}

// Add adds n locations to the chromosome.
func (self *Chromosome) Add(locations ...Location) {
	//Add each location
	for i := range locations {
		self.locations = append(self.locations, locations[i])
	}
}

// Remove removes element i from the chromosome.
// Returns an error if i is out of bounds.
func (self *Chromosome) Remove(i int) error {
	//if the index is in bounds
	if i < len(self.locations) {
		//remove the specified element
		self.locations = append(self.locations[:i], self.locations[i + 1:]...)
		return nil
	}
	return errors.New("Index ouf of bounds.")
}

// Swap swaps element first and element second. If either first
// or second is out of bounds, an error is returned.
func (self *Chromosome) Swap(first, second int) error {
	if first < len(self.locations) && second < len(self.locations) {
		//To a simple swap
		temp := self.locations[first]
		self.locations[first] = self.locations[second]
		self.locations[second] = temp

		return nil
	}

	return errors.New("Swap indeces out of bounds.")
}

// Length returns a count of the number of elements in
// the chromosome.
func (self *Chromosome) Length() (location int) {
	location = len(self.locations)
	return
}