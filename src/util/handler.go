package util

import (
	"github.com/gin-gonic/gin"
	"github.com/sevenNt/wzap"
	"net/http"
)

//OpFailFn
var OpFailFn = func(c *gin.Context, e error) {
	wzap.Err(e)
	c.JSON(http.StatusInternalServerError, gin.H{
		"code": "500",
		"data": e.Error(),
	})
}
