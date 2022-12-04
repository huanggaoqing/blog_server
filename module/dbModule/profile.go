package dbModule

type Profile struct {
	ProfileId int    `json:"profileId" gorm:"column:profile_id;primaryKey;not null;PRECISION:6;AUTO_INCREMENT"`
	UserId    int    `json:"userId" gorm:"column:user_id;not null"`
	Content   string `json:"content" gorm:"column:content;type:text;not null"`
	Base
}

func (p *Profile) TableName() string {
	return "blog_profile"
}
