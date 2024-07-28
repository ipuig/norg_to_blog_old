package main

import (
	"blog/routes"
	"log"
	"net/http"
)

func main() {
    log.Println("Server running at port 8080")
    http.ListenAndServe(":8080", routes.Routes())
}
