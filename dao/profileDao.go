package dao

import (
	"blog_server/db"
	"blog_server/dto"
	"blog_server/module/dbModule"
	"blog_server/resp"
	"gorm.io/gorm"
)

type ProfileDao struct {
	db *gorm.DB
}

func (p *ProfileDao) Save(params *dbModule.Profile) (int, error) {
	err := p.db.Create(params).Error
	if err != nil {
		return -1, err
	}
	return params.ProfileId, err
}

func (p *ProfileDao) Get(params *dto.GetProfileRequest) (*dbModule.Profile, error) {
	profile := &dbModule.Profile{}
	err := p.db.Where("user_id = ?", params.UserId).First(profile).Error
	if err == gorm.ErrRecordNotFound {
		return nil, resp.NOT_PROFILE
	}
	if err != nil {
		return nil, err
	}
	return profile, err
}

func (p *ProfileDao) Set(params *dto.SetProfileRequest) (int, error) {
	profile := &dbModule.Profile{}
	err := p.db.Model(profile).Where("profile_id = ?", params.ProfileId).Update("content", params.Content).Error
	if err != nil {
		return -1, err
	}
	return profile.ProfileId, nil
}

func NewProfileDao() *ProfileDao {
	return &ProfileDao{
		db: db.GetDataBase(),
	}
}
