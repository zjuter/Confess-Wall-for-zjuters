package articleController

import (
	"BBQ/app/utils"
	"BBQ/app/services/articleService"
	"BBQ/app/models"

	"github.com/gin-gonic/gin"
)

type UpdateArticleData struct {                   
	ID       string `json:"id" binding:"required"`
	Title    string `json:"title" binding:"required"`
	CateID   string `json:"cate_id" binding:"required"`
	Content  string `json:"content" binding:"required"`
	CoverImg string `json:"cover_img"`
	State    string `json:"state" binding:"required"`
}
func UpdateArticle(c *gin.Context) {
	var data UpdateArticleData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}
	//获取表白信息
	var article *models.Article
	article, err = articleService.GetArticleByID(data.ID)
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
    //判断信息是否修改
	flag := articleService.CompareCon(data.Content, article.Content)
	if flag {
		utils.JsonErrorResponse(c, 1, "编辑文章失败!")
		return
	}
    //更新文章
	err = articleService.UpdateArticle(models.Article{
		ID:       data.ID,
		Title:    data.Title,
	    CateID:   data.CateID,
	    Content:  data.Content,
		CoverImg: data.CoverImg,
	    State:    data.State,
	})
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	utils.JsonErrorResponse(c, 0, "修改文章成功")
}
