//Package genetic TSP provides structures to approximate a solution
//to the traveling salesman problem using a genetic algorithm.
package geneticTSP

import (
)
import (
	"math/rand"
	"github.com/kr/pretty"
)

type TSPAlgorithm struct {
	Pop *Population
	Elitism  bool           // Boolean flag for keeping fittest chromosome
	Google   bool           // Uses data for google maps to parse
	Locations []Location    // Locations for the tour
	PopSize  int
	Matrix FitnessMatrix
}

// NewTSPAlgorithm initializes an algorithm handler
func NewTSPAlgorithm(locations []Location, google bool, elitism bool, popSize int) (*TSPAlgorithm, error) {
	TA := TSPAlgorithm{
		Locations: locations,
		Google: google,
		Elitism: elitism,
		PopSize: popSize,
	}

	// Load the matrix
	TA.Matrix = *NewFitnessMatrix()

	// If we have a one for Google maps
	if TA.Google {
		err := TA.Matrix.LoadGoogleMapsMatrix("AIzaSyDExb4usTvy3QNZSuEo-CvcHtcRAoI2-7U", locations)
		if err != nil {
			return &TA, err
		}
	} else {
		err := TA.Matrix.LoadPointMatrix(locations)
		if err != nil {
			return &TA, err
		}
	}

	// Create a random population
	TA.Pop = TA.RandomPop()

	return &TA, nil
}

// Evolve evolves the generation by crossover and then mutation. Mutate the population in place.
func (self *TSPAlgorithm) Evolve() {
	// New empty slice of Chromosomes
	newChromosomes := make([]Chromosome, 0)

	// Offset for elitism
	offset := 0

	// Keep the best if elitism is on
	if self.Elitism {
		newChromosomes = append(newChromosomes, *self.Pop.GetFittest())
		offset++
	}

	// Crossover
	numCross := int(float32(self.PopSize) * self.Pop.CrossThreshold)
	numLive := self.PopSize - numCross

	for i := offset; i < numCross; i++ {
		// Select parent chromosomes
		parent1 := self.Pop.TournamentSelect(5)
		parent2 := self.Pop.TournamentSelect(5, parent1.Distance())

		// Do the crossover and add to the new generation
		child, error := self.Pop.OrderedCrossover(parent1, parent2)

		if error != nil {
			pretty.Println(error)
		}

		newChromosomes = append(newChromosomes, *child)
	}

	// Bring over survivors to the new generation
	for i := 0; i < numLive; i++ {
		newChromosome := self.Pop.TournamentSelect(15)
		newChromosomes = append(newChromosomes, *newChromosome)
	}

	// Initialize a new population
	nextGen := &Population{
		Chromosomes: newChromosomes,
		IDCounter: self.Pop.IDCounter,
		MutThreshold: self.Pop.MutThreshold,
	}

	// Mutation
	nextGen.RSMutate()

	// Assign the next population
	self.Pop.Chromosomes = make([]Chromosome, len(nextGen.Chromosomes))
	copy(self.Pop.Chromosomes, nextGen.Chromosomes)

	//pretty.Println(self.Pop.GetFittest().Distance())
}

// RandomPop returns a random population of popSize
func (self *TSPAlgorithm) RandomPop() *Population {
	p := Population{
		MutThreshold: 0.5,
		CrossThreshold: 0.95,
	}

	p.Chromosomes = make([]Chromosome, 0)

	for i := 0; i < self.PopSize; i++ {
		newChromo := &Chromosome{
			Locations: self.Locations,
			Matrix: &self.Matrix,
			Id: i + 1,
		}

		p.Chromosomes = append(p.Chromosomes, *newChromo)
	}

	// Randomize
	for i, _ := range p.Chromosomes {
		swap := rand.Intn(15)

		for j := 0; j < swap; j++ {
			p.Chromosomes[i].RandSwap()
		}
	}

	p.IDCounter = self.PopSize + 1

	return &p
}
