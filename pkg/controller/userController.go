package controller

import (
	"net/http"
	"example/hcmut-lib/config"
	"example/hcmut-lib/pkg/model"

	"github.com/gin-gonic/gin"
)

func GetAllUserAccount(c *gin.Context) {
	var user []model.User

	if err := config.DB.Find(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, &user)
}

func GetUserAccount(c *gin.Context) {
	var user []model.User

	if err := config.DB.Where("id_account = ?", c.Param("id")).Find(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, &user)
}

func UpdateUserAccount(c *gin.Context) {
	var user model.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}

	if err := config.DB.Model(&user).Where("id_account = ?", c.Param("id")).Updates(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]string{"message": "user updated successfully"})
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
