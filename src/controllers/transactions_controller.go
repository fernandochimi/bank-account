package controllers

import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/fernandochimi/bank-account/src/models"
	u "github.com/fernandochimi/bank-account/src/utils"
)

var CreateTransaction = func(w http.ResponseWriter, r *http.Request) {

	transaction := &models.Transaction{}
	err := json.NewDecoder(r.Body).Decode(transaction)
	if err != nil {
		fmt.Println(err)
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := transaction.Create()
	u.Respond(w, resp)
}
