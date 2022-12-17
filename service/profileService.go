package service

import (
	"blog_server/dao"
	"blog_server/dto/request"
	"blog_server/module/dbModule"
)

type ProfileService struct{}

func (p *ProfileService) Save(params *request.SaveProfileRequest) (int, error) {
	profileDao := dao.NewProfileDao()
	profileData := &dbModule.Profile{
		UserId:  params.UserId,
		Content: params.Content,
	}
	profileId, err := profileDao.Save(profileData)
	if err != nil {
		return -1, err
	}
	return profileId, nil
}

func (p *ProfileService) Get(params *request.GetProfileRequest) (*dbModule.Profile, error) {
	profileDao := dao.NewProfileDao()
	profile, err := profileDao.Get(params)
	if err != nil {
		return nil, err
	}
	return profile, nil
}

func (p *ProfileService) Set(params *request.SetProfileRequest) (int, error) {
	profileDao := dao.NewProfileDao()
	profile, err := profileDao.Set(params)
	if err != nil {
		return -1, err
	}
	return profile, nil
}

func NewProfileService() *ProfileService {
	return &ProfileService{}
}
