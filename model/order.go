package model

import (
	"time"
)

type Order struct {
	ID        OrderID `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Token     string // 口令
	CoinFrom  string // 发送的币
}

// OrderID 订单ID
type OrderID string

func (t *OrderStatus) OrderID() string {
	return "varchar(20)"
}

// OrderStatus 订单状态
type OrderStatus string

func (t *OrderStatus) GormDataType() string {
	return "varchar(16)"
}

const (
	OrderStatusNew        = OrderStatus("new")        // 新订单
	OrderStatusPending    = OrderStatus("pending")    // 资金已收到，正在确认
	OrderStatusProcessing = OrderStatus("processing") // 正在处理
	OrderStatusDone       = OrderStatus("done")       // 已完成
)
