package controllers

import (
	"net/http"

	u "github.com/skantuz/goreson/utils"
)

var Index = func(w http.ResponseWriter, r *http.Request) {
	u.Respond(w, u.Message(true, "System Run"))
}
