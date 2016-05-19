package test

import (
	"testing"
	"runtime"
	"path/filepath"
	"optiroute/lib"
	_ "optiroute/routers"

	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/kr/pretty"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".." + string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

func getPopulation() geneticTSP.Population {
	loc := geneticTSP.Location{
		Long: 12.0,
		Lat: 23.0,
		Id: 0,
	}

	loc2 := geneticTSP.Location{
		Long: 23.0,
		Lat: 45.0,
		Id: 1,
	}

	loc3 := geneticTSP.Location{
		Long: 455.0,
		Lat: 45,
		Id: 2,
	}

	locations := []geneticTSP.Location{
		loc,
		loc2,
		loc3,
	}

	matrix := geneticTSP.NewFitnessMatrix()
	matrix.LoadPointMatrix(locations)

	chromo := geneticTSP.NewChromosome(locations, matrix)

	locations[0].Long = 23.0

	b := make([]geneticTSP.Location, len(locations))
	copy(b, locations)

	chromoTwo := geneticTSP.NewChromosome(b, matrix)

	chromo.Id = 0
	chromoTwo.Id = 1

	chromosomes := []geneticTSP.Chromosome {
		*chromo,
		*chromoTwo,
	}

	return geneticTSP.Population{
		Chromosomes: chromosomes,
		Recalculate: true,
	}
}

func TestPopulation(t *testing.T) {
	Convey("Population should have an associated length", t, func(){
		pop := getPopulation()
		So(pop.Size(), ShouldEqual, 2)
	})

	Convey("Population should find its fittest chromosome", t, func(){
		pop := getPopulation()

		chromoOne := pop.Chromosomes[0]
		chromoTwo := pop.Chromosomes[1]

		fitest := chromoOne

		if chromoOne.Fitness() < chromoTwo.Fitness() {
			fitest = chromoTwo
		}

		So(pop.GetFittest().Fitness(), ShouldEqual, fitest.Fitness())
	})

	Convey("Population should allow chromosomes to be added to it and retreived", t, func(){
		pop := getPopulation()

		initialCount := pop.Size()

		pop.Add(&pop.Chromosomes[0])

		So(pop.Size(), ShouldEqual, initialCount + 1)

		// get a chromosome
		chromo, error := pop.Get(0)

		So(error, ShouldBeNil)
		So(chromo.Fitness(), ShouldEqual, pop.Chromosomes[0].Fitness())
	})

	Convey("Population should perform a tournament selection",t, func() {
		pop := getPopulation()

		chromo := pop.TournamentSelect(2)
		pretty.Print(chromo)
	})

	Convey("Population should have the ablility to mutate", t, func(){
		popOne := getPopulation()
		popTwo := getPopulation()

		//Always mutate
		popTwo.MutThreshold = 2
		popTwo.Mutate()

		pretty.Println(popOne, "\n==========================")
		pretty.Println(popTwo)
	})
}



