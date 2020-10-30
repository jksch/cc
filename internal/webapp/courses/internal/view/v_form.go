package view

import (
	"fmt"
	"strconv"

	"github.com/jksch/cc/internal/models"
	"github.com/jksch/cc/internal/webapp/courses/internal/controller"
	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

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
