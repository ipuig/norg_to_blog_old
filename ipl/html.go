package ipl

import (
	"fmt"
	"html/template"
	"log"
	"strconv"
	"strings"
)

const (
    p = iota
    span
    em
    b
    br
    hr
    h1
    h2
    h3
    h4
    h5
    h6
    figure
    img
    article
    section
    a
    div
    hgroup
    strong
    li
    head
)

type Date struct {
    Year int
    Month int
    Day int
    Hours int
    Minutes int
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
        case span: toRender = fmt.Sprintf("<span>%s</span>", t.Text)
        case em: toRender = fmt.Sprintf("<em>%s</em>", t.Text)
        case b: toRender = fmt.Sprintf("<b>%s</b>", t.Text)
        case strong: toRender = fmt.Sprintf("<strong>%s</strong>", t.Text)
        case br: toRender = fmt.Sprintf("<br>%s</br>", t.Text)
        case hr: toRender = fmt.Sprintf("<hr>%s</hr>", t.Text)
        case article: toRender = fmt.Sprintf("<article>%s</article>", t.Text)
        case section: toRender = fmt.Sprintf("<section>%s</section>", t.Text)
        case a: toRender = fmt.Sprintf("<a href=\"%s\">%s</a>", t.TagOptions.Path, t.Text)
        case figure: toRender = fmt.Sprintf("<figure>%s</figure>", t.Text)
        case img: toRender = fmt.Sprintf("<img src=\"%s\"/>", t.TagOptions.Path)
        case hgroup: toRender = fmt.Sprintf("<hgroup>%s</hgroup>", t.Text)
        case li: toRender = fmt.Sprintf("<li>%s</li>", t.Text)
        case head:  toRender = parseHeader(&t)
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

func parseHeader(t *Tag) string {
    date := Date{}
    var tarr []string

    for _, line := range strings.Split(t.Text, "\n") {

        if strings.HasPrefix(line, "date ") {
            dateStr := strings.TrimPrefix(line, "date ")
            timedate := strings.Split(dateStr, "/")

            year, ferr := strconv.Atoi(timedate[0]) 
            if ferr != nil {
                log.Panicln("Date format error: years are wrong")
            }
            date.Year = year

            month, ferr := strconv.Atoi(timedate[1]) 
            if ferr != nil {
                log.Panicln("Date format error: months are wrong")
            }
            date.Month = month


            day, ferr := strconv.Atoi(timedate[2]) 
            if ferr != nil {
                log.Panicln("Date format error: days are wrong")
            }
            date.Day = day
            continue
        }

        if strings.HasPrefix(line, "time") {
            timeStr := strings.TrimPrefix(line, "time ")
            time := strings.Split(timeStr, "/")

            hours, ferr := strconv.Atoi(time[0])
            if ferr != nil {
                log.Panicln("Time format error: Hours are wrong")
            }
            minutes, ferr := strconv.Atoi(time[1])
            if ferr != nil {
                log.Panicln("Time format error: Minutes are wrong")
            }
            date.Hours = hours
            date.Minutes = minutes
            continue
        }


        if strings.HasPrefix(line, "tags") {
            tags := strings.TrimPrefix(line, "tags ")
            tarr = strings.Split(tags, ",")
            continue
        }

    }

    sb := strings.Builder{}

    sb.WriteString("<hgroup>\n")
    sb.WriteString(fmt.Sprintf("<b>%d-%d-%d</b>", date.Year, date.Month, date.Day))
    if date.Hours != 0 && date.Minutes != 0 {
        sb.WriteString(fmt.Sprintf( "%d:%d\n", date.Hours, date.Minutes))
    } else {
        sb.WriteString("\n")
    }

    sb.WriteString("<ul class=\"tag_list\">\n")
    for _, ctag := range tarr {
        sb.WriteString(fmt.Sprintf("<li>%s</li>\n", ctag))
    }
    sb.WriteString("</ul>\n</hgroup>\n")
    return sb.String()
}
