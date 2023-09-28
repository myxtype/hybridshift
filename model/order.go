package model

import (
	"github.com/thanhpk/randstr"
	"time"
)

type Order struct {
	ID        OrderID `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Token     string      // 口令
	Status    OrderStatus // 状态
}

// OrderID 订单ID
type OrderID string

func (t *OrderID) GormDataType() string {
	return "varchar(20)"
}

func NewOrderID() string {
	return randstr.Hex(20)
}

// OrderStatus 订单状态
type OrderStatus string

func (t *OrderStatus) GormDataType() string {
	return "varchar(16)"
}

const (
	OrderStatusNew        = OrderStatus("new")        // 新订单
	OrderStatusConfirming = OrderStatus("confirming") // 资金已收到，正在确认
	OrderStatusPending    = OrderStatus("pending")    // 确认完毕，等待处理
	OrderStatusProcessing = OrderStatus("processing") // 已确认，正在处理
	OrderStatusDone       = OrderStatus("done")       // 已完成
)
