package system

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type JwtAPI struct{}

func (j *JwtAPI) JsonInBlacklist(c *gin.Context) {
	token := c.Request.Header.Get("x-token")
	fmt.Println("token", token)
}
