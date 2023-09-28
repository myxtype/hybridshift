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

type EmergencyStatus string

const (
	EmergencyStatusExpired = EmergencyStatus("EXPIRED")
	EmergencyStatusLess    = EmergencyStatus("LESS")
	EmergencyStatusMore    = EmergencyStatus("MORE")
	EmergencyStatusLimit   = EmergencyStatus("LIMIT")
)

type EmergencyChoice string

const (
	EmergencyChoiceNone     = EmergencyChoice("NONE")
	EmergencyChoiceExchange = EmergencyChoice("EXCHANGE")
	EmergencyChoiceRefund   = EmergencyChoice("REFUND")
)
