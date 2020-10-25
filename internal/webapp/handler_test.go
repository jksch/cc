package webapp

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWabAppHandler(t *testing.T) {
	var tests = []struct {
		resource string
	}{
		{
			resource: "/app.js",
		},
		{
			resource: "/web/app.wasm",
		},
		{
			resource: "/web/bootstrap.css",
		},
		{
			resource: "/web/gopher_192.png",
		},
	}
	for ti, tt := range tests {
		ti, tt := ti, tt
		t.Run(fmt.Sprintf("%d. Should be able to request '%s'", ti, tt.resource), func(t *testing.T) {
			t.Parallel()
			mux := http.NewServeMux()
			mux.Handle("/", New())

			rec := httptest.NewRecorder()
			req, err := http.NewRequest("GET", tt.resource, nil)
			logErr(t, err)

			mux.ServeHTTP(rec, req)

			if rec.Code != 200 {
				t.Errorf("%d. Could not request '%s' got: %d", ti, tt.resource, rec.Code)
			}
		})
	}
}

func TestStaticFS(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		call func(*staticFS) string
		path string
	}{
		{
			call: func(fs *staticFS) string {
				return fs.AppResources()
			},
			path: "",
		},
		{
			call: func(fs *staticFS) string {
				return fs.StaticResources()
			},
			path: "/web",
		},
		{
			call: func(fs *staticFS) string {
				return fs.AppWASM()
			},
			path: "/web/app.wasm",
		},
		{
			call: func(fs *staticFS) string {
				return fs.RobotsTxt()
			},
			path: "/web/robots.txt",
		},
		{
			call: func(fs *staticFS) string {
				return fs.AdsTxt()
			},
			path: "/web/ads.txt",
		},
	}
	for ti, tt := range tests {
		ti, tt := ti, tt
		t.Run(fmt.Sprintf("%d. Should load path '%s'", ti, tt.path), func(t *testing.T) {
			t.Parallel()
			fs := &staticFS{}
			path := tt.call(fs)

			if tt.path != path {
				t.Errorf("%d. Exp path: '%s' got: '%s'", ti, tt.path, path)
			}
		})
	}

}

func logErr(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("Unexpected error, %v", err)
	}
}
