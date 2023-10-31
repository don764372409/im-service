package service

import (
	utils "IM-Service/configs/err"
	"IM-Service/configs/log"
	"IM-Service/dto"
	"IM-Service/entity"
	"IM-Service/util"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

// Register 用户注册
func (_self *UserService) Register(username, password string) (result *dto.ResultDTO, err error) {
	if username == "" {
		return nil, log.WithError(utils.ERR_USER_REGISTER_USERNAME_NULL)
	}
	if password == "" {
		return nil, log.WithError(utils.ERR_USER_REGISTER_PASSWORD_NULL)
	}
	if len(username) < 6 || len(username) > 20 {
		return nil, log.WithError(utils.ERR_USER_REGISTER_USERNAME_LENGTH)
	}
	if len(password) < 6 || len(password) > 20 {
		return nil, log.WithError(utils.ERR_USER_REGISTER_PASSWORD_LENGTH)
	}

	params := &entity.RegisterUser{
		Username: username,
		Password: password,
	}
	return util.Post("/api/user/register", params)
}
