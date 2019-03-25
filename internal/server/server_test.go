package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/jksch/cc/internal/models"
)

type testDB struct {
	courses []models.Course
	id      int
	err     error
}

func (t *testDB) LoadCourses() ([]models.Course, error) {
	return t.courses, t.err
}

func (t *testDB) UpdateCourse(models.Course) (ID int, err error) {
	return t.id, t.err
}

func TestNewServer(t *testing.T) {
	var tests = []struct {
		name string

		addr string
		db   Persister
		log  *log.Logger

		err     string
		expAddr string
	}{
		{
			name: "No db given, nil",
			db:   nil,
			err:  "server, no database provided",
		},
		{
			name:    "Use default addr",
			db:      &testDB{},
			expAddr: DefaultAddr,
		},
		{
			name:    "Use 8080 as server address",
			db:      &testDB{},
			addr:    ":8080",
			expAddr: ":8080",
		},
	}

	for ti, tt := range tests {
		ti, tt := ti, tt
		t.Run(fmt.Sprintf("%d. %s", ti, tt.name), func(t *testing.T) {
			t.Parallel()

			srv, err := New(tt.addr, tt.db, tt.log)
			sErr := errStr(err)
			if tt.err != sErr {
				t.Errorf("%d. %s, exp err: '%s' got: %s", ti, tt.name, tt.err, sErr)
			}
			if err != nil {
				return // done
			}

			if tt.expAddr != srv.srv.Addr {
				t.Errorf("%d. %s, exp server addr: '%s' got: '%s'", ti, tt.name, tt.expAddr, srv.srv.Addr)
			}
			if defaultTimeout != srv.srv.WriteTimeout {
				t.Errorf("%d. %s, exp server WriteTimeout: %v got: %v", ti, tt.name, tt.addr, srv.srv.WriteTimeout)
			}
			if defaultTimeout != srv.srv.ReadTimeout {
				t.Errorf("%d. %s, exp server ReadTimeout: %v got: %v", ti, tt.name, tt.addr, srv.srv.ReadTimeout)
			}
			if !reflect.DeepEqual(tt.db, srv.db) {
				t.Errorf("%d. %s, exp server db: %v got: %v", ti, tt.name, tt.db, srv.db)
			}
			if srv.log == nil {
				t.Errorf("%d. %s, exp server log not to be nil", ti, tt.name)
			}
		})
	}
}

func TestServerRunHandlesError(t *testing.T) {
	t.Parallel()

	srv, err := New("", &testDB{}, nil)
	logErr(t, err)

	srv.srv.Addr = "" // break server

	if err := srv.Run(); err == nil {
		t.Errorf("Exp server run to return error")
	}
}

func TestServerRunLogsAddr(t *testing.T) {
	t.Parallel()
	buf := bytes.NewBuffer(nil)
	log := log.New(buf, "", 0)

	srv, err := New("", &testDB{}, log)
	logErr(t, err)

	srv.srv.Addr = "" // break server
	srv.Run()

	exp := "Running server on "
	got := buf.String()
	if !strings.HasPrefix(got, exp) {
		t.Errorf("Exp log prefix: '%s' got: %s", exp, got)
	}
}

func TestShouldDeliverStaticContent(t *testing.T) {
	t.Parallel()

	srv, err := New("", &testDB{}, nil)
	logErr(t, err)

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/index.html", nil)
	srv.srv.Handler.ServeHTTP(rec, req)

	if rec.Code > 399 {
		t.Errorf("Server could not deliver index.html")
	}
}

func TestServerLoggerMiddleware(t *testing.T) {
	t.Parallel()

	buf := bytes.NewBuffer(nil)
	logger := log.New(buf, "", 0)

	srv, err := New("", &testDB{}, logger)
	logErr(t, err)

	rec := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/not/found", nil)
	logErr(t, err)

	srv.srv.Handler.ServeHTTP(rec, req)

	got := buf.String()
	prefix := "<-- GET /not/found"
	if !strings.HasPrefix(got, prefix) {
		t.Errorf("Exp log prefix: '%s' got: '%s'", got, prefix)
	}
}

