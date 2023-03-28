package dao

import (
	"gokyrie/model"
	"gokyrie/service/dto"
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

func (m *UserDao) GetUserByName(userName string) (model.User, error) {
	var iUser model.User
	err := m.Orm.Model(&model.User{}).Where("name=?", userName).Find(&iUser).Error
	return iUser, err
}

func (m *UserDao) GetUserByNameAndPassword(userName string, password string) model.User {
	var iUser model.User
	m.Orm.Model(&iUser).Where("name=? and password=?", userName, password).Find(&iUser)
	return iUser
}

func (m *UserDao) CheckUserNameExist(stUserName string) bool {
	var nTotal int64
	m.Orm.Model(&model.User{}).Where("name = ?", stUserName).Count(&nTotal)
	return nTotal > 0
}

func (m *UserDao) AddUser(iUserAddDTO *dto.UserAddDTO) error {
	var iUser model.User
	// iUserAddDTO.RoleId
	iUserAddDTO.ConvertToModel(&iUser)
	res := m.Orm.Create(&iUser)
	if res.Error == nil {
		iUserAddDTO.ID = iUser.ID
		iUserAddDTO.Password = ""
	}
	return res.Error
}

func (m *UserDao) GetUserById(id uint) (model.User, error) {
	var iUser model.User
	err := m.Orm.First(&iUser, id).Error
	return iUser, err
}

func (m *UserDao) GetUserList(iUserListDto *dto.UserListDTO) ([]model.User, int64, error) {
	var giUserList []model.User
	var nTotal int64
	Db := m.Orm
	Db = Db.Joins("Dept")
	if iUserListDto.Name != "" {
		Db = Db.Where("name = ?", iUserListDto.Name)
	}
	Db = Db.Scopes(Paginate(iUserListDto.Paginate)).Find(&giUserList).Offset(-1).Limit(-1).Count(&nTotal)
	err := Db.Error
	return giUserList, nTotal, err
}

func (m *UserDao) UpdateUser(iUserUpdateDTO *dto.UserUpdateDTO) error {
	var iUser model.User
	m.Orm.First(&iUser, iUserUpdateDTO.ID)
	iUserUpdateDTO.ConvertToModel(&iUser)
	//需判断role_id，查找role表
	return m.Orm.Save(&iUser).Error
}

func (m *UserDao) DeleteUserById(id uint) error {
	return m.Orm.Delete(&model.User{}, id).Error
}
