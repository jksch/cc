package view

import "github.com/maxence-charriere/go-app/v7/pkg/app"

// NewCourses creates a working courses component.
func NewCourses() *Courses {
	return &Courses{
		Popup:    &Popup{},
		Progress: &ProgressBar{Visible: true},
		Table:    &CoursesTable{},
		Form:     &CoursesForm{},
	}
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

// Popup holds a pop up info.
type Popup struct {
	app.Compo
}

func (p *Popup) Render() app.UI {
	return app.Div().Class("mt-4 alert alert-success").Body(
		app.Text("Request successful!"),
	)
}

//ProgressBar holds a infinite progress bar.
type ProgressBar struct {
	app.Compo
	Visible bool
}

func (p *ProgressBar) Render() app.UI {
	return app.Div().Class("mt-4 progress").Body(
		app.Div().
			Class("progress-bar progress-bar-striped progress-bar-animated").
			Aria("valuenow", "75").
			Aria("valuemin", "0").
			Aria("valuemax", "100").
			Style("width", "100%").
			Hidden(!p.Visible),
	)
}

// CoursesTable holds the table view.
type CoursesTable struct {
	app.Compo
}

func (c *CoursesTable) Render() app.UI {
	return app.Div().Body(
		app.Table().Class("mt-4 table table-striped").Body(
			app.THead().Body(
				app.Tr().Body(
					app.Th().Scope("col").Body(
						app.Text("#"),
					),
					app.Th().Scope("col").Body(
						app.Text("Name:"),
					),
					app.Th().Scope("col").Body(
						app.Text("Description:"),
					),
					app.Th().Scope("col").Body(
						app.Text("Instructor:"),
					),
					app.Th().Scope("col").Body(
						app.Text("Price:"),
					),
					app.Th().Scope("col").Body(
						app.Text("Options:"),
					),
				),
			),
			// TODO add loop and model.
			app.TBody().Body(
				app.Td().Body(
					app.Text("1"),
				),
				app.Td().Body(
					app.Text("Go"),
				),
				app.Td().Body(
					app.Text("A Golang course"),
				),
				app.Td().Body(
					app.Text("John Doe"),
				),
				app.Td().Body(
					app.Text("10.47 EUR"),
				),
				app.Td().Body(
					app.Button().
						Class("btn btn-primary").
						Text("Edit"),
				),
			),
		),
	)
}

// CoursesForm holds the input form.
type CoursesForm struct {
	app.Compo
}

func (c *CoursesForm) Render() app.UI {
	return app.Div().Class("mt-4").Body(
		app.Form().Body(
			app.Div().Class("form-group").Body(
				app.Label().For("in-course-name").Body(
					app.Text("Name:"),
				),
				app.Input().
					ID("in-course-name").
					Type("text").
					Class("form-control").
					Aria("describedby", "name").
					Placeholder("Course name").
					Required(true),
			),
			app.Div().Class("form-group").Body(
				app.Label().For("in-course-description").Body(
					app.Text("Description:"),
				),
				app.Input().
					ID("in-course-description").
					Type("text").
					Class("form-control").
					Aria("describedby", "description").
					Placeholder("Course description").
					Required(true),
			),
			app.Div().Class("form-group").Body(
				app.Label().For("in-course-instructor").Body(
					app.Text("Instructor:"),
				),
				app.Input().
					ID("in-course-instructor").
					Type("text").
					Class("form-control").
					Aria("describedby", "instructor").
					Placeholder("Course instructor").
					Required(true),
			),
			app.Div().Class("form-group").Body(
				app.Label().For("in-course-price").Body(
					app.Text("Price:"),
				),
				app.Input().
					ID("in-course-price").
					Type("number").
					Class("form-control").
					Aria("describedby", "price").
					Placeholder("Course price").
					Required(true),
			),
			app.Button().
				Class("btn btn-primary").
				Type("submit").
				Text("Submit"),
		),
	)
}
