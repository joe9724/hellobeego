package controllers

import (
	"github.com/astaxie/beego"
)

type ResourceController struct {
	beego.Controller
}

func (c *ResourceController) Get() {
	c.TplName = "admin_main.tpl"
	/*c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"*/
}

func (c *ResourceController)Post(){
	c.TplName="admin.tpl"
}

//resource

