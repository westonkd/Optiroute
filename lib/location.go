//Package genetic TSP provides structures to approximate a solution
//to the traveling salesman problem using a genetic algorithm.
package geneticTSP

// Represents a location by name or longitude and latitude
type Location struct {
	Long, Lat float32
	Name      string
	Id        int
}
