package ipl

import (
	"fmt"
	"html/template"
	"strings"
)

type Post struct {
    Elements []*TagHTML
    Title string
    URL string
    // Desc string
    Header template.HTML
    Abstract template.HTML
    Tags []string
    Date *Date
    HTML []template.HTML
}
type Posts []*Post

func (a Posts) Len() int {
    return len(a)
}

func (a Posts) Swap(i, j int) {
    a[i], a[j] = a[j], a[i]
}

func (a Posts) Less(i, j int) bool {
    if a[i].Date.Year != a[j].Date.Year {
        return a[i].Date.Year < a[j].Date.Year
    }
    if a[i].Date.Month != a[j].Date.Month {
        return a[i].Date.Month < a[j].Date.Month
    }
    if a[i].Date.Day != a[j].Date.Day {
        return a[i].Date.Day < a[j].Date.Day
    }
    return a[i].Date.Hours < a[j].Date.Hours
}


func (p *Post) Generate() {
    p.HTML = make([]template.HTML, 0)
    p.Date = &Date{}
    for _, element := range p.Elements {
        switch element.TagType {
        case h1:
            p.URL = strings.ReplaceAll(strings.ToLower(element.Text), " ", "_")
            fmt.Println(p.URL)
            p.Title = element.Text
        case head:
            for _, line := range strings.Split(element.Text, "\n") {
                if strings.HasPrefix(line, "tags") {
                    t := strings.TrimPrefix(line, "tags ")
                    p.Tags = strings.Split(t, ",")
                    continue
                }

                p.Date.DateFromString(line)
                p.Header = element.Render()
            }
        case abs:
            p.Abstract = element.Render()
        default:
            p.HTML = append(p.HTML, element.Render())
    }
    }
}
