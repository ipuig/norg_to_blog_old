package routes

import (
	"blog/ipl"
	"fmt"
	"html/template"

	"net/http"
)

func Routes() *http.ServeMux {
    mux := http.NewServeMux()

    /* Add Resources */
    ServeFiles(mux)

    /* Generate Posts */
    posts := DynamicRoutes(mux)

    /* Generate HomePage */
    GenerateHomePage(mux, posts)

    return mux
}

func GenerateHomePage(mux *http.ServeMux, posts ipl.Posts) {
    _ = posts
    mux.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
        tmpl := template.Must(template.ParseFiles("layouts/homepage/index.html"))
        data :=  Homepage { Title: "Home", Posts: []Post {
            {"First Post", []string {}, Date{2024,7,24}, ""},
            // {"Second Post", []string {"htb", "file upload"}, Date{2024,7,24}, "Brief insights of how i pwned the Pilgrimage box from HTB"},
            // {"Second Post", []string {"htb", "file upload"}, Date{2024,7,24}, "Brief insights of how i pwned the Pilgrimage box from HTB"},
            // {"Second Post", []string {"htb", "file upload"}, Date{2024,7,24}, "Brief insights of how i pwned the Pilgrimage box from HTB"},
            // {"Second Post", []string {"htb", "file upload"}, Date{2024,7,24}, "Brief insights of how i pwned the Pilgrimage box from HTB"},
            // {"Second Post", []string {"htb", "file upload"}, Date{2024,7,24}, "Brief insights of how i pwned the Pilgrimage box from HTB"},
            // {"Second Post", []string {"htb", "file upload"}, Date{2024,7,24}, "Brief insights of how i pwned the Pilgrimage box from HTB"},
            // {"Second Post", []string {"htb", "file upload"}, Date{2024,7,24}, "Brief insights of how i pwned the Pilgrimage box from HTB"},
            // {"Second Post", []string {"htb", "file upload"}, Date{2024,7,24}, "Brief insights of how i pwned the Pilgrimage box from HTB"},
            // {"Second Post", []string {"htb", "file upload"}, Date{2024,7,24}, "Brief insights of how i pwned the Pilgrimage box from HTB"},
            // {"Second Post", []string {"htb", "file upload"}, Date{2024,7,24}, "Brief insights of how i pwned the Pilgrimage box from HTB"},
            // {"Third Post", []string {"networking"}, Date{2024,7,24}, ""},
        }}
        tmpl.Execute(w, data)
    })
}


func DynamicRoutes(mux *http.ServeMux) ipl.Posts{
    content := ipl.AnalyseContent()
    posts := make(ipl.Posts, 0)
    tmpl := template.Must(template.ParseFiles("layouts/post/index.html"))
    for _, post := range content {
        fmt.Println(post.URL)
        mux.HandleFunc(fmt.Sprintf("/%d/%s", post.Date.Year, post.URL), func(w http.ResponseWriter, r *http.Request) {
            posts = append(posts, post)
            tmpl.Execute(w, post)
        })
    }
    return posts
}
