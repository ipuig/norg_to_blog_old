package routes

import (
	. "blog/pages"
	"net/http"
)

func Routes() *http.ServeMux {
    mux := http.NewServeMux()

    /* Add Resources */
    ServeFiles(mux)

    /* Generate Posts */
    posts := GeneratePosts(mux)

    /* Generate HomePage */
    GenerateHomePage(mux, posts)

    return mux
}

func GenerateHomePage(mux *http.ServeMux, posts []Post) {
    home := &Homepage{
        Page: Page{
            Title: "Recent Posts",
        },
        Posts: posts,
    }
    mux.HandleFunc("/", home.Template())
}

func GeneratePosts(mux *http.ServeMux) []Post {

    posts := []Post{
        {
            Page: Page{
                Title: "First Post",
            },
            PostTags: []string{ "yap", "yup"},
            Date: Date{Year: 2024, Month: 7, Day: 24},
            Abstract: "stuff",
        },

        {
            Page: Page{
                Title: "Second Post",
            },
            PostTags: []string{ "yap", "yup"},
            Date: Date{Year: 2024, Month: 8, Day: 24},
            Abstract: "Lorem ipsum dolor sit amet, officia excepteur ex fugiat reprehenderit enim labore culpa sint ad nisi Lorem pariatur mollit ex esse exercitation amet. Nisi anim cupidatat excepteur officia. Reprehenderit nostrud nostrud ipsum Lorem est aliquip amet voluptate voluptate dolor minim nulla est proident. Nostrud officia pariatur ut officia. Sit irure elit esse ea nulla sunt ex occaecat reprehenderit commodo officia dolor Lorem duis laboris cupidatat officia voluptate. Culpa proident adipisicing id nulla nisi laboris ex in Lorem sunt duis officia eiusmod. Aliqua reprehenderit commodo ex non excepteur duis sunt velit enim. Voluptate laboris sint cupidatat ullamco ut ea consectetur et est culpa et culpa duis.",
        },
    }

    all := ProcessPosts(posts)

    for _, post := range all[2024] {
        mux.HandleFunc(post.URL(), post.Template())
    }

    return posts
}

//
// func DynamicRoutes(mux *http.ServeMux) ipl.Posts{
//     content := ipl.AnalyseContent()
//     posts := make(ipl.Posts, 0)
//     tmpl := template.Must(template.ParseFiles("layouts/post/index.html"))
//     for _, post := range content {
//         fmt.Println(post.URL)
//         mux.HandleFunc(fmt.Sprintf("/%d/%s", post.Date.Year, post.URL), func(w http.ResponseWriter, r *http.Request) {
//             posts = append(posts, post)
//             tmpl.Execute(w, post)
//         })
//     }
//     return posts
// }
