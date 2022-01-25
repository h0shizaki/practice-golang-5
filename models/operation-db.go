package models

import (
	"context"
	"time"
)

func (m *DBModel) GetAllOperation() ([]*Operation, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT * FROM operation ; `

	rows, err := m.DB.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	var operations []*Operation

	for rows.Next() {

		var operation Operation

		err := rows.Scan(
			&operation.ID,
			&operation.Op_name,
		)

		if err != nil {
			return nil, err
		}

		operations = append(operations, &operation)
	}

	return operations, nil

}

func (m *DBModel) GetOneOperation(id int) (*Operation, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT * FROM operation WHERE op_id = $1`

	row := m.DB.QueryRowContext(ctx, query, id)

	var operation Operation

	err := row.Scan(
		&operation.ID,
		&operation.Op_name,
	)

	if err != nil {
		return nil, err
	}

	return &operation, nil

}

func (m *DBModel) InsertOpearion(op Operation) error {

	ctx, cancel := context.WithTimeout(context.TODO(), 3*time.Second)
	defer cancel()

	stmt := ` INSERT INTO operation(op_name) VALUES($1) ;`

	_, err := m.DB.ExecContext(ctx, stmt, op.Op_name)

	if err != nil {
		return err
	}

	return nil
}

func (m *DBModel) DeleteOperation(id int) error {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `DELETE FROM operation WHERE op_id = $1 `

	_, err := m.DB.ExecContext(ctx, stmt, id)

	if err != nil {
		return err
	}

	return nil

}
