package controllers

import (
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	"github.com/fernandochimi/bank-account/src/models"
	u "github.com/fernandochimi/bank-account/src/utils"
)


var GetOperations = func(w http.ResponseWriter, r *http.Request) {
	data := models.GetOperations()
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var GetOperation = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}
	
	data := models.GetOperation(uint(id))
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}