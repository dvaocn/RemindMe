package main

import (
	"RemindMe/core"
	"RemindMe/global"
	"RemindMe/initialize"
	"fmt"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// @title Swagger Example API
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
// @BasePath /
func main() {
	global.VP = core.Viper()      // 初始化Viper
	fmt.Println(global.Config.Zap)
	global.Log = core.Zap()       // 初始化zap日志库
	global.DB = initialize.Gorm() // gorm连接数据库
	initialize.Timer()
	if global.DB != nil {
		initialize.MysqlTables(global.DB) // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.DB.DB()
		defer db.Close()
	}
	core.RunWindowsServer()
}
