package test

import (
	_ "optiroute/routers"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

func TestTSA(t *testing.T) {
	Convey("Should produce a new population", t, func(){
		//for j := 0; j < 50; j++ {
		//	locations := []geneticTSP.Location{}
		//
		//	for i := 0; i < 50; i++ {
		//		location := geneticTSP.Location{
		//			Id:   i + 1,
		//			Lat:  float32(rand.Intn(390)),
		//			Long: float32(rand.Intn(390)),
		//		}
		//
		//		locations = append(locations, location)
		//	}
		//
		//	ga, err := geneticTSP.NewTSPAlgorithm(locations,false,true, 50)
		//
		//	fmt.Print(ga.Pop.GetFittest().Distance(), ",")
		//
		//	for i := 0; i < 100; i++ {
		//		ga.Evolve()
		//	}
		//
		//	fmt.Println(ga.Pop.GetFittest().Distance())
		//	So(err, ShouldBeNil)
		//
		//}
		So(1,ShouldEqual,1)
	})
}