package responder

// Factory creates responders depending on the mime type.
type Factory interface {
	Create(mimeType string) Responder
}

type factory struct{}

func (fac factory) Create(mimeType string) Responder {
	switch mimeType {
	case "application/json":
		return applicationJSONResponder{}
	case "application/xml":
		return applicationXMLResponder{}
	case "image/gif":
		return imageGIFResponder{}
	case "image/png":
		return imagePNGResponder{}
	case "text/html":
		return textHTMLResponder{}
	case "text/plain":
		return textPlainResponder{}
	case "*/*":
		return textPlainResponder{}
	default:
		return nil
	}
}

// NewFactory returns a new factory.
func NewFactory() Factory {
	return factory{}
}
