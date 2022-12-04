package dbModule

type Role struct {
	RoleId   int    `gorm:"column:role_id;primaryKey;not null;PRECISION:6;AUTO_INCREMENT"`
	RoleName string `gorm:"column:role_name;not null"`
	Base
}

func (r *Role) TableName() string {
	return "blog_role"
}
