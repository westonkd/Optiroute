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

// NewTSPAlgorithm initializes an algorithm
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
		err := TA.Matrix.LoadGoogleMapsMatrix("asdfasdfasdf", locations)
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
	for i := offset; i < self.PopSize; i++ {
		// Select parent chromosomes
		parent1 := self.Pop.TournamentSelect(5)
		parent2 := self.Pop.TournamentSelect(5, parent1.Distance())

		//if parent1.Distance() == parent2.Distance() {
		//	pretty.Println("<<<<<<<<<<<<<<<<<<<<<<<<<<<<")
		//	pretty.Println(parent1.Locations, parent1.Id)
		//	pretty.Println(parent2.Locations, parent2.Id)
		//	return
		//}

		// Do the crossover and add to the new generation
		child, error := self.Pop.SimpleCrossover(parent1, parent2)

		//pretty.Println("parent 1: ", parent1.Distance(), " ", " p2: ", parent2.Distance(), " ", " child: ", child.Distance(), "Min: ", self.Pop.GetFittest().Distance())

		if error != nil {
			pretty.Println("PANICCCC!")
			pretty.Println(error)
		}

		newChromosomes = append(newChromosomes, *child)
	}

	// Initialize a new population
	nextGen := &Population{
		Chromosomes: newChromosomes,
		IDCounter: self.Pop.IDCounter,
		MutThreshold: self.Pop.MutThreshold,
	}

	// Mutation
	nextGen.Mutate()

	//fmt.Println("Distance: ", nextGen.GetFittest().Distance(), " ", self.Pop.GetFittest().Distance())

	// Assign the next population
	self.Pop.Chromosomes = make([]Chromosome, len(nextGen.Chromosomes))
	copy(self.Pop.Chromosomes, nextGen.Chromosomes)

	//pretty.Println(self.Pop.GetFittest().Distance())
}

func (self *TSPAlgorithm) RandomPop() *Population {
	p := Population{
		MutThreshold: .5,
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
