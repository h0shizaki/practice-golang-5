package main

import "net/http"

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
