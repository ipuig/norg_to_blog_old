package main

import (
	"blog/routes"
	"net/http"
)

type PageData = struct {
    Title string
    Styles []string
}

func main() {

    http.ListenAndServe(":8080", routes.Routes())

}
