package controllers

import (
	"chenblog/models"
	"fmt"
	"github.com/astaxie/beego"
)

type CategoryController struct {
	beego.Controller
}

func (this *CategoryController) Get() {

	op := this.Input().Get("op")
	name1 := this.Input().Get("name")
	fmt.Println("-------------op:", op)
	fmt.Println("-------------name1:", name1)

	switch op {
	case "add":
		name := this.Input().Get("name")
		if len(name) == 0 {
			break
		}
		fmt.Println("1---------")
		err := models.AddCategory(name)
		fmt.Println("2---------")
		if err != nil {
			fmt.Println("3---------")
			fmt.Println(err)
			beego.Error(err)
		}
		this.Redirect("/category", 301)
		return
	case "del":
		id := this.Input().Get("id")
		if len(id) == 0 {
			break
		}
		err := models.DelCategory(id)
		if err != nil {
			beego.Error(err)
		}
		this.Redirect("/category", 301)
		return
	}

	this.Data["IsCategory"] = true
	this.TplName = "category.html"
	var err error
	this.Data["Categories"], err = models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}

}

func (this *CategoryController) Post() {
	op := this.Input().Get("op")
	name1 := this.Input().Get("name")
	fmt.Println("-------------op:", op)
	fmt.Println("-------------name1:", name1)

	switch op {
	case "add":
		name := this.Input().Get("name")
		if len(name) == 0 {
			break
		}
		fmt.Println("1---------")
		err := models.AddCategory(name)
		fmt.Println("2---------")
		fmt.Println(err)
		if err != nil {
			fmt.Println("3---------")
			beego.Error(err)
		}
		this.Redirect("/category", 301)
		return
	case "del":
		id := this.Input().Get("id")
		if len(id) == 0 {
			break
		}
		err := models.DelCategory(id)
		if err != nil {
			beego.Error(err)
		}
		this.Redirect("/category", 301)
		return
	}

}
