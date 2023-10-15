package articleService

import (
	"BBQ/app/models"
	"BBQ/config/database"
	"fmt"
	"gorm.io/gorm"
	"time"
)

func CheckArticleExistByArticle(ID string) error {
    var article models.Article
    result := database.DB.Where("id = ?", ID).First(&article)
    if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
        fmt.Printf("CheckArticleExistByArticle: failed to check article existence for ID %s: %s\n", ID, result.Error)
        return result.Error
    }

    if result.RowsAffected > 0 {
        fmt.Printf("CheckArticleExistByArticle: found article with ID %s: %+v\n", ID, article)
    } else {
        fmt.Printf("CheckArticleExistByArticle: article not found for ID %s\n", ID)
    }

    return nil
}

func Article(article models.Article) error {
    // 处理日期字段
    if article.PubDate.IsZero() {
        // 如果 PubDate 为零值，设置一个默认的日期值
        article.PubDate = time.Now()
    }

    result := database.DB.Create(&article)
    if result.Error != nil {
        fmt.Printf("Article: failed to insert article: %s\n", result.Error)
        return result.Error
    }
    fmt.Printf("Article: inserted article: %+v\n", article)
    return nil
}



func UpdateArticle(article models.Article)error{
	result := database.DB.Save(&article)
	return result.Error
}

func CompareCon(pwd1 string, pwd2 string) bool {
	return pwd1 == pwd2
}

func GetArticleByID(ID string) (*models.Article, error) {
	fmt.Println("1")
    var article models.Article
    result := database.DB.Where("id = ?", ID).First(&article)
    if result.Error != nil {
        fmt.Printf("GetArticleByID: failed to get article for ID %s: %s\n", ID, result.Error)
        return nil, result.Error
    }
    fmt.Printf("GetArticleByID: retrieved article for ID %s: %+v\n", ID, article)
    return &article, nil
}


func DeleteArticle(id string) error {
	result := database.DB.Where("id = ?", id).Delete(&models.Article{})
	return result.Error
}
