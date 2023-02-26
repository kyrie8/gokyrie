package service

import (
	"errors"
	"gokyrie/dao"
	"gokyrie/model"
	"gokyrie/service/dto"
)

type UserService struct {
	BaseService
	Dao *dao.UserDao
}

var userService *UserService

func NewUserService() *UserService {
	if userService == nil {
		userService = &UserService{
			Dao: dao.NewUserDao(),
		}
	}
	return userService
}

func (m *UserService) Login(iUserDTO dto.UserLoginDTO) (model.User, error) {
	var errResult error
	iUser := m.Dao.GetUserByNameAndPassword(iUserDTO.Name, iUserDTO.Password)
	if iUser.ID == 0 {
		errResult = errors.New("账号或者密码错误")
	}
	return iUser, errResult
}
