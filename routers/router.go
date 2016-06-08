package routers

import (
	"optiroute/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/about", &controllers.AboutController{})
	beego.Router("/route", &controllers.MapServicesController{})
}
