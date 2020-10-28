package webmodels

import "time"

// Popup holds the popup message.
type Popup struct {
	ID    time.Time
	Text  string
	Error bool
}
