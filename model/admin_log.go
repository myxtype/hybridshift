package model

import "time"

// 管理员操作日志
type AdminLog struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	AdminId   uint   `gorm:"index"` // 管理员ID
	Notes     string // 备注
	Ip        string // IP
}
