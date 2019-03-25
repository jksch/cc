// Package db holds the database implementation.
package db

import (
	"fmt"
	"sync"

	"github.com/jksch/cc/internal/models"
)

// Database is a mock implementation of the app database.
type Database struct {
	mux *sync.Mutex // protects the courses list and index.

	index   int
	courses []models.Course
}

// New crates a new instance of the database.
func New(path string) (*Database, error) {
	if path == "" {
		return nil, fmt.Errorf("db, no db path provided")
	}
	db := &Database{
		mux: &sync.Mutex{},
	}
	return db, nil
}

// LoadCourses returns all stored courses.
// Error for this mock is allways nil.
func (db *Database) LoadCourses() ([]models.Course, error) {
	db.mux.Lock()
	defer db.mux.Unlock()
	return db.courses, nil
}

// UpdateCourse creates or updates a given course in the database.
// If the course ID is nil the course will be created.
// On success the course ID form the db will be returned.
func (db *Database) UpdateCourse(course models.Course) (ID int, err error) {
	// Real implementation should check course validity.
	db.mux.Lock()
	defer db.mux.Unlock()
	if course.ID == 0 {
		db.index++
		course.ID = db.index
		db.courses = append(db.courses, course)
		return course.ID, nil
	}
	for pos := 0; pos < len(db.courses); pos++ {
		if db.courses[pos].ID == course.ID {
			db.courses[pos] = course
			return course.ID, nil
		}
	}
	return 0, fmt.Errorf("db, no course entry to update found")
}
