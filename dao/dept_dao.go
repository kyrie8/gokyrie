package dao

import (
	"gokyrie/model"
	"gokyrie/service/dto"
)

type DeptDao struct {
	BaseDao
}

var deptDao *DeptDao

func NewDeptDao() *DeptDao {
	if deptDao == nil {
		deptDao = &DeptDao{
			BaseDao: NewBaseDao(),
		}
	}
	return deptDao
}

func (m *DeptDao) AddDept(iDeptDto *dto.DeptAddDTO) error {
	var iDept model.Dept
	iDeptDto.ConvertToModel(&iDept)
	return m.Orm.Create(&iDept).Error
}

func (m *DeptDao) CheckDeptNameExist(name string) bool {
	var noTotal int64
	m.Orm.Model(&model.Dept{}).Where("dept_name = ?", name).Count(&noTotal)
	return noTotal > 0
}

func (m *DeptDao) UpdateDept(iDeptDto *dto.DeptUpdateDTO) error {
	var iDept *model.Dept
	return m.Orm.Model(iDept).Where("dept_id = ?", iDeptDto.DeptId).Select("*").Omit("CreatedAt").Updates(iDeptDto).Error
}

func (m *DeptDao) GetDeptByName(name string) (model.Dept, error) {
	var dept model.Dept
	err := m.Orm.Model(&model.Dept{}).Where("dept_name =?", name).Find(&dept).Error
	return dept, err
}

func (m *DeptDao) GetDeptList(dto *dto.DeptListDTO) ([]model.Dept, int64, error) {
	var nTotal int64
	var iDept []model.Dept
	Db := m.Orm
	if dto.Name != "" {
		Db = Db.Where("dept_name = ?", dto.Name)
	}
	err := Db.Scopes(Paginate(dto.Paginate)).Find(&iDept).Offset(-1).Limit(-1).Count(&nTotal).Error
	return iDept, nTotal, err
}

func (m *DeptDao) GetUserByDept(dto *dto.DeptListDTO) (model.Dept, int64, error) {
	var nTotal int64
	var iDept model.Dept
	err := m.Orm.Preload("Users").Where("dept_name = ?", dto.Name).Scopes(Paginate(dto.Paginate)).Find(&iDept).Offset(-1).Limit(-1).Count(&nTotal).Error
	return iDept, nTotal, err
}

func (m *DeptDao) DeleteDeptById(id uint) error {
	err := m.Orm.Model(&model.User{}).Where("dept_id =?", id).Updates(model.User{DeptId: 0}).Error
	if err != nil {
		return err
	}
	return m.Orm.Where("dept_id =?", id).Delete(&model.Dept{}).Error
}
