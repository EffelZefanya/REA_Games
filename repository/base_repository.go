package repository

import (
	"database/sql"
	"rea_games/config"
)

type BaseRepository struct {
	db *sql.DB
}

func NewBaseRepository() *BaseRepository {
	return &BaseRepository{
		db: config.ConnectDB(),
	}
}
