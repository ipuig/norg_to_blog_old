package pages

type BlogPost = struct {
    Title string
    abstract string
}

type HomeData = struct {
    Title string
    About string
    Styles []string
}

type NavigationBar = struct {
}

type Footer = struct {
    Title string
    Links []string
}
