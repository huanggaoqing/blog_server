package dbModule

type Article struct {
	ArticleId      string `gorm:"column:article_id;primaryKey;not null;"`
	UserId         int    `gorm:"column:user_id;not null"`
	ArticleContent string `gorm:"column:article_content;not null;type:text;"`
	ArticleTitle   string `gorm:"column:article_title;not null;"`
	IsStick        int    `gorm:"column:is_stick;not nll;default:0"`
	GroupId        int    `gorm:"column:group_id;"`
	*Base
}

type ArticleItem struct {
	UserName string
	*Article
}

func (a *Article) TableName() string {
	return "blog_article"
}
