package parser

import (
    "fmt"
    "os"
    "strings"
    "blog/config"
)


func FetchPostPaths() []string {
    contentLocationPath := config.SiteConfig.ContentPath
    paths := make([]string, 0)

    dirs, err := os.ReadDir(contentLocationPath)
    if err != nil {
        return nil
    }


    for _, dir := range dirs {
        if dir.Name() == ".git" {
            continue
        }
        files, err := os.ReadDir(contentLocationPath + "/" + dir.Name())
        if err != nil {
            fmt.Println("failed to read posts at", contentLocationPath + "/" + dir.Name())
            continue
        }

        for _, file := range files {
            if file.IsDir() {
                paths = append(paths, contentLocationPath + "/" + dir.Name() + "/" + file.Name())
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
    Content string
    Title string
    Date string
    ReleaseDate string
    Tags []string
    Logo string
    Abstract string
}

func (fp *FSPost) Metadata() {
    if fp.Title != "" {
        return
    }

    lines := strings.Split(fp.Content, "\n")
    for idx, line := range lines {

        if strings.HasPrefix(line, "logo") {
            fp.Logo = extractMetadata("logo:", line)
            continue
        }

        if strings.HasPrefix(line, "release_date") {
            fp.ReleaseDate = extractMetadata("release_date:", line)
            continue
        }

        if strings.HasPrefix(line, "abstract") {
            fp.Abstract = extractMetadata("abstract:", line)
            continue
        }

        if strings.HasPrefix(line, "title:") {
            fp.Title = extractMetadata("title:", line)
            continue
        }

        if strings.HasPrefix(line, "date:") {
            fp.Date = extractMetadata("date:", line)
            continue
        }

        if strings.HasPrefix(line, "tags:") {
            tags := extractMetadata("tags:", line)
            fp.Tags = strings.Split(tags, ",")
            continue
        }

        if strings.HasPrefix(line, "@end") {
            fp.Content = strings.Join(lines[idx+1:], "\n")
            break;
        }
    }
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
    fpost.Content = string(content)
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
