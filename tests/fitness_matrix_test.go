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

func TestFitnessMatrix(t *testing.T) {
	Convey("Should create a sizexsize matrix\n", t, func() {
		matrix := geneticTSP.NewFitnessMatrix(50)
		So(len(matrix.Matrix), ShouldEqual, 50)
		So(len(matrix.Matrix[0]), ShouldEqual, 50)
		So(len(matrix.Matrix[25]), ShouldEqual, 50)
		So(len(matrix.Matrix[49]), ShouldEqual, 50)
	})

	Convey("Should laod JSON in the google maps API distance matrix format\n", t, func() {
		matrix := geneticTSP.NewFitnessMatrix(2)

		montpelier := geneticTSP.Location{
			Name: "Montpelier, ID",
			Id: 0,
		}

		rexburg := geneticTSP.Location{
			Name: "Rexburg, ID",
			Id: 1,
		}

		paris := geneticTSP.Location{
			Name: "Paris, ID",
			Id: 2,
		}

		geneva := geneticTSP.Location{
			Name: "Geneva, ID",
			Id: 3,
		}

		dummy := []geneticTSP.Location{
			montpelier,
			rexburg,
			paris,
			geneva,
		}

		matrix.LoadGoogleMapsMatrix("AIzaSyDExb4usTvy3QNZSuEo-CvcHtcRAoI2-7U", dummy)
	})
}


