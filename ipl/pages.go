package ipl

type HomeData = struct {
    Title string
    About string
    Styles []string
    Posts *map[int][]Post
}

type Footer = struct {
    Title string
    Links []string
}

func Example() Post {
    post := Post{
        Date: &Date{
            Year: 2024,
            Month: 4,
            Day: 16,
            Hours: 15,
            Minutes: 13,
        },
        PostTitle: "Exemple 1",
        Tags: []string{"test", "pomiscles"},
        Sections: []Section {
            {
                SectionTitle: "Primera seccio",
                Tags: tagsToHtml([]Tag{
                    {
                        TagType: p, 
                        Text: "Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat.",
                    },
                    {
                        TagType: p, 
                        Text: "Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat.",
                    },
                    {
                        TagType: h6, 
                        Text: "PROUBA",
                    },
                }),
            },
            {
                SectionTitle: "PONSIAS seccio",
                Tags: tagsToHtml([]Tag{
                    {
                        TagType: p, 
                        Text: "Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat.",
                    },
                    {
                        TagType: p, 
                        Text: "Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat.",
                    },
                    {
                        TagType: h6, 
                        Text: "PROUBA",
                    },
                }),
            },
            {
                SectionTitle: "Primera seccio",
                Tags: tagsToHtml([]Tag{
                    {
                        TagType: p, 
                        Text: "Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat.",
                    },
                    {
                        TagType: p, 
                        Text: "Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat.",
                    },
                    {
                        TagType: h6, 
                        Text: "PROUBA",
                    },
                }),
            },
            {
                SectionTitle: "Primera seccio",
                Tags: tagsToHtml([]Tag{
                    {
                        TagType: p, 
                        Text: "Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat.",
                    },
                    {
                        TagType: p, 
                        Text: "Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat.",
                    },
                    {
                        TagType: h6, 
                        Text: "PROUBA",
                    },
                }),
            },
        },
    }
    return post
}
