package models

import (
	"context"
	"time"
)

func (m *DBModel) GetAllCrewOP() ([]*CrewOp, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := ` SELECT * FROM crew_op ; `

	rows, _ := m.DB.QueryContext(ctx, query)

	var crew_op_list []*CrewOp

	for rows.Next() {
		var crew_op CrewOp

		err := rows.Scan(
			&crew_op.ID,
			&crew_op.Crew_ID,
			&crew_op.OP_ID,
		)

		if err != nil {
			return nil, err
		}

		crew_op_list = append(crew_op_list, &crew_op)

	}

	return crew_op_list, nil

}

func (m *DBModel) DeleteCrewOP(crew_id int, op_id int) error {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `DELETE FROM crew_op WHERE crew_id = $1 AND op_id = $2 ;`

	_, err := m.DB.ExecContext(ctx, stmt, crew_id, op_id)

	if err != nil {
		return err
	}

	return nil

}

func (m *DBModel) AddCrewOP(crew_id int, op_id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `INSERT INTO crew_op(crew_id , op_id) VALUES($1,$2) ;`

	_, err := m.DB.ExecContext(ctx, stmt, crew_id, op_id)

	if err != nil {
		return err
	}

	return nil

}
