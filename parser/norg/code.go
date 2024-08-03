package norg

import (
	e "blog/error"
	"fmt"
	"regexp"
	"strings"
)


type CodeRegex struct {
    Compiled bool
    Expr *regexp.Regexp
}

var bashKeywords []string = []string{
    "curl", "grep", "awk", "ls", "echo", "sudo", "nmap", "tee", "ffuf",
}

var BashRegex CodeRegex

type CodeBlockParser struct {
    Content string
    Language string
}

func (cbp CodeBlockParser) RenderMeta() (string, string, error) {
    lines := strings.Split(cbp.Content, "\n")
    err := cbp.findLanguage(lines[0])
    if err != nil {
        return "", cbp.Content, err
    }

    block, eIdx, err := cbp.makeCodeBlock(lines[1:])
    if err != nil {
        return "", "", err
    }

    return block, strings.Join(lines[eIdx:], "\n"), nil
}

func (cbp *CodeBlockParser) makeCodeBlock(lines []string) (string, int, error) {
    sb := &strings.Builder{}
    sb.WriteString("\n\n<code class=\"" + cbp.Language + "\">\n")

    for idx, line := range lines {
        if (strings.Contains(line, "@end")) {
            sb.WriteString("</code>\n\n\n")
            return applyStyles(sb.String(), cbp.Language), idx+2, nil
        }
        sb.WriteString(line + "\n")
    }
    return "", 0, e.ParserErrorMetadataMissingEnd
}

func (cbp *CodeBlockParser) findLanguage(line string) error {
    languageInfo := strings.Split(line, " ")

    if len(languageInfo) == 0 {
        return e.ParserErrorInvalidMetadata
    }

    cbp.Language = languageInfo[0]
    // TODO: handle modifiers after language
    return nil
}

func applyStyles(block, lang string) string{

    if lang == "bash" {

        if !BashRegex.Compiled {
            BashRegex.Compiled = true
            expr := fmt.Sprintf("(%s)", strings.Join(bashKeywords, "|"))
            BashRegex.Expr = regexp.MustCompile(expr)
        }

        visited := make(map[string]struct{})
        for _, matches := range BashRegex.Expr.FindAllStringSubmatch(block, -1) {
            for _, match := range matches {
                _, ok := visited[match]
                if !ok {
                    visited[match] = struct{}{}
                    block = strings.ReplaceAll(block, match, "<span class=\"code_keyword\">" + match + "</span>")
                }
            }
        }
    }

    return block
}
