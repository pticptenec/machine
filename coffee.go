package machine

import "fmt"

func main() {

}

type CmdReader interface {
	GetCmd() Command
}

type MemCommands struct {
	commands []Command
}

func NewMemCommands() *MemCommands {
	mc := MemCommands{}
	mc.commands = make([]Command, 0)
	return &mc
}

func (mc *MemCommands) Push(cmd Command) {
	mc.commands = append(mc.commands, cmd)
}

func (mc *MemCommands) GetCmd() Command {
	cmd := mc.commands[0]
	mc.commands = mc.commands[1:]
	return cmd
}

type Machine struct {
	devices []Device
	maps    map[Command]func()
}

func NewMachine() *Machine {
	m := &Machine{}
	_maps := map[Command]func(){
		OffCommand: m.Off,
	}
	m.maps = _maps
	return m
}

func (m *Machine) On(input CmdReader) (exit string) {
	defer func() {
		if r := recover(); r.(Command) == OffCommand {
			exit = fmt.Sprintf("Exit by %d Command", OffCommand)
		} else {
			panic(r)
		}
	}()

	for _, d := range m.devices {
		d.Check()
	}

	for {
		cmd := input.GetCmd()
		exec, ok := m.maps[cmd]
		if ok == false {
			panic("something wrong")
		}
		exec()
	}
}

func (m *Machine) Off() {
	panic(OffCommand)
}

type Device interface {
	Check()
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
