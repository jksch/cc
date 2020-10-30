package view

import (
	"fmt"

	"github.com/jksch/cc/internal/models"
	"github.com/jksch/cc/internal/webapp/courses/internal/controller"
	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

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
