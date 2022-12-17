package request

type PagingRequest struct {
	Page int `json:"page" form:"page" binding:"required"`
	Size int `json:"size" form:"size" binding:"required"`
}
