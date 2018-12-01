package controllers

import (
	"github.com/astaxie/beego"
)

// DemoController Operations about Bank
type DemoController struct {
	beego.Controller
}

// DemoApi ...
// @Title DemoApi
// @Description DemoAPI for test
// @Param	randomstring		query 	string	true	"random string"
// @Success 200
// @router /demoapi [get]
func (c *DemoController) DemoApi() {
	var rs = c.GetString("randomstring", "")

	c.Data["json"] = map[string]string{
		"returnstring": rs + "_back",
	}
	c.ServeJSON()
	return
}
