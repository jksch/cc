package db

import (
	"fmt"
	"reflect"
	"sync"
	"testing"

	"github.com/jksch/cc/internal/models"
)

func TestNewDatabase(t *testing.T) {
	var tests = []struct {
		name string
		path string
		exp  string
	}{
		{
			name: "No db path given",
			path: "",
			exp:  "db, no db path provided",
		},
		{
			name: "Successfully create db",
			path: "app.db",
			exp:  "",
		},
	}

	for ti, tt := range tests {
		ti, tt := ti, tt
		t.Run(fmt.Sprintf("%d. %s", ti, tt.name), func(t *testing.T) {
			t.Parallel()
			db, err := New(tt.path)
			got := errStr(err)
			if tt.exp != got {
				t.Errorf("%d. %s, exp err: '%s' got: '%s'", ti, tt.name, tt.exp, got)
			}
			if err != nil {
				return // done
			}
			if db == nil {
				t.Errorf("%d. %s, exp db not nil", ti, tt.name)
			}
		})
	}
}

func TestLoadCourses(t *testing.T) {
	var tests = []struct {
		name string
		init []models.Course
		exp  []models.Course
	}{
		{
			name: "No entry should be empty list",
			exp:  nil,
		},
		{
			name: "Return entry",
			init: []models.Course{
				{
					ID:          1,
					Name:        "Go",
					Description: "Go course",
					Instructor:  "John",
					CentPrice:   1000,
				},
			},
			exp: []models.Course{
				{
					ID:          1,
					Name:        "Go",
					Description: "Go course",
					Instructor:  "John",
					CentPrice:   1000,
				},
			},
		},
	}

	for ti, tt := range tests {
		ti, tt := ti, tt
		t.Run(fmt.Sprintf("%d. %s", ti, tt.name), func(t *testing.T) {
			t.Parallel()
			db, err := New("app.db")
			logErr(t, err)
			db.courses = tt.init

			got, err := db.LoadCourses()
			if err != nil {
				t.Errorf("%d. %s, exp err: nil got: %v", ti, tt.name, err)
			}
			if !reflect.DeepEqual(tt.exp, got) {
				t.Errorf("%d. %s, exp loaded courses: %v got: %v", ti, tt.name, tt.exp, got)
			}
		})
	}
}

func TestUpdateCourse(t *testing.T) {
	var tests = []struct {
		name  string
		init  []models.Course
		given models.Course
		id    int
		err   string
		exp   []models.Course
	}{
		{
			name: "Add course",
			given: models.Course{
				Name:        "Go",
				Description: "Go course",
				Instructor:  "John",
				CentPrice:   1000,
			},
			id: 1,
			exp: []models.Course{
				{
					ID:          1,
					Name:        "Go",
					Description: "Go course",
					Instructor:  "John",
					CentPrice:   1000,
				},
			},
		},
		{
			name: "Update course",
			init: []models.Course{
				{
					ID:          1,
					Name:        "Go",
					Description: "Go course",
					Instructor:  "John",
					CentPrice:   1000,
				},
			},
			given: models.Course{
				ID:          1,
				Name:        "Go II",
				Description: "Go course",
				Instructor:  "John Smith",
				CentPrice:   1000,
			},
			id: 1,
			exp: []models.Course{
				{
					ID:          1,
					Name:        "Go II",
					Description: "Go course",
					Instructor:  "John Smith",
					CentPrice:   1000,
				},
			},
		},
		{
			name: "Update course, no entry found",
			given: models.Course{
				ID:          1,
				Name:        "Go II",
				Description: "Go course",
				Instructor:  "John Smith",
				CentPrice:   1000,
			},
			id:  0,
			err: "db, no course entry to update found",
			exp: nil,
		},
	}

	for ti, tt := range tests {
		ti, tt := ti, tt
		t.Run(fmt.Sprintf("%d. %s", ti, tt.name), func(t *testing.T) {
			t.Parallel()
			db, err := New("app.db")
			logErr(t, err)
			db.courses = tt.init

			id, err := db.UpdateCourse(tt.given)
			sErr := errStr(err)
			if tt.err != sErr {
				t.Errorf("%d. %s, exp err: '%s' got: '%s'", ti, tt.name, tt.err, sErr)
			}
			if tt.id != id {
				t.Errorf("%d. %s, exp ID: %d got: %d", ti, tt.name, tt.id, id)
			}
			if !reflect.DeepEqual(tt.exp, db.courses) {
				t.Errorf("%d. %s, exp stored courses: %v got: %v", ti, tt.name, tt.exp, db.courses)
			}
		})
	}
}

// This test should be run with: 'go test -race'
func TestDatabaseIsThreadSave(t *testing.T) {
	t.Parallel()

	db, err := New("app.db")
	logErr(t, err)

	wg := &sync.WaitGroup{}
	wg.Add(3)
	go func() {
		_, _ = db.UpdateCourse(models.Course{
			Name:        "Go II",
			Description: "Go course",
			Instructor:  "John Smith",
			CentPrice:   1000,
		})
		wg.Done()
	}()
	go func() {
		_, _ = db.LoadCourses()
		wg.Done()
	}()
	go func() {
		_, _ = db.UpdateCourse(models.Course{
			Name:        "Go",
			Description: "Go course",
			Instructor:  "John Smith",
			CentPrice:   1000,
		})
		wg.Done()
	}()

	// This test is only usefull when run with go test -race
	wg.Wait()
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
