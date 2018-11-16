package handler

import (
	"demo/src/util"
	"github.com/gin-gonic/gin"
	"github.com/sevenNt/wzap"
)

type LoginController struct{}

func NewLoginController() *LoginController {
	return &LoginController{}
}

type authRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (login *LoginController) Login(c *gin.Context) {
	var auth authRequest
	if err := c.BindJSON(&auth); err != nil {
		util.OpFailFn(c, err)
		return
	}
	wzap.Debug("login user ", "user", auth)
	c.JSON(200, gin.H{
		"code": 20000,
		"data": map[string]interface{}{
			"token": "testtssssssss",
		},
	})
}

// 校验token并返回当前用户所拥有的权限
func (login *LoginController) UserInfo(c *gin.Context) {
	token := c.Query("token")
	if len(token) <= 0 {
		wzap.Error("token is null", "token", token)
		return
	}
	//TODO 校验token从token中获取用户id
	var tmpuser = 1
	c.JSON(200, gin.H{
		"code": 20000,
		"data": map[string]interface{}{
			"id":     tmpuser,
			"name":   "admin",
			"avatar": "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif",
			"roles": []string{
				"admin",
			},
		},
	})
}

func (login *LoginController) LoginOut(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 20000,
		"data": "success",
	})
}
