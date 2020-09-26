package controllers

import (
	"Project922/models"
	"fmt"
	"github.com/astaxie/beego"
)
//该结构体用于处理register请求
type RegisterController struct {
	beego.Controller
}
//该方法用于处理post请求
func (r*RegisterController)Post()  {
	fmt.Println(r ==nil)
	fmt.Println(r.Ctx==nil)
	fmt.Println(r.Ctx.Request==nil)
// 1.接收前端传递的数据
////	 body,err:=r.Ctx.Request.GetBody()
//	////	if err!=nil {
//	////		r.Ctx.WriteString("数据接收错误")
//	////		return
//	}
//	bodyBytes,err:=ioutil.ReadAll(r.Ctx.Request.Body)
//	if err!=nil {
//		r.Ctx.WriteString("数据接收错误，请重试")
//		return
//	}
//	var user models.User
//	err=json.Unmarshal(bodyBytes,&user)
//	if err != nil {
//		r.Ctx.WriteString("数据解析错误，请重试")
//		return
//	}

var user models.User
	 err:=r.ParseForm(user)
	if err!=nil {
		r.Ctx.WriteString("数据解析错误")
	}
	fmt.Println(user.User)
	fmt.Println(user.Nick)
	r.Ctx.WriteString("数据接收成功，并成功解析。")
}
