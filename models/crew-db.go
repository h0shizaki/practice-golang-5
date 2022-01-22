package models

import (
	"context"
	"time"
)

func (m *DBModel) GetAllCrew() ([]*Crew, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT * FROM crew ORDER BY id ; `

	rows, err := m.DB.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var allCrews []*Crew

	for rows.Next() {
		var crew Crew
		err := rows.Scan(
			&crew.ID,
			&crew.NAME,
			&crew.BIRTH_DATE,
		)

		if err != nil {
			return nil, err
		}

		allCrews = append(allCrews, &crew)

	}

	return allCrews, nil
}
