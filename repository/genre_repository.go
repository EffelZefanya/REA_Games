package repository

import (
	"rea_games/entity"
)

type GenreRepository struct {
	*BaseRepository
}

func NewGenreRepository() *GenreRepository {
	return &GenreRepository{
		BaseRepository: NewBaseRepository(),
	}
}

func (r *GenreRepository) GetAllGenre() ([]entity.Genre, error) {
	query := `
		SELECT
			genre_id,
			genre_name
		FROM genre
		WHERE deleted_at IS NULL
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var genres []entity.Genre
	for rows.Next() {
		var genre entity.Genre
		err := rows.Scan(
			&genre.Genre_id,
			&genre.Genre_Name,
		)
		if err != nil {
			return nil, err
		}
		genres = append(genres, genre)
	}

	return genres, nil
}
