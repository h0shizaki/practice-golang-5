package models

import (
	"context"
	"time"
)

func (m *DBModel) GetAllCrew() ([]*Crew, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT * FROM crew ;`

	rows, err := m.DB.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	var crews []*Crew

	for rows.Next() {

		var crew Crew

		err := rows.Scan(
			&crew.ID,
			&crew.Name,
			&crew.Birth_date,
		)

		if err != nil {
			return nil, err
		}

		crews = append(crews, &crew)

	}

	return crews, nil
}
