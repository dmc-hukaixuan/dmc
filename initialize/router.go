package initialize

import (
	v1 "dmc/api/v1"
	//	"dmc/api/v1/system"
	"dmc/global/log"
	"dmc/kernel/middleware"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Routers() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	// 跨域，如需跨域可以打开下面的注释
	// router.Use(middleware.Cors()) // 直接放行全部跨域请求
	// router.Use(middleware.CorsByRules()) // 按照配置的规则放行跨域请求
	router.Use(Cors())
	router.Use(Recovery)
	// Router.Use(middleware.LoadTls())  // 打开就能玩https了

	fmt.Println("router init *--------1111")
	// 获取路由组件实例

	group := router.Group("")
	group.Use(middleware.JWTAuth())
	{
		group.GET("/user/admin/ticketTemplate", v1.APIGroupApp.Admin.TicketTemplateApi.Base)
		//group.GET("/user/login", system.Auth)
		group.GET("/user/admin/processManagement", v1.APIGroupApp.Admin.ProcessManagementApi.ProcessOverview)
		group.GET("/user/admin/processManagement/typeSave", v1.APIGroupApp.Admin.ProcessManagementApi.ProcessTypeSave)
		group.GET("/user/admin/processManagement/edit", v1.APIGroupApp.Admin.ProcessManagementApi.ProcessManagement)
		group.GET("/user/admin/processManagement/typeEdit/*id", v1.APIGroupApp.Admin.ProcessManagementApi.ProcessTypeEdit)
		//group.GET("/user/test", system.Test)
		// group.GET("/user/:uuid", v1.GetUserDetails)
	}
	return router

}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") //请求头部
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		// 允许类型校验
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "ok!")
		}

		defer func() {
			if err := recover(); err != nil {
				log.Logger.Error("HttpError", zap.Any("HttpError", err))
				fmt.Println("err", err)
			}
		}()

		c.Next()
	}
}

func Recovery(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("err", r)
			log.Logger.Error("gin catch error: ", log.Any("gin catch error: ", r))
			// c.JSON(http.StatusOK, response.FailMsg("系统内部错误"))
		}
	}()
	c.Next()
}
