// Package server holds the application web server.
package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	"log"
	"net/http"
	"reflect"
	"time"

	"github.com/jksch/cc/internal/models"
	"github.com/jksch/cc/internal/webapp"
	"github.com/jksch/rest/v2"
)

// DefaultAddr will be use when no addres is provided.
const DefaultAddr = ":3021"
const defaultTimeout = 10 * time.Second

// Persister is the db interface for this server.
type Persister interface {
	LoadCourses() ([]models.Course, error)
	UpdateCourse(models.Course) (ID int, err error)
}

// Server holds the web server for this application.
type Server struct {
	log *log.Logger
	db  Persister

	srv http.Server
	mux *http.ServeMux
}

// New creates a new instace of this http server.
func New(addr string, db Persister, logger *log.Logger) (*Server, error) {
	if addr == "" {
		addr = DefaultAddr
	}
	if db == nil || reflect.ValueOf(db).IsNil() {
		return nil, fmt.Errorf("server, no database provided")
	}
	if logger == nil {
		logger = log.New(ioutil.Discard, "", 0)
	}
	s := &Server{
		log: logger,
		db:  db,
		mux: http.NewServeMux(),
	}

	s.mux.Handle("/", webapp.New())
	s.mux.HandleFunc("/api/courses", s.corses)

	// A production server should have some kind of logger middleware.
	handler := rest.ChainHandler(s.mux, rest.RequestLoggerMiddleware(s.log, s.log))

	s.srv = http.Server{
		Handler:      handler,
		Addr:         addr,
		WriteTimeout: defaultTimeout,
		ReadTimeout:  defaultTimeout,
	}

	return s, nil
}

// Run blocks and starts the server.
func (s *Server) Run() error {
	s.log.Printf("Running server on %s", s.srv.Addr)
	return s.srv.ListenAndServe()
}

func (s *Server) corses(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.loadCorses(w, r)
	case http.MethodPost:
		s.updateCorse(w, r)
	default:
		http.Error(w, "Only GET or POST supported.", http.StatusMethodNotAllowed)
	}
}

func (s *Server) loadCorses(w http.ResponseWriter, r *http.Request) {
	all, err := s.db.LoadCourses()
	if err != nil {
		s.log.Printf("Could not load courses, %v", err)
		http.Error(w, "Could not load courses", http.StatusInternalServerError)
		return
	}

	rest.JSON(w, all, http.StatusOK)
}

func (s *Server) updateCorse(w http.ResponseWriter, r *http.Request) {
	var update models.Course
	d := json.NewDecoder(r.Body)
	defer r.Body.Close()

	if err := d.Decode(&update); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	if err := update.IsValid(); err != nil {
		http.Error(w, fmt.Sprintf("Invalid req, %v", err), http.StatusBadRequest)
		return
	}

	ID, err := s.db.UpdateCourse(update)
	if err != nil {
		s.log.Printf("Could not save course, %v", err)
		http.Error(w, "Could not save course", http.StatusInternalServerError)
		return
	}

	rest.String(w, strconv.Itoa(ID), http.StatusCreated)
}
