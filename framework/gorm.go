//	数据库处理,使用Grom
//  依赖情况:
//			"github.com/jinzhu/gorm"

//	_ "github.com/jinzhu/gorm/dialects/mssql"
//	_ "github.com/jinzhu/gorm/dialects/mysql"
//	_ "github.com/jinzhu/gorm/dialects/postgres"
//	_ "github.com/jinzhu/gorm/dialects/sqlite"

package framework

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
)

var Gorm  map[string]*gorm.DB

func init() {
	Gorm = make(map[string]*gorm.DB)
}

// 初始化Gorm
func NewDB(dbname string) {

	var orm *gorm.DB
	var err error

	//默认配置
	AppConfig.SetDefault(dbname, map[string]interface{}{
		"dbHost"          : "127.0.0.1",
		"dbName"          : "phalgo",
		"dbUser"          : "root",
		"dbPasswd"        : "",
		"dbPort"          : 3306,
		"dbIdleconns_max" : 0,
		"dbOpenconns_max" : 20,
		"dbType"          : "mysql",
	})
	dbHost := AppConfig.GetString(dbname + ".dbHost")
	dbName := AppConfig.GetString(dbname + ".dbName")
	dbUser := AppConfig.GetString(dbname + ".dbUser")
	dbPasswd := AppConfig.GetString(dbname + ".dbPasswd")
	dbPort := AppConfig.GetString(dbname + ".dbPort")
	dbType := AppConfig.GetString(dbname + ".dbType")

	orm, err = gorm.Open(dbType, dbUser + ":" + dbPasswd + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8")
	//开启sql调试模式
	//GDB.LogMode(true)
	if err != nil {
		fmt.Println("数据库连接异常!")
	}
	//连接池的空闲数大小
	orm.DB().SetMaxIdleConns(AppConfig.GetInt(dbname + ".idleconns_max"))
	//最大打开连接数
	orm.DB().SetMaxIdleConns(AppConfig.GetInt(dbname + ".openconns_max"))
	Gorm[dbname] = orm
	//defer Gorm[dbname].Close()
}

// 通过名称获取Gorm实例
func GetORMByName(dbname string) *gorm.DB {

	return Gorm[dbname]
}

// 获取默认的Gorm实例
func GetORM() *gorm.DB {

	return Gorm["dbDefault"]
}
