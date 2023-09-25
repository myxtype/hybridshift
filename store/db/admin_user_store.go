package db

import (
	"frame/model"
	"gorm.io/gorm"
)

// 获取管理员
func (s *Store) GetAdminUserById(id uint) (*model.AdminUser, error) {
	var admin model.AdminUser
	err := s.db.First(&admin, id).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &admin, nil
}

// 通过用户名获取管理员
func (s *Store) GetAdminUserByUsername(username string) (*model.AdminUser, error) {
	var admin model.AdminUser
	err := s.db.First(&admin, `"username" = ?`, username).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &admin, nil
}

// 更新管理员登录版本
func (s *Store) UpdateAdminUserLoginVersion(admin *model.AdminUser) error {
	return s.db.Model(admin).UpdateColumn("login_version", admin.LoginVersion).Error
}

// 保存用户
func (s *Store) SaveAdminUser(admin *model.AdminUser) error {
	return s.db.Save(admin).Error
}

// 查询管理员
func (s *Store) QueryAdminUsers(id, roleId int64, username string, page, limit int) ([]*model.AdminUser, int64, error) {
	db := s.db

	if id > 0 {
		db = db.Where(`"id" = ?`, id)
	}
	if roleId > 0 {
		db = db.Where(`"role_id" = ?`, roleId)
	}
	if username != "" {
		db = db.Where(`"username" = ?`, username)
	}

	var data []*model.AdminUser
	var count int64
	err := db.Offset((page - 1) * limit).Limit(limit).Order(`"id" DESC`).Find(&data).Offset(-1).Limit(-1).Count(&count).Error

	return data, count, err
}
