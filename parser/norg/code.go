package norg

import "strings"

type CodeBlockParser struct {
    Content string
    Language string
}

// RenderMeta implements Meta.
func (cbp CodeBlockParser) RenderMeta() string {
    panic("unimplemented")
}

func (cbp *CodeBlockParser) findLanguage() error {
    sb := &strings.Builder{}
    for idx, c := range wp {

    }
}

type LanguageConfiguration struct {
}
