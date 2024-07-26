package pages

import (
	"fmt"
	"html/template"
	"net/http"
	"regexp"
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
