package repository

import (
	"rea_games/entity"
)

type DeveloperRepository struct {
	*BaseRepository
}

func NewDeveloperRepository() *DeveloperRepository {
	return &DeveloperRepository{
		BaseRepository: NewBaseRepository(),
	}
}

func (r *DeveloperRepository) GetAllDevelopers() ([]entity.Developer, error) {
	query := `
		SELECT
			developer_id,
			developer_name
		FROM developers
		WHERE deleted_at IS NULL
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var developers []entity.Developer
	for rows.Next() {
		var developer entity.Developer
		err := rows.Scan(
			&developer.Developer_id,
			&developer.Developer_Name,
		)
		if err != nil {
			return nil, err
		}
		developers = append(developers, developer)
	}

	return developers, nil
}
