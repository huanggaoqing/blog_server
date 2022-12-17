package request

type SaveProfileRequest struct {
	UserId  int    `json:"userId" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type GetProfileRequest struct {
	UserId int `form:"userId" binding:"required"`
}

type SetProfileRequest struct {
	ProfileId int    `json:"profileId" binding:"required"`
	Content   string `json:"content" binding:"required"`
}
