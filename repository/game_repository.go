package repository

import (
	"fmt"
	"rea_games/entity"
	"time"
)

type GameRepository struct {
	*BaseRepository
}

func NewGameRepository() *GameRepository {
	return &GameRepository{
		BaseRepository: NewBaseRepository(),
	}
}

func (r *GameRepository) GetAllGamesDisplay() ([]entity.GameDisplay, error) {
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
		ORDER BY g.game_id ASC
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var games []entity.GameDisplay
	for rows.Next() {
		var game entity.GameDisplay
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

func (r *GameRepository) GetGameDisplayByID(gameID int) (*entity.GameDisplay, error) {
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

	var game entity.GameDisplay
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

func (r *GameRepository) GetGameByID(gameID int) (*entity.Game, error) {
	query := `
		SELECT
			game_id,
			title,
			price,
			release_date,
			game_quantity,
			developer_id,
			created_at
		FROM games 
		WHERE game_id = $1
		AND deleted_at IS NULL
	`

	var game entity.Game
	err := r.db.QueryRow(query, gameID).Scan(
		&game.GameID,
		&game.Title,
		&game.Price,
		&game.ReleaseDate,
		&game.GameQuantity,
		&game.DeveloperName,
		&game.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &game, nil
}

func (r *GameRepository) GetDeveloperByID(DevID int) ([]string, error) {
	var developers []string
	query := `
		SELECT developer_name
		FROM developers
		WHERE developer_id = $1
		AND deleted_at IS NULL
	`
	rows, err := r.db.Query(query, DevID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var developer string
		if err := rows.Scan(&developer); err != nil {
			return nil, err
		}
		developers = append(developers, developer)
	}

	return developers, nil
}

func (r *GameRepository) GetGenreByID(genreID []int) ([]string, error) {
	var genres []string
	var temp string
	query := `
		SELECT genre_name
		FROM genre
		WHERE genre_id = $1
		AND deleted_at IS NULL
	`
	for _, genre := range genreID {
		err := r.db.QueryRow(query, genre).Scan(
			&temp,
		)
		if err != nil {
			return nil, err
		}
		genres = append(genres, temp)
	}

	return genres, nil
}

func (r *GameRepository) CreateGame(game *entity.Game) error {

	query := `
		INSERT INTO games (
		developer_id,
		title,
		price,
		release_date,
		game_quantity,
		game_detail_id
	)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING game_id, created_at
	`
	err := r.db.QueryRow(
		query,
		game.DeveloperName,
		game.Title,
		game.Price,
		game.ReleaseDate,
		game.GameQuantity,
		1,
	).Scan(
		&game.GameID,
		&game.CreatedAt,
	)
	if err != nil {
		fmt.Println("error Inserting into games")
		return err
	}

	descriptionquery := `
		INSERT INTO games_detail
	(
		description, 
		game_detail_id
	)
		VALUES ($1,$2)
	`
	_, err = r.db.Exec(
		descriptionquery,
		game.Description,
		game.GameID,
	)
	if err != nil {
		fmt.Println("error Inserting into description")
		return err
	}

	updateGameDetailIDquery := `
		UPDATE games
		SET game_detail_id = $1
		WHERE game_id = $2
	`
	_, err = r.db.Exec(updateGameDetailIDquery, game.GameID, game.GameID)
	if err != nil {
		fmt.Println("error GAME DECSRIPTION in database")
		return err
	}

	insertGenreGamequery := `
		INSERT INTO genre_game
		(
			game_id,
			genre_id
		)
		VALUES ($1, $2)
	`
	for _, genre := range game.Genre {
		_, err = r.db.Exec(
			insertGenreGamequery,
			game.GameID,
			genre,
		)
		if err != nil {
			fmt.Println("error GENRE in database")
			return err
		}
	}

	return err

}

func (r *GameRepository) UpdateGame(genre []int, release_date time.Time, developer_id int, title string, price float64, game_quantity int, game_id int) error {
	query := `
		UPDATE games
		SET release_date = $1,
			developer_id = $2,
			title = $3,
			price = $4,
			game_quantity = $5,
			updated_at = NOW()
		WHERE game_id = $6
		AND deleted_at IS NULL
	`
	_, err := r.db.Exec(query, release_date, developer_id, title, price, game_quantity, game_id)
	if err != nil {
		fmt.Println("Error Updating games")
		return err
	}
	querydeletegenre := `
		UPDATE genre_game
			SET deleted_at = NOW()
		WHERE game_id = $1
	`
	_, err = r.db.Exec(querydeletegenre, game_id)
	if err != nil {
		fmt.Println("Error Updating Genre game")
		return err
	}

	queryinsert := `
			INSERT INTO genre_game
		(
			game_id,
			genre_id
		)
		VALUES ($1, $2)
	`
	for _, genre := range genre {
		_, err = r.db.Exec(
			queryinsert,
			game_id,
			genre,
		)
		if err != nil {
			fmt.Println("error GENRE in database")
			return err
		}
	}

	return nil
}

func (r *GameRepository) UpdateDescription(description string, game_id int) error {
	query := `
		UPDATE games_detail
		SET description = $1,
			updated_at = NOW()
		WHERE game_detail_id = $2
		AND deleted_at IS NULL
	`
	_, err := r.db.Exec(query, description, game_id)
	return err
}

func (r *GameRepository) DeleteGame(game_id int) error {
	query := `
		UPDATE games
		SET deleted_at = NOW()
		WHERE game_id = $1
	`
	_, err := r.db.Exec(query, game_id)
	if err != nil {
		fmt.Println("Error Deleting Game")
		return err
	}
	querydeletegenre := `
		UPDATE genre_game
		SET deleted_at = NOW()
		WHERE game_id = $1
	`
	_, err = r.db.Exec(querydeletegenre, game_id)
	if err != nil {
		fmt.Println("Error Deleting Genre_game")
		return err
	}
	querydeletedetail := `
		UPDATE games_detail
		SET deleted_at = NOW()
		WHERE game_detail_id = $1
	`
	_, err = r.db.Exec(querydeletedetail, game_id)
	if err != nil {
		fmt.Println("Error Deleting Games_Detail")
		return err
	}
	return nil
}
