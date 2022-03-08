package request

import "io"

type CustomRequest struct {
	Method      string
	URL         string
	AccessToken string
	Value       interface{}
	Body        io.ReadCloser
}
