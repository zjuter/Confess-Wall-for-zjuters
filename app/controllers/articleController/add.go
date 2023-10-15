package articleController

import (
	"BBQ/app/models"
	"BBQ/app/services/articleService"
	"BBQ/app/utils"
	"github.com/gin-gonic/gin"
)

//由于前端部分是form，所以我把这里改成了form类型使得可以对接成功

type CreateArticleData struct {
	ID       string                `form:"id" binding:"-"`
	Title    string                `form:"title" binding:"required"`
	CateID   string                `form:"cate_id" binding:"required"`
	Content  string                `form:"content" binding:"required"`
	State    string                `form:"state" binding:"required"`
}

// 发布表白
func CreateArticle(c *gin.Context) {
	var data CreateArticleData
	err := c.ShouldBind(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 2, "参数错误")
		return
	}

	// 判断内容是否为空
	flag := articleService.CompareCon(data.Content, "")
	if flag {
		utils.JsonErrorResponse(c, 2, "\"content\" is required")
		return
	}

	// 判断状态是否为空
	flag = articleService.CompareCon(data.State, "")
	if flag {
		utils.JsonErrorResponse(c, 2, "\"state\" is required")
		return
	}

	// 决定表白是否要发布
	flag1 := articleService.CompareCon(data.State, "草稿")
	flag2 := articleService.CompareCon(data.State, "发布")
	if flag1 || flag2 {
		article := models.Article{
			Title:    data.Title,
			CateID:   data.CateID,
			Content:  data.Content,
			State:    data.State,
		}
		
		err = articleService.Article(article)
		if err != nil {
			utils.JsonInternalServerErrorResponse(c)
			return
		}

	} else {
		utils.JsonErrorResponse(c, 2, "\"state\" must be one of [草稿, 发布]")
	}
}
