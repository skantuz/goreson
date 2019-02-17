package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/skantuz/goreson/app"
	c "github.com/skantuz/goreson/controllers"
)

type Route struct {
	Name       string
	Method     string
	Pattern    string
	HandleFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.Use(app.JwtAuthentication) //attach JWT auth middleware
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Handler(route.HandleFunc)
	}
	return router
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/api",
		c.Index,
	},
	Route{
		"New Role",
		"POST",
		"/api/roles",
		c.CreateRole,
	},
	Route{
		"New User",
		"POST",
		"/api/users",
		c.CreateUser,
	},
	Route{
		"New User",
		"GET",
		"/api/users/id",
		c.GetUsers,
	},
	Route{
		"Index",
		"GET",
		"/api/users/login",
		c.Index,
	},
}
