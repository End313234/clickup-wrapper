package request

type CustomRequest struct {
	Method      string
	URL         string
	AccessToken string
	Value       interface{}
}
