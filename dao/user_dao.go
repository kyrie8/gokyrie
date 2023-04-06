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
	tx := m.Orm.Begin()
	if err := tx.Create(&iUser).Error; err != nil {
		tx.Rollback()
		return err
	}
	iUserAddDTO.ID = iUser.ID
	iUserAddDTO.Password = ""
	if len(iUserAddDTO.RoleId) > 0 { // 需要加入事务防止用户创建成功，但没有和角色建立关系
		var roles []model.Role
		for _, v := range iUserAddDTO.RoleId {
			roles = append(roles, model.Role{
				RoleId: uint(v),
			})
		}
		err := tx.Model(&iUser).Association("Roles").Append(&roles)
		if err != nil {
			return err
		}
	}
	return tx.Commit().Error
}

func (m *UserDao) GetUserById(id uint) (model.User, error) {
	var iUser model.User
	err := m.Orm.Joins("Dept").First(&iUser, id).Error
	return iUser, err
}

func (m *UserDao) GetUserList(iUserListDto *dto.UserListDTO) ([]model.User, int64, error) {
	var giUserList []model.User
	var nTotal int64
	Db := m.Orm
	Db = Db.Joins("Dept") //一对多 带出 dept
	if iUserListDto.Name != "" {
		Db = Db.Where("name = ?", iUserListDto.Name)
	}
	if iUserListDto.DeptId > 0 {
		Db = Db.Where(&model.User{DeptId: iUserListDto.DeptId})
	}
	Db = Db.Scopes(Paginate(iUserListDto.Paginate)).Find(&giUserList).Offset(-1).Limit(-1).Count(&nTotal)
	err := Db.Error
	return giUserList, nTotal, err
}

func (m *UserDao) UpdateUser(iUserUpdateDTO *dto.UserUpdateDTO) error {
	var iUser model.User
	var roles []model.Role
	iUser.ID = iUserUpdateDTO.ID
	tx := m.Orm.Begin()
	if len(iUserUpdateDTO.RoleId) == 0 {
		err := tx.Model(&iUser).Association("Roles").Clear()
		if err != nil {
			tx.Rollback()
			return err
		}
	} else {
		for _, v := range iUserUpdateDTO.RoleId {
			roles = append(roles, model.Role{
				RoleId: uint(v),
			})
		}
		err := tx.Model(&iUser).Association("Roles").Replace(&roles)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.First(&iUser, iUserUpdateDTO.ID)
	iUserUpdateDTO.ConvertToModel(&iUser)
	var err error
	if iUserUpdateDTO.DeptId > 0 {
		err = tx.Save(&iUser).Error
	} else {
		err = tx.Omit("dept_id").Updates(&iUser).Error
	}
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (m *UserDao) DeleteUserById(id uint) error {
	return m.Orm.Delete(&model.User{}, id).Error
}
