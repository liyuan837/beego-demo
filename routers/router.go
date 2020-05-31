package routers

import (
	"beego-demo/controllers"
	"github.com/astaxie/beego"
)
//路由注册beego.Router函数，映射URL到controller
func init() {
    beego.Router("/", &controllers.MainController{})

    beego.Router("/user", &controllers.UserController{},"post:Post")
	beego.Router("/user/list", &controllers.UserController{},"get:GetAll")
	beego.Router("/user/:id", &controllers.UserController{},"get:Get")
	beego.Router("/user/:id", &controllers.UserController{},"put:Put")
	beego.Router("/user/:id", &controllers.UserController{},"delete:Delete")
	beego.Router("/user/login", &controllers.UserController{},"get:Login")
	beego.Router("/user/login", &controllers.UserController{},"get:Logout")
}
