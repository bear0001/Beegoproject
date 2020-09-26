package routers

import (
	"Project922/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//register: 注册
	beego.Router("/register",&controllers.RegisterController{})
    beego.Router("/", &controllers.MainController{})
}
