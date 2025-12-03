package entity

import "time"

type GameDisplay struct {
	GameID        int
	Title         string
	Price         float64
	ReleaseDate   time.Time
	GameQuantity  int
	DeveloperName string
	Gamedetail_id int
	Genres        []string
	Description   string
	CreatedAt     time.Time
}

type Game struct {
	GameID        int
	Title         string
	Price         float64
	ReleaseDate   time.Time
	DeveloperName int
	Gamedetail_id int
	Genre         []int
	Description   string
	CreatedAt     time.Time
	GameQuantity  int
}
