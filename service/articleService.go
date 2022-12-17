package service

import (
	"blog_server/dao"
	"blog_server/dto/request"
	"blog_server/dto/response"
	"blog_server/module/dbModule"
	"blog_server/tools"
)

type ArticleService struct{}

func (a *ArticleService) Save(params *request.SaveArticleRequest) (string, error) {
	// 存储文章数据
	articleData := &dbModule.Article{
		ArticleId:      tools.GetXid(),
		UserId:         params.UserId,
		ArticleContent: params.ArticleContent,
		ArticleTitle:   params.ArticleTitle,
		GroupId:        params.GroupId,
	}
	articleDao := dao.NewArticleDao()
	articleId, err := articleDao.Save(articleData)
	if err != nil {
		return "", err
	}
	// 存储文章与文章关联数据
	err = a.saveTagAssociate(articleId, params.Tag)
	if err != nil {
		return "", err
	}
	return articleId, nil
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

func (a *ArticleService) GetList(params *request.GetArticleListRequest) (response.GetArticleListResponseByPaging, error) {
	articleDao := dao.NewArticleDao()
	articleList, err := articleDao.GetList(params)
	if err != nil {
		return nil, err
	}
	data := make([]*response.GetArticleListResponse, 0)
	for _, v := range articleList {
		data = append(data, response.ToArticleListResponse(v))
	}
	total, err := articleDao.GetListCount(params.UserId)
	return response.NewPaging[[]*response.GetArticleListResponse](params.Page, params.Size, data, total), nil
}

func (a *ArticleService) GetDetailByArticleId(articleId string) (*response.GetArticleDetailResponse, error) {
	articleDao := dao.NewArticleDao()
	articleDetail, err := articleDao.GetDetail(articleId)
	if err != nil {
		return nil, err
	}
	tagService := NewTagService()
	tags, err := tagService.getTagsByArticleId(articleId)
	if err != nil {
		return nil, err
	}
	data := response.ToGetArticleDetailResponse(articleDetail, tags)
	return data, nil
}

func NewArticleService() *ArticleService {
	return &ArticleService{}
}
