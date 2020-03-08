package main

import (
	"github.com/astaxie/beego"
)

func main() {
	beego.Info("项目启动")
	beego.Run("localhost:8080")
}
