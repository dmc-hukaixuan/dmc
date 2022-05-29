package system

import (
	"dmc/global/log"
	"dmc/kernel/model/common/response"
	"dmc/kernel/model/user"
	"dmc/kernel/service"
	"dmc/kernel/service/auth"
	"fmt"

	"github.com/gin-gonic/gin"
)

type AuthApi struct{}

type login struct {
	Username string `json:"username"form:"username" binding:"required"`  // 用户名
	Password string `json:"password" form:"password" binding:"required"` // 密码
	// Captcha   string `json:"captcha"`   // 验证码
	// CaptchaId string `json:"captchaId"` // 验证码ID
}

func Test(c *gin.Context) {
	fmt.Println("gin test ---------------------")
}

// auth user
func Auth(c *gin.Context) {
	fmt.Println("system auth : ")
	var login_param login
	name := c.Query("username")
	password := c.Query("password")
	login_param = login{
		Username: name,
		Password: password,
	}
	fmt.Println("login_param : ", login_param, " name ", name, password)
	// if err := c.ShouldBindJSON(&login_param); err != nil {
	// 	// 返回错误信息
	// 	// gin.H封装了生成json数据的工具
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }
	log.Logger.Debug("user", log.Any("user", login_param))
	fmt.Println("login_param : ", login_param, " name ", name)
	if login_param.Username == "" || login_param.Password == "" {
		return
	}
	u := &user.LoginParam{Login: login_param.Username, PW: login_param.Password}
	fmt.Println("login_param : ", login_param)
	if err, user := auth.UserAuth(u); err != nil {
		fmt.Println("auth error ")
	} else {
		// create user seesion
		tokenCreate(c, *user)
	}
}

// login starand return structure
type LoginResponse struct {
	User      user.User `json:"user"`
	Token     string    `json:"token"`
	ExpiresAt int64     `json:"expiresAt"`
}

// 用户登陆以后返回 token
func tokenCreate(c *gin.Context, user user.User) {
	fmt.Println("User : ", c.Request.Header.Get("User-Agent"))
	err, token := service.CreateSessionID(&user)
	if err != nil {
		log.Logger.Error("Get seesion fail", log.Any("serverError", err))
		response.FailWithMessage("Get seesion fail", c)
		return
	}

	// 检查系统配置中是否配置允许多点登陆，

	// 给浏览器设置 cookie
	c.SetCookie("dmc-token", token, 3600, "/", "localhost", false, true)

	// 写入缓存
	response.SuccessWithDetailed(LoginResponse{
		User:      user,
		Token:     token,
		ExpiresAt: 7200,
	}, "Login successful", c)
}
