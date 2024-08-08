package norg

import (
    "fmt"
    "strings"
    e "blog/error"
)

type TableParser struct {
    Content string
    WithHeader bool

}

func (tp TableParser) RenderMeta() (string, string, error) {
    lines := strings.Split(tp.Content, "\n")
    if lines[0] == "header" {
        tp.WithHeader = true
    }

    block, eIdx, err := tp.makeTable(lines[1:])
    if err != nil {
        return "", "", err
    }

    return block, strings.Join(lines[eIdx:], "\n"), nil
}

func (tp *TableParser) makeTable(rows []string) (string, int, error) {
    sb := &strings.Builder{}
    sb.WriteString("<table>\n")
    htype := "td"
    if tp.WithHeader {
        sb.WriteString("<thead>\n")
        htype = "th"
    }

    headings, err := makeCols(rows[1])
    if err != nil {
        panic(err)
    }

    sb.WriteString("<tr>\n")
    for _, heading := range headings {
        sb.WriteString(fmt.Sprintf("<%s>%s</%s>\n", htype, heading, htype))
    }
    sb.WriteString("</tr>\n")

    if tp.WithHeader {
        sb.WriteString("</thead>\n")
    }

    sb.WriteString("<tbody>\n")
    for idx, row := range rows[2:] {

        if (strings.Contains(row, "@end")) {
            sb.WriteString("<tbody>\n</table>\n")
            return sb.String(), idx + 4, nil; 
        }
        cols, err := makeCols(row)
        if err != nil {
            continue
        }
        sb.WriteString("<tr>\n")

        for _, col := range cols {
            sb.WriteString("<td>" + col + "</td>\n")
        }
        sb.WriteString("</tr>\n")
    }

    return sb.String(), 0, nil
}

func makeCols(line string) ([]string, error) {
    elements := strings.Split(line, "|")
    if len(elements) >= 3 {
        return elements[1:len(elements) - 1], nil
    }
    return nil, e.NorgParserErrorMetadataUnknownType
}

