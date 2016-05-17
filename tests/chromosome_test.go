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

func TestChromosome(t *testing.T) {
	Convey("Locations should be able to be added to chromosomes\n", t, func() {
		chromo := geneticTSP.Chromosome{}
		So(chromo.Length() , ShouldEqual, 0)

		//add a location
		chromo.Add(geneticTSP.Location{}, geneticTSP.Location{})

		So(chromo.Length(), ShouldEqual, 2)
	})

	Convey("Locations should be able to be removed from chromosomes\n", t, func() {
		chromo := geneticTSP.Chromosome{}
		loc1 := geneticTSP.Location{0,0,"zero", 0}
		loc2 := geneticTSP.Location{0,0,"one", 1}
		loc3 := geneticTSP.Location{0,0,"two", 2}
		chromo.Add(loc1, loc2, loc3)

		//Remove first value
		chromoTest := chromo
		chromoTest.Remove(2)
		val, _ := chromoTest.Get(1)
		So(val.Name, ShouldEqual, "one")
	})

	Convey("Chromosome should support gene swappint\n", t, func() {
		chromo := geneticTSP.Chromosome{}
		loc1 := geneticTSP.Location{0,0,"zero", 0}
		loc2 := geneticTSP.Location{0,0,"one", 1}
		loc3 := geneticTSP.Location{0,0,"two", 2}

		chromo.Add(loc1, loc2, loc3)

		chromo.Swap(0,1)
		val, _ := chromo.Get(0)
		val2, _ := chromo.Get(1)

		So(val.Name, ShouldEqual, "one")
		So(val2.Name, ShouldEqual, "zero")
	})

	Convey("Chromosome should have a non-default constructor\n", t, func(){
		loc := geneticTSP.Location{
			Long: 12.0,
			Lat: 23.0,
		}

		loc2 := geneticTSP.Location{
			Long: 23.0,
			Lat: 45.0,
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

	Convey("Chromosome should calculate its fitness using its distance matrix", t, func(){
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
		So(chromo.Distance(), ShouldEqual, 48)
		So(chromo.Fitness(), ShouldEqual, 1.0 / 48.0)
	})
}
