package service
/*
 * @Author : Duc Hung Ho
 * @Description : Service for user controllers
 * @Param: *gin.Context, *model.User
 * @return *model, error
 */
import (
	"example/hcmut-lib/config"
	"example/hcmut-lib/pkg/model"

	"github.com/gin-gonic/gin"
)
func GetAllUser(c *gin.Context, obj *[]model.User) ([]model.User, error) {
	if err := config.DB.Find(&obj).Error; err != nil {
		return *obj, err
	}
	return *obj, nil
}

func GetUserByID(c *gin.Context, obj *model.User, id string) (model.User, error) {
	if err := config.DB.Where("id_account = ?", id).Find(&obj).Error; err != nil {
		return *obj, err
	}
	return *obj, nil
}

func UpdateUserByID(c *gin.Context, obj *model.User, id string) (model.User, error) {
	var user model.User
	if err := c.BindJSON(&user); err != nil {
		return user, err
	}
	if err := config.DB.Where("id_account = ?", id).Updates(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func DeleteUserByID(c *gin.Context, obj *model.User, id string) {

}

func CreateUser(c *gin.Context, obj *model.User) {

}
