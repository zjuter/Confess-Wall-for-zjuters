package articleController

import (
	"BBQ/app/services/articleService"
	"BBQ/app/utils"

	"github.com/gin-gonic/gin"
)

type DeleteArticleData struct {
	ID string `form:"id" binding:"required"`
}

func DeleteArticle(c *gin.Context) {
	var data DeleteArticleData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}
	//判断表白是否存在
	err = articleService.CheckArticleExistByArticle(data.ID)
	if err != nil {
		utils.JsonErrorResponse(c, 1, "您要删除的表白不存在！")
		return
	}
	err = articleService.DeleteArticle(data.ID)
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	utils.JsonErrorResponse(c, 0, "删除成功")
}
