package main

import (
	"net/http"
)

type Status struct {
	Status      string
	Environment string
	Version     string
}

func (app *Application) statusHandler(w http.ResponseWriter, r *http.Request) {
	currentStatus := Status{
		Status:      "Available",
		Environment: app.Config.Environment,
		Version:     app.Config.Version,
	}

	app.writeJSON(w, 200, currentStatus, "status")

}
