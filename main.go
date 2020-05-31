package main

import (
	//这边引入routers包下的init函数，注册URL映射到Controller的关联
	_ "beego-demo/routers"
	"github.com/astaxie/beego"
	"fmt"
	"github.com/astaxie/beego/context"
)
//main 函数是项目的入口函数
func main() {
	//将static包下的资源作为静态资源暴露出去
	//beego.BConfig.WebConfig.StaticDir["/static"] = "static"
	beego.SetStaticPath("/views","views")
	mysqluser := beego.AppConfig.String("mysqluser")
	mysqlpass := beego.AppConfig.String("mysqlpass")
	mysqlurls := beego.AppConfig.String("mysqlurls")
	mysqlport := beego.AppConfig.String("mysqlport")
	mysqldb := beego.AppConfig.String("mysqldb")
	fmt.Println(mysqluser,mysqlpass,mysqlurls,mysqlport,mysqldb)

	//支持swagger文档
	beego.BConfig.WebConfig.DirectoryIndex = true
	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"

	//用户登录状态校验
	var FilterUser = func(ctx *context.Context) {
		token := ctx.Request.Header.Get("token")
		if token=="" && ctx.Request.RequestURI != "/login" {
			//ctx.Redirect(302, "/login")
			ctx.ResponseWriter.Status = 201
			ctx.ResponseWriter.Write([]byte("请先登录"))
		}
	}

	beego.InsertFilter("/*",beego.BeforeRouter,FilterUser)
	beego.Run()
}

