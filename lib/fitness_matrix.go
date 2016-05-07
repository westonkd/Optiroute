package geneticTSP

import (
	"googlemaps.github.io/maps"
	"github.com/kr/pretty"
	"golang.org/x/net/context"
	"errors"
)

type FitnessMatrix struct {
	distanceJson map[string]interface{}
	Matrix [][]int
}

func NewFitnessMatrix(size int) *FitnessMatrix {
	//Create the new matrix
	fm := FitnessMatrix{}

	// Allocate the top-level slice.
	fm.Matrix = make([][]int, size)

	// Loop over the rows, allocating the slice for each row.
	for i := range fm.Matrix {
		fm.Matrix[i] = make([]int, size)
	}

	return &fm
}

func (self *FitnessMatrix) generateOrigDest(locations []Location) (Location, Location) {
	var origins []Location
	var destinations []Location

	//generate a mapping of all!

	return  origins, destinations
}

func (self *FitnessMatrix) LoadMatrix(apiKey string, locations []Location) error {

	origins, destinations := self.generateOrigDest(locations)

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