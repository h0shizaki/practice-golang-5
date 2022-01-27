package main

import (
	"encoding/json"
	"net/http"
	"server/models"
)

func (app *Application) getAllCrewOP(w http.ResponseWriter, r *http.Request) {
	var crew_op []*models.CrewOp

	crew_op, err := app.Models.DB.GetAllCrewOP()

	if err != nil {
		app.Logger.Fatal(err)
		app.writeError(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, crew_op, "crew_op")

	if err != nil {
		app.Logger.Fatal(err)
		app.writeError(w, err)
		return
	}
}

type crewOPPayload struct {
	Crew_ID int `json:"crew_id"`
	OP_ID   int `json:"op_id"`
}

func (app *Application) insertCrewOP(w http.ResponseWriter, r *http.Request) {
	var payload crewOPPayload

	err := json.NewDecoder(r.Body).Decode(&payload)

	if err != nil {
		app.Logger.Println(err)
		app.writeError(w, err)
		return
	}

	err = app.Models.DB.AddCrewOP(payload.Crew_ID, payload.OP_ID)

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

func (app *Application) deleteCrewOP(w http.ResponseWriter, r *http.Request) {
	var payload crewOPPayload

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		app.Logger.Println(err)
		app.writeError(w, err)
		return
	}

	err = app.Models.DB.DeleteCrewOP(payload.Crew_ID, payload.OP_ID)
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
