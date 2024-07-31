package pages

import (
	"html/template"
	"net/http"
)

const ErrorPagePath = "layouts/error/index.html"
type ErrorPage struct {
    URL string
}

func (ep *ErrorPage) WriteError(w http.ResponseWriter) {
    w.WriteHeader(http.StatusNotFound)
    tmpl := template.Must(template.ParseFiles(ErrorPagePath))
    tmpl.Execute(w, ep)
}
