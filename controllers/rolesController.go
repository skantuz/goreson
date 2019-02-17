package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/skantuz/goreson/models"
	u "github.com/skantuz/goreson/utils"
)

var CreateRole = func(w http.ResponseWriter, r *http.Request) {

	role := &models.Role{}
	err := json.NewDecoder(r.Body).Decode(role) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := role.Create() //Create user
	u.Respond(w, resp)
}
