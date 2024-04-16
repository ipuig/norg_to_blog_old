package ipl

import "errors"

type TagReader struct {
    head *TagNode
    table map[string]int
}

type TagNode struct {
    tag *Tag
    prev *TagNode
}

func (tr *TagReader) Read(t Tag) error {
    v, ok := tr.table[t.name]

    if !ok && !t.open {
        return errors.New("Found %s closing tag before any opening tag")
    }

    if !ok {
        tr.table[t.name] = 1
        return nil
    }

    if t.open {
        tr.table[t.name] = v + 1
        tnode := &TagNode{tag: &t, prev: tr.head}
        tr.head = tnode
    }

    if v > 1 {
        tr.table[t.name] = v - 1
    } else {
        delete(tr.table, t.name)
    }
    
    tr.head = tr.head.prev
    return nil
} 

type Tag struct {
    name string
    open bool
    opts string
    content string
}



