package routers

import (
	"github.com/astaxie/beego"
	"github.com/kevinxu001/houserent/controllers"
)

func init() {
	beego.Router("/", &controllers.IndexController{})
	beego.Router("/organization", &controllers.OrganizationController{})
}
