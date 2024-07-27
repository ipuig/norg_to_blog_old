package routes

import "net/http"

var files = map[string]string {
    "/static/css/homepage/styles.css": "./assets/css/homepage/homepage.css",
    "/static/css/base/styles.css": "./assets/css/base/base.css",
    "/static/css/post/styles.css": "./assets/css/posts/posts.css",
    "/static/css/base/colours.css": "./assets/css/base/colours.css",
    "/static/css/base/fonts.css": "./assets/css/base/fonts.css",
    "/static/fonts/base/font.ttf": "./assets/fonts/BlexMonoNerdFont-Text.ttf",

    "/static/social/github.svg": "./assets/images/svg/github.svg",
    "/static/social/linkedin.svg": "./assets/images/svg/linkedin.svg",
    "/static/social/htb.svg": "./assets/images/svg/htb.svg",
}

func fileHandler(w http.ResponseWriter, r *http.Request) {
    requestedFile := r.URL.Path
    filepath, ok := files[requestedFile]
    if ok {
        http.ServeFile(w, r, filepath)
    } else {
        http.NotFound(w, r)
    }
}

func ServeFiles(mux *http.ServeMux) {
    for path := range files {
        mux.HandleFunc(path, fileHandler)
    }
}
