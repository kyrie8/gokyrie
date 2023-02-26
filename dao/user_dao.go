package dao

import (
	"gokyrie/model"
)

type UserDao struct {
	BaseDao
}

var userDao *UserDao

func NewUserDao() *UserDao {
	if userDao == nil {
		userDao = &UserDao{
			BaseDao: NewBaseDao(),
		}
	}
	return userDao
}

func (m *UserDao) GetUserByNameAndPassword(userName string, password string) model.User {
	var iUser model.User
	m.Orm.Model(&iUser).Where("name=? and password=?", userName, password).Find(&iUser)
	return iUser
}
