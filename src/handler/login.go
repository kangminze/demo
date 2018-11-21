package handler

import (
	"demo/src/middleware/jwtpath"
	"demo/src/service"
	"demo/src/util"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sevenNt/wzap"
	"time"
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
	user, err := service.User.Validate(auth.Username, auth.Password)
	//校验用户名密码
	if err != nil {
		util.OpFailFn(c, err)
		return
	}
	//创建token信息
	j := jwtpath.NewJWT()
	claim := jwtpath.CustomClaims{
		ID:   user.ID,
		Name: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}
	token, err := j.CreateToken(claim)
	if err != nil {
		util.OpFailFn(c, err)
		return
	}
	c.JSON(200, gin.H{
		"code": 20000,
		"data": map[string]interface{}{
			"token": token,
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
