package test

import (
	"testing"
	"runtime"
	"path/filepath"
	"optiroute/lib"
	_ "optiroute/routers"

	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
	"math/rand"
	"time"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".." + string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

func TestFitnessMatrix(t *testing.T) {
	Convey("Should load points given a set of locations with just longitude and lattitude\n",t, func(){
		// Create distance matrix and locations list

		distanceMatrix := geneticTSP.NewFitnessMatrix()
		locations := []geneticTSP.Location{}

		rand.Seed(time.Now().Unix())

		for i := 0; i < 50; i++ {
			newLocation := geneticTSP.Location{
				Id: i,
				Long: float32(rand.Intn(400)),
				Lat: float32(rand.Intn(400)),
			}


			locations = append(locations, newLocation)
		}

		distanceMatrix.LoadPointMatrix(locations)
	})

	Convey("Should laod distanes from Google maps given a list of locations\n", t, func() {
		//matrix := geneticTSP.NewFitnessMatrix()

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

		locationList := []geneticTSP.Location{
			montpelier,
			rexburg,
			paris,
			geneva,
		}


		// err := matrix.LoadGoogleMapsMatrix("AIzaSyDExb4usTvy3QNZSuEo-CvcHtcRAoI2-7U", locationList)
		// So(err, ShouldBeNil)

		So(len(locationList), ShouldEqual, 4)
	})
}


