package dao

import (
	"blog_server/constant"
	"blog_server/db"
	"blog_server/dto"
	"blog_server/module/dbModule"
	"blog_server/resp"
	"gorm.io/gorm"
)

type UserDao struct {
	db *gorm.DB
}

// FindUserByPhone 通过手机号查找用户
func (u *UserDao) FindUserByPhone(phoneNumber string) (*dbModule.User, error) {
	user := &dbModule.User{}
	err := u.db.Model(user).Where("phone_number = ?", phoneNumber).First(user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

func (u *UserDao) FindUserByPassword(params *dto.UserLoginRequest) (*dbModule.User, error) {
	user := &dbModule.User{}
	var err error
	if params.Type == constant.ADMIN {
		err = u.db.Model(user).Where("role != ? phone_number = ? AND password = ?", 3, params.Phone, params.Password).First(user).Error
	} else {
		err = u.db.Model(user).Where("phone_number = ? AND password = ?", params.Phone, params.Password).First(user).Error
	}
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, resp.USER_NOT_EXISTS
		}
		return nil, err
	}
	return user, nil
}

func (u *UserDao) Create(params *dto.UserRegisterRequest) (int, error) {
	user := &dbModule.User{
		UserName:    params.UserName,
		Password:    params.Password,
		PhoneNumber: params.Phone,
	}
	err := u.db.Create(&user).Error
	if err != nil {
		return -1, err
	}
	return user.UserId, nil
}

func NewUserDao() *UserDao {
	return &UserDao{
		db: db.GetDataBase(),
	}
}
