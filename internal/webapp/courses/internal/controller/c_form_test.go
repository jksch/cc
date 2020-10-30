package controller

import (
	"fmt"
	"testing"

	"github.com/jksch/cc/internal/models"
)

type testPersister struct {
	ID  int64
	err error
}

func (t *testPersister) SaveCourse(course models.Course) (ID int64, err error) {
	return t.ID, t.err
}

func TestFormAddCourse(t *testing.T) {
	t.Parallel()

	form := &CoursesForm{}
	if form.Course != (models.Course{}) {
		t.Errorf("Exp empty course got: %+v", form.Course)
	}

	course := models.Course{
		ID:          1,
		Name:        "Go",
		Description: "How to go...",
		Instructor:  "John Doe",
		CentPrice:   2050,
	}
	form.EditCourse(course)

	if form.Course != course {
		t.Errorf("Exp course: %+v got: %+v", course, form.Course)
	}
}

func TestFormOnSubmitError(t *testing.T) {
	t.Parallel()

	expCallSequece := []string{
		"ShowProgress;true",
		"ShowProgress;false",
		"ShowPopup;Could not save course!;true",
	}
	callSequece := []string{}
	course := models.Course{
		ID:          1,
		Name:        "Go",
		Description: "How to go...",
		Instructor:  "John Doe",
		CentPrice:   2050,
	}

	form := &CoursesForm{
		Client: &testPersister{err: fmt.Errorf("Boom!")},
		Course: course,
		AddCourse: func(models.Course) {
			t.Errorf("AddCourses() should not have been called.")
		},
		ShowProgress: func(b bool) {
			callSequece = append(callSequece, fmt.Sprintf("ShowProgress;%t", b))
		},
		ShowPopup: func(text string, error bool) {
			callSequece = append(callSequece, fmt.Sprintf("ShowPopup;%s;%t", text, error))
		},
	}

	success := form.OnSubmit()
	if success {
		t.Errorf("Exp OnSubmit() to fail.")
	}
	if len(expCallSequece) != len(callSequece) {
		t.Fatalf("Exp calls: %d got: %d", len(expCallSequece), len(callSequece))
	}
	for i, exp := range expCallSequece {
		if exp != callSequece[i] {
			t.Errorf("%d. Expected call: '%s' got: '%s'", i, exp, callSequece[i])
		}
	}
	if form.Course != course {
		t.Errorf("Exp course: '%+v' got: '%+v'", course, form.Course)
	}
}

func TestFormOnSubmitSuccess(t *testing.T) {
	t.Parallel()

	expCallSequece := []string{
		"ShowProgress;true",
		"ShowProgress;false",
		"AddCourse;1",
		"ShowPopup;Successfully save course!;false",
	}
	callSequece := []string{}
	course := models.Course{
		ID:          0,
		Name:        "Go",
		Description: "How to go...",
		Instructor:  "John Doe",
		CentPrice:   2050,
	}

	form := &CoursesForm{
		Client: &testPersister{ID: 1},
		Course: course,
		AddCourse: func(course models.Course) {
			callSequece = append(callSequece, fmt.Sprintf("AddCourse;%d", course.ID))
		},
		ShowProgress: func(b bool) {
			callSequece = append(callSequece, fmt.Sprintf("ShowProgress;%t", b))
		},
		ShowPopup: func(text string, error bool) {
			callSequece = append(callSequece, fmt.Sprintf("ShowPopup;%s;%t", text, error))
		},
	}

	success := form.OnSubmit()
	if !success {
		t.Errorf("Exp OnSubmit() to succeed.")
	}
	if len(expCallSequece) != len(callSequece) {
		t.Fatalf("Exp calls: %d got: %d", len(expCallSequece), len(callSequece))
	}
	for i, exp := range expCallSequece {
		if exp != callSequece[i] {
			t.Errorf("%d. Expected call: '%s' got: '%s'", i, exp, callSequece[i])
		}
	}
	if form.Course != (models.Course{}) {
		t.Errorf("Exp course: '%+v' got: '%+v'", course, form.Course)
	}
}
