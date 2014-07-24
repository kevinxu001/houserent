package controllers

import (
	"github.com/astaxie/beego"
)

type IndexController struct {
	CommonController
}

func (this *IndexController) Get() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplNames = "index.tpl"

	this.Data["DBHost"] = beego.AppConfig.String("db::host")
}
