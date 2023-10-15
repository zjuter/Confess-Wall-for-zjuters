package userService

import (
	"BBQ/app/models"
	"BBQ/config/database"
)

func CheckUserExistByUsername(Username string) error {
	result := database.DB.Where("username= ?", Username).First(&models.User{})
	return result.Error
}

func GetUserByUsername(Username string) (*models.User, error) {
	var user models.User
	result := database.DB.Where("username = ?", Username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func ComparePwd(pwd1 string, pwd2 string) bool {
	return pwd1 == pwd2
}

func Register(user models.User) error {
	result := database.DB.Create(&user)
	return result.Error
}
