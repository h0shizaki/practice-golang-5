package main

import (
	"net/http"
)

func (app *Application) getAllCrew(w http.ResponseWriter, r *http.Request) {
	crews, err := app.Models.DB.GetAllCrew()

	if err != nil {
		app.writeError(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, crews, "crew_list")

	if err != nil {
		app.writeError(w, err)
		return
	}
}
