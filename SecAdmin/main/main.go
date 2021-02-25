package main

import (
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/zbrechave/seckill/SecAdmin/router"
)

func main() {
	err := initAll()
	if err != nil {
		panic(fmt.Sprintf("init database failed, err:%v", err))
		return
	}
	beego.Run()
}
