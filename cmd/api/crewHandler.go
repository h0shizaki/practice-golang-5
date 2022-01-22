package main

import (
	"fmt"
	"net/http"
)

func (app *Application) getAllCrew(w http.ResponseWriter, r *http.Request) {
	crews, err := app.Models.DB.GetAllCrew()

	if err != nil {
		app.writeError(w, err)
		return
	}

	fmt.Println(crews)
	// err = app.writeJSON(w, 200, crews, "crewlist")

	if err != nil {
		app.writeError(w, err)
		return
	}

}
