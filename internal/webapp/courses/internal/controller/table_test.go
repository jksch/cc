package controller

import (
	"fmt"
	"testing"

	"github.com/jksch/cc/internal/models"
)

func TestTableAddCourse(t *testing.T) {
	t.Parallel()

	course := models.Course{
		ID:          1,
		Name:        "Go",
		Description: "How to go...",
		Instructor:  "John Doe",
		CentPrice:   2050,
	}

	table := &CoursesTable{}
	count := len(table.Table)
	if count != 0 {
		t.Errorf("Exp table to be empty got: %d", count)
	}

	table.AddCourse(course)
	table.AddCourse(course) // The same course should be updated.

	count = len(table.Table)
	if count != 1 {
		t.Fatalf("Exp table count: 1 got: %d", count)
	}

	if table.Table[0] != course {
		t.Errorf("Exp course: %+v got: %+v", course, table.Table[0])
	}
}

type testRequester struct {
	courses []models.Course
	err     error
}

func (t *testRequester) LoadCourses() (courses []models.Course, err error) {
	return t.courses, t.err
}

func TestTableFailToLoadCourses(t *testing.T) {
	t.Parallel()

	expCallSequece := []string{
		"ShowProgress;true",
		"ShowProgress;false",
		"ShowPopup;Requesting courses list failed!;true",
	}
	callSequece := []string{}

	table := &CoursesTable{
		Client: &testRequester{err: fmt.Errorf("Boom!")},
		ShowProgress: func(b bool) {
			callSequece = append(callSequece, fmt.Sprintf("ShowProgress;%t", b))
		},
		ShowPopup: func(text string, error bool) {
			callSequece = append(callSequece, fmt.Sprintf("ShowPopup;%s;%t", text, error))
		},
	}

	success := table.LoadCourses()
	if success {
		t.Errorf("Expected LoadCourses() to fail.")
	}
	if len(expCallSequece) != len(callSequece) {
		t.Fatalf("Exp calls: %d got: %d", len(expCallSequece), len(callSequece))
	}
	for i, exp := range expCallSequece {
		if exp != callSequece[i] {
			t.Errorf("%d. Expected call: '%s' got: '%s'", i, exp, callSequece[i])
		}
	}
	count := len(table.Table)
	if count != 0 {
		t.Errorf("Exp table to be empty got: %d", count)
	}
}

func TestTableSuccessfullyLoadCourses(t *testing.T) {
	t.Parallel()

	expCallSequece := []string{
		"ShowProgress;true",
		"ShowProgress;false",
		"ShowPopup;Successfully loaded courses!;false",
	}
	callSequece := []string{}

	course := models.Course{
		ID:          1,
		Name:        "Go",
		Description: "How to go...",
		Instructor:  "John Doe",
		CentPrice:   2050,
	}
	table := &CoursesTable{
		Client: &testRequester{courses: []models.Course{course}},
		ShowProgress: func(b bool) {
			callSequece = append(callSequece, fmt.Sprintf("ShowProgress;%t", b))
		},
		ShowPopup: func(text string, error bool) {
			callSequece = append(callSequece, fmt.Sprintf("ShowPopup;%s;%t", text, error))
		},
	}

	success := table.LoadCourses()
	if !success {
		t.Errorf("Expected LoadCourses() to succeed.")
	}
	if len(expCallSequece) != len(callSequece) {
		t.Fatalf("Exp calls: %d got: %d", len(expCallSequece), len(callSequece))
	}
	for i, exp := range expCallSequece {
		if exp != callSequece[i] {
			t.Errorf("%d. Expected call: '%s' got: '%s'", i, exp, callSequece[i])
		}
	}
	count := len(table.Table)
	if count != 1 {
		t.Fatalf("Exp table count: 1 got: %d", count)
	}
	if table.Table[0] != course {
		t.Errorf("Exp course: %+v got: %+v", course, table.Table[0])
	}
}
