package db

import "frame/model"

func (s *Store) AddAdminLog(v *model.AdminLog) error {
	return s.db.Create(v).Error
}
