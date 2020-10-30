package view

import (
	"github.com/jksch/cc/internal/webapp/courses/internal/client"
	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

// NewCourses creates a working courses component.
func NewCourses() *Courses {
	client := &client.App{}
	cs := &Courses{
		Popup:    &Popup{},
		Progress: &ProgressBar{},
		Table:    &CoursesTable{},
		Form:     &CoursesForm{},
	}
	cs.Table.CoursesTable.EditCourse = cs.Form.EditCourse
	cs.Table.CoursesTable.ShowProgress = cs.Progress.ShowProgress
	cs.Table.CoursesTable.ShowPopup = cs.Popup.ShowPopup
	cs.Table.Client = client

	cs.Form.CoursesForm.AddCourse = cs.Table.AddCourse
	cs.Form.CoursesForm.ShowProgress = cs.Progress.ShowProgress
	cs.Form.CoursesForm.ShowPopup = cs.Popup.ShowPopup
	cs.Form.Client = client

	cs.Popup.Popup.Rerender = cs.Popup.Compo.Update
	return cs
}

// Courses view holds all components and functions to add/edit courses.
type Courses struct {
	app.Compo
	Popup    *Popup
	Progress *ProgressBar
	Table    *CoursesTable
	Form     *CoursesForm
}

// Render renders the view.
func (c *Courses) Render() app.UI {
	return app.Div().Body(
		app.Main().Body(
			app.Nav().Class("navbar navbar-dark bg-primary").Body(
				app.A().Class("navbar-brand").Body(
					app.Text("Courses App"),
				),
			),
			app.Div().Class("container").Body(
				c.Popup,
				c.Progress,
				c.Table,
				c.Form,
			),
		),
	)
}

//ProgressBar holds a infinite progress bar.
type ProgressBar struct {
	app.Compo
	Visible bool
}

// Render renders the progress bar.
func (p *ProgressBar) Render() app.UI {
	return app.Div().
		Class("progress-bar progress-bar-striped progress-bar-animated").
		Aria("valuenow", "75").
		Aria("valuemin", "0").
		Aria("valuemax", "100").
		Style("width", "100%").
		Hidden(!p.Visible)
}

// ShowProgress show or hides the progress bar.
func (p *ProgressBar) ShowProgress(visible bool) {
	p.Visible = visible
	p.Update()
}
