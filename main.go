package main

import (
	"fmt"
	"net/http"
	"time"

	config "dmc/config/database"
	"dmc/global"
	"dmc/global/log"
	"dmc/initialize"
	db "dmc/initialize/database"
)

func main() {
	// 实例化日志配置
	log.InitLogger(config.GetConfig().Log.Path, config.GetConfig().Log.Level)
	log.Logger.Info("config", log.Any("config", config.GetConfig()))

	log.Logger.Info("start server", log.String("start", "start web sever..."))

	router := initialize.Routers()
	address := fmt.Sprintf(":%d", config.GetConfig().SysInfo.Port)
	global.GVA_DB = db.Gorm()
	fmt.Println("address *---------------", address)
	s := &http.Server{
		Addr:           ":8503",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	if nil != err {
		log.Logger.Error("server error", log.Any("serverError", err))
	}
}
