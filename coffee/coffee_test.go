package coffee

import (
	"testing"
)

func TestNewErrTank(t *testing.T) {
	e1 := NewErrTank("not defined")
	e2 := NewErrTank("not ready")
	e3 := NewErrTank("no such err")
	if e1 != ErrTankNotDefined ||
		e2 != ErrTankNotReady ||
		e3 != ErrTankNotDefined {
		t.Errorf("ErrTank constuctor error")
	}
}

var c Config

func TestNewMachine(t *testing.T) {

	c = Config{
		Water:       10,
		Beans:       20,
		Grind:       30,
		WaterHandle: 4,
		BeansHandle: 5,
	}
	m := NewMachine(c)
	if m.waterTank.water != 10 ||
		m.beansTank.beans != 20 ||
		m.grindTank.grind != 30 ||
		m.waterHandle.value != 4 ||
		m.beansHandle.value != 5 ||
		m.ready != false {
		t.Errorf("error with Machine constucotr, %v", m)
	}
}

func TestCircullarHandle(t *testing.T) {
	ch := circullarHandle{5}
	if ch.Get() != 5 {
		t.Errorf("wrong Get Method: %v", ch)
	}
	ch.Set(6)
	if ch.Get() != 6 {
		t.Errorf("wrong Set Method: %v", ch)
	}
	ch.Set(11)
	if ch.Get() != 6 {
		t.Errorf("wrong Set Method: %v", ch)
	}
}

func TestOnMachineChecksDone(t *testing.T) {
	c = Config{
		Water:       50,
		Beans:       30,
		Grind:       30,
		WaterHandle: 5,
		BeansHandle: 5,
	}
	m := NewMachine(c)

	m.On()

	lamps := [...]bool{m.beansTank.Status(), m.waterTank.Status(),
		m.grindTank.Status()}
	for _, l := range lamps {
		if l != false {
			t.Errorf("wrong Check command")
		}
	}

	if m.ready != true {
		t.Errorf("machine not ready")
	}
}

func TestOnMachineChecksDoneLampsOn(t *testing.T) {
	c = Config{
		Water:       5,
		Beans:       5,
		Grind:       110,
		WaterHandle: 5,
		BeansHandle: 5,
	}
	m := NewMachine(c)

	m.On()

	lamps := [...]bool{m.beansTank.Status(), m.waterTank.Status(),
		m.grindTank.Status()}
	for i, l := range lamps {
		if l != true {
			t.Errorf("%d wrong Check command", i)
		}
	}
	if m.ready != false {
		t.Errorf("machine not ready")
	}
}

func TestMakeEspresso(t *testing.T) {
	c = Config{
		Water: 50,
		Beans: 50,
		Grind: 50,
	}

	m := NewMachine(c)
	m.SetWaterHandle(50)
	m.SetGrindHandle(50)
	m.On()
	cof := m.Espresso()

}
