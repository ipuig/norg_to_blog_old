package error

import "errors"

var (
    ExtensionNotSupported  = errors.New("The extension used to write the blog is not supported")
    MissingFileExtension  = errors.New("Missing extension to parse the post")
    ParserErrorInvalidMetadata = errors.New("Invalid metadata declaration")
    ParserErrorMetadataMissingEnd = errors.New("Metadata is missing closing '@end'")
    ParserErrorMetadataUnknownType = errors.New("Metadata type not recognised")
)
