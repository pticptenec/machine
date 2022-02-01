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
	// TODO
	t.Errorf("%v", m)
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

	// TODO
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

	// TODO
}

func TestMakeEspresso(t *testing.T) {
	c = Config{
		Water: 50,
		Beans: 50,
		Grind: 50,
	}

	m := NewMachine(c)
	m.On()
	coffee := m.Espresso()
	if coffee.String() != "b: 50, w: 50, Espresso coffee" {
		t.Errorf("method Espresso not works")
	}
}
