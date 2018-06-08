package main

import (
	"goframework/framework"
	"goframework/router"
        "fmt"
)

func main() {
     // 初始化配置文件
     framework.InitConfig("Config", "conf")
     // 初始化数据库连接
     framework.NewDB("dbDefault")
     // 载入并初始化路由      
     instance:=routers.NewRouter()
     var port string = framework.AppConfig.GetString("system.port")
     if (port==":80"){
         port = ""
     }

    instance.Run(fmt.Sprintf("0.0.0.0%s", port))
}
