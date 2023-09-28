package model

import "gorm.io/gorm"

type Coin struct {
	gorm.Model
	Name    string // 名称
	Network string // 网络
	Logo    string // 图标
}
