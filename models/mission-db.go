package models

import (
	"context"
	"log"
	"time"
)

//Create CRUD here

func (m *DBModel) GetAllMission() ([]*Mission, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
		SELECT m.id , o.op_name , m.rocket ,m.crew_size , o.op_id , m.launch_site , m.launch_date  FROM 
			mission m INNER JOIN operation o ON o.op_id = m.op_id
		;
	`

	rows, _ := m.DB.QueryContext(ctx, query)
	defer rows.Close()

	var missions []*Mission

	for rows.Next() {
		var mission Mission

		err := rows.Scan(
			&mission.ID,
			&mission.Operation_name,
			&mission.Rocket,
			&mission.Crew_size,
			&mission.Operation_id,
			&mission.Launch_site,
			&mission.Launch_date,
		)

		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		crewQuery := `SELECT c.id , c.name , c.birth_date FROM crew c RIGHT JOIN crew_op co ON c.id = co.crew_id WHERE co.op_id = $1 ;`
		crewRows, _ := m.DB.QueryContext(ctx, crewQuery, mission.Operation_id)

		crew_list := make(map[int]string)

		for crewRows.Next() {
			var crew Crew
			err := crewRows.Scan(
				&crew.ID,
				&crew.Name,
				&crew.Birth_date,
			)

			if err != nil {
				log.Fatal(err)
				return nil, err
			}

			crew_list[crew.ID] = crew.Name
		}
		crewRows.Close()
		mission.Crew_list = crew_list

		missions = append(missions, &mission)
	}

	return missions, nil

}
