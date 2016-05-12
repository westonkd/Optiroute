package geneticTSP

import (
	"googlemaps.github.io/maps"
	"github.com/kr/pretty"
	"golang.org/x/net/context"
	"errors"
	"fmt"
)

type FitnessMatrix struct {
	distanceJson map[string]interface{}
	Matrix [][]int
}

func NewFitnessMatrix(size int) *FitnessMatrix {
	// Create the new matrix
	fm := FitnessMatrix{}

	// Allocate the top-level slice.
	fm.Matrix = make([][]int, size)

	// Loop over the rows, allocating the slice for each row.
	for i := range fm.Matrix {
		fm.Matrix[i] = make([]int, size)
	}

	return &fm
}

func (self *FitnessMatrix) fillMatrix(locations []Location) ([]string, []string) {
	var origins []string
	var destinations []string

	for i, origin := range locations {
		for _, destination := range locations[i:] {
			// if we are not going to the same location
			if (origin.Id != destination.Id) {
				fmt.Println(origin.Name + " -> " + destination.Name)
			}
		}

		fmt.Println("======")
	}
	// generate a mapping of all!

	return  origins, destinations
}

// LoadGoogleMapMatrix takes an API key and alist of locations. It loads the locations
// into the matrix (indexed by location id) and sets the value of the distance from a
// location to each other location
func (self *FitnessMatrix) LoadGoogleMapsMatrix(apiKey string, locations []Location) error {

	origins, destinations := self.fillMatrix(locations)

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

	//pretty.Println(resp.Rows[0].Elements[0].Distance.Meters)
	pretty.Println(resp)

	return nil
}