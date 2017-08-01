package routers

import (
	"hellobeego/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/resource/list",&controllers.ResourceController{})
}
