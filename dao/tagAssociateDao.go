package dao

import (
	"blog_server/db"
	"blog_server/module/dbModule"
	"gorm.io/gorm"
)

type TagAssociateDao struct {
	db *gorm.DB
}

func (t *TagAssociateDao) Save(data ...*dbModule.TagAssociate) error {
	err := t.db.Create(data).Error
	if err != nil {
		return err
	}
	return nil
}

func NewTagAssociateDao() *TagAssociateDao {
	return &TagAssociateDao{
		db: db.GetDataBase(),
	}
}
