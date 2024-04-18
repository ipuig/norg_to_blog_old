package main

import (
	"blog/routes"
	"net/http"
)

func main() {
    http.ListenAndServe(":8080", routes.Routes())
}

