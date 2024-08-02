package norg

type MetadataParser struct {
    Content string
    Meta Meta 
}

type Meta interface {
    RenderMeta() string
}

func (mp *MetadataParser) findType() {
    n := len(mp.Content)
    if n >= 4 && mp.Content[0:4] == "code" {
        mp.Meta = CodeBlockParser{
            Content: mp.Content[5:], // skip the aswell whitespace
        }
    }
}

func (mp *MetadataParser) Process() {
    mp.Meta.RenderMeta()
}
