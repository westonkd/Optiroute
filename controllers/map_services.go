package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"fmt"
	"optiroute/lib"
)

// Number of evolutions to do
const EvolveNum = 100

type MapServicesController struct {
	beego.Controller
}

type mystruct struct {
	Route []string `json:"route"`
}

func (c *MapServicesController) Post() {

	// Get the locations
	req := c.Ctx.Request
	var locations []string
	p := make([]byte, req.ContentLength)
	_, err := c.Ctx.Request.Body.Read(p)
	if err == nil {
		err1 := json.Unmarshal(p, &locations)
		if err1 != nil {
			fmt.Println("Unable to unmarshall the JSON request", err1);
		}
	}

	// Create the array of location objects
	routeLocations := []geneticTSP.Location{}

	// Create the locations
	for i, val := range locations {
		location := geneticTSP.Location{
			Id: i + 1,
			Name: val,
		}

		routeLocations = append(routeLocations, location)
	}

	// Create the new genetic TSP Handler
	ga, err := geneticTSP.NewTSPAlgorithm(routeLocations,true,true, 50)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Initial Distance: ", ga.Pop.GetFittest().Distance(), "Meters")

	// Find a good route
	for i := 0; i < EvolveNum; i++ {
		ga.Evolve()
	}

	fmt.Println("Final Distance: ", ga.Pop.GetFittest().Distance(), " Meters")

	// Get the final routes
	finalRoute := make([]string, 0)

	for _, val := range ga.Pop.GetFittest().Locations {
		finalRoute = append(finalRoute, val.Name)
	}

	test := mystruct{Route: finalRoute}
	c.Data["json"] = &test
	c.ServeJSON()
}