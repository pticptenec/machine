package coffee

import "fmt"

// 1. interfaces
// 2. structs
// 3. errors
// 4. methods

// 1. interfaces

type handle interface {
	Set(int)
	Get() int
}

type light interface {
	On()
	Off()
}

type doesAction interface {
	Do(int)
	Check() bool
}

// 2. structs

type Config struct {
	Water       int
	Beans       int
	Grind       int
	WaterHandle int
	BeansHandle int
}

const (
	waterTankKey = iota
	beansTankKey
	grindTankKey
	waterHandleKey
	beansHandleKey
)

type Machine struct {
	ready       bool
	lamps [3]*lamp
	tanks map[int]doesAction
	handles map[int]handle
}

type circullarHandle struct {
	value int
}

type grindTank struct {
	lamp  *lamp
	grind int
}

type beansTank struct {
	lamp  *lamp
	beans int
}

type waterTank struct {
	lamp  *lamp
	water int
}

type lamp struct {
	on bool
}

type Coffee struct {
	beans int
	water int
	grind int
}

// 3. errors

type ErrTank struct {
	slug string
}

func (et ErrTank) Error() string {
	return et.slug
}

var (
	ErrTankNotDefined = ErrTank{"not defined"}
	ErrTankNotReady   = ErrTank{"not ready"}
)

func NewErrTank(s string) ErrTank {
	switch s {
	case ErrTankNotReady.slug:
		return ErrTankNotReady
	}
	return ErrTankNotDefined
}

// 4. methods

// 4.0 machine

func NewMachine(c Config) *Machine {

	lamps := [...]*lamp{&lamp{}, &lamp{}, &lamp{}}

	tanks := map[int]doesAction{
		grindTankKey: grindTank{lamp: lamps[0], grind: c.Grind},
		beansTankKey: beansTank{lamp: lamps[1], beans: c.Beans},
		waterTankKey: waterTank{lamp: lamps[2], water: c.Water},
	}

	handles := map[int]handle{
		beansHandleKey: circullarHandle{c.BeansHandle},
		waterHandleKey: circullarHandle{c.WaterHandle},
	}

	return &Machine{
		ready: false,
		lamps: lamps,
		tanks: tanks,
		handles: handles,
	}
}

func (m *Machine) On() {
	var ready bool = true
	for _, d := range m.tanks {
		ready = ready && d.Check()
	}
	m.ready = ready
}

func (m *Machine) Off() {
	for _, l := range m.lamps {
		l.Off()
	}
	m.ready = false
}

func (m *Machine) Espresso() (*Coffee, error) {
	if m.ready == false {
		return nil, ErrTankNotReady
	}

	

	return Coffee{
		Beans: m.beans,
		Water: m.water,
		Name: "espresso"
	}
}

func (m *Machine) Lungo() Coffee {

}

// 4.1 handle
const CircullarHandleMax = 10
const CircullarHandleMin = 1

func (ch *circullarHandle) Set(qty int) {
	if qty < CircullarHandleMax && qty > CircullarHandleMin {
		ch.value = qty
	}
}

func (ch circullarHandle) Get() int {
	return ch.value
}

// 4.2 grindTank

const GrindTankMax = 100

func (gt grindTank) Check() bool {
	if gt.grind > GrindTankMax {
		gt.lamp.On()
		return false
	}
	return true
}

func (gt grindTank) Do(grind int) error {
	if gt.Check() == false {
		return ErrTankNotReady
	}
	gt.grind += grind
	gt.Check()
	return nil
}

// 4.3 beansTank

const BeansTankMin = 10

func (bt BeansTank) Check() bool {
	if bt.beans < BeansTankMin {
		bt.lamp.On()
		return false
	}
	return true
}

func (bt BeansTank) Do(quantity int) error {
	if bt.Check() == false {
		return ErrTankNotReady
	}
	bt.beans -= quantity
	_ = bt.Check()
	return nil
}

// 4.4 waterTank

const WaterTankMin = 10

func (wt WaterTank) Check() bool {
	if wt.water < WaterTankMin {
		wt.lamp.On()
		return false
	}
	return true
}

func (wt WaterTank) Do(quantity int) error {
	if wt.Check() == false {
		return ErrTankNotReady
	}
	wt.water -= quantity
	_ = wt.Check()
	return nil
}

func (l *Lamp) On() {
	l.on = true
}

func (l *Lamp) Off() {
	l.on = false
}

// 4.5 Coffee

func (c Coffee) String() string {
	return fmt.Sprintf("b: %d, w: %d, name: %s coffee",
		c.Beans, c.Water, c.Name)
}
