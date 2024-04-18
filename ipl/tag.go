package ipl

import (
	"fmt"
	"html/template"
	"strconv"
	"strings"
)

const (
    p = iota; span; em;
    b; br; hr; h1; h2; h3;
    h4; h5; h6; img; strong;
    a; li; div; hgroup; head;
    figure; aside;
    article; section;
    st
)

type Tag interface {
    String() string
    Render() template.HTML
    Container() bool
}

type TagOptions struct {
    Path string
    X int
    Y int
    Size float32
    Classes []string
    InUse bool
}

type TagHTML struct {
    TagType int
    Text string
    Children []*TagHTML
    TagOptions
}

func (t *TagHTML) Render() template.HTML {
    return template.HTML(t.String())
}

func (t *TagHTML) Container() bool {
    if t.TagType >= div {
        return true
    }
    return false
}

func (t *TagHTML) String() string {
    if t.TagType == head {
        return parseHeader(t)
    }

    if t.Container() {
        return wrapContainer(t)
    }

    return wrapSimpleTag(t.TagType, t.Text, t.TagOptions)
}

func wrapSimpleTag(tag int, content string, opts TagOptions) string {
    var tname string
    sb := strings.Builder{}

    switch tag {
    case h1: tname = "h1"
    case h2: tname = "h2"
    case h3: tname = "h3"
    case h4: tname = "h4"
    case h5: tname = "h5"
    case h6: tname = "h6"
    case p: tname = "p"
    case b: tname = "b"
    case strong: tname = "strong"
    case em: tname = "em"
    case li: tname = "li"
    case img: tname = "img" 
    case br: tname = "br"
    case hr: tname = "hr"
    }

    sb.WriteString("<" + tname)

    if opts.InUse {

        switch {
        case opts.Classes != nil:
            sb.WriteString(" class=\"")
            for _, c := range opts.Classes {
                sb.WriteString(" " + c)
            }
            sb.WriteString("\"")
        case opts.Path != "":
            if tag == img {
                sb.WriteString(" src=\"" + opts.Path + "\"")
            } else {
                sb.WriteString(" href=\"" + opts.Path + "\"")
            }
    // TODO: X, Y and Size
        }
    }

    if tag == img || tag == br || tag == hr {
        sb.WriteString("/>")
    } else {
        sb.WriteString(">" + strings.TrimSpace(content) + "</" + tname + ">")
    }

    return sb.String()
}

func wrapContainer(tag *TagHTML) string {
    var tname string
    sb := strings.Builder{}

    switch tag.TagType {
    case div: tname = "div"
    case hgroup: tname = "hgroup"
    case article: tname = "article"
    case section: tname = "section"
    case figure: tname = "figure"
    case aside: tname = "aside"
    case head: tname = "hgroup"
    case st: tname = "section"
    }

    sb.WriteString("<" + tname)

    if tag.TagOptions.Classes != nil {
        sb.WriteString(" class=\"")
        for _, c := range tag.TagOptions.Classes {
            sb.WriteString(" " + c)
        }
        sb.WriteString("\"")
    }

    sb.WriteString(">\n")

    if tag.TagType == st {
        sb.WriteString("<h3>" + strings.TrimSpace(tag.Text) + "</h3>\n")
    }

    if tag.Children != nil {
        for _, child := range tag.Children {
            sb.WriteString(child.String() + "\n")
        }
    }

    sb.WriteString("</" + tname + ">\n")
    return sb.String()
}

type Date struct {
    Year int
    Month int
    Day int
    Hours int
    Minutes int
}

func (d *Date) Date() string {
    return fmt.Sprintf("%d-%d-%d", d.Year, d.Month, d.Day)
}

func (d *Date) DateTime() string {
    return fmt.Sprintf("%d-%d-%d | %d:%d", d.Year, d.Month, d.Day, d.Hours, d.Minutes)
}

func parseHeader(t *TagHTML) string {
    date := Date{}
    var tarr []string

    for _, line := range strings.Split(t.Text, "\n") {
        if strings.HasPrefix(line, "date ") {
            dateStr := strings.TrimPrefix(line, "date ")
            timedate := strings.Split(dateStr, "/")

            day, ferr := strconv.Atoi(timedate[0]) 
            if ferr != nil {
                return "Date format error: days are wrong"
            }
            date.Day = day

            month, ferr := strconv.Atoi(timedate[1]) 
            if ferr != nil {
                return "Date format error: months are wrong"
            }
            date.Month = month


            year, ferr := strconv.Atoi(timedate[2][:len(timedate[2])-1]) 
            if ferr != nil {
                return "Date format error: years are wrong "
            }
            date.Year = year
            continue
        }

        if strings.HasPrefix(line, "time") {
            timeStr := strings.TrimPrefix(line, "time ")
            time := strings.Split(timeStr, ":")

            hours, ferr := strconv.Atoi(time[0])
            if ferr != nil {
                return "Time format error: Hours are wrong"
            }
            minutes, ferr := strconv.Atoi(time[1][:len(time[1]) - 1])
            if ferr != nil {
                return "Time format error: Minutes are wrong"
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
    sb.WriteString(fmt.Sprintf("<b>%d-%d-%d</b>  ", date.Year, date.Month, date.Day))
    if date.Hours != 0 && date.Minutes != 0 {
        sb.WriteString(fmt.Sprintf( "%d:%d\n", date.Hours, date.Minutes))
    } else {
        sb.WriteString("\n")
    }

    sb.WriteString("<ul class=\"tag_list\">\n")
    for _, ctag := range tarr {
        sb.WriteString(fmt.Sprintf("<li>%s</li>\n", ctag))
    }
    sb.WriteString("</ul>\n</hgroup>")
    return sb.String()
}
