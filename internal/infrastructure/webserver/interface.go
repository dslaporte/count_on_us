package webserver

import "net/http"

type HTTPMethod string

const (
	GET    HTTPMethod = "GET"
	POST   HTTPMethod = "POST"
	PUT    HTTPMethod = "PUT"
	DELETE HTTPMethod = "DELETE"
)

type WebServerInterface interface {
	AddHandler(method HTTPMethod, path string, handler http.HandlerFunc)
	Start()
}

type WebValidator interface {
	IsValid() error
}
