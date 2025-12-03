package repository

import (
	"rea_games/entity"
)

type GameRepository struct {
	*BaseRepository
}

func NewGameRepository() *GameRepository {
	return &GameRepository{
		BaseRepository: NewBaseRepository(),
	}
}

// GetAllGames returns all games without aggregating genres
func (r *GameRepository) GetAllGames() ([]entity.Game, error) {
	query := `
		SELECT
			g.game_id,
			g.title,
			g.price,
			g.release_date,
			g.game_quantity,
			d.developer_name,
			gd.description,
			g.created_at
		FROM games g
		JOIN developers d ON g.developer_id = d.developer_id
		JOIN games_detail gd ON g.game_detail_id = gd.game_detail_id
		WHERE g.deleted_at IS NULL
		ORDER BY g.created_at DESC
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var games []entity.Game
	for rows.Next() {
		var game entity.Game
		err := rows.Scan(
			&game.GameID,
			&game.Title,
			&game.Price,
			&game.ReleaseDate,
			&game.GameQuantity,
			&game.DeveloperName,
			&game.Description,
			&game.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		games = append(games, game)
	}

	return games, nil
}

// GetGameByID returns a single game without genres
func (r *GameRepository) GetGameByID(gameID int) (*entity.Game, error) {
	query := `
		SELECT
			g.game_id,
			g.title,
			g.price,
			g.release_date,
			g.game_quantity,
			d.developer_name,
			gd.description,
			g.created_at
		FROM games g
		JOIN developers d ON g.developer_id = d.developer_id
		JOIN games_detail gd ON g.game_detail_id = gd.game_detail_id
		WHERE g.game_id = $1
		AND g.deleted_at IS NULL
	`

	var game entity.Game
	err := r.db.QueryRow(query, gameID).Scan(
		&game.GameID,
		&game.Title,
		&game.Price,
		&game.ReleaseDate,
		&game.GameQuantity,
		&game.DeveloperName,
		&game.Description,
		&game.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &game, nil
}

// GetGenresByGameID fetches all genres for a given game
func (r *GameRepository) GetGenresByGameID(gameID int) ([]string, error) {
	query := `
		SELECT ge.genre_name
		FROM genre_game gg
		JOIN genre ge ON gg.genre_id = ge.genre_id
		WHERE gg.game_id = $1
		AND gg.deleted_at IS NULL
		AND ge.deleted_at IS NULL
	`

	rows, err := r.db.Query(query, gameID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var genres []string
	for rows.Next() {
		var genre string
		if err := rows.Scan(&genre); err != nil {
			return nil, err
		}
		genres = append(genres, genre)
	}

	return genres, nil
}

// UpdateGameQuantity remains the same
func (r *GameRepository) UpdateGameQuantity(gameID int, quantity int) error {
	query := `
		UPDATE games
		SET game_quantity = $1,
		    updated_at = NOW()
		WHERE game_id = $2
		AND deleted_at IS NULL
	`

	_, err := r.db.Exec(query, quantity, gameID)
	return err
}
