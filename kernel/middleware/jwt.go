package middleware

import (
	"dmc/kernel/model/common/response"
	"dmc/kernel/service"
	"fmt"

	"github.com/gin-gonic/gin"
)

// var jwtService = service.ServiceGroupApp.SystemServiceGroup.JwtService

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 这里 jwt 鉴权取头部信息 x-token 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localStorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
		token := c.Request.Header.Get("dmc-token")
		remoteAddr := c.ClientIP()
		remoteUserAgent := c.GetHeader("User-Agent")
		sess, err := c.Cookie("serviceCoolToken")

		fmt.Println("token ^^^^^^^^^^^^", token, sess, err)
		if token == "" {
			response.FailWithDetailed(gin.H{"reload": true}, "未登录或非法访问", c)
			c.Abort()
			return
		}
		// if jwtService.IsBlacklist(token) {
		// 	response.FailWithDetailed(gin.H{"reload": true}, "您的帐户异地登陆或令牌失效", c)
		// 	c.Abort()
		// 	return
		// }
		// check session id

		j, Message, sessionData := service.CheckSessionID(token, remoteAddr, remoteUserAgent)
		if j == 0 {
			response.FailWithDetailed(gin.H{"reload": true}, Message, c)
			c.Abort()
			return
		}
		// get session data
		// j, err := service.GetSessionIDData(sess)

		//fmt.Println(" j --------------------: ", j, "err", err)
		// j := utils.NewJWT()
		// // parseToken 解析token包含的信息
		// claims, err := j.ParseToken(token)
		// if err != nil {
		// 	if err == utils.TokenExpired {
		// 		response.FailWithDetailed(gin.H{"reload": true}, "授权已过期", c)
		// 		c.Abort()
		// 		return
		// 	}
		// 	response.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
		// 	c.Abort()
		// 	return
		// }
		// 用户被删除的逻辑 需要优化 此处比较消耗性能 如果需要 请自行打开
		//if err, _ = userService.FindUserByUuid(claims.UUID.String()); err != nil {
		//	_ = jwtService.JsonInBlacklist(system.JwtBlacklist{Jwt: token})
		//	response.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
		//	c.Abort()
		//}
		// if claims.ExpiresAt-time.Now().Unix() < claims.BufferTime {
		// 	claims.ExpiresAt = time.Now().Unix() + global.GVA_CONFIG.JWT.ExpiresTime
		// 	newToken, _ := j.CreateTokenByOldToken(token, *claims)
		// 	newClaims, _ := j.ParseToken(newToken)
		// 	c.Header("new-token", newToken)
		// 	c.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt, 10))
		// 	if global.GVA_CONFIG.System.UseMultipoint {
		// 		err, RedisJwtToken := jwtService.GetRedisJWT(newClaims.Username)
		// 		if err != nil {
		// 			global.GVA_LOG.Error("get redis jwt failed", zap.Error(err))
		// 		} else { // 当之前的取成功时才进行拉黑操作
		// 			//_ = jwtService.JsonInBlacklist(system.JwtBlacklist{Jwt: RedisJwtToken})
		// 		}
		// 		// 无论如何都要记录当前的活跃状态
		// 		_ = jwtService.SetRedisJWT(newToken, newClaims.Username)
		// 	}
		// }

		c.Set("sessionData", sessionData)
		c.Set("userID", sessionData.UserID)
		c.Set("token", j)
		//	fmt.Println("j ", j)
		c.Next()
	}
}
