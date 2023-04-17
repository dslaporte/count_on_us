package webserver

import (
	"count_on_us/internal/infrastructure/webserver"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type WebServer struct {
	router        chi.Router
	webServerPort string
}

func NewGoChiWebServer(serverPort string) *WebServer {
	server := chi.NewRouter()
	server.Use(middleware.Logger)
	server.Use(middleware.Recoverer)
	return &WebServer{
		router:        server,
		webServerPort: serverPort,
	}
}

func (s *WebServer) AddHandler(method webserver.HTTPMethod, path string, handler http.HandlerFunc) {
	switch method {
	case
		webserver.GET:
		s.router.Get(path, handler)
	case webserver.POST:
		s.router.Post(path, handler)
	case webserver.PUT:
		s.router.Put(path, handler)
	case webserver.DELETE:
		s.router.Delete(path, handler)
	default:
		panic(errors.New("invalid http method (POST|PUT|DELETE|GET)"))
	}
}

func (s *WebServer) Start() {
	fmt.Printf("Server is running at :%s", s.webServerPort)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", s.webServerPort), s.router); err != nil {
		panic(err)
	}
}
