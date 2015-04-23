package im

import (
	"database/sql"
	"im-go/im/common"
	"log"
)

/*
全局变量
*/
var (
	Database *sql.DB          = nil // 数据库操作对象
	Config   *common.IMConfig       //配置对象
)

/*
启动服务方法
*/
func Start(config *common.IMConfig) {
	Config = config

	// 连接db
	var err error
	Database, err = config.DBConfig.Connect()
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Printf("获取到得数据库连接: %s", config.DBConfig.Name)
	defer Database.Close()

	go func() {
		// 启动HTTP服务
		err := StartHttpServer(*config)
		log.Fatalf("HttpServer", err)
	}()

	// 启动IM服务
	err = StartIMServer(*config)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("获取到得数据库连接: %s", config.DBConfig.Name)
}
