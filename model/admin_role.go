package model

import (
	"frame/pkg/sql/sqltypes"
	"time"
)

// 管理员角色
type AdminRole struct {
	ID          uint `gorm:"primaryKey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string               // 角色名称
	Permissions sqltypes.StringArray // 权限列表
	Disabled    bool                 // 是否禁用
}
