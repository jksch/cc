package view

import (
	"fmt"
	"strconv"

	"github.com/jksch/cc/internal/models"
	"github.com/jksch/cc/internal/webapp/courses/internal/client"
	"github.com/jksch/cc/internal/webapp/courses/internal/controller"
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

// Popup holds a pop up info.
type Popup struct {
	app.Compo
	controller.Popup
}

// Render renders the popup.
func (p *Popup) Render() app.UI {
	messages := p.RenderMessages()
	popups := make([]app.UI, 0, len(messages))
	for _, msg := range messages {
		pop := app.If(msg.Error,
			app.Div().Class("mt-4 alert alert-danger").Body(
				app.A().Class("alert-link").Body(
					app.Text(msg.Text),
				),
			)).Else(
			app.Div().Class("mt-4 alert alert-success").Body(
				app.A().Class("alert-link").Body(
					app.Text(msg.Text),
				),
			))
		popups = append(popups, pop)
	}
	return app.Div().Body(popups...)

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

// CoursesTable holds the table view.
type CoursesTable struct {
	app.Compo
	controller.CoursesTable
}

func (c *CoursesTable) OnMount(ctx app.Context) {
	go func() {
		if c.LoadCourses() {
			app.Dispatch(func() {
				c.Update()
			})
		}
	}()
}

func (c *CoursesTable) Render() app.UI {
	if len(c.Table) == 0 {
		return app.Div().Class("mt-4 alert alert-light").Body(
			app.A().Class("alert-link").Body(
				app.Text("Course table is empty."),
			),
		)
	}
	table := make([]app.UI, 0, len(c.Table))
	for index, course := range c.Table {
		tr := app.Tr().Body(
			app.Td().Body(
				app.Text(fmt.Sprintf("% d", index+1)),
			),
			app.Td().Body(
				app.Text(course.Name),
			),
			app.Td().Body(
				app.Text(course.Description),
			),
			app.Td().Body(
				app.Text(course.Instructor),
			),
			app.Td().Body(
				app.Text(fmt.Sprintf("%.2f €", float64(course.CentPrice)/100)),
			),
			app.Td().Body(
				app.Button().
					Class("btn btn-primary").
					Text("Edit").
					OnClick(func(app.Context, app.Event) {
						c.EditCourse(course)
					}),
			),
		)
		table = append(table, tr)
	}
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
			app.TBody().Body(table...),
		),
	)
}

// AddCourse allows to add a course to the table view.
func (c *CoursesTable) AddCourse(course models.Course) {
	c.CoursesTable.AddCourse(course)
	c.Update()
}

// CoursesForm holds the input form.
type CoursesForm struct {
	app.Compo
	controller.CoursesForm
}

// Render renders the form view.
func (c *CoursesForm) Render() app.UI {
	return app.Div().Class("mt-4").Body(
		app.Form().OnSubmit(func(ctx app.Context, e app.Event) {
			e.PreventDefault()
			go func() {
				if c.OnSubmit() {
					app.Dispatch(func() {
						ctx.JSSrc.Call("reset")
						c.Update() // Without is the reset dose not work after edit.
					})
				}
			}()
		}).Body(
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
					Value(c.Course.Name).
					Required(true).
					OnChange(func(ctx app.Context, e app.Event) {
						c.Course.Name = ctx.JSSrc.Get("value").String()
					}).
					AutoFocus(true),
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
					Value(c.Course.Description).
					OnChange(func(ctx app.Context, e app.Event) {
						c.Course.Description = ctx.JSSrc.Get("value").String()
					}).
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
					Value(c.Course.Instructor).
					OnChange(func(ctx app.Context, e app.Event) {
						c.Course.Instructor = ctx.JSSrc.Get("value").String()
					}).
					Required(true),
			),
			app.Div().Class("form-group").Body(
				app.Label().For("in-course-price").Body(
					app.Text("Price:"),
				),
				app.Input().
					ID("in-course-price").
					Type("number").
					Step(0.01).
					Class("form-control").
					Aria("describedby", "price").
					Placeholder("Course price").
					Value(fmt.Sprintf("%.2f", float64(c.Course.CentPrice)/100)).
					OnChange(func(ctx app.Context, e app.Event) {
						price, err := strconv.ParseFloat(ctx.JSSrc.Get("value").String(), 64)
						if err == nil {
							c.Course.CentPrice = int64(price * 100)
						}
					}).
					Required(true),
			),
			app.Button().
				Class("btn btn-primary").
				Type("submit").
				Text("Submit"),
		),
	)
}

// EditCourse allows to edit a given course.
func (c *CoursesForm) EditCourse(course models.Course) {
	c.CoursesForm.EditCourse(course)
	c.Update()
}
