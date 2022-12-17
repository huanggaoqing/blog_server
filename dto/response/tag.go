package response

import "blog_server/module/dbModule"

type GetTagsResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
}

func ToTagsResponse(tags *dbModule.Tag) *GetTagsResponse {
	return &GetTagsResponse{
		Id:   tags.Id,
		Name: tags.Name,
		Url:  tags.Url,
	}
}
