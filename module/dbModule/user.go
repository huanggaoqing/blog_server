package dbModule

type User struct {
	UserId      int    `gorm:"column:user_id;primaryKey;not null;PRECISION:6;AUTO_INCREMENT"`
	UserName    string `gorm:"column:user_name;not null"`
	Password    string `gorm:"column:password;not null"`
	PhoneNumber string `gorm:"column:phone_number;not null"`
	Avatar      string `gorm:"column:avatar"`
	Role        int    `gorm:"column:role;default:1"`
	Base
}

func (u *User) TableName() string {
	return "blog_user"
}
