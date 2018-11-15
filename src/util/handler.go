package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//OpFailFn
var OpFailFn  = func(c *gin.Context, e error) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"code": "500",
		"data": e.Error(),
	})
}
