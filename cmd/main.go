package main

import (
	"fmt"
	"gin_frame/config"
	"gin_frame/global"
	"gin_frame/router"
	"go.uber.org/zap"
)

func main() {
	//读取配置
	err := config.Init("config.yaml")
	if err != nil {
		return
	}

	//初始化日志
	err = global.InitLogger(config.Conf)
	if err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	defer zap.L().Sync()
	zap.L().Debug("logging init success")

	//初始化mysql
	err = global.InitMysql(config.Conf.DataBase)
	if err != nil {
		zap.L().Error("初始化mysql失败 %v", zap.Error(err))
		return
	}

	defer func() {
		if err = global.MysqlClose(); err != nil {
			zap.L().Error("%v", zap.Error(err))
			return
		}
	}()
	zap.L().Debug("mysql init success")

	//初始化路由
	r := router.InitRouter()
	r.Run(config.Conf.Http.Addr)
}
