package main

import (
	"fmt"
	"net/http"
	"os"

	r "github.com/skantuz/goreson/routes"
)

func main() {

	router := r.NewRouter()
	port := os.Getenv("sys_port") //Get port from .env file, we did not specify any port so this should return an empty string when tested locally
	if port == "" {
		port = "8000" //localhost
	}

	fmt.Println("Listen" + port)

	err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}
