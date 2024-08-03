package norg

import (
	e "blog/error"
	"strings"
)

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
            return sb.String(), idx+2, nil
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
