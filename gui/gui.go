package gui

import (
	"fmt"
	"log"
	"strings"

	"github.com/jroimartin/gocui"
)

type Layout struct {
	g          *gocui.Gui
	components map[string]*Component
	curX       int
	curY       int
}

type Component struct {
	Pos     [2]int
	Size    [2]int
	Color   gocui.Attribute
	Handler func() error
	paint   func(g *gocui.Gui) error
}

func (c *Component) Layout(g *gocui.Gui) error {
	if c.paint == nil {
		panic("wrong usage")
	}
	return c.paint(g)
}

func NewComponent(posX, posY, w, h int, color gocui.Attribute,
	handler func() error,
	paint func(*gocui.Gui) error) *Component {
	return &Component{
		Pos:     [2]int{posX, posY},
		Size:    [2]int{w, h},
		Color:   color,
		Handler: handler,
		paint:   paint,
	}
}

func NewLayout() (*Layout, error) {
	components := make(map[string]*Component)
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		return nil, err
	}

	g.Highlight = true
	g.SelFgColor = gocui.ColorRed

	return &Layout{
		g:          g,
		components: components,
		curX:       1,
		curY:       1,
	}, nil
}

type Position struct {
	slug int
}

func (p Position) Int() int {
	return p.slug
}

var (
	PositionCenter = Position{0}
	PositionLine   = Position{1}
)

func (l *Layout) CurPos() (int, int) {
	return l.curX, l.curY
}

func (l *Layout) Add(e *Element, p Position, color gocui.Attribute,
	handler func() error) {
	startX, startY := l.CurPos()
	w, h := e.ContentSize()
	maxX, maxY := l.g.Size()
	endX, endY := l.Locate(w, h, maxX, maxY, p)

	paintF := func(g *gocui.Gui) error {
		v, err := g.SetView(e.name, startX, startY, endX, endY)
		if err != nil && err != gocui.ErrUnknownView {
			return err
		}
		v.Clear()
		// g.SelBgColor = color
		// v.Highlight = true
		fmt.Fprint(v, e.content)
		return nil
	}
	c := NewComponent(startX+w, startY+h, w, h, color, nil, paintF)
	l.components[e.name] = c
}

func (l *Layout) Locate(startX, startY, endX, endY int, p Position) (int, int) {
	l.curX += startX
	l.curY += startY
	return l.curX, l.curY
}

func (l *Layout) OnClose() {
	l.g.Close()
}

func (l *Layout) Listen() {

	arr := make([]gocui.Manager, 0, 10)
	for _, val := range l.components {
		arr = append(arr, gocui.Manager(val))
	}

	l.g.SetManager(arr...)

	if err := l.g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}
	if err := l.g.SetKeybinding("", gocui.KeyTab, gocui.ModNone, l.makeSortThrough()); err != nil {
		log.Panicln(err)
	}
	if err := l.g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func (l *Layout) makeSortThrough() func(*gocui.Gui, *gocui.View) error {
	arrNames := make([]string, 0, 10)
	for k, _ := range l.components {
		arrNames = append(arrNames, k)
	}
	var i int
	return func(g *gocui.Gui, v *gocui.View) error {
		name := arrNames[i]
		i++
		i %= len(arrNames)
		if v != nil && v.Name() == name {
			_, err := g.SetCurrentView(name)
			return err
		}
		return nil
	}
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func main() {
	l, err := NewLayout()
	if err != nil {
		log.Panicln(err)
	}
	defer l.OnClose()

	header := NewElement("header", strings.ToUpper("Office Coffee Machine"))
	help := NewElement("help", "Press Tab to switch button\nPress Ctrl+C to Exit\nPress Enter to Push Button\n")
	layout, _ := NewLayout()

	layout.Add(header, PositionCenter, gocui.ColorBlue, nil)
	layout.Add(help, PositionCenter, gocui.ColorBlue, nil)

	layout.Listen()
}

type Element struct {
	name    string
	content string
}

func NewElement(name, content string) *Element {
	return &Element{
		name:    name,
		content: content,
	}
}

func (e *Element) ContentSize() (width, height int) {
	lines := strings.Split(e.content, "\n")
	var maxLen int
	for _, s := range lines {
		if _len := len(s); _len > maxLen {
			maxLen = _len
		}
	}

	borderWidth := 1
	width = maxLen + borderWidth
	height = len(lines) + borderWidth
	return
}
