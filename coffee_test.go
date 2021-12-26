package machine

import (
	"fmt"
	"testing"
)

func TestOnOffMachine(t *testing.T) {
	m := NewMachine()

	off := OffCommand
	cmds := NewMemCommands()
	cmds.Push(off)
	exit := m.On(cmds)
	if exit != fmt.Sprintf("Exit by %d Command", off) {
		t.Error("exit not work")
	}
}
