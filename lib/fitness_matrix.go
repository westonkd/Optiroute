package geneticTSP

import (
	"googlemaps.github.io/maps"
	"github.com/kr/pretty"
	"golang.org/x/net/context"
	"errors"
	"strconv"
	"math"
)

// FitnessMatrix represents a distance matrix from
// a single destination to every other destination.
type FitnessMatrix struct {
	distanceJson map[string]interface{}
	DistanceMap map[string]int
}

// NewFintessMatrix is the constructor for
// fitness matrix.
func NewFitnessMatrix() *FitnessMatrix {
	// Create the new matrix
	fm := FitnessMatrix{}

	// Make the map
	fm.DistanceMap = make(map[string]int)

	return &fm
}

func (self *FitnessMatrix) LoadPointMatrix(locations []Location) error {
	for _, rVal := range locations {
		for _, cVal := range locations {
			// Create the keys
			key := strconv.Itoa(rVal.Id) + strconv.Itoa(cVal.Id)
			keyTwo := strconv.Itoa(cVal.Id) + strconv.Itoa(rVal.Id)

			value := math.Sqrt(math.Pow(float64(rVal.Long - cVal.Long), 2.0) + math.Pow(float64(rVal.Lat - cVal.Lat), 2.0))

			// Add to the map
			self.DistanceMap[key] = int(math.Abs(value))
			self.DistanceMap[keyTwo] = int(math.Abs(value))
		}
	}

	return nil
}

// fillMatrix takes a string of locations and returns the mapping of sources
// and destinations for the Google maps distance matrix
func (self *FitnessMatrix) getMapping(locations []Location) ([]string, []string) {
	var origins []string
	var destinations []string

	for _, val := range locations {
		origins = append(origins, val.Name)
		destinations = append(destinations, val.Name)
	}

	return  origins, destinations
}

// LoadGoogleMapMatrix takes an API key and a list of locations. It loads the locations
// into the matrix (indexed by location id) and sets the value of the distance from a
// location to each other location
func (self *FitnessMatrix) LoadGoogleMapsMatrix(apiKey string, locations []Location) error {

	origins, destinations := self.getMapping(locations)

	if len(origins) != len(destinations) {
		return errors.New("Error parsing locations")
	}

	c, err := maps.NewClient(maps.WithAPIKey(apiKey))
	if err != nil {
		return errors.New("Uanble to create new map with given api key.")
	}

	r := &maps.DistanceMatrixRequest{
		Origins: origins,
		Destinations: destinations,
	}

	resp, err := c.DistanceMatrix(context.Background(), r)
	if err != nil {
		return errors.New("Error making request.")
	}


	// Get the distances from the response
	distances := resp.Rows

	// Set up the distance matrix
	for origin, valO := range distances {
		for dest, valD  := range valO.Elements {
			key := strconv.Itoa(locations[origin].Id) + strconv.Itoa(locations[dest].Id)
			keyTwo := strconv.Itoa(locations[dest].Id) + strconv.Itoa(locations[origin].Id)
			value := valD.Distance.Meters
			self.DistanceMap[key] = value
			self.DistanceMap[keyTwo] = value
		}
	}

	pretty.Println(self.DistanceMap)

	return nil
}