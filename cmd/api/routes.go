package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *Application) routes() http.Handler {

	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/status", app.statusHandler)

	//Get and Search method
	router.HandlerFunc(http.MethodGet, "/v1/api/crewlist", app.getAllCrew)
	router.HandlerFunc(http.MethodGet, "/v1/api/crewlist/:id", app.getOneCrew)

	//Delete crew
	router.HandlerFunc(http.MethodGet, "/v1/api/delete/crew/:id", app.deleteCrew)

	//Post method insert
	router.HandlerFunc(http.MethodPost, "/v1/api/addcrew", app.addCrew)

	return router
}
