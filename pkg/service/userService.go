package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func BindJsonService(c *gin.Context, obj interface{}) {
	if err := c.BindJSON(&obj); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
}
