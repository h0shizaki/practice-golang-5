package models

import (
	"context"
	"time"
)

//Post
func (m *DBModel) InsertCrew(crew Crew) error {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := ` INSERT INTO crew(name, birth_date) VALUES($1,$2) ;`

	_, err := m.DB.ExecContext(ctx, stmt,
		crew.Name,
		crew.Birth_date,
	)

	if err != nil {
		return err
	}

	return nil
}

//Search by id
func (m *DBModel) GetOneCrew(id int) (*Crew, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT * FROM crew WHERE id = $1 ;`

	rows := m.DB.QueryRowContext(ctx, query, id)

	var crew Crew

	err := rows.Scan(
		&crew.ID,
		&crew.Name,
		&crew.Birth_date,
	)

	if err != nil {
		return nil, err
	}

	return &crew, nil
}

//Get all crew member
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
