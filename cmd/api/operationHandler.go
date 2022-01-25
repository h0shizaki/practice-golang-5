package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"server/models"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *Application) getOneOperation(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))

	if err != nil {
		app.Logger.Println(errors.New("invalid id parameter"))
		app.writeError(w, err)
		return
	}

	operation, err := app.Models.DB.GetOneOperation(id)

	if err != nil {
		app.writeError(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, operation, "operation")

	if err != nil {
		app.writeError(w, err)
		return
	}

}

func (app *Application) getAllOperation(w http.ResponseWriter, r *http.Request) {

	operations, err := app.Models.DB.GetAllOperation()

	if err != nil {
		app.writeError(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, operations, "operation-list")

	if err != nil {
		app.writeError(w, err)
		return
	}

}

type opPayload struct {
	OP_name string `json:"operation_name"`
}

func (app *Application) insertOperation(w http.ResponseWriter, r *http.Request) {
	var payload opPayload

	err := json.NewDecoder(r.Body).Decode(&payload)

	if err != nil {
		app.writeError(w, err)
		return
	}

	op := models.Operation{
		Op_name: payload.OP_name,
	}

	// app.Logger.Println(op)

	err = app.Models.DB.InsertOpearion(op)

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

func (app *Application) deleteOperation(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))

	if err != nil {
		app.writeError(w, err)
		return
	}

	err = app.Models.DB.DeleteOperation(id)

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
