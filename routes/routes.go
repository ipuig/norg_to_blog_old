package routes

import (
	"blog/pages"
	"blog/parser"
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
    error := pages.ErrorPage{ URL: url }
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
    processedPosts := pages.ProcessPosts(posts)

    /* Generate HomePage */
    GenerateHomePage(router, processedPosts)

    /* Generate Year pages */
    GenerateYearPages(router, processedPosts)

    /* Add Resources */
    ServeFiles(router)
    return router
}

func GenerateHomePage(router *Router, processedPosts pages.ProcessedPosts) {
    home := &pages.Homepage{
        Page: pages.Page{
            Title: "Recent Posts",
        },
        Posts: processedPosts.Posts(),
        Author: "Ivan B. Puig",
    }
    router.AddRoute("/", home.Template())
    router.AddRoute("/posts", home.Template())
}

func GenerateYearPages(router *Router, processedPosts pages.ProcessedPosts) {
    for year, ps := range processedPosts.ByYear {

        yearPage := &pages.Homepage{
            Page: pages.Page{
                Title: fmt.Sprintf("Posts from %d", year),
                PathControl: pages.PathControl{
                    URL: fmt.Sprintf("/posts/%d", year),
                },
            },
            Posts: ps,
        }

        _, next := processedPosts.ByYear[year + 1]
        _, previous := processedPosts.ByYear[year - 1]

        if previous {
            control := pages.MakeLateralControl(">", fmt.Sprintf("/posts/%d", year - 1), fmt.Sprintf("Posts from %d", year - 1))
            yearPage.Page.RightLateralControl = control
        }

        if next {
            control := pages.MakeLateralControl("<", fmt.Sprintf("/posts/%d", year + 1), fmt.Sprintf("Posts from %d", year + 1))
            yearPage.Page.LeftLateralControl = control
        }

        router.AddRoute(fmt.Sprintf("/posts/%d", year), yearPage.Template())
    }
}

func GeneratePosts(router *Router) []pages.Post {
    paths := parser.FetchPostPaths()
    posts := make([]pages.Post, 0)

    for _, path := range paths {
        fpost := parser.PostFromPath(path)
        ServePostResources(&fpost)
        post, err := parser.ParseContent(&fpost)
        if err != nil {
            continue
        }
        posts = append(posts, post)
    }
    processed := pages.ProcessPosts(posts)
    posts = processed.Posts()

    for _, post := range posts {
        router.AddRoute(post.URL(), post.Template())
    }

    return posts
}
