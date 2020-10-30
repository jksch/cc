package view

import (
	"github.com/jksch/cc/internal/webapp/courses/internal/controller"
	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

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
