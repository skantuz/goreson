package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/skantuz/goreson/models"
	u "github.com/skantuz/goreson/utils"
)

var CreateUser = func(w http.ResponseWriter, r *http.Request) {

	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user) //decode the request body into struct and failed if any error occur
	log.Printf(user.Email)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request, Verify the json format "+err.Error()))
		return
	}

	resp := user.Create() //Create user
	u.Respond(w, resp)
}

var Authenticate = func(w http.ResponseWriter, r *http.Request) {

	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := models.Login(user.Username, user.Password)
	u.Respond(w, resp)
}

var GetUsers = func(w http.ResponseWriter, r *http.Request) {
	data := models.ListUsers()
	log.Panicf("6")
	resp := u.Message(true, "success")
	log.Panicf("7")
	resp["data"] = data
	log.Panicf("8")
	u.Respond(w, resp)

	return
}
