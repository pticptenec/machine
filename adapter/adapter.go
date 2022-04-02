package adapter

import (
	"fmt"
	"strings"

	"github.com/jroimartin/gocui"
)

const (
	NEWLINE = iota
	CENTER
	NEXT
)

type LayoutManager struct {
	*gocui.Gui
	names     []string
	positions map[string][2]int
}

func NewLayoutManager() (*LayoutManager, error) {
	g, err := gocui.NewGui(gocui.OutputNormal)
	names := make([]string, 0)
	positions := make(map[string][2]int, 0)
	l := &LayoutManager{g, names, positions}
	return l, err
}

func (l *LayoutManager) Add(c *Component, position int) {
	var err error
	if position == NEWLINE {
		l.addNewLine(c)
	} else if position == CENTER {
		err = l.addCenter(c)
	} else if position == NEXT {
		err = l.addNext(c)
	}

	if err != nil {
		l.addNewLine(c)
	}
}

func (l *LayoutManager) addNewLine(c *Component) {

}

type Component struct {
	Name           string
	Title          string
	body           string
	startX, startY int
	lastX, lastY   int
	handler        gocui.ManagerFunc
	layout         *LayoutManager
}

func NewComponent(name, body string, layout *LayoutManager) *Component {
	c := &Component{
		Name:    name,
		body:    body,
		handler: nil,
		layout:  layout,
	}
	//TODO Add pos
	layout.Add(c, 0)
	return c
}

func (c *Component) Layout(g *gocui.Gui) error {
	v, err := g.SetView(c.Name, c.startX, c.startY, c.lastX, c.lastY)
	// TODO
	// v.Title = "status:"
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprint(v, c.body)
	}
	return nil
}

func (c *Component) Size() (int, int) {
	lines := strings.Split(c.body, "\n")
	height := len(lines) + 1
	width := -1
	for _, line := range lines {
		cur := len(line)
		if cur > width {
			width = cur
		}
	}
	return width, height
}

func (c *Component) SetStartPos(x, y int) {
	c.startX = x
	c.startY = y
}

func (c *Component) SetEndPos(x, y int) {
	c.lastX = x
	c.lastY = y
}

func (sw *StatusWidget) Layout(g *gocui.Gui) error {
	v, err := g.SetView(sw.name, sw.x, sw.y, sw.x+sw.w, sw.y+sw.h)
	v.Title = "status:"
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprint(v, sw.body)
	}
	return nil
}

func (l *LayoutManager) addNewLine(c *Component, lastName string) {
	l.names = append(l.names, c.Name)
	lastPos := l.positions[lastName]
}

func (l *LayoutManager) Add(c *Component, pos int) {
	var err error
	if pos == NEWLINE {
		l.AddNewLine(c)
	} else if pos == CENTER {
		err = l.AddCenter(c)
	} else if pos == NEXT {
		err = l.AddNext(c)
	}
	if err != nil {
		l.AddNewLine(c)
	}
	if l.names == nil {
		l.names = make([]string, 0)
		l.positions = make(map[string][2]int, 0)
	}
	var last string
	if len(l.names) > 0 {
		last = l.names[len(l.names)-1]
	}

	l.names = append(l.names, c.Name)
	lastPos := l.positions[last]
	gapNewLine := 1
	x := lastPos[0]
	y := lastPos[1] + gapNewLine
	sizeX, sizeY := c.Size()
	l.positions[c.Name] = [2]int{x + sizeX, y + sizeY}
	c.SetEndPos(x+sizeX, y+sizeY)
	fmt.Println(c.startX, c.startY, c.lastX, c.lastY)
}

func (l *LayoutManager) Size(prevName string) (int, int) {
	if l.names == nil {
		return 1, 1
	}
	var prev string
	for _, name := range l.names {
		if name == prevName {
			prev = name
		}
	}
	pos := l.positions[prev]
	return pos[0], pos[1]