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
    "/static/css/post/code.css": "./assets/css/posts/code.css",
    "/static/css/error/styles.css": "./assets/css/error/error.css",
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

func ServePostResources(fpost *parser.FSPost) {
    fpost.Metadata()
    if fpost.Title == "" || fpost.Date == "" {
        fmt.Println("no metadata", fpost.RootPath)
        return 
    }

    url := regexp.MustCompile(" ").ReplaceAllString(strings.ToLower(fpost.Title), "-")
    path := fpost.RootPath
    alias := "/static/posts"

    re := regexp.MustCompile(`[2][0][2-9][0-9]`)
    year := re.FindStringSubmatch(fpost.Date)[0]
    if year == "" {
        fmt.Println("Date wrong format")
        return
    }

    if fpost.CSSFiles != nil {
        updatedCSS := make([]string, 0)
        for _, css := range fpost.CSSFiles {
            if css == "" {
                continue
            }
            newpath := alias + "/" + year + "/" + url + "/css/" + css
            updatedCSS = append(updatedCSS, newpath)
            files[newpath] = path + "/css/" + css 
        }
        fpost.CSSFiles = updatedCSS
    }

    if fpost.ImagesPath != nil {
        updatedImages := make([]string, 0)
        for _, img := range fpost.ImagesPath {
            if img == "" {
                continue
            }
            newpath := alias + "/" + year + "/" + url + "/img/" + img
            updatedImages = append(updatedImages, newpath)
            files[newpath] = path + "/img/" + img 
        }
        fpost.ImagesPath = updatedImages
    }

    if fpost.Logo != "" {
        files[alias + "/" + year + "/" + url + "/img/" + fpost.Logo] = path + "/img/" + fpost.Logo
        fpost.Logo = alias + "/" + year + "/" + url + "/img/" + fpost.Logo
    }
}

func ServeFiles(router *Router) {
    for path := range files {
        router.AddRoute(path, fileHandler)
    }
}
