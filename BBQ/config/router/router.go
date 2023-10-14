package router

import (
    "BBQ/app/controllers/userController"
    "BBQ/app/controllers/articleController"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine){
	const pre1 = "/api"

	api :=r.Group(pre1)
	{
		api.POST("/login", userController.Login)
		api.POST("/reg", userController.Register)

	}

	const pre2 = "/my"
	my  :=r.Group(pre2)
	{
		my.POST("/article/add", articleController.CreateArticle)
		my.PUT("/article/info", articleController.UpdateArticle)
		my.DELETE("/article/info",articleController.DeleteArticle)
		my.GET("/article/info", articleController.GetAritical)
		my.GET("/article/list", articleController.GetList)
	}


}
