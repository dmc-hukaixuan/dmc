package system

import (
	"github.com/gin-gonic/gin"
)

type JWTRouter struct{}

func (s *JWTRouter) InitJwtRouter(Router *gin.RouterGroup) {
	jwtRouter := Router.Group("jwt")

}
