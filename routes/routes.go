package routes

import (
	"blog/pages"
	"fmt"
	"html/template"
	"net/http"
)

func Routes() *http.ServeMux {
    mux := http.NewServeMux()
    mux.HandleFunc("/", HomePage)
    mux.HandleFunc("/test", RandomCrap("<h2> THIS IS A TEST!</h2>"))
    mux.HandleFunc("/secret", RandomCrap("sjsldfkjlk[j4klj2lkjdfkljslakjflj"))

    /* Handle resources*/
    resources := http.FileServer(http.Dir("./assets"))
    mux.Handle("/resources/", http.StripPrefix("/resources/", resources))

    return mux
}

func HomePage(w http.ResponseWriter, r *http.Request) {

    tmpl := template.Must(template.ParseFiles("index.html"))
    data := pages.HomeData { Title: "Home", Styles: []string {"resources/css/base.css"} }
    tmpl.Execute(w, data)
}

func RandomCrap(text string) func(http.ResponseWriter, *http.Request) {

    return func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, text)
    }
}
