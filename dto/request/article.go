package request

type SaveArticleRequest struct {
	UserId         int    `json:"userId" binding:"required"`
	ArticleContent string `json:"articleContent" binding:"required"`
	ArticleTitle   string `json:"articleTitle" binding:"required"`
	GroupId        int    `json:"groupId" binding:"required"`
	Tag            *[]int `json:"tag" binding:"required"`
}

type GetArticleListRequest struct {
	UserId int `form:"userId" binding:"required"`
	*PagingRequest
}

type GetArticleDetailRequest struct {
	ArticleId string `form:"articleId" binding:"required"`
}
