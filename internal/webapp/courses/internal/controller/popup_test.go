package controller

import (
	"testing"
	"time"
)

func TestPopup(t *testing.T) {
	t.Parallel()
	popupDuration = 10 * time.Millisecond
	rerenderCalled := 0

	popup := &Popup{Rerender: func() { rerenderCalled++ }}
	popup.ShowPopup("POP", false)

	all := popup.RenderMessages()
	if len(all) != 1 {
		t.Errorf("Exp massages length: 1 got: %d", len(all))
	}
	if all[0].Text != "POP" {
		t.Errorf("Exp text: 'POP' got: '%s'", all[0].Text)
	}

	time.Sleep(11 * time.Millisecond)

	all = popup.RenderMessages()
	if len(all) != 0 {
		t.Errorf("Exp massages length: 0 got: %d", len(all))
	}
	if rerenderCalled != 2 {
		t.Errorf("Exp rerender calls: 2 got: %d", rerenderCalled)
	}

}
