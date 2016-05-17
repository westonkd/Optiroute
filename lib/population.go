package geneticTSP

// Represents a population of Chromosomes (tours)
type Population struct {
	Chromosomes []Chromosome
	Recalculate bool
}

// Size returns the population count
func (self *Population) Size() (size int) {
	size = len(self.Chromosomes)
	return
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
