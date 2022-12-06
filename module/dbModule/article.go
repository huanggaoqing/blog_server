package dbModule

type Article struct {
	ArticleId      int    `gorm:"column:article_id;primaryKey;not null;PRECISION:6;AUTO_INCREMENT"`
	UserId         int    `gorm:"column:user_id;not null"`
	ArticleContent string `gorm:"column:article_content;not null;type:text;"`
	ArticleTitle   string `gorm:"column:article_title;not null;"`
	IsStick        int    `gorm:"column:is_stick;not nll;default:0"`
	GroupId        int    `gorm:"column:group_id;"`
	*Base
}

func (a *Article) TableName() string {
	return "blog_article"
}
