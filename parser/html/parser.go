package parser

import (
    "html/template"
    "regexp"
    "strings"
)

func HTML(body string, styles, images []string) template.HTML {

    for _, newpath := range styles {
        positions := strings.Split(newpath, "/")
        filename := "css/" + positions[len(positions) - 1]
        body = regexp.MustCompile(filename).ReplaceAllString(body, newpath)
    }

    for _, newpath := range images {
        positions := strings.Split(newpath, "/")
        filename := "img/" + positions[len(positions) - 1]
        body = regexp.MustCompile(filename).ReplaceAllString(body, newpath)
    }

    return template.HTML(body)
}
