package articleController

import (
	"BBQ/app/models"
	"BBQ/app/services/articleService"
	"BBQ/app/utils"

	"github.com/gin-gonic/gin"
)

type CreateArticleData struct {
	ID       string `json:"id" binding:"required"`
	Title    string `json:"title" binding:"required"`
	CateID   string `json:"cate_id" binding:"required"`
	Content  string `json:"content" binding:"required"`
	CoverImg string `json:"cover_img"`
	State    string `json:"state" binding:"required"`
	Author   string `json:"author" binding:"required"`
	Anonymous string `json:"anonymous" binding:"required"`
}

//发布表白
func CreateArticle(c *gin.Context) {
	var data CreateArticleData
	err := c.ShouldBindJSON(&data)
	if err != nil {
        //判断内容是否为空
	    flag := articleService.CompareCon(data.Content, "")
	    if flag {
		    utils.JsonErrorResponse(c, 2, "\"content\" is required")
		    return
	    }
		//判断状态是否为空
	    flag = articleService.CompareCon(data.State, "")
	    if flag {
			utils.JsonErrorResponse(c, 2,  "\"state\" is required")
			return
		}
	}
	//决定表白是否要发布
	flag1 := articleService.CompareCon(data.State, "草稿")
	flag2 := articleService.CompareCon(data.State, "发布")
	//草稿
	if  flag1  {
		err = articleService.Article(models.Article{
			ID:       data.ID,
			Title:    data.Title,
			CateID:   data.CateID,
			Content:  data.Content,
			CoverImg: data.CoverImg,
			State:    data.State,
			Author:   data.Author,
			Anonymous:data.Anonymous,
		})
		if err != nil {
			utils.JsonInternalServerErrorResponse(c)
			return
		}
		utils.JsonErrorResponse(c, 0, "发布文章成功")
		return
	}else if flag2  {
		//草稿变成发布    
		err = articleService.CheckArticleExistByArticle(data.ID)
		if err == nil{
			err = articleService.UpdateArticle(models.Article{
				ID:       data.ID,
				Title:    data.Title,
				CateID:   data.CateID,
				Content:  data.Content,
				CoverImg: data.CoverImg,
				State:    data.State,
				Author:   data.Author,
				Anonymous:data.Anonymous,
			})
			if err != nil {
				utils.JsonInternalServerErrorResponse(c)
				return
			}
			utils.JsonErrorResponse(c, 0, "发布文章成功")
		    return
		}
		//发布
		err = articleService.Article(models.Article{
			ID:       data.ID,
			Title:    data.Title,
			CateID:   data.CateID,
			Content:  data.Content,
			CoverImg: data.CoverImg,
			State:    data.State,
			Author:   data.Author,
			Anonymous:data.Anonymous,
		})
		if err != nil {
			utils.JsonInternalServerErrorResponse(c)
			return
		}
		utils.JsonErrorResponse(c, 0, "发布文章成功")
		return
	}else {
		utils.JsonErrorResponse(c, 2,  "\"state\" must be one of [草稿, 发布]")
		return
	}
}
