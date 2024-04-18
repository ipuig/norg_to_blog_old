package routes

import (
	"blog/ipl"
	"fmt"
	"html/template"

	"net/http"
)

func Routes() *http.ServeMux {
    mux := http.NewServeMux()

    /* Handle resources */
    resources := http.FileServer(http.Dir("./assets"))
    mux.Handle("/resources/", http.StripPrefix("/resources/", resources))
    table := make(map[int]ipl.Posts)

    /* Generate Posts */
    DynamicRoutes(mux, table)

    /* Generate HomePage */
    GenerateHomePage(mux, table)

    return mux
}

func GenerateHomePage(mux *http.ServeMux, table map[int]ipl.Posts) {
    mux.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
        tmpl := template.Must(template.ParseFiles("layouts/base/index.html"))
        data :=  HomePage{ Title: "Home", Styles: []string {"resources/css/base.css"}, Posts: table }
        tmpl.Execute(w, data)
    })
}


func DynamicRoutes(mux *http.ServeMux, table map[int]ipl.Posts) {
    posts := ipl.AnalyseContent()
    tmpl := template.Must(template.ParseFiles("layouts/post/index.html"))
    for _, post := range posts {
        mux.HandleFunc(fmt.Sprintf("/%d/%s", post.Date.Year, post.URL), func(w http.ResponseWriter, r *http.Request) {
            ps, ok := table[post.Date.Year]
            if !ok {
                ps = make(ipl.Posts, 0)
            }
            table[post.Date.Year] = append(ps, post)
            tmpl.Execute(w, post)
        })
    }
}

type HomePage struct {
    Title string
    Styles []string
    Posts map[int]ipl.Posts
}
