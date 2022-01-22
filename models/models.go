package models

import (
	"database/sql"
	"time"
)

type Models struct {
	DB DBModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		DB: DBModel{DB: db},
	}
}

type DBModel struct {
	DB *sql.DB
}

//Create models here

type Crew struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Birth_date time.Time `json:"birth_date"`
}

type Operation struct {
	ID      int    `json:"id"`
	Op_name string `json:"operation_name"`
}

type CrewOperation struct {
	ID             int            `json:"id"`
	Crew_id        int            `json:"crew_id"`
	Crew_list      map[int]string `json:"crew"`
	Crew_size      int            `json:"crew_size"`
	Operation_id   int            `json:"-"`
	Operation_name string         `json:"operation_name"`
	Rocket         string         `json:"rocket"`
	Launch_site    string         `json:"launch_site"`
	Launch_date    time.Time      `json:"launch_date"`
}
