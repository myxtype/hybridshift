package model

import "gorm.io/gorm"

type OrderShift struct {
	gorm.Model
	OrderID OrderID          `gorm:"index"`
	Status  OrderShiftStatus `gorm:"index"`
}

type OrderShiftStatus string

func (t *OrderShiftStatus) GormDataType() string {
	return "varchar(16)"
}

const (
	OrderShiftStatusNew        = OrderShiftStatus("new")        // 等待创建订单
	OrderShiftStatusPending    = OrderShiftStatus("pending")    // 订单已创建等待发款
	OrderShiftStatusConfirming = OrderShiftStatus("confirming") // 已发款等待确认
	OrderShiftStatusProcessing = OrderShiftStatus("processing") // 确认成功，等待处理
)
