package error

import "errors"

var (
    ParserErrorExtensionNotSupported  = errors.New("The extension used to write the blog is not supported")
    ParserErrorMissingFileExtension  = errors.New("Missing extension to parse the post")
    NorgParserErrorInvalidMetadata = errors.New("Invalid metadata declaration")
    NorgParserErrorMetadataMissingEnd = errors.New("Metadata is missing closing '@end'")
    NorgParserErrorMetadataUnknownType = errors.New("Metadata type not recognised")
    NorgParserErrorInvalidTableColumn = errors.New("Table columns are malformed")
)
