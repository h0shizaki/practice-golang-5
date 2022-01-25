package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *Application) routes() http.Handler {

	router := httprouter.New()

	//Check Status
	router.HandlerFunc(http.MethodGet, "/status", app.statusHandler)

	//Get and Search method
	router.HandlerFunc(http.MethodGet, "/v1/api/crewlist", app.getAllCrew)
	router.HandlerFunc(http.MethodGet, "/v1/api/crewlist/:id", app.getOneCrew)
	router.HandlerFunc(http.MethodGet, "/v1/api/operationlist", app.getAllOperation)
	router.HandlerFunc(http.MethodGet, "/v1/api/operationlist/:id", app.getOneOperation)

	//Delete crew
	router.HandlerFunc(http.MethodDelete, "/v1/api/delete/crew/:id", app.deleteCrew)
	router.HandlerFunc(http.MethodDelete, "/v1/api/delete/operation/:id", app.deleteOperation)

	//Post method insert
	router.HandlerFunc(http.MethodPost, "/v1/api/insert/crew", app.insertCrew)
	router.HandlerFunc(http.MethodPost, "/v1/api/insert/operation", app.insertOperation)

	//Put
	router.HandlerFunc(http.MethodPut, "/v1/api/update/crew", app.editCrew)

	return router
}
