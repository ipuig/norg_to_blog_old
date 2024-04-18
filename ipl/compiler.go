package ipl

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"sort"
)

func AnalyseContent() Posts {

    dir, err := os.ReadDir("content/posts")

    if err != nil {
        os.MkdirAll("content/posts", 0755)
        log.Println("Place your posts at \"content/posts/\"")
        return nil
    }

    posts := make(Posts, 0, 10)

    for _, f := range dir {
        if f.Type().IsRegular() {

            tags, err := ProcessFile("content/posts/" + f.Name())

            if err != nil {
                continue // ignore invalid files TODO: add warning
            }
            post := &Post{ Elements: tags}
            post.Generate()
            posts = append(posts, post)
        }
    }

    if posts.Len() > 1 {
        sort.Sort(posts)
    }

    return posts
}

type FileState struct {
    Tags []*TagHTML
    Head *ByteNode
}

type ByteNode struct {
    Prev *ByteNode
    Nested []*ByteNode
    Content []byte
    Tag *TagHTML
}

func (bn *ByteNode) Container() bool {
    return bn.Tag.Container()
}

func (bn *ByteNode) Compile() *TagHTML {
    bn.Tag.Text = string(bytes.TrimSpace(bn.Content))
    return bn.Tag
}

func (fs *FileState) addToContent(stream []byte) {
    if fs.Head == nil {
        return
    }

    if fs.Head.Container() && len(fs.Head.Nested) > 0 {
        fs.Head.Nested[len(fs.Head.Nested)-1].Content = append(fs.Head.Nested[len(fs.Head.Nested)-1].Content, stream...)
        return
    }

    fs.Head.Content = append(fs.Head.Content, stream...)
}

func (fs *FileState) Consume(tag int) {
    if fs.Head == nil {
        fmt.Println("error unexpected closing")
        return
    }

    if fs.Head.Container() {
        // closing container
        if fs.Head.Tag.TagType == tag {
            fs.Tags = append(fs.Tags, fs.Head.Compile())
            fs.Head = fs.Head.Prev
            return
        }

        // element within container
        if len(fs.Head.Nested) >= 1 {
            child := fs.Head.Nested[len(fs.Head.Nested)-1].Compile()
            fs.Head.Tag.Children = append(fs.Head.Tag.Children, child)
            fs.Head.Nested = fs.Head.Nested[:len(fs.Head.Nested) - 1]
            return
        }
    }

    // simple tags
    fs.Tags = append(fs.Tags, fs.Head.Compile())
    fs.Head = fs.Head.Prev
}

func (fs *FileState) AppendChild(tag *TagHTML) {
    bn := ByteNode{
        Tag: tag,
        Prev: fs.Head,
        Content: make([]byte, 0),
    }
    fs.Head.Nested = append(fs.Head.Nested, &bn)
}

func (fs *FileState) Push(tag int) {
    t := TagHTML{ TagType: tag}

    if fs.Head != nil && fs.Head.Container() {
        fs.AppendChild(&t)
        return
    }

    bn := ByteNode{
        Tag: &t,
        Prev: fs.Head,
        Content: make([]byte, 0),
    }

    if bn.Container() {
        bn.Nested = make([]*ByteNode, 0)
    }

    fs.Head = &bn
}

func (fs *FileState) Load(tag int, open bool) {
    if !open {
        fs.Consume(tag)
        return
    }
    fs.Push(tag)
}


func ProcessFile(path string) ([]*TagHTML, error) {
    fstate := FileState{
        Tags: make([]*TagHTML, 0),
    }

    fbytes, err := os.ReadFile(path)
    if err != nil {
        return nil, nil
    }

    for _, line := range bytes.Split(fbytes, []byte("\n")) {
        processLine(line, &fstate)
        fstate.addToContent([]byte("\n"))
    }

    return fstate.Tags, nil
}

func processLine(line []byte, fstate *FileState) {
    words := bytes.Split(line, []byte(" "))
    for _, word := range words {
        processWord(word, fstate)
        fstate.addToContent([]byte(" "))
    }
}

func processWord(word []byte, fstate *FileState) {
    l := len(word)
    if l <= 2 {
        fstate.addToContent(word)
        return
    }
    if word[0] == ':' && word[1] == ':' {
        processCommand(word[2:], true, fstate)
        return
    }
    if word[l - 1] == ':' && word[l - 2] == ':' {
        processCommand(word[:l-2], false, fstate)
        return
    }
    fstate.addToContent(word)
}

func processCommand(command []byte, open bool, fstate *FileState) {
    switch string(command) {
    case "mt": fstate.Load(h1, open)
    case "p": fstate.Load(p, open)
    case "head": fstate.Load(head, open)
    case "s": fstate.Load(st, open)
    case "abs": fstate.Load(abs, open)
    }
}
