package model

import (
	"frame/pkg/sql/sqltypes"
	"time"
)

// 管理员
type AdminUser struct {
	ID           uint `gorm:"primaryKey"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Username     string             `gorm:"uniqueIndex"`
	Password     *sqltypes.Password // 登录密码
	RoleId       uint               // 角色
	Name         string             // 昵称
	Avatar       string             // 头像
	Disabled     bool               // 是否禁用
	LoginVersion int64              // 登录版本
}
