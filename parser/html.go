package parser

import (
	"html/template"
	"regexp"
	"strings"
)

func HTML(fp FSPost) template.HTML {
    for _, newpath := range fp.CSSFiles {
        positions := strings.Split(newpath, "/")
        filename := positions[len(positions) - 1]
        fp.Content = regexp.MustCompile(filename).ReplaceAllString(fp.Content, newpath)
    }

    for _, newpath := range fp.ImagesPath {
        positions := strings.Split(newpath, "/")
        filename := positions[len(positions) - 1]
        fp.Content = regexp.MustCompile(filename).ReplaceAllString(fp.Content, newpath)
    }

    return template.HTML(fp.Content)
}

