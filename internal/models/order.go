package models

type Side int

const (
	Buy  Side = 1
	Sell Side = 0
)

func (s Side) String() string {
	switch s {
	case Buy:
		return "buy"
	case Sell:
		return "sell"
	default:
		return "buy"
	}
}

type Order struct {
	UserID  int64
	Amount  int64
	Price   int64
	Side    Side
	OrderID int64
}
