package articleController

import (
	"BBQ/app/models"
	"BBQ/app/services/articleService"
	"BBQ/app/utils"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GetArticleData struct {
	ID string `form:"id" binding:"required"`
}

// 获取详情
func GetAritical(c *gin.Context) {
	var data GetArticleData
	err := c.ShouldBindQuery(&data)
	fmt.Println(err)
	if err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}

	var article *models.Article

article, err = articleService.GetAriticalContext(data.ID)
//fmt.Println(article)
if err != nil {
    if err == gorm.ErrRecordNotFound {
        utils.JsonErrorResponse(c, 1, "没有查到对应数据！")
        return
    } else {
        utils.JsonInternalServerErrorResponse(c)
        return
    }
}


	c.JSON(200, gin.H{
		"code":    0,
		"message": "获取文章成功！",
		"data":    article,
	})
}
