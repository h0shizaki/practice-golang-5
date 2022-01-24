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

type jsonResponse struct {
	OK bool
}

func (app *Application) getOneCrew(w http.ResponseWriter, r *http.Request) {

	param := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(param.ByName("id"))

	if err != nil {
		app.Logger.Println(errors.New("invalid id parameter"))
		app.writeError(w, err)
		return
	}

	crew, err := app.Models.DB.GetOneCrew(id)

	if err != nil {
		app.Logger.Println(err)
		app.writeError(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, crew, "crew")

	if err != nil {
		app.Logger.Println(err)
		app.writeError(w, err)
		return
	}
}

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

type crewPayload struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Birth_date string `json:"birth_date"`
}

func (app *Application) addCrew(w http.ResponseWriter, r *http.Request) {

	var payload crewPayload

	err := json.NewDecoder(r.Body).Decode(&payload)

	if err != nil {
		app.Logger.Println(err)
		app.writeError(w, err)
		return
	}

	// log.Println(payload)

	var crew models.Crew
	crew.Name = payload.Name
	crew.Birth_date, _ = time.Parse("2006-01-02", payload.Birth_date)

	err = app.Models.DB.InsertCrew(crew)

	if err != nil {
		app.Logger.Println(err)
		app.writeError(w, err)
		return
	}

	var res jsonResponse

	res.OK = true

	err = app.writeJSON(w, http.StatusOK, res, "response")

	if err != nil {
		app.Logger.Println(err)
		app.writeError(w, err)
		return
	}

}

func (app *Application) deleteCrew(w http.ResponseWriter, r *http.Request) {
	params := r.Context().Value("params").(httprouter.Params)

	id, err := strconv.Atoi(params.ByName("id"))

	if err != nil {
		app.writeError(w, err)
		return
	}

	err = app.Models.DB.DeleteCrew(id)

	ok := jsonResponse{
		OK: true,
	}

	err = app.writeJSON(w, http.StatusOK, ok, "response")
	if err != nil {
		app.writeError(w, err)
		return
	}

}

//Update crew

func (app *Application) editCrew(w http.ResponseWriter, r *http.Request) {

	var payload crewPayload

	err := json.NewDecoder(r.Body).Decode(&payload)

	if err != nil {
		app.Logger.Println(err)
		app.writeError(w, err)
		return
	}

	var crew models.Crew

	crew.ID, _ = strconv.Atoi(payload.ID)
	crew.Name = payload.Name
	crew.Birth_date, _ = time.Parse("2006-01-02", payload.Birth_date)

	err = app.Models.DB.UpdateCrew(crew)

	if err != nil {
		app.writeError(w, err)
		return
	}

	ok := jsonResponse{
		OK: true,
	}

	err = app.writeJSON(w, http.StatusOK, ok, "response")

	if err != nil {
		app.writeError(w, err)
		return
	}

}
