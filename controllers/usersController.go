package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

var GetUserFor = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		//The passed path parameter is not an integer
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}
	data := models.GetUser(uint(id))
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)

	return
}

var GetUsers = func(w http.ResponseWriter, r *http.Request) {
	data := models.ListUsers()
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)

	return
}
