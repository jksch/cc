package controller

import "github.com/jksch/cc/internal/models"

// CoursesRequester allows to request all courses.
type CoursesRequester interface {
	LoadCourses() (courses []models.Course, err error)
}

// CoursesTable holds the table controller.
type CoursesTable struct {
	Client       CoursesRequester
	Table        []models.Course
	ShowProgress func(bool)
	ShowPopup    func(text string, error bool)
	EditCourse   func(models.Course)
}

// LoadCourses loads all saved courses.
// Returns true when successful.
func (c *CoursesTable) LoadCourses() bool {
	c.ShowProgress(true)

	courses, err := c.Client.LoadCourses()

	c.ShowProgress(false)
	if err != nil {
		c.ShowPopup("Requesting courses list failed!", true)
		return false
	}
	c.ShowPopup("Successfully loaded courses!", false)
	c.Table = courses
	return true
}

// AddCourse adds a course to the table.
func (c *CoursesTable) AddCourse(course models.Course) {
	for index, saved := range c.Table {
		if saved.ID == course.ID {
			c.Table[index] = course
			return
		}
	}
	c.Table = append(c.Table, course)
}
