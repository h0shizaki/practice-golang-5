package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *Application) routes() http.Handler {

	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/status", app.statusHandler)

	router.HandlerFunc(http.MethodGet, "/v1/api/crewlist", app.getAllCrew)

	return router
}
