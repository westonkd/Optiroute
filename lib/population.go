package geneticTSP

import (
	"time"
	"math/rand"
	"errors"
	"github.com/kr/pretty"
)

// Represents a population of Chromosomes (tours)
type Population struct {
	Chromosomes  []Chromosome
	Recalculate  bool
	MutThreshold float32
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
	if index < len(self.Chromosomes) - 1 && index > -1 {
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
	rand.Seed(time.Now().Unix())

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
// if a randomly chosen number is bellow the mutation threshold level.
func (self *Population) Mutate() {
	// Seed the random number generator
	rand.Seed(time.Now().Unix())

	// Loop through each chromosome
	for _, val := range self.Chromosomes {
		// Generate a random number between 0 and 1
		mutVal := rand.Float32()

		// Mutate if the random number is bellow threshold
		if mutVal < self.MutThreshold {
			val.RandSwap()
		}
	}

}

