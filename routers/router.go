package routers

import (
	"ccxt.com/nginxconf/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "*:Index")
	beego.Router("/login", &controllers.MainController{}, "*:Login")
	beego.Router("/domain", &controllers.MainController{}, "*:Domain")
	beego.Router("/addserver", &controllers.MainController{}, "*:Addomain")
	beego.Router("/create/:id", &controllers.MainController{}, "*:Create")
	beego.Router("/delete/:id", &controllers.MainController{}, "*:Delete")
	beego.Router("/addconfdir", &controllers.MainController{}, "*:Addconfdir")
	beego.Router("/test", &controllers.MainController{}, "*:Tes")

}
