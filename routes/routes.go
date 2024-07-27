package routes

import (
	. "blog/pages"
	"fmt"
	"net/http"
)

func Routes() *http.ServeMux {
    mux := http.NewServeMux()

    /* Add Resources */
    ServeFiles(mux)

    /* Generate Posts */
    posts := GeneratePosts(mux)
    processedPosts := ProcessPosts(posts)

    /* Generate HomePage */
    GenerateHomePage(mux, processedPosts)

    /* Generate Year pages */
    GenerateYearPages(mux, processedPosts)
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

    posts := []Post{
        {
            Page: Page{
                Title: "My wife told me 28h til next kisu in pipi",
            },
            PostTags: []string{ "yap", "yup"},
            Date: Date{Year: 2024, Month: 7, Day: 24},
            Abstract: "stuff",
        },
        {
            Page: Page{
                Title: "We'll eat wings today",
            },
            PostTags: []string{ "yap", "yup"},
            Date: Date{Year: 2024, Month: 8, Day: 24},
            Abstract: "Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat.",
        },

        {
            Page: Page{
                Title: "Babsuuuu",
            },
            PostTags: []string{ "yap", "yup"},
            Date: Date{Year: 2023, Month: 8, Day: 24},
            Abstract: "Lorem ipsum dolor sit amet, officia excepteur ex fugiat reprehenderit enim labore culpa sint ad nisi Lorem pariatur mollit ex esse exercitation amet. Nisi anim cupidatat excepteur officia. Reprehenderit nostrud nostrud ipsum Lorem est aliquip amet voluptate voluptate dolor minim nulla est proident. Nostrud officia pariatur ut officia. Sit irure elit esse ea nulla sunt ex occaecat reprehenderit commodo officia dolor Lorem duis laboris cupidatat officia voluptate. Culpa proident adipisicing id nulla nisi laboris ex in Lorem sunt duis officia eiusmod. Aliqua reprehenderit commodo ex non excepteur duis sunt velit enim. Voluptate laboris sint cupidatat ullamco ut ea consectetur et est culpa et culpa duis.",
        },
        CreatePostFromHTML("Post from html", "This is a test post that I created from a random html file", []string{"testing", "webdev", "golang"}, Date{Year: 2024, Month: 7, Day: 27}, "content/test.html"),
    }

    processedPosts := ProcessPosts(posts)
    for _, post := range processedPosts.Posts() {
        mux.HandleFunc(post.URL(), post.Template())
    }

    return posts
}
