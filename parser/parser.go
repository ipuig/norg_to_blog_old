package parser

import (
	e "blog/error"
	"blog/pages"
	h "blog/parser/html"
	"blog/parser/norg"
	"html/template"
	"strings"
)

func ParseContent(fpost *FSPost) (pages.Post, error) {
    post := pages.Post{}
    fd := strings.Split(fpost.Filename, ".")
    if len(fd) < 2 {
        return post, e.ParserErrorMissingFileExtension
    }

    var html template.HTML

    switch(fd[1]) {
    case "norg":
        norgParser := norg.Parser{ Content: fpost.Content }
        fpost.Content = norgParser.Parse()
        html = h.HTML(fpost.Content, fpost.CSSFiles, fpost.ImagesPath)

    case "html":
        html = h.HTML(fpost.Content, fpost.CSSFiles, fpost.ImagesPath)

    default:
        return post, e.ParserErrorExtensionNotSupported
    }

    post = pages.Post{
        Page: pages.Page{
            Title: fpost.Title,
        },
        HTML: html,
        Date: pages.DateFromString(fpost.Date),
        PostTags: fpost.Tags,
        AdditionalCSS: fpost.CSSFiles,
        Abstract: fpost.Abstract,
        Logo: fpost.Logo,
    }

    return post, nil
} 

func suportedExtension(ext string) bool {
    if ext == "norg" {
        return true
    }

    if ext == "html" {
        return true
    }

    return false
}
