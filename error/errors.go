package error

import "errors"

var (
    ExtensionNotSupported  = errors.New("The extension used to write the blog is not supported")
    MissingFileExtension  = errors.New("Missing extension to parse the post")
)
