package coffee

import (
	"fmt"
	"strings"
)

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
	Do(int) error
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
	ready   bool
	lamps   [3]*lamp
	tanks   map[int]doesAction
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
	Beans int
	Water int
	Name  string
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

	lamps := [...]*lamp{{}, {}, {}}

	tanks := map[int]doesAction{
		grindTankKey: &grindTank{lamp: lamps[0], grind: c.Grind},
		beansTankKey: &beansTank{lamp: lamps[1], beans: c.Beans},
		waterTankKey: &waterTank{lamp: lamps[2], water: c.Water},
	}

	handles := map[int]handle{
		beansHandleKey: &circullarHandle{c.BeansHandle},
		waterHandleKey: &circullarHandle{c.WaterHandle},
	}

	return &Machine{
		ready:   false,
		lamps:   lamps,
		tanks:   tanks,
		handles: handles,
	}
}

func (m *Machine) On() {
	var ready bool = true
	for _, d := range m.tanks {
		ready = d.Check() && ready
	}
	m.ready = ready
}

func (m *Machine) Off() {
	for _, l := range m.lamps {
		l.Off()
	}
	m.ready = false
}

func (m *Machine) Espresso() *Coffee {
	const espressoMult = 1
	return m.makeCoffee(espressoMult)
}

func (m *Machine) Lungo() *Coffee {
	const lungoMult = 2
	return m.makeCoffee(lungoMult)
}

func (m *Machine) makeCoffee(espressoOrLungo int) *Coffee {
	if m == nil || m.ready == false {
		return nil
	}

	// TODO bug here
	var beans int = m.handles[beansHandleKey].Get()
	var water int = m.handles[waterHandleKey].Get() * espressoOrLungo

	m.ready = false
	isReady := make(chan bool, len(m.tanks))
	go func() {
		err := m.tanks[beansTankKey].Do(beans)
		if err != nil {
			isReady <- false
			return
		}
		isReady <- true
	}()
	go func() {
		err := m.tanks[waterTankKey].Do(water)
		if err != nil {
			isReady <- false
			return
		}
		isReady <- true
	}()
	// defer reverse order
	defer func() {
		var flag bool
		flag = <-isReady
		flag = <-isReady && flag
		flag = <-isReady && flag
		if flag == true {
			m.On()
		}
	}()
	defer func() {
		err := m.tanks[grindTankKey].Do(beans)
		if err != nil {
			isReady <- false
			return
		}
		isReady <- true
	}()

	name := "Espresso"
	if espressoOrLungo != 1 {
		name = "Lungo"
	}

	return &Coffee{
		Beans: beans,
		Water: water,
		Name:  name,
	}
}

func (m *Machine) Status() string {
	res := []string{}
	prefixes := [...]string{"g: ", "b: ", "w: "}
	for i, l := range m.lamps {
		if l.on == true {
			res = append(res, prefixes[i]+"on; ")
		} else {
			res = append(res, prefixes[i]+"off; ")
		}
	}
	status := strings.Join(res, "")
	return status[:len(status)-len(" ")]
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
	if gt.grind >= GrindTankMax {
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

func (bt beansTank) Check() bool {
	if bt.beans <= BeansTankMin {
		bt.lamp.On()
		return false
	}
	return true
}

func (bt beansTank) Do(quantity int) error {
	if bt.Check() == false {
		return ErrTankNotReady
	}
	bt.beans -= quantity
	_ = bt.Check()
	return nil
}

// 4.4 waterTank

const WaterTankMin = 10

func (wt waterTank) Check() bool {
	if wt.water <= WaterTankMin {
		wt.lamp.On()
		return false
	}
	return true
}

func (wt waterTank) Do(quantity int) error {
	if wt.Check() == false {
		return ErrTankNotReady
	}
	wt.water -= quantity
	_ = wt.Check()
	return nil
}

func (l *lamp) On() {
	l.on = true
}

func (l *lamp) Off() {
	l.on = false
}

// 4.5 Coffee

func (c Coffee) String() string {
	return fmt.Sprintf("b: %d, w: %d, name: %s coffee",
		c.Beans, c.Water, c.Name)
}
