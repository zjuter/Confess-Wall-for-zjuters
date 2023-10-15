package articleService

import (
	"BBQ/app/models"
	"BBQ/config/database"

	"gorm.io/gorm"
)

func CheckArticleExistByArticle(ID string) error {
	result := database.DB.Where("id= ? ", ID).First(&models.Article{})
	return result.Error
}

func Article(article models.Article) error {
	result := database.DB.Create(&article)
	return result.Error
}

func UpdateArticle(article models.Article) error {
	result := database.DB.Save(&article)
	return result.Error
}

func CompareCon(pwd1 string, pwd2 string) bool {
	return pwd1 == pwd2
}

func GetArticleByID(ID string) (*models.Article, error) {
	var article models.Article
	result := database.DB.Where("id = ?", ID).First(&article)
	if result.Error != nil {
		return nil, result.Error
	}
	return &article, nil
}

func DeleteArticle(id string) error {
	result := database.DB.Where("id = ?", id).Delete(&models.Article{})
	return result.Error
}

func GetAriticalContext(ID string) ([]models.Article, error) {
	var articleList []models.Article
	result := database.DB.Where("id = ?", ID).Find(&articleList)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return nil, result.Error
	}
	return articleList, nil
}

func GetAriticalList(Cateid string, State string) ([]models.Article, int, error) {
	var List []models.Article
	var result *gorm.DB
	if Cateid != "" && State != "" {
		result = database.DB.Where("cate_id = ? AND state = ?", Cateid, State).Find(&List)
	} else if Cateid != "" {
		result = database.DB.Where("cate_id = ?", Cateid).Find(&List)
	} else if State != "" {
		result = database.DB.Where("state = ?", State).Find(&List)
	} else {
		result = database.DB.Find(&List)
	}

	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return nil, 0, result.Error
	}
	return List, int(result.RowsAffected), nil
}
