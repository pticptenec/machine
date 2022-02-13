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

func TestLamp(t *testing.T) {
	l := &lamp{}
	l.On()
	if l.on != true {
		t.Errorf("wrong lamp On methond")
	}
	l.Off()
	if l.on != false {
		t.Errorf("wrong lamp Off method")
	}
}

func TestCoffee(t *testing.T) {
	c := Coffee{Beans: 10, Water: 20, Name: "Lungo"}
	if s := c.String(); s != "b: 10, w: 20, name: Lungo coffee" {
		t.Errorf("wrong Coffee String method: %s", s)
	}
}

var c Config

func TestNewMachine(t *testing.T) {
	c = Config{
		Water:       50,
		Beans:       20,
		Grind:       30,
		WaterHandle: 4,
		BeansHandle: 5,
	}
	m := NewMachine(c)

	if m.ready != false {
		t.Errorf("wrong NewMachineMethod, m.ready")
	}

	for _, l := range m.lamps {
		if l == nil || l.on != false {
			t.Errorf("wrong lamps configured")
		}
	}

	// tanks

	wtPtr, ok := m.tanks[waterTankKey].(*waterTank)
	if ok != true && wtPtr.water != c.Water {
		t.Errorf("wrong Water Tank configured, %t, %v", ok, wtPtr)
	}

	btPtr, ok := m.tanks[beansTankKey].(*beansTank)
	if ok != true && btPtr.beans != c.Beans {
		t.Errorf("wrong Beans Tank configured, %t, %v", ok, btPtr)
	}

	gtPtr, ok := m.tanks[grindTankKey].(*grindTank)
	if ok != true && gtPtr.grind != c.Grind {
		t.Errorf("wrong Gring Tank configured, %t, %v", ok, btPtr)
	}

	// handles
	bh, ok := m.handles[beansHandleKey].(*circullarHandle)
	if ok != true && bh.value != c.BeansHandle {
		t.Errorf("wrong handle configured for Beans")
	}

	wh, ok := m.handles[waterHandleKey].(*circullarHandle)
	if ok != true && wh.value != c.WaterHandle {
		t.Errorf("wrong handle configured for Water")
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

	if m.ready != true {
		t.Errorf("wrong On method on ready sign")
	}

	for _, l := range m.lamps {
		if l.on != false {
			t.Errorf("wrong On methond on lamps")
		}
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

	if m.ready != false {
		t.Errorf("wrong On method on ready sign")
	}

	for _, l := range m.lamps {
		if l.on != true {
			t.Errorf("wrong On methond on lamps")
		}
	}
}

func TestOnMachineChecksDoneLampsPartial(t *testing.T) {
	c = Config{
		Water:       50,
		Beans:       50,
		Grind:       110,
		WaterHandle: 5,
		BeansHandle: 5,
	}
	m := NewMachine(c)

	m.On()

	if m.ready != false {
		t.Errorf("wrong On method on ready sign")
	}

	for _, l := range m.lamps[1:] {
		if l.on != false {
			t.Errorf("wrong On methond on lamps")
		}
	}

	if m.lamps[0].on != true {
		t.Errorf("wrong On methond on lamps (2) grind lamp")
	}
}

func TestMakeEspresso(t *testing.T) {
	c = Config{
		Water:       50,
		Beans:       50,
		Grind:       50,
		WaterHandle: 7,
		BeansHandle: 7,
	}

	m := NewMachine(c)
	m.On()
	coffee := m.Espresso()
	if coffee != nil && coffee.String() != "b: 7, w: 7, name: Espresso coffee" {
		t.Errorf("method Espresso not works: %v", coffee)
	}
}

func TestMakeLungo(t *testing.T) {
	c = Config{
		Water:       50,
		Beans:       50,
		Grind:       50,
		WaterHandle: 7,
		BeansHandle: 7,
	}

	m := NewMachine(c)
	m.On()
	coffee := m.Lungo()
	if coffee != nil && coffee.String() != "b: 7, w: 14, name: Lungo coffee" {
		t.Errorf("method Lungo not works: %v", coffee)
	}
}

func TestMakeTriple(t *testing.T) {
	c = Config{
		Water:       99,
		Beans:       99,
		Grind:       1,
		WaterHandle: 5,
		BeansHandle: 5,
	}

	m := NewMachine(c)
	m.On()
	lungos := [3]*Coffee{m.Lungo(), m.Lungo(), m.Lungo()}
	espressos := [3]*Coffee{m.Espresso(), m.Espresso(), m.Espresso()}
	for _, l := range lungos {
		if l != nil && l.String() != "b: 5, w: 10, name: Lungo coffee" {
			t.Errorf("method Lungo not works: %v", l)
		}
	}
	for _, e := range espressos {
		if e != nil && e.String() != "b: 5, w: 5, name: Espresso coffee" {
			t.Errorf("method Lungo not works: %v", e)
		}
	}
}

func TestMachineLamps(t *testing.T) {
	c = Config{
		Water:       5,
		Beans:       99,
		Grind:       1,
		WaterHandle: 5,
		BeansHandle: 5,
	}

	m := NewMachine(c)
	m.On()
	stat := m.Status()
	if stat != "g: off; b: off; w: on;" {
		t.Errorf("method Status on Machine not work: [%s]", stat)
	}
}
