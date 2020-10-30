package controller

import (
	"github.com/jksch/cc/internal/models"
)

// CoursesPersister allows to persist a course.
type CoursesPersister interface {
	SaveCourse(course models.Course) (ID int64, err error)
}

// CoursesForm holds the form controller.
type CoursesForm struct {
	Client       CoursesPersister
	Course       models.Course
	AddCourse    func(models.Course)
	ShowProgress func(bool)
	ShowPopup    func(text string, error bool)
}

// EditCourse sets the form values to the given course.
func (c *CoursesForm) EditCourse(course models.Course) {
	c.Course = course
}

// OnSubmit handles the form submit.
// Returns true if successful.
func (c *CoursesForm) OnSubmit() bool {
	c.ShowProgress(true)

	ID, err := c.Client.SaveCourse(c.Course)

	c.ShowProgress(false)
	if err != nil {
		c.ShowPopup("Could not save course!", true)
		return false
	}

	c.Course.ID = int(ID)
	c.AddCourse(c.Course)
	c.Course.Reset()
	c.ShowPopup("Successfully save course!", false)
	return true
}
