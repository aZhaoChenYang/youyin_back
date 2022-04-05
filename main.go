package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"youyin/common"
	"youyin/middleware"
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
	common.InitDB()

	//创建DB实例进行数据库操作
	db := common.GetDB()
	//延迟关闭数据库
	defer db.Close()

	//r.Use(middleware.Cors())
	r = router.InitRouter(r)

	addr := common.Conf.APP.Addr
	port := common.Conf.APP.Port

	r.Run(fmt.Sprintf("%s:%s", addr, port))
}
