package pages

import (
	"html/template"
	"net/http"
)

const HomepagePath = "layouts/homepage/index.html"
type Homepage struct {
    Posts []Post
    Page Page
}

func (h *Homepage) Template() func (w http.ResponseWriter, r *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
        tmpl := template.Must(template.ParseFiles(HomepagePath, LateralControlPath, PathControlPath, FooterPath))
        tmpl.Execute(w, h)
    }
}
