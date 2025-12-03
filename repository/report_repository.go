package repository

import "rea_games/entity"

type ReportRepository struct {
	*BaseRepository
}

func NewReportRepository() *ReportRepository {
	return &ReportRepository{
		BaseRepository: NewBaseRepository(),
	}
}

func (r *ReportRepository) GetDevelopersReport() ([]entity.DeveloperGameCountReport, error){
	query := `
		SELECT
			d.developer_name,
			COUNT(g.game_id) AS total_games_developed,
			STRING_AGG(g.title, ', ') AS list_of_games
		FROM
			developers d
		JOIN
			games g ON d.developer_id = g.developer_id
		GROUP BY
			d.developer_id, d.developer_name
		ORDER BY
			d.developer_id ASC;
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var DeveloperGameCounts []entity.DeveloperGameCountReport

	for rows.Next() {
		var DeveloperGameCount entity.DeveloperGameCountReport

		err := rows.Scan(
			&DeveloperGameCount.DeveloperName,
			&DeveloperGameCount.TotalGamesDeveloped,
			&DeveloperGameCount.ListOfGames,
		)

		if err != nil {
			return nil, err
		}

		DeveloperGameCounts = append(DeveloperGameCounts, DeveloperGameCount)
	}

	return DeveloperGameCounts, nil
}