var errTest = fmt.Errorf("IO error")
var testCourses = []models.Course{
	{
		ID:          1,
		Name:        "Go I",
		Description: "Go course",
		Instructor:  "John",
		CentPrice:   1000,
	},
}

func TestWebServiceCoursesPaths(t *testing.T) {
	var tests = []struct {
		name string
		db   Persister
		req  *http.Request

		code int
		body string
	}{
		{
			name: "Invalid method PUT",
			db:   &testDB{},
			req: httptest.NewRequest(
				http.MethodPut,
				"/api/courses",
				nil,
			),

			code: http.StatusMethodNotAllowed,
			body: "Only GET or POST supported.\n",
		},
		{
			name: "DB error on loading courses",
			db:   &testDB{err: errTest},
			req: httptest.NewRequest(
				http.MethodGet,
				"/api/courses",
				nil,
			),

			code: http.StatusInternalServerError,
			body: "Could not load courses\n",
		},
		{
			name: "Successfully load course",
			db:   &testDB{courses: testCourses},
			req: httptest.NewRequest(
				http.MethodGet,
				"/api/courses",
				nil,
			),

			code: http.StatusOK,
			body: strings.Trim(toJSON(t, testCourses).String(), "\n"),
		},
		{
			name: "Invalid JSON passed to update course.",
			db:   &testDB{},
			req: httptest.NewRequest(
				http.MethodPost,
				"/api/courses",
				bytes.NewBufferString("{..."),
			),

			code: http.StatusBadRequest,
			body: "Invalid JSON body\n",
		},
		{
			name: "Invalid course passed to update course",
			db:   &testDB{},
			req: httptest.NewRequest(
				http.MethodPost,
				"/api/courses",
				toJSON(t, models.Course{}),
			),

			code: http.StatusBadRequest,
			body: "Invalid req, course, name is required\n",
		},
		{
			name: "DB error on updating course",
			db:   &testDB{err: errTest},
			req: httptest.NewRequest(
				http.MethodPost,
				"/api/courses",
				toJSON(t, models.Course{
					ID:          1,
					Name:        "Go I",
					Description: "Go course",
					Instructor:  "John",
					CentPrice:   1000,
				}),
			),

			code: http.StatusInternalServerError,
			body: "Could not save course\n",
		},
		{
			name: "Successfully update course",
			db:   &testDB{id: 1},
			req: httptest.NewRequest(
				http.MethodPost,
				"/api/courses",
				toJSON(t, models.Course{
					ID:          1,
					Name:        "Go I",
					Description: "Go course",
					Instructor:  "John",
					CentPrice:   1000,
				}),
			),

			code: http.StatusCreated,
			body: "1\n",
		},
	}

	for ti, tt := range tests {
		ti, tt := ti, tt
		t.Run(fmt.Sprintf("%d.[%s %s] %s", ti, tt.req.Method, tt.req.URL.String(), tt.name), func(t *testing.T) {
			t.Parallel()

			srv, err := New("", tt.db, nil)
			logErr(t, err)

			rec := httptest.NewRecorder()
			srv.srv.Handler.ServeHTTP(rec, tt.req)

			if tt.code != rec.Code {
				t.Errorf("%d. %s, exp status code: %d got: %d", ti, tt.name, tt.code, rec.Code)
			}
			body := rec.Body.String()
			if tt.body != body {
				t.Errorf("%d. %s, exp res body:\n'%s'\ngot:\n'%s'\n", ti, tt.name, tt.body, body)
			}
		})
	}
}

func logErr(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("Unexpected err, %+v", err)
	}
}

func errStr(err error) string {
	if err != nil {
		return err.Error()
	}
	return ""
}

func toJSON(t *testing.T, val interface{}) *bytes.Buffer {
	t.Helper()
	buf := bytes.NewBuffer(nil)
	e := json.NewEncoder(buf)
	err := e.Encode(val)
	logErr(t, err)
	return buf
}
