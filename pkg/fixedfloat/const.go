package fixedfloat

type Direction string

const (
	DirectionTo   = Direction("to")
	DirectionFrom = Direction("from")
)

type OrderType string

const (
	OrderTypeFloat = OrderType("float")
	OrderTypeFixed = OrderType("fixed")
)

type OrderStatus string

const (
	OrderStatusNew       = OrderStatus("NEW")
	OrderStatusPending   = OrderStatus("PENDING")
	OrderStatusExchange  = OrderStatus("EXCHANGE")
	OrderStatusWithdraw  = OrderStatus("WITHDRAW")
	OrderStatusDone      = OrderStatus("DONE")
	OrderStatusExpired   = OrderStatus("EXPIRED")
	OrderStatusEmergency = OrderStatus("EMERGENCY")
)
