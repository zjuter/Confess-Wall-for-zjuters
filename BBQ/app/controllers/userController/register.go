package userController

import (
	"BBQ/app/models"
	"BBQ/app/services/userService"
	"BBQ/app/utils"

	"github.com/gin-gonic/gin"
)

type RegisterData struct {
	Username        string `json:"username" binding:"required"`
	Password        string `json:"password" binding:"required"`
	RePassword      string `json:"repassword" binding:"required"`
}

// 注册
func Register(c *gin.Context){
	var data RegisterData
	err := c.ShouldBindJSON(&data)
	
	if err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}
	// 判断账号是否已经注册
	err = userService.CheckUserExistByUsername(data.Username)
	if err == nil {
		utils.JsonErrorResponse(c, 1, "用户名被占用，请更换其他用户名！")
		return
	}
	// 判断密码是否一致
	flag := userService.ComparePwd(data.Password, data.RePassword)
	if !flag {
		utils.JsonErrorResponse(c, 200505, "密码不一致")
		return
	}

	// 注册用户
	err = userService.Register(models.User{
		Username: data.Username,
		Password: data.Password,
	})
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	utils.JsonErrorResponse(c, 0, "注册成功")
}