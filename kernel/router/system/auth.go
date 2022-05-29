package system

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type AuthRouter struct{}

func (s *AuthRouter) AuthUserRouter(Router *gin.RouterGroup) {
	initRouter := Router.Group("user")
	fmt.Println("auto user : ")
	dbApi := v1.ApiGroupApp.SystemApiGroup.DBApi
	{
		initRouter.POST("initdb", dbApi.InitDB)   // 创建Api
		initRouter.POST("checkdb", dbApi.CheckDB) // 创建Api
	}
}
