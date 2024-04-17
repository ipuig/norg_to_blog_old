package routes

import (
	"html/template"
	"net/http"
        "blog/ipl"
)

func Routes() *http.ServeMux {
    mux := http.NewServeMux()
    mux.HandleFunc("/", HomePage)

    /* Handle resources*/
    resources := http.FileServer(http.Dir("./assets"))
    mux.Handle("/resources/", http.StripPrefix("/resources/", resources))

    return mux
}

func HomePage(w http.ResponseWriter, r *http.Request) {

    tmpl := template.Must(template.ParseFiles("layouts/base/index.html"))

    table := ipl.GeneratePostsTable([]ipl.Post {ipl.Example()})
    posts := []ipl.Post { ipl.Example() }
    table[2024] = posts

    data := ipl.HomeData { Title: "Home", Styles: []string {"resources/css/base.css"}, Posts: &table }
    tmpl.Execute(w, data)
}
