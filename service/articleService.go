package service

import (
	"blog_server/dao"
	"blog_server/dto"
	"blog_server/module/dbModule"
)

type ArticleService struct{}

func (a *ArticleService) Save(params *dto.SaveArticleRequest) (string, error) {
	// 存储文章数据

	// 存储文章与文章关联数据
	err := a.saveTagAssociate("", params.Tag)
	if err != nil {
		return "", err
	}
	return "", nil
}

func (a *ArticleService) saveTagAssociate(articleId string, tagIds *[]int) error {
	tagDao := dao.NewTagAssociateDao()
	data := make([]*dbModule.TagAssociate, 0)
	for _, v := range *tagIds {
		data = append(data, &dbModule.TagAssociate{
			TagId:     v,
			ArticleId: articleId,
		})
	}
	err := tagDao.Save(data...)
	if err != nil {
		return err
	}
	return nil
}

func NewArticleService() *ArticleService {
	return &ArticleService{}
}
