package pages

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
