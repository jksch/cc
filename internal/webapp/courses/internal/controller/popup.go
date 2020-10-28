package controller

import (
	"sync"
	"time"

	"github.com/jksch/cc/internal/webapp/courses/internal/webmodels"
)

var popupDuration = 5 * time.Second

// Popup holds the pop up logic.
type Popup struct {
	mux      sync.Mutex // Protects the messages slice.
	messages []webmodels.Popup
	Rerender func()
}

// RenderMessages returns the messages that should be rendered.
func (p *Popup) RenderMessages() []webmodels.Popup {
	p.mux.Lock()
	defer p.mux.Unlock()
	return p.messages
}

// ShowPopup creates a new popup that will be shown for 5 Sec.
func (p *Popup) ShowPopup(text string, error bool) {
	p.mux.Lock()
	defer p.mux.Unlock()
	pop := webmodels.Popup{
		ID:    time.Now(),
		Text:  text,
		Error: error,
	}
	p.messages = append(p.messages, pop)
	p.removeLater(pop.ID)
	defer p.Rerender()
}

func (p *Popup) removeLater(ID time.Time) {
	time.AfterFunc(popupDuration, func() {
		p.mux.Lock()
		defer p.mux.Unlock()
		for index, message := range p.messages {
			if message.ID == ID {
				p.messages = append(p.messages[:index], p.messages[index+1:]...)
				defer p.Rerender()
				return
			}
		}
	})
}
