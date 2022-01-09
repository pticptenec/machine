package machine

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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

func main() {
	commands := make(chan int)
	reader := bufio.NewReader(os.Stdin)
	go ListenCommands(commands)
	for {
		fmt.Print(">>>")
		str, _ := reader.ReadString('\n')
		num, err := strconv.Atoi(str)
		if err != nil {
			continue
		}
		commands <- num
	}
}

func ListenCommands(commands chan int) {
	for {
		fmt.Println("<<<", <-commands)
	}
}

type Config struct {
	Water       int
	Beans       int
	Grind       int
	WaterHandle int
	BeansHandle int
}

type Machine struct {
	ready       bool
	waterTank   WaterTank
	beansTank   BeansTank
	grindTank   GrindTank
	waterHandle CircullarHandle
	beansHandle CircullarHandle
	Display     RunLine
}

type CircullarHandle struct {
	value int
}

const CircullarHandleMax = 10
const CircullarHandleMin = 1

func (ch CircullarHandle) Set(qty int) {
	if qty < CircullarHandleMax && qty > CircullarHandleMin {
		ch.value = qty
	}
}

func (ch CircullarHandle) Get() int {
	return ch.value
}

func NewMachine(c Config) *Machine {
	lamps := [3]Lamp{{}, {}, {}}
	wt := WaterTank{water: c.Water, lamp: &lamps[0]}
	bt := BeansTank{beans: c.Beans, lamp: &lamps[1]}
	gt := GrindTank{grind: c.Grind, lamp: &lamps[2]}

	wh := CircullarHandle{c.WaterHandle}
	bh := CircullarHandle{c.BeansHandle}

	m := &Machine{
		waterTank:   wt,
		beansTank:   bt,
		grindTank:   gt,
		waterHandle: wh,
		beansHandle: bh,
	}
	return m
}

type RunLine struct {
	slug string
}

func (r *RunLine) Hello() {
	r.slug = fmt.Sprintf("Coffee?\n")
}

func (r *RunLine) Ready(water, beans bool) {
	r.slug = fmt.Sprintf("Status:\nW: %t, B: %t\n", water, beans)
}

func (r *RunLine) NotReady() {
	return
}

type GrindTank struct {
	lamp  *Lamp
	grind int
}

const GrindTankMax = 100

func (gt GrindTank) Check() bool {
	if gt.grind > GrindTankMax {
		gt.lamp.On()
		return false
	}
	return true
}

func (gt GrindTank) Status() bool {
	return gt.lamp.on
}

func (gt GrindTank) Do(grind int) error {
	if gt.Check() == false {
		return ErrTankNotReady
	}
	gt.grind += grind
	gt.Check()
	return nil
}

type BeansTank struct {
	lamp  *Lamp
	beans int
}

const BeansTankMin = 10

func (bt BeansTank) Check() bool {
	if bt.beans < BeansTankMin {
		bt.lamp.On()
		return false
	}
	return true
}

func (bt BeansTank) Status() bool {
	return bt.lamp.on
}

func (bt BeansTank) Do(quantity int) error {
	if bt.Check() == false {
		return ErrTankNotReady
	}
	bt.beans -= quantity
	_ = bt.Check()
	return nil
}

type WaterTank struct {
	lamp  *Lamp
	water int
}

const WaterTankMin = 10

func (wt WaterTank) Check() bool {
	if wt.water < WaterTankMin {
		wt.lamp.On()
		return false
	}
	return true
}

func (wt WaterTank) Status() bool {
	return wt.lamp.on
}

func (wt WaterTank) Do(quantity int) error {
	if wt.Check() == false {
		return ErrTankNotReady
	}
	wt.water -= quantity
	_ = wt.Check()
	return nil
}

type Lamp struct {
	on bool
}

func (l *Lamp) On() {
	l.on = true
}

func (l *Lamp) Off() {
	l.on = false
}

type Tank interface {
	Check() bool
	Do(int) error
	Status() bool
}

func (m *Machine) On() {
	m.Display.Hello()
	devices := [...]Tank{m.beansTank, m.waterTank, m.grindTank}
	var ready bool = false
	for _, d := range devices {
		ready = d.Check()
	}
	if !ready {
		m.Display.NotReady()
	} else {
		m.ready = true
		m.Display.Ready(m.waterTank.Status(), m.beansTank.Status())
	}
}

func (m *Machine) Off() {
	panic(OffCommand)
}

type Command struct {
	slug int
}

func (c Command) Int() int {
	return c.slug
}

var (
	OffCommand = Command{1}
	OnCommand  = Command{2}
)

func (m *Machine) MakeEspresso() {
	if !m.ready {
		return
	}

	done := make(chan struct{})
	go func() {
		m.waterTank.Do(m.waterHandle.Get())
		done <- struct{}{}
	}()
	go func() {
		m.beansTank.Do(m.beansHandle.Get())
		done <- struct{}{}
	}()
	go func() {
		m.grindTank.Do(m.beansHandle.Get())
		done <- struct{}{}
	}()
	var i int
	for range done {
		i++
		if i == 2 {
			fmt.Println("Done!")
		}
	}
}
