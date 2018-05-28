package controllers

import(
	"fmt"
	//"net/url"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type LoginController struct{
	beego.Controller
}

func (this *LoginController)Get(){
	fmt.Println("-----------------")
	//this.Data["IsLogin"] = true
	isExit := this.Input().Get("exit") == "true"
	fmt.Println("-------------------")
	fmt.Println(isExit)
	if isExit{
		this.Ctx.SetCookie("uname","",-1,"/")
		this.Ctx.SetCookie("pwd","",-1,"/")
		this.Redirect("/",301)
		//return
	}
	this.TplName = "login.html"
}

func (this *LoginController)Post(){
	fmt.Println("---------------------")
	//this.Ctx.WriteString(fmt.Sprint(this.Input()))
	uname := this.Input().Get("uname")
	pwd := this.Input().Get("pwd")
	autoLogin := this.Input().Get("autoLogin") == "on"

	if beego.AppConfig.String("uname") == uname &&
	beego.AppConfig.String("pwd") == pwd{
		maxAge:=0
		if autoLogin{
			maxAge = 10
		}
		this.Ctx.SetCookie("uname",uname,maxAge,"/")
		this.Ctx.SetCookie("pwd",pwd,maxAge,"/")

	}
	this.Redirect("/",301)
	return
}

func checkAccount(ctx *context.Context) bool{
	ck,err := ctx.Request.Cookie("uname")
	
	if err!=nil{
		return false
	}
	uname:= ck.Value
	// fmt.Println(uname,err)
	ck,err = ctx.Request.Cookie("pwd")
	
	if err!=nil{
		return false
	}
	pwd:= ck.Value
	// fmt.Println(pwd,err)
	return beego.AppConfig.String("uname")==uname &&
	beego.AppConfig.String("pwd")==pwd

}