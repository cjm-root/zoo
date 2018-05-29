package routers

import (
	"zoo/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})

    beego.Router("/animals", &controllers.AnimalController{}, "get:ListAnimals;post:AddAnimal")
	beego.Router("/animals/:id:int", &controllers.AnimalController{}, "get:ListAnimal;put:UpdateAnimal")

    // Add Staff here as well
	beego.Router("/staff", &controllers.StaffController{}, "get:ListAllStaff;post:AddStaff")
	beego.Router("/staff/:id:int", &controllers.StaffController{}, "get:ListStaff;put:UpdateStaff")


}
