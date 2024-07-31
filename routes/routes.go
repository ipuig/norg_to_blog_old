package routes

import (
	. "blog/pages"
	. "blog/parser"
	"fmt"
	"net/http"
)

type Router struct {
    Routes map[string]struct{}
    Mux *http.ServeMux
}

func NewRouter() *Router {
    return &Router{
        Routes: make(map[string]struct{}),
        Mux: http.NewServeMux(),
    }
}

func (r *Router) AddRoute(endpoint string, resource func(w http.ResponseWriter, r *http.Request)) {
    r.Routes[endpoint] = struct{}{}
    r.Mux.HandleFunc(endpoint, resource)
}

func (r *Router) IsValidRoute(endpoint string) bool {
    _, ok := r.Routes[endpoint]
    return ok
}

func (r *Router) NotFoundPage(url string, w http.ResponseWriter, req *http.Request) {
    error := ErrorPage{ URL: url }
    error.WriteError(w)
    fmt.Fprint(w)
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    if r.IsValidRoute(req.URL.Path) {
        r.Mux.ServeHTTP(w, req)
    } else {
        r.NotFoundPage(req.URL.Path, w, req)
    }
}

func Routes() *Router {
    router := NewRouter()

    /* Generate Posts */
    posts := GeneratePosts(router)
    processedPosts := ProcessPosts(posts)

    /* Generate HomePage */
    GenerateHomePage(router, processedPosts)

    /* Generate Year pages */
    GenerateYearPages(router, processedPosts)

    /* Add Resources */
    ServeFiles(router)
    return router
}

func GenerateHomePage(router *Router, processedPosts ProcessedPosts) {
    home := &Homepage{
        Page: Page{
            Title: "Recent Posts",
        },
        Posts: processedPosts.Posts(),
        Author: "Ivan B. Puig",
    }
    router.AddRoute("/", home.Template())
}

func GenerateYearPages(router *Router, processedPosts ProcessedPosts) {
    for year, ps := range processedPosts.ByYear {

        yearPage := &Homepage{
            Page: Page{
                Title: fmt.Sprintf("Posts from %d", year),
                PathControl: PathControl{
                    URL: fmt.Sprintf("/posts/%d", year),
                },
            },
            Posts: ps,
        }

        _, next := processedPosts.ByYear[year + 1]
        _, previous := processedPosts.ByYear[year - 1]

        if previous {
            control := MakeLateralControl(">", fmt.Sprintf("/posts/%d", year - 1), fmt.Sprintf("Posts from %d", year - 1))
            yearPage.Page.RightLateralControl = control
        }

        if next {
            control := MakeLateralControl("<", fmt.Sprintf("/posts/%d", year + 1), fmt.Sprintf("Posts from %d", year + 1))
            yearPage.Page.LeftLateralControl = control
        }

        router.AddRoute(fmt.Sprintf("/posts/%d", year), yearPage.Template())
    }
}

func GeneratePosts(router *Router) []Post {
    paths := FetchPostPaths()
    posts := make([]Post, 0)

    for _, path := range paths {
        fpost := PostFromPath(path)
        ServePostResources(&fpost)
        post := Post{
            Page: Page{
                Title: fpost.Title,
            },
            HTML: HTML(fpost),
            Date: DateFromString(fpost.Date),
            PostTags: fpost.Tags,
            AdditionalCSS: fpost.CSSFiles,
            Abstract: fpost.Abstract,
        }
        posts = append(posts, post)
    }
    processed := ProcessPosts(posts)
    posts = processed.Posts()

    for _, post := range posts {
        router.AddRoute(post.URL(), post.Template())
    }

    return posts
}
