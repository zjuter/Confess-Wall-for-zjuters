package articleService

import (
	"BBQ/app/models"
	"BBQ/config/database"
	"fmt"
	"gorm.io/gorm"
)

// import (
// 	"learngo/CONFESS/app/models"
// 	"learngo/CONFESS/config/database"

// 	"gorm.io/gorm"
// )

func GetAriticalContext(ID string) (*models.Article, error) {
    var article models.Article
    result := database.DB.Where("id=?", ID).First(&article)
    if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
        fmt.Printf("GetAriticalContext: failed to get article for ID %s: %s\n", ID, result.Error)
        return nil, result.Error
    }
    fmt.Printf("GetAriticalContext: retrieved article for ID %s: %+v\n", ID, article)
    return &article, nil
}

func GetAriticalList(Pagenum int, Pagesize string) ([]models.Article, error) {
	var List []models.Article
	result := database.DB.Where("Pagenum = ?", Pagenum).Find(&List)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return nil, result.Error
	}
	return List, nil
}

// func GetAriticList(ID uint) ([]models.Articles, error) {
// 	result := database.DB.Where("id = ?", ID).First(&models.Articles{})
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}
// 	var articleList []models.Articles
// 	result = database.DB.Where("id = ?", ID).Find(&articleList)
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}
// 	return articleList, nil
// }
