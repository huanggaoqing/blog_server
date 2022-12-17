package response

import "blog_server/module/dbModule"

type GetArticleListResponseByPaging *PagingResponse[[]*GetArticleListResponse]

type GetArticleListResponse struct {
	UserName     string `json:"userName"`
	ArticleId    string `json:"articleId""`
	ArticleTitle string `json:"articleTitle"`
	IsStick      int    `json:"isStick"`
	*Base
}

func ToArticleListResponse(article *dbModule.ArticleItem) *GetArticleListResponse {
	return &GetArticleListResponse{
		UserName:     article.UserName,
		ArticleId:    article.ArticleId,
		ArticleTitle: article.ArticleTitle,
		IsStick:      article.IsStick,
		Base: &Base{
			CreateTime: article.CreateTime,
			UpdateTime: article.UpdateTime,
		},
	}
}

type GetArticleDetailResponse struct {
	ArticleContent string             `json:"articleContent"`
	GroupId        int                `json:"groupId"`
	Tag            []*GetTagsResponse `json:"tag"`
	*GetArticleListResponse
}

func ToGetArticleDetailResponse(article *dbModule.ArticleItem, tag []*GetTagsResponse) *GetArticleDetailResponse {
	return &GetArticleDetailResponse{
		ArticleContent: article.ArticleContent,
		GroupId:        article.GroupId,
		Tag:            tag,
		GetArticleListResponse: &GetArticleListResponse{
			UserName:     article.UserName,
			ArticleId:    article.ArticleId,
			ArticleTitle: article.ArticleTitle,
			IsStick:      article.IsStick,
			Base: &Base{
				CreateTime: article.CreateTime,
				UpdateTime: article.UpdateTime,
			},
		},
	}
}
