package service

import (
	"blog_server/dao"
	"blog_server/dto/response"
)

type TagService struct{}

func (t *TagService) getTagsByArticleId(articleId string) ([]*response.GetTagsResponse, error) {
	tagDao := dao.NewTagAssociateDao()
	tags, err := tagDao.GetTagByArticle(articleId)
	if err != nil {
		return nil, err
	}
	data := make([]*response.GetTagsResponse, 0)
	for _, v := range tags {
		data = append(data, response.ToTagsResponse(v))
	}
	return data, nil
}

func NewTagService() *TagService {
	return &TagService{}
}
