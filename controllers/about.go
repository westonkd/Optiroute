package controllers

import (
	"github.com/astaxie/beego"
	"optiroute/lib"
	"encoding/json"
	"fmt"
	"math/rand"
)

type AboutController struct {
	beego.Controller
}

type AboutResponse struct {
	Initial geneticTSP.Chromosome
	Final geneticTSP.Chromosome
	InitialDistance int
	FinalDistance int
}

func (c *AboutController) Get() {
	//Layout Info
	c.Layout = "layout_about.tpl"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["HeadScripts"] = "Shared/head_scripts.tpl"
	c.LayoutSections["HeadStyles"] = "Shared/head_styles.tpl"
	c.LayoutSections["Header"] = "Shared/header.tpl"
	c.LayoutSections["Footer"] = "Shared/footer.tpl"

	//View Info
	c.TplName = "about.tpl"
}

func (c *AboutController) Post() {
	// Get the locations
	req := c.Ctx.Request
	var locations []geneticTSP.Location
	p := make([]byte, req.ContentLength)
	_, err := c.Ctx.Request.Body.Read(p)
	if err == nil {
		err1 := json.Unmarshal(p, &locations)
		if err1 != nil {
			fmt.Println("Unable to unmarshall the JSON request", err1);
		}
	}

	response := AboutResponse{}

	ga, err := geneticTSP.NewTSPAlgorithm(locations,false,true, 50)

	response.Initial = ga.Pop.Chromosomes[rand.Intn(ga.PopSize - 1)]
	response.InitialDistance = response.Initial.Distance()

	for i := 0; i < 200; i++ {
		ga.Evolve()
	}

	response.Final = *ga.Pop.GetFittest()
	response.FinalDistance = response.Final.Distance()

	fmt.Println(response.Initial.Length(), ", ", response.Final.Length())

	c.Data["json"] = response
	c.ServeJSON()
}
