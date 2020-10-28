package webapp

import (
	"net/http"

	"github.com/jksch/cc/internal/webapp/static"
	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

// New creates a new courses app http handler.
func New() http.Handler {
	return &app.Handler{
		Title:  "Courses app",
		Author: "jksch",
		Resources: &staticFS{
			Handler: http.FileServer(static.FS(false)),
		},
		Icon: app.Icon{
			Default: "/gopher_192.png",
			Large:   "/gopher_512.png",
		},
		LoadingLabel: "App is loading! Gophers are working on it...",
		Styles: []string{
			"/bootstrap.css",
		},
	}
}

//go:generate make assable
type staticFS struct {
	http.Handler
}

func (s staticFS) AppResources() string {
	return ""
}

func (s staticFS) StaticResources() string {
	return "/web"
}

func (s staticFS) AppWASM() string {
	return s.StaticResources() + "/app.wasm"
}

func (s staticFS) RobotsTxt() string {
	return s.StaticResources() + "/robots.txt"
}

func (s staticFS) AdsTxt() string {
	return s.StaticResources() + "/ads.txt"
}
