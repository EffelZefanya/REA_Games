package entity

import "time"

type Game struct {
	GameID        int
	Title         string
	Price         float64
	ReleaseDate   time.Time
	GameQuantity int
	DeveloperName string
	Genres        []string
	Description   string
	CreatedAt     time.Time
}
