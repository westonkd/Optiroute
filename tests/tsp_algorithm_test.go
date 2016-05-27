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
	"time"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

func TestTSA(t *testing.T) {
	Convey("Should produce a new population", t, func(){
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

		ga, err := geneticTSP.NewTSPAlgorithm(locations,false,true, 50)

		for i := 0; i < 5; i++ {
			ga.Evolve()
		}

		So(err, ShouldBeNil)
	})
}