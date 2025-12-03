package entity

import "time"

type Order struct {
	OrderID   int
	UserID    int
	GameID    int
	OrderDate time.Time
	GameQuantity  int
	TotalPrice     float64
	CreatedAt time.Time
	DeletedAt *time.Time
	UserEmail string
	GameTitle string
}
