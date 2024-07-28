package parser

import (
    "fmt"
    "os"
    "strings"
)

const contentLocationPath = "content/posts"

func FetchPostPaths() []string {
    years := []string { "2024" }
    paths := make([]string, 0)
    for _, year := range years {
        files, err := os.ReadDir(contentLocationPath + "/" + year)
        if err != nil {
            fmt.Println("failed to read posts at", contentLocationPath + "/" + year)
            continue
        }

        for _, file := range files {
            if file.IsDir() {
                paths = append(paths, contentLocationPath + "/" + year + "/" + file.Name())
            }
        }
    }
    return paths
}

type FSPost struct {
    RootPath string
    CSSFiles []string
    ImagesPath []string
    Filename string
    Content []byte
}

func (fp FSPost) Metadata() (string, string) {
    content := string(fp.Content)
    title := ""
    date := ""

    for _, line := range strings.Split(content, "\n") {
        if strings.HasPrefix(line, "title:") {
            title = extractMetadata("title:", line)
            continue
        }

        if strings.HasPrefix(line, "date:") {
            date = extractMetadata("date:", line)
        }

        if line != "" && date != "" {
            break
        }
    }
    return title, date
}

func extractMetadata(prefix, line string) string {
    values := strings.Split(line, prefix)
    if len(values) > 1 {
        return strings.TrimSpace(values[1])
    }
    return ""
}

func PostFromPath(path string) FSPost {
    fs, err := os.ReadDir(path)
    if err != nil {
        panic(err)
    }

    fpost := FSPost{
        RootPath: path,
    }


    for _, file := range fs {
        if !file.IsDir() {
            fpost.Filename = file.Name()
            continue
        }

        if file.Name() == "img" {
            fpost.ImagesPath = listFile(path + "/img")
        }

        if file.Name() == "css" {
            fpost.CSSFiles = listFile(path + "/css")
        }
    }
    content, err := os.ReadFile(fpost.RootPath + "/" + fpost.Filename)
    if err != nil {
        panic(err)
    }
    fpost.Content = content
    return fpost
}

func listFile(dir string) []string {
    fs, err := os.ReadDir(dir)
    if err != nil {
        return nil
    }

    files := make([]string, len(fs))
    for _, f := range fs {
        files = append(files, f.Name())
    }
    return files
}
