package models

import (
	"database/sql"
	"time"
)

type DBModel struct {
	DB *sql.DB
}

type Models struct {
	DB DBModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		DB: DBModel{DB: db},
	}
}

//Create models here

type Crew struct {
	ID         int       `json:"id"`
	NAME       string    `json:"name"`
	BIRTH_DATE time.Time `json:"birth_date"`
}

type Operation struct {
	ID      int    `json:"id"`
	OP_NAME string `json:"operation_name"`
}

type CrewOperation struct {
	ID             int            `json:"id"`
	CREW_ID        int            `json:"crew_id"`
	CREW_LIST      map[int]string `json:"crew"`
	CREW_SIZE      int            `json:"crew_size"`
	OPERATION_ID   int            `json:"-"`
	OPERATION_NAME string         `json:"operation_name"`
	ROCKET         string         `json:"rocket"`
	LAUNCH_SITE    string         `json:"launch_site"`
	LAUNCH_DATE    time.Time      `json:"launch_date"`
}
