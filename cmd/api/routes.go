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
	router.HandlerFunc(http.MethodGet, "/v1/api/missionlist", app.getAllMission)
	router.HandlerFunc(http.MethodGet, "/v1/api/crew_op/", app.getAllCrewOP)

	//Delete crew
	router.HandlerFunc(http.MethodDelete, "/v1/api/delete/crew/:id", app.deleteCrew)
	router.HandlerFunc(http.MethodDelete, "/v1/api/delete/operation/:id", app.deleteOperation)
	router.HandlerFunc(http.MethodDelete, "/v1/api/delete/crew_op", app.deleteCrewOP)
	router.HandlerFunc(http.MethodDelete, "/v1/api/delete/mission/:id", app.deleteMission)

	//Post method insert
	router.HandlerFunc(http.MethodPost, "/v1/api/insert/crew", app.insertCrew)
	router.HandlerFunc(http.MethodPost, "/v1/api/insert/operation", app.insertOperation)
	router.HandlerFunc(http.MethodPost, "/v1/api/insert/crew_op", app.insertCrewOP)
	router.HandlerFunc(http.MethodPost, "/v1/api/insert/mission", app.insertMission)

	//Put
	router.HandlerFunc(http.MethodPut, "/v1/api/update/crew", app.editCrew)

	return router
}
