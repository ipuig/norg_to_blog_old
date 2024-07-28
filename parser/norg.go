package parser

import (
	"fmt"
	"regexp"
	"strings"
)

type NorgParser struct {
    Idx int
    LastIdx int
    Content string
    Builder strings.Builder
}

func removeLeadingWhiteSpace(content string) string {
    re := regexp.MustCompile(`(?m)^ +`)
    return re.ReplaceAllString(content, "")
}

func (np *NorgParser) parseSimple() {
    np.Idx = 0
    np.Content = removeLeadingWhiteSpace(np.Content)
    np.LastIdx = len(np.Content) - 1
    np.Builder = strings.Builder{}

    for np.Idx <= np.LastIdx {

        if np.Content[np.Idx] == '*' && np.HasNext() && np.Content[np.Idx + 1] == '*' {
            np.Builder.WriteString(np.parseHeading())
            continue
        }

        if np.Content[np.Idx] == '*' && np.HasNext() {
            np.Builder.WriteString(np.parseBold())
            continue
        }

        if np.Content[np.Idx] == '_' && np.HasNext() {
            np.Builder.WriteString(np.parseUnderline())
            continue
        }

        if np.Content[np.Idx] == '`' && np.HasNext() {
            np.Builder.WriteString(np.parseSpan())
            continue
        }

        if np.Content[np.Idx] == '.' && np.HasNext() {
            np.Builder.WriteString(np.parseImage())
            continue
        }

        np.Builder.WriteByte(np.Content[np.Idx])
        np.Idx++
    }
}

func (np *NorgParser) Parse() string {
    np.parseSimple()
    np.Content = np.Builder.String()
    np.Idx = 0
    np.LastIdx = len(np.Content) - 1
    np.Builder.Reset()
    np.parseList()
    np.parseParagraph()
    return np.Content
}



func (np *NorgParser) HasNext() bool {
    return np.Idx < np.LastIdx
}

func (np *NorgParser) parseHeading() string {
    start := np.Idx
    for np.Content[np.Idx] == '*' {
        np.Idx++
    }
    level := np.Idx - start

    start = np.Idx
    for np.Content[np.Idx] != '\n' {
        np.Idx++
    }
    heading := np.Content[start:np.Idx]
    np.Idx++
    return fmt.Sprintf("\n<h%d>%s</h%d>\n\n", level, heading, level)
}

func (np *NorgParser) parseBold() string {
    np.Idx++
    start := np.Idx
    for np.Content[np.Idx] != '*' {
        np.Idx++
    }
    boldText := np.Content[start:np.Idx]
    np.Idx++
    return "<strong>" + boldText + "</strong>"
}

func (np *NorgParser) parseUnderline() string {
    np.Idx++
    start := np.Idx
    for np.Content[np.Idx] != '_' {
        np.Idx++
    }
    boldText := np.Content[start:np.Idx]
    np.Idx++
    return "<u>" + boldText + "</u>"
}

func (np *NorgParser) parseSpan() string {
    np.Idx++
    start := np.Idx
    for np.Content[np.Idx] != '`' {
        np.Idx++
    }
    boldText := np.Content[start:np.Idx]
    np.Idx++
    return `<span class="fmt">` + boldText + "</span>"
}

func (np *NorgParser) parseImage() string {
    np.Idx++

    if (np.Idx + 6 ) < np.LastIdx && np.Content[np.Idx:np.Idx+6] == "image " {
        np.Idx += 6
    } else {
        return "."
    }

    path := strings.Builder{}
    for np.Content[np.Idx] != '\n' {
        path.WriteByte(np.Content[np.Idx])
        np.Idx++
    }
    np.Idx++
    return "\n<img src=\"" + path.String() + "\">\n"
}

func (np *NorgParser) parseList() {
    re := regexp.MustCompile(`(?m)(^- .*)(\n^[^-\n].*)*`)
    matches := re.FindAllString(np.Content, -1)

    for _, match := range matches {
        wrapped := fmt.Sprintf("<li>%s</li>", match[2:])
        np.Content = strings.ReplaceAll(np.Content, match, wrapped)
    }

    re = regexp.MustCompile(`(?m)(^<li>.*)(\n^[^\n].*)*</li>`)
    matches = re.FindAllString(np.Content, -1)
    for _, match := range matches {
        wrapped := fmt.Sprintf("\n<ul>\n<l%s\n</ul>\n", match[2:])
        fmt.Println(wrapped)
        np.Content = strings.ReplaceAll(np.Content, match, wrapped)
    }
}

func (np *NorgParser) parseParagraph() {
    for _, paragraph := range strings.Split(np.Content, "\n\n") {
        p := strings.TrimSpace(paragraph)
        if len(p) == 0 || strings.HasPrefix(p, "<") {
            continue;
        }
        np.Content = strings.ReplaceAll(np.Content, p, fmt.Sprintf("<p>%s</p>\n", p))
    }
}
