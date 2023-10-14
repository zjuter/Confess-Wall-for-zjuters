package main

import (
	"BBQ/config/database"
	"BBQ/config/router"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Init()
	r := gin.Default()
	
	//这里是设置跨域访问，我也不知道是不是必要的，反正就是问chatgpt，不知不觉就能访问了。

	// 设置跨域请求头信息
	r.Use(func(c *gin.Context) {
		// 设置允许的请求来源
		c.Header("Access-Control-Allow-Origin", "http://localhost:5173")
		
		// 设置允许的请求方法
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		
		// 设置允许的请求头
		allowedHeaders := "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Token"
		requestHeaders := c.Request.Header.Get("Access-Control-Request-Headers")
		if len(requestHeaders) > 0 {
			allowedHeaders += ", " + requestHeaders
		}
		c.Header("Access-Control-Allow-Headers", allowedHeaders)
		
		// 如果是预检请求 OPTIONS，直接返回成功状态码和空响应体
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		
		// 继续处理其他请求
		c.Next()
	})

	router.Init(r)

	err := r.Run("127.0.0.1:8080")
	if err != nil {
		log.Fatal("Server start failed: ", err)
	}
}
