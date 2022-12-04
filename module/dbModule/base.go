package dbModule

import "time"

type Base struct {
	CreateTime *time.Time `json:"createTime" gorm:"column:create_time;autoCreateTime;not null"`
	UpdateTime *time.Time `json:"updateTime" gorm:"column:update_time;autoUpdateTime;not null"`
}
