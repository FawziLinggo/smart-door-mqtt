package router

import (
	"RESTful/app/handlers"
	"net/http"
)

func SetupRoutes() {
	http.HandleFunc("/login", handlers.LoginHandler)
}
