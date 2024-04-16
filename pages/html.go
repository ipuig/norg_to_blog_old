package pages

import (
    "fmt"
    "html/template"
)

const (
    p = iota
    h1
    h2
    h3
    h4
    h5
    h6
    figure
    img
    a
    article
    section
)

type Date struct {
    Year int
    Month int
    Day int
    hours int
    minutes int
    seconds int
}

type Post struct {
    PostTitle string
    Sections []Section
    Date *Date
    Tags []string
}

type Section struct {
    SectionTitle string
    Tags []template.HTML
} 

type Tag struct {
    TagType int
    Text string
    TagOptions
}

type TagOptions = struct {
    Path string
    X int
    Y int
    Size float32
}

func (t Tag) HTML() template.HTML {

    var toRender string

    switch t.TagType {

    case p: toRender = fmt.Sprintf("<p>%s</p>", t.Text)
    case h1: toRender = fmt.Sprintf("<h1>%s</h1>", t.Text)
    case h2: toRender = fmt.Sprintf("<h2>%s</h2>", t.Text)
    case h3: toRender = fmt.Sprintf("<h3>%s</h3>", t.Text)
    case h4: toRender = fmt.Sprintf("<h4>%s</h4>", t.Text)
    case h5: toRender = fmt.Sprintf("<h5>%s</h5>", t.Text)
    case h6: toRender = fmt.Sprintf("<h6>%s</h6>", t.Text)

    }

    return template.HTML(toRender)
}

func GeneratePostsTable(posts []Post) map[int][]Post {
    table := make(map[int][]Post)

    for _, post := range posts {

        year := post.Date.Year
        yearPosts, ok := table[year]

        if !ok {
            yearPosts = []Post{ post }
            table[year] = yearPosts
            continue
        }

        yearPosts = append(yearPosts, post)
        table[year] = yearPosts
    }

    return table
}

func tagsToHtml(tags []Tag) []template.HTML {
    var htmlTags []template.HTML
    for _, tag := range tags {
        htmlTags = append(htmlTags, tag.HTML())
    }
    return htmlTags
}
