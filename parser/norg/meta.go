package norg

import (
	e "blog/error"
)

type MetadataParser struct {
    Content string
    Meta Meta 
    Ptr *int
}

type Meta interface {
    RenderMeta() (string, string, error)
}

func (mp *MetadataParser) findType() error {
    n := len(mp.Content)
    if n >= 4 && mp.Content[0:4] == "code" {
        mp.Meta = CodeBlockParser{
            Content: mp.Content[5:], // skip also the whitespace
        }
        return nil
    }
    return e.ParserErrorMetadataUnknownType
}

func (mp *MetadataParser) Process() (string, error) {
    err := mp.findType()
    if err != nil {
        return "", err
    }

    parsedMeta, remaining, err := mp.Meta.RenderMeta()
    if err != nil {
        return "", err
    }

    progress := len(mp.Content) - len(remaining)
    *mp.Ptr += progress

    return parsedMeta, nil
}
