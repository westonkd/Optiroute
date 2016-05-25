package main

import (
	"github.com/astaxie/beego"
	_ "optiroute/routers"
	"os"
	"strconv"
)

func main() {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err == nil {
		beego.BConfig.Listen.HTTPPort = port
	}

	beego.Run()
}
