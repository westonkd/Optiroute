package test

import (
	"testing"
	"runtime"
	"path/filepath"
	"optiroute/lib"
	_ "optiroute/routers"

	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
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

	locations := []geneticTSP.Location{
		loc,
		loc2,
	}

	matrix := geneticTSP.NewFitnessMatrix()
	matrix.LoadPointMatrix(locations)

	chromo := geneticTSP.NewChromosome(locations, matrix)

	locations[0].Long = 23.0

	chromoTwo := geneticTSP.NewChromosome(locations, matrix)

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
}



