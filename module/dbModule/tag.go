package dbModule

type Tag struct {
	Id   int    `gorm:"column:id;primaryKey;not null;PRECISION:6;AUTO_INCREMENT"`
	Name string `gorm:"column:name;not null"`
	Url  string `gorm:"column:url;not null"`
	Base
}

func (t *Tag) TableName() string {
	return "blog_tag"
}
