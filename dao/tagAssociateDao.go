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

func (t *TagAssociateDao) GetTagByArticle(articleId string) ([]*dbModule.Tag, error) {
	data := make([]*dbModule.Tag, 0)
	err := t.db.Raw(
		"SELECT id, `name`, url FROM blog_tag WHERE id IN (SELECT tag_id FROM blog_tag_associate WHERE article_id = ?)",
		articleId,
	).Scan(&data).Error
	if err != nil {
		if err == gorm.ErrNotImplemented {
			return data, nil
		}
		return nil, err
	}
	return data, nil

}

func NewTagAssociateDao() *TagAssociateDao {
	return &TagAssociateDao{
		db: db.GetDataBase(),
	}
}
