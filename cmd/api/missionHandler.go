package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"server/models"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

func (app *Application) getAllMission(w http.ResponseWriter, r *http.Request) {

	missions, err := app.Models.DB.GetAllMission()

	if err != nil {
		app.Logger.Println(err)
		app.writeError(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, missions, "mission")

	if err != nil {
		app.Logger.Println(err)
		app.writeError(w, err)
		return
	}

}

type MissionPayload struct {
	Operation_id int    `json:"op_id"`
	Crew_size    int    `json:"crew_size"`
	Rocket       string `json:"rocket"`
	Launch_site  string `json:"launch_site"`
	Launch_date  string `json:"launch_date"`
}

func (app *Application) insertMission(w http.ResponseWriter, r *http.Request) {
	var payload MissionPayload

	err := json.NewDecoder(r.Body).Decode(&payload)

	if err != nil {
		app.Logger.Println(err)
		app.writeError(w, err)
		return
	}

	var mission models.Mission
	mission.Operation_id = payload.Operation_id
	mission.Crew_size = payload.Crew_size
	mission.Rocket = payload.Rocket
	mission.Launch_site = payload.Launch_site
	mission.Launch_date, _ = time.Parse("2006-01-02", payload.Launch_date)

	err = app.Models.DB.InsertMission(mission)

	if err != nil {
		app.Logger.Println(err)
		app.writeError(w, err)
		return
	}

	res := jsonResponse{
		OK: true,
	}

	err = app.writeJSON(w, http.StatusOK, res, "response")

	if err != nil {
		app.Logger.Println(err)
		app.writeError(w, err)
		return
	}
}

func (app *Application) deleteMission(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))

	if err != nil {
		app.Logger.Println(errors.New("invalid id parameter"))
		app.writeError(w, err)
		return
	}

	err = app.Models.DB.DeleteMission(id)

	if err != nil {
		app.Logger.Println(err)
		app.writeError(w, err)
		return
	}

	res := jsonResponse{
		OK: true,
	}

	err = app.writeJSON(w, http.StatusOK, res, "response")

	if err != nil {
		app.Logger.Println(err)
		app.writeError(w, err)
		return
	}
}
