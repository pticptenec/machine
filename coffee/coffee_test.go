package coffee

import (
	"fmt"
	"testing"
)

var c Config

func TestOnMachineChecksDone(t *testing.T) {
	c = Config{Water: 50, Beans: 30, Grind: 30}
	m := NewMachine(c)
	m.On()

	lamps := [...]bool{m.beansTank.Status(), m.waterTank.Status(),
		m.grindTank.Status()}
	for _, l := range lamps {
		if l != false {
			t.Error("wrong Check command")
		}
	}
}

func TestOnMachineChecksDoneLampsOn(t *testing.T) {
	c = Config{Water: 5, Beans: 5, Grind: 110}
	m := NewMachine(c)
	m.On()

	lamps := [...]bool{m.beansTank.Status(), m.waterTank.Status(),
		m.grindTank.Status()}
	for i, l := range lamps {
		if l != true {
			t.Error(fmt.Sprintf("%d wrong Check command", i))
		}
	}
}
