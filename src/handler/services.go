package handler

import (
	"demo/src/service"
	"demo/src/util"
	"github.com/gin-gonic/gin"
)

type ServiceController struct{}

func NewServiceController() *ServiceController {
	return &ServiceController{}
}

// 获取所有service
func (svc *ServiceController) ServiceList(c *gin.Context) {
	namespace := c.Query("namespace")
	svcService := service.NewSvcService()
	services, err := svcService.GetServiceList(namespace)
	if err != nil {
		util.OpFailFn(c, err)
		return
	}

	c.JSON(200, gin.H{
		"code": 20000,
		"data": services,
	})

}
