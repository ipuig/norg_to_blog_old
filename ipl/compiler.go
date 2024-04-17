package ipl

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func AnalyseContent() {
    dir, err := os.ReadDir("content/posts")

    if err != nil {
        os.MkdirAll("content/posts", 0755)
    }

    for _, f := range dir {
        if f.Type().IsRegular() {
            ProcessFile("content/posts/" + f.Name())
        }
    }
}


type LineAnalyser struct {
    head *Node
    order int
    processed map[int]Tag
}

type Node struct {
    val string
    prev *Node
    order int
    accStr []byte
}

func (la *LineAnalyser) load(tagType int, open bool) {

    if !open && la.head != nil{
        la.consume(tagType)
        return
    }


    node := Node {
        prev: la.head,
        order: la.order,
        accStr: []byte{},
    }

    la.head = &node
    la.order++
}

func (la *LineAnalyser) consume(tagType int) {
    tag := Tag{
        TagType: tagType,
        Text: string(la.head.accStr[:len(la.head.accStr) - 1]),
        Children: []Tag{} ,
    }

    if tagType >= hgroup {
        for id, st := range la.processed {
            if id != la.head.order {
                tag.Children = append(tag.Children, st)
                delete(la.processed, id)
            }
        }
    }

    la.processed[la.head.order] = tag
    la.head = la.head.prev
}

func (la *LineAnalyser) addToTagBuffer(word []byte) {
    if la.head == nil {
        return
    }
    if len(la.head.accStr) > 0 && la.head.accStr[len(la.head.accStr)-1] != ' ' && la.head.accStr[len(la.head.accStr)-1] != '\n' {
        la.head.accStr = append(la.head.accStr, ' ')
    }
    la.head.accStr = append(la.head.accStr, word...)
}

func ProcessFile(path string) (Post, error) {
    state := LineAnalyser{
        head: nil,
        order: 0,
        processed: make(map[int]Tag),
    }

    filecontent, err := os.ReadFile(path)
    post := Post{}

    if err != nil {
        log.Printf("couldn't read the file %s\n", path)
        return post, err
    }


    lines := bytes.Split(filecontent, []byte{'\n'})

    for _, line := range lines {
        processLine(line, &state)

        if state.head != nil && len(state.head.accStr) >= 1 {
            state.head.accStr = append(state.head.accStr, '\n')
        }
    }

    for _, p := range state.processed {
        fmt.Println(p.String())
    }
    return post, nil
}

func processLine(line []byte, state *LineAnalyser) {
    for _, word := range bytes.Split(line, []byte { ' ' }) {
        processWord(word, state)
    }
}

func processWord(word []byte, state *LineAnalyser) {
    l := len(word)
    if l <= 2 {
        state.addToTagBuffer(word)
        return
    }
    if word[0] == ':' && word[1] == ':' {
        processCommand(word[2:], true, state)
        return
    }
    if word[l - 1] == ':' && word[l - 2] == ':' {
        processCommand(word[:l-2], false, state)
        return
    }
    state.addToTagBuffer(word)
}

func processCommand(command []byte, open bool, state *LineAnalyser) {
    switch string(command) {
    case "mt": state.load(h2, open)
    case "p": state.load(p, open)
    case "head": state.load(head, open)
    case "s": state.load(section, open)
    }
}
