package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/skantuz/goreson/models"
	u "github.com/skantuz/goreson/utils"
)

var CreateUser = func(w http.ResponseWriter, r *http.Request) {

	user := &models.Users{}
	err := json.NewDecoder(r.Body).Decode(user) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := user.Create() //Create user
Usernamespond(w, resp)
}

var Authenticate = func(w http.ResponseWriter, r *http.Request) {

	user Usernameodels.Users{}
	err := json.NewDecoder(r.Body).Decode(user)Usernamecode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := models.Login(user.Username, user.Password)
	u.Respond(w, resp)
}


