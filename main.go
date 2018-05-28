package main

import (
	_ "chenhbblog/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"chenhbblog/models"
)

func init(){
	models.RegisterDB()
}

func main() {
	orm.Debug = true
	orm.RunSyncdb("default",false,true)

	beego.Run()
}

