package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/jroimartin/gocui"
)

type HeaderWidget struct {
	name string
	x, y int
	w, h int
	body string
}

func NewHeaderWidget(x, y int) *HeaderWidget {
	var body = strings.ToUpper("Office Coffee Machine")
	w := len(body) + 1
	h := 2
	return &HeaderWidget{
		name: "header",
		x:    x,
		y:    y,
		w:    w,
		h:    h,
		body: body,
	}
}

func (hw *HeaderWidget) Layout(g *gocui.Gui) error {
	v, err := g.SetView(hw.name, hw.x, hw.y, hw.x+hw.w, hw.y+hw.h)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprint(v, hw.body)
	}
	return nil
}

type DescriptionWidget struct {
	name string
	x, y int
	w, h int
	body string
}

func NewDescriptionWidget(x, y int) *DescriptionWidget {
	const body = `Tab: move betwen buttons
Enter: push button
Num Cell: enter 1-10
^C, Exit Btn: Exit`
	maxLen := 0
	lines := strings.Split(body, "\n")
	for _, l := range lines {
		curLen := len(l)
		if maxLen < curLen {
			maxLen = curLen
		}
	}

	w := maxLen + 1
	h := len(lines) + 1
	return &DescriptionWidget{
		name: "description",
		x:    x,
		y:    y,
		w:    w,
		h:    h,
		body: body,
	}
}

func (dw *DescriptionWidget) Layout(g *gocui.Gui) error {
	v, err := g.SetView(dw.name, dw.x, dw.y, dw.x+dw.w, dw.y+dw.h)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprint(v, dw.body)
	}
	return nil
}

type EspressoWidget struct {
	name string
	x, y int
	w, h int
	body string
}

func NewEspressoWidget(x, y int) *EspressoWidget {
	var body = "Espresso"
	return &EspressoWidget{
		name: "espresso",
		x:    x,
		y:    y,
		w:    len(body) + 1,
		h:    2,
		body: body,
	}
}

func (ew *EspressoWidget) Layout(g *gocui.Gui) error {
	v, err := g.SetView(ew.name, ew.x, ew.y, ew.x+ew.w, ew.y+ew.h)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprint(v, ew.body)
	}
	return nil
}

type LungoWidget struct {
	name string
	x, y int
	w, h int
	body string
}

func NewLungoWidget(x, y int) *LungoWidget {
	var body = "Lungo"
	return &LungoWidget{
		name: "lungo",
		x:    x,
		y:    y,
		w:    len(body) + 1,
		h:    2,
		body: body,
	}
}

func (lw *LungoWidget) Layout(g *gocui.Gui) error {
	v, err := g.SetView(lw.name, lw.x, lw.y, lw.x+lw.w, lw.y+lw.h)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprint(v, lw.body)
	}
	return nil
}

type OffWidget struct {
	name    string
	x, y    int
	w, h    int
	body    string
	handler func(*gocui.Gui, *gocui.View) error
}

func NewOffWidget(x, y int, handler func(*gocui.Gui, *gocui.View) error) *OffWidget {
	var body = "Off"
	return &OffWidget{
		name:    "off",
		x:       x,
		y:       y,
		w:       len(body) + 1,
		h:       2,
		body:    body,
		handler: handler,
	}
}

func (ow *OffWidget) Layout(g *gocui.Gui) error {
	v, err := g.SetView(ow.name, ow.x, ow.y, ow.x+ow.w, ow.y+ow.h)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		if err := g.SetKeybinding(ow.name, gocui.KeyEnter, gocui.ModNone, ow.handler); err != nil {
			return err
		}
		fmt.Fprint(v, ow.body)
	}
	return nil
}

type WaterWidget struct {
	name    string
	x, y    int
	w, h    int
	body    string
	handler func(*gocui.Gui, *gocui.View) error
	val     int
}

func NewWaterWidget(x, y int, handler func(*gocui.Gui, *gocui.View) error) *WaterWidget {
	var body = " 4 "
	return &WaterWidget{
		name:    "water",
		x:       x,
		y:       y,
		w:       len(body) + 1,
		h:       2,
		body:    body,
		handler: handler,
		val:     4,
	}
}

func (ww *WaterWidget) Layout(g *gocui.Gui) error {
	v, err := g.SetView(ww.name, ww.x, ww.y, ww.x+ww.w, ww.y+ww.h)
	v.Title = "w:"
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		if err := g.SetKeybinding(ww.name, gocui.KeyEnter, gocui.ModNone, ww.handler); err != nil {
			return err
		}
		fmt.Fprint(v, ww.body)
	}
	return nil
}

type BeansWidget struct {
	name    string
	x, y    int
	w, h    int
	body    string
	handler func(*gocui.Gui, *gocui.View) error
	val     int
}

func NewBeansWidget(x, y int, handler func(*gocui.Gui, *gocui.View) error) *BeansWidget {
	var body = " 9 "
	return &BeansWidget{
		name:    "beans",
		x:       x,
		y:       y,
		w:       len(body) + 1,
		h:       2,
		body:    body,
		handler: handler,
		val:     9,
	}
}

func (bw *BeansWidget) Layout(g *gocui.Gui) error {
	v, err := g.SetView(bw.name, bw.x, bw.y, bw.x+bw.w, bw.y+bw.h)
	v.Title = "b:"
	v.Editable = true
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		if err := g.SetKeybinding(bw.name, gocui.KeyEnter, gocui.ModNone, bw.handler); err != nil {
			return err
		}
		fmt.Fprint(v, bw.body)
	}
	return nil
}

type StatusWidget struct {
	name string
	x, y int
	w, h int
	body string
}

func NewStatusWidget(x, y int) *StatusWidget {
	var body = " coffee status "
	return &StatusWidget{
		name: "status",
		x:    x,
		y:    y,
		w:    len(body) + 1,
		h:    2,
		body: body,
	}
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

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Highlight = true
	g.SelFgColor = gocui.ColorRed

	header := NewHeaderWidget(1, 1)
	descr := NewDescriptionWidget(1, 4)
	status := NewStatusWidget(27, 4)
	water := NewWaterWidget(1, 10, nil)
	beans := NewBeansWidget(7, 10, nil)
	espresso := NewEspressoWidget(14, 10)
	lungo := NewLungoWidget(24, 10)
	off := NewOffWidget(31, 10, quit)

	g.SetManager(header, descr, status, beans, water, espresso, lungo, off)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("", gocui.KeyTab, gocui.ModNone, toggleButton); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func toggleButton(g *gocui.Gui, v *gocui.View) error {
	const num = 5
	names := [num]string{"water", "beans", "espresso", "lungo", "off"}
	nextview := names[0]
	for i, name := range names {
		if v != nil && v.Name() == name {
			nextview = names[(i+1)%num]
			break
		}
	}
	_, err := g.SetCurrentView(nextview)
	return err
}
