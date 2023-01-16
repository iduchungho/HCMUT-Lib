package controller

import (
	"example/hcmut-lib/config"
	"example/hcmut-lib/pkg/model"
	"example/hcmut-lib/pkg/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllUserAccount(c *gin.Context) {
	var user []model.User
	users , err := service.GetAllUser(c, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, &users)
}

func GetUserAccount(c *gin.Context) {
	var user model.User
	users, err := service.GetUserByID(c, &user, c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, &users)
}

func UpdateUserAccount(c *gin.Context) {
	var user model.User
	users, err := service.UpdateUserByID(c, &user, c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, &users)
}

func CreateUserAccount(c *gin.Context) {
	var newUser model.User

	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}

	if err := config.DB.Create(&newUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, &newUser)
}

func DeleteUserAccount(c *gin.Context) {
	var user model.User

	if err := config.DB.Where("id_account = ?", c.Param("id")).Delete(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]string{"message": "user delete successfully"})
}
