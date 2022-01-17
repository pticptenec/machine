package machine

import (
	"fmt"
	"testing"

	"github.com/jroimartin/gocui"
)

func TestAddOneElement(t *testing.T) {
	e := NewElement("tmp", "this is for\ntest\n")
	l, err := NewLayout()
	if err != nil {
		t.Errorf(fmt.Sprintf("%v", err))
	}

	l.Add(e, Position.Center, gocui.ColorBlue, nil)

	added, ok := l.components["tmp"]
}
