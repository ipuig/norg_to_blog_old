package parser

import (
	e "blog/error"
	"strings"
)

func ParseContent(fpost *FSPost) error {
    f := strings.Split(fpost.Filename, ".")
    if len(f) < 2 {
        return e.MissingFileExtension
    }

    switch f[1] {
    case "html": return nil
    case "norg": 
        np := NorgParser{ Content: fpost.Content }
        fpost.Content = np.Parse()
        return nil

    default: return e.ExtensionNotSupported
    }
} 
