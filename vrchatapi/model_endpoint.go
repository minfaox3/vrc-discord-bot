package vrchatapi

type MethodType string

const (
	GET    MethodType = "GET"
	POST   MethodType = "POST"
	PUT    MethodType = "PUT"
	DELETE MethodType = "DELETE"
)

type Endpoint struct {
	url    string     `json:"url"`
	method MethodType `json:"method"`
}
