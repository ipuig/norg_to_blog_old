package routes

import (
	. "blog/pages"
	. "blog/parser"
	"fmt"
	"net/http"
)

func Routes() *http.ServeMux {
    mux := http.NewServeMux()

    /* Generate Posts */
    posts := GeneratePosts(mux)
    processedPosts := ProcessPosts(posts)

    /* Generate HomePage */
    GenerateHomePage(mux, processedPosts)

    /* Generate Year pages */
    GenerateYearPages(mux, processedPosts)

    /* Add Resources */
    ServeFiles(mux)
    return mux
}

func GenerateHomePage(mux *http.ServeMux, processedPosts ProcessedPosts) {
    home := &Homepage{
        Page: Page{
            Title: "Recent Posts",
        },
        Posts: processedPosts.Posts(),
        Author: "Ivan B. Puig",
    }
    mux.HandleFunc("/", home.Template())
}

func GenerateYearPages(mux *http.ServeMux, processedPosts ProcessedPosts) {
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

        mux.HandleFunc(fmt.Sprintf("/posts/%d", year), yearPage.Template())
    }
}

func GeneratePosts(mux *http.ServeMux) []Post {
    paths := FetchPostPaths()
    posts := make([]Post, len(paths) - 1)

    for _, path := range paths {
        fpost := PostFromPath(path)
        ServePostResources(&fpost)
        post := Post{
            Page: Page{
                Title: fpost.Title,
            },
            HTML: HTML(fpost),
            Date: DateFromString(fpost.Date),
            
        }
        posts = append(posts, post)
    }
    processed := ProcessPosts(posts)
    posts = processed.Posts()

    for _, post := range posts {
        mux.HandleFunc(post.URL(), post.Template())
    }

    return posts
}
