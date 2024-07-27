package pages

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"regexp"
	"slices"
	"sort"
	"strings"
)

type Date struct {
    Year int
    Month int
    Day int
}


func (d Date) Format() string {
    return fmt.Sprintf("%d-%d-%d", d.Day, d.Month, d.Year)
}

const PostPath = "layouts/post/index.html"
type Post struct {
    Abstract string
    PostTags []string
    Date Date
    Page Page
    HTML template.HTML
}

func (p *Post) Template() func (w http.ResponseWriter, r *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
        tmpl := template.Must(template.ParseFiles(PostPath, LateralControlPath, PathControlPath, FooterPath))
        tmpl.Execute(w, p)
    }
}

func (p Post) URL() string {
    re := regexp.MustCompile(`(\s+|\s*,\s*)`)
    urlSafeTitle := re.ReplaceAllString(p.Page.Title, "-")
    return fmt.Sprintf("/posts/%d/%s", p.Date.Year, urlSafeTitle)
}

func (p Post) Tags() string {
    return strings.Join(p.PostTags, ", ")
}

func (p Post) HasTags() bool {
    return p.PostTags != nil && len(p.PostTags) >= 1;
}

func (p Post) HasAbstract() bool {
    return p.Abstract != ""
}

func CreatePostFromHTML(title, abstract string, tags []string, date Date, filepath string) Post {
    post := Post{
        Page: Page{ Title: title },
        Abstract: abstract,
        PostTags: tags,
        Date: date,
    }

    content, err := os.ReadFile(filepath)
    if err != nil {
        panic("Error while reading the file on CreatePostFromHTML")
    }

    post.HTML = template.HTML(string(content))
    return post
}

type Posts []Post
func (ps Posts) Len() int { return len(ps) }
func (ps Posts) Swap(i, j int) { ps[i], ps[j] = ps[j], ps[i] }
func (ps Posts) Less(i, j int) bool { 
    ya := ps[i].Date.Year
    yb := ps[j].Date.Year
    if ya < yb {
        return true
    }
    ma := ps[i].Date.Month
    mb := ps[j].Date.Month
    if ya == yb && ma < mb {
        return true
    }

    da := ps[i].Date.Day
    db := ps[j].Date.Day
    if ya == yb && ma == mb && da < db {
        return true
    }

    return false
}

type ProcessedPosts struct {
    Len int
    ByYear map[int][]Post
}

func (pp ProcessedPosts) Posts() []Post {
    years := make([]int, 0, len(pp.ByYear))
    posts := make([]Post, 0, pp.Len)

    for year := range pp.ByYear {
        years = append(years, year)
    }

    sort.Ints(years)
    slices.Reverse(years)

    for _, year := range years {
        posts = append(posts, pp.ByYear[year]...)
    }

    return posts
}

func ProcessPosts(posts []Post) ProcessedPosts {
    classified := make(map[int][]Post)
    sort.Sort(sort.Reverse(Posts(posts)))

    for idx, post := range posts {

        if idx > 0 {
            previous := posts[idx - 1]
            lc := MakeLateralControl("<", previous.URL(), previous.Page.Title)
            post.Page.LeftLateralControl = lc;
        }

        if idx < (len(posts) - 1) {
            next := posts[idx + 1]
            lc := MakeLateralControl(">", next.URL(), next.Page.Title)
            post.Page.RightLateralControl = lc;

        }

        post.Page.PathControl = PathControl{ post.URL() }

        ps, ok := classified[post.Date.Year]
        if !ok {
            classified[post.Date.Year] = []Post{ post }
            continue
        }

        classified[post.Date.Year] = append(ps, post)
    }
    return ProcessedPosts{Len: len(posts), ByYear: classified}
} 
