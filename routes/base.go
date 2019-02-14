package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name      string
	Method    string
	Pattern   string
	HandleFun http.HandleFunc
}

type Routes []Route

func NewRouter() *mux.Router {

}
