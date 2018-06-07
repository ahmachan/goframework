// 使用viper配置文件获取工具作为配置文件工具
// 依赖情况:
// "github.com/spf13/viper"

package framework

import (
	"github.com/spf13/viper"
	"fmt"
)

var AppConfig *viper.Viper

//初始化配置文件
func InitConfig(filePath string, fileName string) {

	AppConfig = viper.New()
	AppConfig.WatchConfig()
	AppConfig.SetConfigName(fileName)
	//filePath支持相对路径和绝对路径 etc:"/a/b" "b" "./b"
	if (filePath[:1] != "/"){
		AppConfig.AddConfigPath(GetPath() + "/" + filePath + "/")
	}else{
		AppConfig.AddConfigPath(filePath + "/")
	}
	// 找到并读取配置文件并且 处理错误读取配置文件
	if err := AppConfig.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err).Error())
	}
}


