package pages

import (
	"fmt"
	"html/template"
	"strings"
)

const (
    LateralControlPath = "layouts/components/lateral_control.html"
    PathControlPath = "layouts/components/path_control.html"
)
type LateralControl struct {
    Symbol string
    MessageOnHover string
    URL string
}
func MakeLateralControl(symbol, url, message string) LateralControl {
    return LateralControl{ Symbol: symbol, URL: url, MessageOnHover: message}
}

type PathControl struct {
    URL string
}
func MakePathControl(url string) PathControl {
    return PathControl{ url }
}

func (pc PathControl) BuildPath() template.HTML {
    html := &strings.Builder{}
    lpath := &strings.Builder{}
    for idx, path := range strings.Split(pc.URL, "/") {
        lpath.WriteString(path)
        if idx > 0 {
            html.WriteString("/")
        }
        html.WriteString(fmt.Sprintf(`<a href="%s">%s</a>`, lpath.String(), path))
        lpath.WriteString("/")
    }
    return template.HTML(html.String())
}
