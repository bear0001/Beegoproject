package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "https://www.baidu.com/"
	c.Data["Email"] = "930987934@qq.com"
	c.TplName = "index.tpl"
}
