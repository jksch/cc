package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/jksch/cc/internal/models"
)

func TestGetCourses(t *testing.T) {
	var tests = []struct {
		name        string
		handler     func(http.ResponseWriter, *http.Request)
		noConnecton bool
		exp         []models.Course
		err         string
	}{
		{
			name:        "Should error on no connection",
			handler:     func(w http.ResponseWriter, r *http.Request) {},
			noConnecton: true,
			err:         "could not request courses,",
		},
		{
			name: "Should error if response code is not 200",
			handler: func(w http.ResponseWriter, r *http.Request) {
				http.Error(w, "Boom!", 500)
			},
			err: "could not request courses, server responded with: 500",
		},
		{
			name: "Should error on invalid JSON",
			handler: func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintf(w, "{...}")
			},
			err: "could not decode courses JSON,",
		},
		{
			name: "Should successfully request courses",
			handler: func(w http.ResponseWriter, r *http.Request) {
				courses := []models.Course{{
					ID:          1,
					Name:        "Go",
					Description: "How to go...",
					Instructor:  "John Doe",
					CentPrice:   2050,
				}}
				logErr(t, json.NewEncoder(w).Encode(courses))
			},
			exp: []models.Course{{
				ID:          1,
				Name:        "Go",
				Description: "How to go...",
				Instructor:  "John Doe",
				CentPrice:   2050,
			}},
		},
	}
	for ti, tt := range tests {
		ti, tt := ti, tt
		t.Run(fmt.Sprintf("%d. %v", ti, tt.name), func(t *testing.T) {
			t.Parallel()

			mux := http.NewServeMux()
			mux.HandleFunc("/api/courses", tt.handler)
			srv := httptest.NewServer(mux)
			if tt.noConnecton {
				srv.Close()
			} else {
				defer srv.Close()
			}

			client := App{Addr: srv.URL}
			gCources, err := client.LoadCourses()

			sErr := errStr(err)
			if !strings.HasPrefix(sErr, tt.err) {
				t.Errorf("%d. %s: Exp err prefix: '%v' got: '%v'", ti, tt.name, tt.err, sErr)
			}
			exp := toJson(t, tt.exp)
			got := toJson(t, gCources)
			if exp != got {
				t.Errorf("%d. %s: Exp courses:\n'%v'\ngot:\n'%v'", ti, tt.name, tt.exp, got)
			}
		})
	}
}

func TestSaveCourses(t *testing.T) {
	var tests = []struct {
		name        string
		handler     func(http.ResponseWriter, *http.Request)
		noConnecton bool
		exp         int64
		err         string
	}{
		{
			name:        "Should error on no connection",
			handler:     func(w http.ResponseWriter, r *http.Request) {},
			noConnecton: true,
			err:         "could not save course,",
		},
		{
			name:    "Should error on not status 201",
			handler: func(w http.ResponseWriter, r *http.Request) {},
			err:     "could not save course, server responded with 200",
		},
		{
			name: "Should error if no ID is returned or readable",
			handler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(201)
				fmt.Fprintf(w, "nope ;)")
			},
			err: "could not extract saved courses ID from response,",
		},
		{
			name: "Should successfully save course",
			handler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(201)
				fmt.Fprintf(w, "1\n")
			},
			exp: 1,
		},
	}
	for ti, tt := range tests {
		ti, tt := ti, tt
		t.Run(fmt.Sprintf("%d. %v", ti, tt.name), func(t *testing.T) {
			t.Parallel()

			mux := http.NewServeMux()
			mux.HandleFunc("/api/courses", tt.handler)
			srv := httptest.NewServer(mux)
			if tt.noConnecton {
				srv.Close()
			} else {
				defer srv.Close()
			}

			course := models.Course{
				ID:          1,
				Name:        "Go",
				Description: "How to go...",
				Instructor:  "John Doe",
				CentPrice:   2050,
			}
			client := App{Addr: srv.URL}
			got, err := client.SaveCourse(course)

			sErr := errStr(err)
			if !strings.HasPrefix(sErr, tt.err) {
				t.Errorf("%d. %s: Exp err prefix: '%v' got: '%v'", ti, tt.name, tt.err, sErr)
			}
			if tt.exp != got {
				t.Errorf("%d. %s: Exp ID: '%d' got: '%d'", ti, tt.name, tt.exp, got)
			}
		})
	}
}

func errStr(err error) string {
	if err != nil {
		return err.Error()
	}
	return ""
}

func logErr(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("Unexpected error, %v", err)
	}
}

func toJson(t *testing.T, val interface{}) string {
	t.Helper()
	b, err := json.MarshalIndent(val, "", "\t")
	logErr(t, err)
	return string(b)
}
