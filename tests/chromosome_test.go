package test

import (
	"optiroute/lib"
	_ "optiroute/routers"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
	"math/rand"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

func TestChromosome(t *testing.T) {
	Convey("Locations should be able to be added to chromosomes\n", t, func() {
		chromo := geneticTSP.Chromosome{}
		So(chromo.Length(), ShouldEqual, 0)

		//add a location
		chromo.Add(geneticTSP.Location{}, geneticTSP.Location{})

		So(chromo.Length(), ShouldEqual, 2)
	})

	Convey("Locations should be able to be removed from chromosomes\n", t, func() {
		chromo := geneticTSP.Chromosome{}
		loc1 := geneticTSP.Location{0, 0, "zero", 0}
		loc2 := geneticTSP.Location{0, 0, "one", 1}
		loc3 := geneticTSP.Location{0, 0, "two", 2}
		chromo.Add(loc1, loc2, loc3)

		//Remove first value
		chromoTest := chromo
		chromoTest.Remove(2)
		val, _ := chromoTest.Get(1)
		So(val.Name, ShouldEqual, "one")
	})

	Convey("Chromosome should support gene swappint\n", t, func() {
		chromo := geneticTSP.Chromosome{}
		loc1 := geneticTSP.Location{0, 0, "zero", 0}
		loc2 := geneticTSP.Location{0, 0, "one", 1}
		loc3 := geneticTSP.Location{0, 0, "two", 2}

		chromo.Add(loc1, loc2, loc3)

		chromo.Swap(0, 1)
		val, _ := chromo.Get(0)
		val2, _ := chromo.Get(1)

		So(val.Name, ShouldEqual, "one")
		So(val2.Name, ShouldEqual, "zero")
	})

	Convey("Chromosome should have a non-default constructor\n", t, func() {
		loc := geneticTSP.Location{
			Long: 12.0,
			Lat:  23.0,
		}

		loc2 := geneticTSP.Location{
			Long: 23.0,
			Lat:  45.0,
		}

		locations := []geneticTSP.Location{
			loc,
			loc2,
		}

		matrix := geneticTSP.NewFitnessMatrix()
		matrix.LoadPointMatrix(locations)

		chromo := geneticTSP.NewChromosome(locations, matrix)

		So(len(chromo.Locations), ShouldEqual, len(locations))
	})

	Convey("Chromosome should calculate its fitness using its distance matrix", t, func() {
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

		locations := []geneticTSP.Location{
			loc,
			loc2,
		}

		matrix := geneticTSP.NewFitnessMatrix()
		matrix.LoadPointMatrix(locations)

		chromo := geneticTSP.NewChromosome(locations, matrix)
		So(chromo.Distance(), ShouldEqual, 48)
		So(chromo.Fitness(), ShouldEqual, 1.0/48.0)
	})

	Convey("Should return the index of a location by ID", t, func() {
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
		distanceMatrix.LoadPointMatrix(locations)

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

		So(chromo1.IndexOf(1), ShouldEqual,0)
		So(chromo1.IndexOf(5), ShouldEqual,1)
		So(chromo1.IndexOf(2), ShouldEqual,6)
		So(chromo1.IndexOf(3), ShouldEqual,3)
		So(chromo1.IndexOf(9), ShouldEqual,-1)
	})
}
