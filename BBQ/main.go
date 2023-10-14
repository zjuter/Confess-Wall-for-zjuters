package userController

import (
	"BBQ/app/models"
    "BBQ/app/services/userService"
	"BBQ/app/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LoginData struct{
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context){
	//接收参数
	var data LoginData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}

	 //判断用户是否存在
	err = userService.CheckUserExistByUsername(data.Username)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.JsonErrorResponse(c, 1, "账号或密码错误，请重新输入叭！")
	    }
		return
	}
 
	 //获取用户信息
	var user *models.User
	user, err = userService.GetUserByUsername(data.Username)
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	 //判断密码是否正确
	flag := userService.ComparePwd(data.Password, user.Password)
	if !flag {
		utils.JsonErrorResponse(c, 1, "账号或密码错误，请重新输入叭！")
		return
	}
	// 返回用户信息
	utils.JsonErrorResponse(c, 0, "登录成功")

}
