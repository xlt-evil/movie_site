package controllers

type UserController struct {
	BaseController
}

func (this *UserController)ToPerson(){
	this.TplName = "person.html"
}
