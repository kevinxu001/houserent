package routers

import (
	"github.com/kevinxu001/houserent/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
