package route

import (
	"example/hcmut-lib/pkg/controller"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine){
	r.GET("/api/getUserAccount/:id", controller.GetUserAccount)
	r.GET("/api/getAllUserAccount", controller.GetAllUserAccount)
	r.POST("/api/createUserAccount", controller.CreateUserAccount)
	r.DELETE("/api/deleteUserAccount/:id", controller.DeleteUserAccount)
	r.PUT("/api/UpdateUserAccount/:id", controller.UpdateUserAccount)
}