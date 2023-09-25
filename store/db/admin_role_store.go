package db

import (
	"frame/model"
	"gorm.io/gorm"
)

// 查询管理员
func (s *Store) GetAdminRoleById(id uint) (*model.AdminRole, error) {
	var role model.AdminRole
	err := s.db.First(&role, id).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &role, err
}

// 查询
func (s *Store) QueryAdminRoles(page, limit int) ([]*model.AdminRole, int64, error) {
	db := s.db

	var data []*model.AdminRole
	var count int64
	err := db.Offset((page - 1) * limit).Limit(limit).Find(&data).Offset(-1).Limit(-1).Count(&count).Error

	return data, count, err
}

func (s *Store) SaveAdminRole(v *model.AdminRole) error {
	return s.db.Save(v).Error
}
