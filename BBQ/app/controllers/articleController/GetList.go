package articleController

import (
	"BBQ/app/models"
	"BBQ/app/services/articleService"
	"BBQ/app/utils"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GetListData struct {
	ID       uint   `form:"id"`
	Pagenum  int    `form:"pagenum" binding:"required"`
	Pagesize string `form:"pagesize" binding:"required"`
	CateId   string `form:"cate_id"`
	State    string `form:"state"`
}

// 获取
func GetList(c *gin.Context) {
	var data GetListData
	err := c.ShouldBindQuery(&data)
	fmt.Println(err)
	if err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}

	var List []models.Article
	List, count, err := articleService.GetAriticalList(data.CateId, data.State)
	fmt.Println(err)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.JsonErrorResponse(c, 1, "没有查到数据！")
			return
		} else {
			utils.JsonInternalServerErrorResponse(c)
			return
		}
	}
	fmt.Println(count)
	fmt.Println(List)

	pageSize, err := strconv.Atoi(data.Pagesize)
	if err != nil {
		utils.JsonErrorResponse(c, 2, "数据条数有误！(可能没有输入整数)")
		return
	}

	if data.Pagenum > count/pageSize {
		utils.JsonErrorResponse(c, 3, "请求页码有误！(可能是大于能显示数)")
		return
	}

	c.JSON(200, gin.H{
		"status":  0,
		"message": "获取文章列表成功！",
		"data":    List[pageSize*(data.Pagenum-1) : pageSize*data.Pagenum],
	})
}

