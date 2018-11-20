package handler

import (
	"demo/src/model"
	"demo/src/service"
	"demo/src/util"
	"github.com/gin-gonic/gin"
	"github.com/sevenNt/wzap"
	"net/http"
	"strconv"
)

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

func (u *UserController) AddUser(c *gin.Context) {
	var tmpuser model.User
	err := c.BindJSON(&tmpuser)
	if err != nil {
		wzap.Err(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 1,
			"data": err.Error(),
		})
	}
	wzap.Debug("user message", &tmpuser)
	service.User.AddUser(&tmpuser)
	c.JSON(200, "hello")
}

func (u *UserController) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		wzap.Error("类型转换异常")
		util.OpFailFn(c, err)
		return
	}
	wzap.Debug("delete user by id ", "id", id)
	if err := service.User.Delete(id); err != nil {
		wzap.Err(err)
		util.OpFailFn(c, err)
		return
	}

	c.JSON(200, "删除成功")
}

//List user by page
func (u *UserController) ListByPage(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	content, ok := c.GetQuery("content")
	if !ok {
		content = ""
	}
	users, total, err := service.User.ListByPage(page, pageSize, content)

	if err != nil {
		wzap.Err(err)
		util.OpFailFn(c, err)
		return
	}
	c.JSON(200, gin.H{
		"code": 20000,
		"data": map[string]interface{}{
			"rows":  users,
			"total": total,
		},
	})

}

func (u *UserController) Details(c *gin.Context) {

}

func (u *UserController) Update(c *gin.Context) {

}
