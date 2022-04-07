package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"youyin/common"
	"youyin/middleware"
	"youyin/model"
	"youyin/router"
)

func main() {
	common.InitConfig("config/config.yaml")
	err := middleware.Init()
	if err != nil {
		return
	}
	r := gin.Default()
	r.Use(middleware.GinLogger(), middleware.GinRecovery(true))
	model.InitDB()

	//r.Use(middleware.Cors())
	r = router.InitRouter(r)

	addr := common.Conf.APP.Addr
	port := common.Conf.APP.Port

	err = r.Run(fmt.Sprintf("%s:%s", addr, port))
	if err != nil {
		fmt.Println("启动失败")
		return
	}
}
