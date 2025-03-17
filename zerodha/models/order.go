package models

type Order struct {
	OrderId         string
	OrderType       OrderType
	Quantity        int
	TotalPrice      float64
	OrderStatus     OrderStatus
	Exchange        Exchange
	TransactionType TransactionType
}

type TransactionType int32

const (
	BUY TransactionType = iota
	SELL
)

type OrderType int32

const (
	MARKET OrderType = iota
	LIMIT
)

type Exchange int32

const (
	NSE Exchange = iota
	BSE
)

type OrderStatus int32

const (
	PENDING OrderStatus = iota
	FAILED
	SUCCESSFUL
)
