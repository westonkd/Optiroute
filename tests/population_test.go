package test

import (
	"optiroute/lib"
	_ "optiroute/routers"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/astaxie/beego"
	"github.com/kr/pretty"
	. "github.com/smartystreets/goconvey/convey"
	"math/rand"
	"time"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

func getPopulation() geneticTSP.Population {
	loc := geneticTSP.Location{
		Long: 12.0,
		Lat:  23.0,
		Id:   0,
	}

	loc2 := geneticTSP.Location{
		Long: 23.0,
		Lat:  45.0,
		Id:   1,
	}

	loc3 := geneticTSP.Location{
		Long: 455.0,
		Lat:  45,
		Id:   2,
	}

	loc4 := geneticTSP.Location{
		Long: 45.0,
		Lat:  45,
		Id:   3,
	}

	loc5 := geneticTSP.Location{
		Long: 55.0,
		Lat:  45,
		Id:   4,
	}

	loc6 := geneticTSP.Location{
		Long: 455.0,
		Lat:  5,
		Id:   5,
	}

	loc7 := geneticTSP.Location{
		Long: 4.0,
		Lat:  45,
		Id:   6,
	}

	loc8 := geneticTSP.Location{
		Long: 355.0,
		Lat:  25,
		Id:   7,
	}

	locations := []geneticTSP.Location{
		loc,
		loc2,
		loc3,
		loc4,
		loc5,
		loc6,
		loc7,
		loc8,
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

	chromosomes := []geneticTSP.Chromosome{
		*chromo,
		*chromoTwo,
	}

	return geneticTSP.Population{
		Chromosomes: chromosomes,
		Recalculate: true,
	}
}

func TestPopulation(t *testing.T) {
	Convey("Population should have an associated length", t, func() {
		pop := getPopulation()
		So(pop.Size(), ShouldEqual, 2)
	})

	Convey("Population should find its fittest chromosome", t, func() {
		pop := getPopulation()

		chromoOne := pop.Chromosomes[0]
		chromoTwo := pop.Chromosomes[1]

		fitest := chromoOne

		if chromoOne.Fitness() < chromoTwo.Fitness() {
			fitest = chromoTwo
		}

		So(pop.GetFittest().Fitness(), ShouldEqual, fitest.Fitness())
	})

	Convey("Population should allow chromosomes to be added to it and retreived", t, func() {
		pop := getPopulation()

		initialCount := pop.Size()

		pop.Add(&pop.Chromosomes[0])

		So(pop.Size(), ShouldEqual, initialCount+1)

		// get a chromosome
		chromo, error := pop.Get(0)

		So(error, ShouldBeNil)
		So(chromo.Fitness(), ShouldEqual, pop.Chromosomes[0].Fitness())
	})

	Convey("Population should perform a tournament selection", t, func() {
		pop := getPopulation()

		chromo := pop.TournamentSelect(2)
		pretty.Print(chromo)
	})

	Convey("Population should have the ablility to mutate", t, func() {
		popOne := getPopulation()
		popTwo := getPopulation()

		//Always mutate
		popTwo.MutThreshold = 2
		popTwo.Mutate()

		pretty.Println(popOne, "\n==========================")
		pretty.Println(popTwo)
	})

	Convey("RSM mutation should work", t, func() {
		popOne := getPopulation()
		popTwo := getPopulation()

		//Always mutate
		popTwo.MutThreshold = 2
		popTwo.RSMutate()

		pretty.Println(popOne, "\n==========================")
		pretty.Println(popTwo)
	})

	Convey("Crossover should produce the expected child chromosome using the SCX method", t, func() {
		rand.Seed(time.Now().Unix())

		loc1 := geneticTSP.Location{
			Id:   1,
			Lat:  float32(rand.Intn(400)),
			Long: float32(rand.Intn(400)),
		}

		loc2 := geneticTSP.Location{
			Id:   2,
			Lat:  float32(rand.Intn(400)),
			Long: float32(rand.Intn(400)),
		}

		loc3 := geneticTSP.Location{
			Id:   3,
			Lat:  float32(rand.Intn(400)),
			Long: float32(rand.Intn(400)),
		}

		loc4 := geneticTSP.Location{
			Id:   4,
			Lat:  float32(rand.Intn(400)),
			Long: float32(rand.Intn(400)),
		}

		loc5 := geneticTSP.Location{
			Id:   5,
			Lat:  float32(rand.Intn(400)),
			Long: float32(rand.Intn(400)),
		}

		loc6 := geneticTSP.Location{
			Id:   6,
			Lat:  float32(rand.Intn(400)),
			Long: float32(rand.Intn(400)),
		}

		loc7 := geneticTSP.Location{
			Id:   7,
			Lat:  float32(rand.Intn(400)),
			Long: float32(rand.Intn(400)),
		}

		locations := []geneticTSP.Location{
			loc1,
			loc2,
			loc3,
			loc4,
			loc5,
			loc6,
			loc7,
		}

		distanceMatrix := geneticTSP.NewFitnessMatrix()
		loadError := distanceMatrix.LoadPointMatrix(locations)

		chromo1 := geneticTSP.Chromosome{
			Locations: []geneticTSP.Location{
				loc1,
				loc5,
				loc7,
				loc3,
				loc6,
				loc4,
				loc2,
			},
			Matrix: distanceMatrix,
		}

		chromo2 := geneticTSP.Chromosome{
			Locations: []geneticTSP.Location{
				loc1,
				loc6,
				loc2,
				loc4,
				loc3,
				loc5,
				loc7,
			},
			Matrix: distanceMatrix,
		}

		pop := geneticTSP.Population{
			Chromosomes: []geneticTSP.Chromosome{
				chromo1,
				chromo2,
			},
		}

		res, err := pop.Crossover(&chromo1, &chromo2)

		pretty.Println("\n======\n", res)
		pretty.Println("\n==dist results==\n", chromo1.Distance(), ", ", chromo2.Distance(), ", ", res.Distance())

		So(loadError, ShouldBeNil)
		So(err, ShouldBeNil)
	})
}
