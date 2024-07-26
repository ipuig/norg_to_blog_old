package routes

import (
	"fmt"
	"regexp"
	"strings"
)

type Homepage struct {
    Title string
    Posts []Post
}

type Date struct {
    Year int
    Month int
    Day int
}

func (d Date) Format() string {
    return fmt.Sprintf("%d-%d-%d", d.Day, d.Month, d.Year)
}

type Post struct {
    Title string
    PostTags []string
    Date Date
    Abstract string
}

func (p Post) URL() string {
    re := regexp.MustCompile(`(\s+|\s*,\s*)`)
    urlSafeTitle := re.ReplaceAllString(p.Title, "-")
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
