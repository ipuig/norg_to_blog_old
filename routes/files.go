package routes

import (
	"blog/parser"
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

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

func ServePostResource(fpost parser.FSPost) {
    title, date := fpost.Metadata()
    if title == "" || date == "" {
        fmt.Println("no metadata", fpost.RootPath)
        return
    }

    url := regexp.MustCompile(" ").ReplaceAllString(strings.ToLower(title), "-")
    path := fpost.RootPath
    alias := "/static/posts"

    re := regexp.MustCompile(`[2][0][2-9][0-9]`)
    year := re.FindStringSubmatch(date)[0]
    if year == "" {
        return
    }

    if fpost.CSSFiles != nil {
        for _, css := range fpost.CSSFiles {
            if css == "" {
                continue
            }
            files[alias + "/" + year + "/" + url + "/css/" + css] = path + "/css/" + css 
        }
    }

    if fpost.ImagesPath != nil {
        for _, img := range fpost.ImagesPath {
            if img == "" {
                continue
            }
            files[alias + "/" + year + "/" + url + "/img/" + img] = path + "/img/" + img 
        }
    }
}

func ServeFiles(mux *http.ServeMux) {
    for path := range files {
        mux.HandleFunc(path, fileHandler)
    }
}
