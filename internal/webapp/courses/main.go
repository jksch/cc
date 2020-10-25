package main

import (
	"github.com/jksch/cc/internal/webapp/courses/internal/view"
	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

func main() {
	app.Route("/", view.NewCourses())
	app.Route("/courses", view.NewCourses())
	app.Run()
}
