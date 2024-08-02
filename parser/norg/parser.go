package norg

import (
	"fmt"
	"regexp"
	"strings"
)

type Parser struct {
    Idx int
    LastIdx int
    Content string
    Builder strings.Builder
}

func removeLeadingWhiteSpace(content string) string {
    re := regexp.MustCompile(`(?m)^ +`)
    return re.ReplaceAllString(content, "")
}

func (np *Parser) parseSimple() {
    np.Idx = 0
    np.Content = removeLeadingWhiteSpace(np.Content)
    np.LastIdx = len(np.Content) - 1
    np.Builder = strings.Builder{}

    for np.Idx <= np.LastIdx {

        if np.Content[np.Idx] == '*' && np.HasNext() && np.Content[np.Idx + 1] == '*' {
            np.parseHeading()
            continue
        }

        if np.Content[np.Idx] == '*' && np.HasNext() {
            np.parseBold()
            continue
        }

        if np.Content[np.Idx] == '_' && np.HasNext() {
            np.parseUnderline()
            continue
        }

        if np.Content[np.Idx] == '{' && np.HasNext() {
            np.parseLinkLeft()
            continue
        }
        
        if np.Content[np.Idx] == '[' && np.HasNext() {
            np.parseLinkRight()
            continue
        }

        if np.Content[np.Idx] == '`' && np.HasNext() {
            np.parseSpan()
            continue
        }

        if np.Content[np.Idx] == '.' && np.HasNext() {
            np.parseImage()
            continue
        }

        // if np.Content[np.Idx] == '@' && np.HasNext() {
        //     // np.Builder.WriteString(np.parseCode())
        //     continue
        // }

        np.Builder.WriteByte(np.Content[np.Idx])
        np.Idx++
    }
}

func (np *Parser) Parse() string {
    np.parseSimple()
    np.Content = np.Builder.String()
    np.Idx = 0
    np.LastIdx = len(np.Content) - 1
    np.Builder.Reset()
    np.parseList()
    np.parseParagraph()
    return "<article>\n" + np.Content + "\n</article>"
}

func (np *Parser) HasNext() bool {
    return np.Idx < np.LastIdx
}

func (np *Parser) parseHeading() {
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
    np.Builder.WriteString(fmt.Sprintf("\n<h%d>%s</h%d>\n\n", level, heading, level))
}

func (np *Parser) parseBold() {
    np.Idx++
    start := np.Idx
    for np.Content[np.Idx] != '*' {
        np.Idx++
    }
    boldText := np.Content[start:np.Idx]
    np.Idx++
    np.Builder.WriteString("<strong>" + boldText + "</strong>")
}

func (np *Parser) parseUnderline() {
    np.Idx++
    start := np.Idx
    for np.Content[np.Idx] != '_' {
        np.Idx++
    }
    boldText := np.Content[start:np.Idx]
    np.Idx++
    np.Builder.WriteString("<u>" + boldText + "</u>")
}

func (np *Parser) parseSpan() {
    np.Idx++
    start := np.Idx
    for np.Content[np.Idx] != '`' {
        np.Idx++
    }
    boldText := np.Content[start:np.Idx]
    np.Idx++
    np.Builder.WriteString(`<span class="fmt">` + boldText + "</span>")
}

func (np *Parser) parseImage() {
    np.Idx++

    if (np.Idx + 6 ) < np.LastIdx && np.Content[np.Idx:np.Idx+6] == "image " {
        np.Idx += 6
    } else {
        np.Builder.WriteByte('.');
        return 
    }

    path := strings.Builder{}
    for np.Content[np.Idx] != '\n' {
        path.WriteByte(np.Content[np.Idx])
        np.Idx++
    }
    np.Idx++
    np.Builder.WriteString("\n<img src=\"" + path.String() + "\">\n")
}

func (np *Parser) parseLinkRight() {
    np.Idx++
    start := np.Idx
    for np.Content[np.Idx] != ']' && np.HasNext() {
        np.Idx++
    }

    if np.Content[np.Idx] != ']' {
        np.Idx = start
        np.Builder.WriteByte('[')
        return
    }

    alias := np.Content[start:np.Idx]
    np.Idx++

    if np.Content[np.Idx] != '{' {
        np.Builder.WriteString("<u>" + alias + "</u>")
        return
    }
    
    np.Idx++
    start = np.Idx
    for np.Content[np.Idx] != '}' && np.HasNext() {
        np.Idx++;
    }
    link := np.Content[start:np.Idx]
    np.Idx++

    np.Builder.WriteString(fmt.Sprintf(`<a href="%s">%s</a>`, link, alias))
}


func (np *Parser) parseLinkLeft() {
    np.Idx++
    start := np.Idx
    for np.Content[np.Idx] != '}' && np.HasNext() {
        np.Idx++
    }

    if np.Content[np.Idx] != '}' {
        np.Idx = start
        np.Builder.WriteByte('{')
        return
    }

    link := np.Content[start:np.Idx]
    np.Idx++

    if np.Content[np.Idx] != '[' {
        np.Builder.WriteString(fmt.Sprintf(`<a href="%s">%s</a>`, link, link))
        return
    }
    
    np.Idx++
    start = np.Idx
    for np.Content[np.Idx] != ']' && np.HasNext() {
        np.Idx++;
    }
    alias := np.Content[start:np.Idx]
    np.Idx++

    np.Builder.WriteString(fmt.Sprintf(`<a href="%s">%s</a>`, link, alias))
}

func (np *Parser) parseMeta() {
    md := MetadataParser{
        Content: np.Content,
    }
}

func (np *Parser) parseList() {
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
        np.Content = strings.ReplaceAll(np.Content, match, wrapped)
    }
}

func (np *Parser) parseParagraph() {
    for _, paragraph := range strings.Split(np.Content, "\n\n") {
        p := strings.TrimSpace(paragraph)
        if len(p) == 0 || strings.HasPrefix(p, "<") {
            continue;
        }
        np.Content = strings.ReplaceAll(np.Content, p, fmt.Sprintf("<p>%s</p>\n", p))
    }
}
