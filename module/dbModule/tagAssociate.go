package dbModule

type TagAssociate struct {
	ArticleId string `gorm:"column:article_id;not null;index:article_id"`
	TagId     int    `gorm:"column:tag_id;not null"`
}

func (t *TagAssociate) TableName() string {
	return "blog_tag_associate"
}
