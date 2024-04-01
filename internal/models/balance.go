package models

type BalanceChange struct {
	UserID   int64  // UserID
	Value    int64  // Amount to add or subtract from the user's balance
	Currency string // Currency
}
