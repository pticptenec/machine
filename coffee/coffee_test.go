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

func TestGrindTank(t *testing.T) {
	gt := grindTank{
		lamp:  &lamp{},
		grind: 70,
	}

	if gt.Check() != true && gt.lamp.on != false {
		t.Errorf("wrong GrindTank Check method")
	}

	gt.grind = GrindTankMax
	if gt.Check() != false && gt.lamp.on != true {
		t.Errorf("wrong GrindTank Check method")
	}

	if gt.Do(10) != ErrTankNotReady {
		t.Errorf("wrong GrindTank Do method")
	}

	gt.grind = GrindTankMax - 1
	if gt.Do(10) != nil && gt.lamp.on != true {
		t.Errorf("wrong GrindTank Do method")
	}

	gt.grind = GrindTankMax - 20
	if gt.Do(10) != nil && gt.lamp.on != false &&
		gt.grind != GrindTankMax-10 {
		t.Errorf("wrong GrindTank Do method")
	}
}

func TestBeansTank(t *testing.T) {
	bt := beansTank{
		lamp:  &lamp{},
		beans: 70,
	}

	if bt.Check() != true && bt.lamp.on != false {
		t.Errorf("wrong BeansTank Check method")
	}

	bt.beans = BeansTankMin
	if bt.Check() != false && bt.lamp.on != true {
		t.Errorf("wrong GrindTank Check method")
	}

	if bt.Do(10) != ErrTankNotReady {
		t.Errorf("wrong BeansTank Do method")
	}

	bt.beans = BeansTankMin + 1
	if bt.Do(10) != nil && bt.lamp.on != true {
		t.Errorf("wrong BeansTank Do method")
	}

	bt.beans = BeansTankMin + 20
	if bt.Do(10) != nil && bt.lamp.on != false &&
		bt.beans != BeansTankMin+10 {
		t.Errorf("wrong BeansTank Do method")
	}
}

func TestWaterTank(t *testing.T) {
	wt := waterTank{
		lamp:  &lamp{},
		water: 70,
	}

	if wt.Check() != true && wt.lamp.on != false {
		t.Errorf("wrong WaterTank Check method")
	}

	wt.water = WaterTankMin
	if wt.Check() != false && wt.lamp.on != true {
		t.Errorf("wrong WaterTank Check method")
	}

	if wt.Do(10) != ErrTankNotReady {
		t.Errorf("wrong WaterTank Do method")
	}

	wt.water = WaterTankMin + 1
	if wt.Do(10) != nil && wt.lamp.on != true {
		t.Errorf("wrong WaterTank Do method")
	}

	wt.water = WaterTankMin + 20
	if wt.Do(10) != nil && wt.lamp.on != false &&
		wt.water != WaterTankMin+20 {
		t.Errorf("wrong WaterTank Do method")
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
