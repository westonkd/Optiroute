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
	"github.com/kr/pretty"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

func TestTSA(t *testing.T) {
	Convey("Should produce a new population", t, func(){
		locations := []geneticTSP.Location{}

		for i := 0; i < 50; i++ {
			location := geneticTSP.Location{
				Id:   i + 1,
				Lat:  float32(rand.Intn(390)),
				Long: float32(rand.Intn(390)),
			}

			locations = append(locations, location)
		}

		ga, err := geneticTSP.NewTSPAlgorithm(locations,false,true, 50)

		pretty.Println("Begin: ", ga.Pop.GetFittest().Distance())

		for i := 0; i < 100; i++ {
			ga.Evolve()
		}

		pretty.Println("End: ", ga.Pop.GetFittest().Distance())

		So(err, ShouldBeNil)
	})
}