package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"zoo/models"
)

type AnimalController struct {
	beego.Controller
}

func (this *AnimalController) ListAnimals() {

}
func (this *AnimalController) AddAnimal() {
	fmt.Println("Hello from AddAnimal")

	var animal models.ZooAnimal

	if err := this.ParseForm(&animal); err != nil {
		this.Ctx.Output.Body([] byte ("No Animal Data Provided"))
		return
	}

	// view
	this.TplName = "animal.tpl"

	// Insert animal into mysql
	t := models.CreateAnimal(animal)

	this.Data["json"] = t
	this.ServeJSON()

}