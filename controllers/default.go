package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

// Get the page
func (c *MainController) Get() {
	//Layout Info
	c.Layout = "layout_main.tpl"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["HeadScripts"] = "Shared/head_scripts.tpl"
	c.LayoutSections["HeadStyles"] = "Shared/head_styles.tpl"
	c.LayoutSections["Header"] = "Shared/header.tpl"
	c.LayoutSections["Footer"] = "Shared/footer.tpl"

	//View Info
	c.TplName = "index.tpl"
}
