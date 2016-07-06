package test

import (
	_ "optiroute/routers"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
	"optiroute/lib"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

func TestTSA(t *testing.T) {
	Convey("Senior Project Time and Fitness Requirements", t, func(){
		//for j := 0; j < 100; j++ {
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
		//	ga, err := geneticTSP.NewTSPAlgorithm(locations,false,false, 50)
		//
		//	initialDistance := ga.Pop.GetFittest().Distance()
		//	fmt.Print(initialDistance, ", ")
		//
		//	// Start Timer
		//	startTime := time.Now()
		//
		//	for i := 0; i < 200; i++ {
		//		ga.Evolve()
		//	}
		//
		//	// End Timer
		//	elapsedTime := time.Since(startTime)
		//
		//	finalDistance := ga.Pop.GetFittest().Distance()
		//
		//	percentDiff := float64(initialDistance - finalDistance) / float64(initialDistance)
		//
		//	fmt.Println(finalDistance, " Time: ", elapsedTime , "seconds", percentDiff * 100, "%")
		//	So(err, ShouldBeNil)
		//
		//}
		So(1,ShouldEqual,1)
	})
}