package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller

}

func (c *MainController) Get() {
	//c就是this，将数据封装到this.Data中，这里的Data是一个map
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	//this.TplName 就是需要渲染的模板
	c.TplName = "index.tpl"
}
