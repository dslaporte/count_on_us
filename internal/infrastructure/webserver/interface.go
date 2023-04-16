package webserver

import "net/http"

type WebServerInterface interface {
	AddHandler(path string, handler http.HandlerFunc)
	Start()
}

type WebValidator interface {
	IsValid() error
}
