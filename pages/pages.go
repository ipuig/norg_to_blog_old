package pages

import "sort"

type Page struct {
    Title string
    LeftLateralControl LateralControl
    RightLateralControl LateralControl
    PathControl PathControl
    Footer Footer
}

func (p *Page) HasRightLateralControls() bool {
    if p.RightLateralControl.URL != "" {
        return true
    }
    return false
}

func (p *Page) HasLeftLateralControls() bool {
    if p.LeftLateralControl.URL != "" {
        return true
    }
    return false
}

func (p *Page) HasPathControl() bool {
        if p.PathControl.URL != "" {
            return true
        }
    return false
}

func (p *Page) HasFooter() bool {
        if p.Footer.Links != nil {
            return true
        }
    return false
}


func ProcessPosts(posts []Post) map[int][]Post {
    classified := make(map[int][]Post)
    sort.Sort(sort.Reverse(Posts(posts)))

    for idx, post := range posts {

        if idx > 0 {
            previous := posts[idx - 1]
            lc := MakeLateralControl("<", previous.URL(), previous.Page.Title)
            post.Page.LeftLateralControl = lc;
        }

        if idx < (len(posts) - 1) {
            next := posts[idx + 1]
            lc := MakeLateralControl(">", next.URL(), next.Page.Title)
            post.Page.RightLateralControl = lc;

        }

        ps, ok := classified[post.Date.Year]
        if !ok {
            classified[post.Date.Year] = []Post{ post }
            continue
        }

        classified[post.Date.Year] = append(ps, post)
    }
    return classified
} 
