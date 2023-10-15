package userController

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"net/http"
)


//和chatgpt交流许久，更新了一下这部分，创建了一个虚拟的token，因为前端那里需要以token来记录登陆状态。
//本来想和别的组搞得一样用JWT，但是我只会问chatgpt，搞不来，有点麻烦，总是麻烦，就直接用随机数当token了。
//仅演示用

type LoginData struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	//接收参数
	var data LoginData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		JsonErrorResponse(c, 200501, "参数错误")
		return
	}

	// 生成虚假的 token
	token := generateRandomToken()

	// 返回用户信息及生成的 token
	JsonSuccessResponse(c, gin.H{
		"token": token,
	})
}

func generateRandomToken() string {
	// 定义 token 长度（字节数）
	tokenLength := 32

	// 生成随机字节片段
	randomBytes := make([]byte, tokenLength)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err)
	}

	// 将字节片段编码为字符串
	token := base64.URLEncoding.EncodeToString(randomBytes)

	return token
}

func JsonSuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "Success",
		"data":    data,
	})
}

func JsonErrorResponse(c *gin.Context, code int, message string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"code":    code,
		"message": message,
	})
}
