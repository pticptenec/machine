package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/jroimartin/gocui"
)

type Content struct {
	Name           string
	Body           string
	startX, startY int
	endX, endY     int
	Title          string
}

func (c Content) Size() (int, int) {
	lines := strings.Split(c.Body, "\n")
	height := len(lines) + 1
	width := -1
	for _, line := range lines {
		cur := len(line)
		if cur > width {
			width = cur
		}
	}
	width += 1
	return width, height
}

type HeaderWidget struct {
	Content
}

func NewHeaderWidget(name, body string) *HeaderWidget {
	hw := &HeaderWidget{
		Content: Content{
			Name:   name,
			Body:   body,
			startX: 1,
			startY: 1,
		},
	}
	w, h := hw.Size()
	hw.endX = hw.startX + w
	lineGap := 1
	hw.endY = hw.endY + h + lineGap
	return hw
}

func (hw *HeaderWidget) Layout(g *gocui.Gui) error {
	v, err := g.SetView(hw.Name, hw.startX, hw.startY, hw.endX, hw.endY)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprint(v, hw.Body)
	}
	return nil
}

type DescriptionWidget struct {
	Content
}

func NewDescriptionWidget(name, body string) *DescriptionWidget {
	const lineHeight = 4
	dw := &DescriptionWidget{
		Content: Content{
			Name:   name,
			Body:   body,
			startX: 1,
			startY: lineHeight,
		},
	}
	w, h := dw.Size()
	dw.endX = dw.startX + w
	dw.endY = dw.startY + h
	return dw
}

func (dw *DescriptionWidget) Layout(g *gocui.Gui) error {
	v, err := g.SetView(dw.Name, dw.startX, dw.startY, dw.endX, dw.endY)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprint(v, dw.Body)
	}
	return nil
}

type StatusWidget struct {
	Content
}

func NewStatusWidget(name, body string) *StatusWidget {
	const lineHeight = 9
	const lineWidth = 27
	dw := &StatusWidget{
		Content: Content{
			Name:  name,
			Body:  body,
			Title: "status",
			endY:  lineHeight,
		},
	}
	w, h := dw.Size()
	dw.endX = lineWidth + w
	dw.startX = lineWidth
	dw.startY = dw.endY - h
	return dw
}

func (sw *StatusWidget) Layout(g *gocui.Gui) error {
	v, err := g.SetView(sw.Name, sw.startX, sw.startY, sw.endX, sw.endY)
	if err != nil {
		v.Title = sw.Title
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprint(v, sw.Body)
	}
	return nil
}

type WaterWidget struct {
	Content
}

func NewWaterWidget(name, body string) *WaterWidget {
	const lineHeight = 10
	ww := &WaterWidget{
		Content: Content{
			Name:   name,
			Body:   body,
			Title:  "w",
			startX: 1,
			startY: lineHeight,
		},
	}
	w, h := ww.Size()
	ww.endX = ww.startX + w
	ww.endY = ww.startY + h
	return ww
}

func (ww *WaterWidget) Layout(g *gocui.Gui) error {
	v, err := g.SetView(ww.Name, ww.startX, ww.startY, ww.endX, ww.endY)
	if err != nil {
		v.Title = ww.Title
		v.Editable = true
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprint(v, ww.Body)
	}
	return nil
}

type BeansWidget struct {
	Content
}

func NewBeansWidget(name, body string) *BeansWidget {
	const lineHeight = 10
	const lineWidth = 7
	bw := &BeansWidget{
		Content: Content{
			Name:   name,
			Body:   body,
			Title:  "b",
			startX: lineWidth,
			startY: lineHeight,
		},
	}
	w, h := bw.Size()
	bw.endX = bw.startX + w
	bw.endY = bw.startY + h
	return bw
}

func (bw *BeansWidget) Layout(g *gocui.Gui) error {
	v, err := g.SetView(bw.Name, bw.startX, bw.startY, bw.endX, bw.endY)
	if err != nil {
		v.Title = bw.Title
		v.Editable = true
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprint(v, bw.Body)
	}
	return nil
}

type EspressoWidget struct {
	Content
}

func NewEspressoWidget(name, body string) *EspressoWidget {
	const lineHeight = 10
	const lineWidth = 13
	ew := &EspressoWidget{
		Content: Content{
			Name:   name,
			Body:   body,
			startX: lineWidth,
			startY: lineHeight,
		},
	}
	w, h := ew.Size()
	ew.endX = ew.startX + w
	ew.endY = ew.startY + h
	return ew
}

func (ew *EspressoWidget) Layout(g *gocui.Gui) error {
	v, err := g.SetView(ew.Name, ew.startX, ew.startY, ew.endX, ew.endY)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprint(v, ew.Body)
	}
	return nil
}

type LungoWidget struct {
	Content
}

func NewLungoWidget(name, body string) *LungoWidget {
	const lineHeight = 10
	const lineWidth = 23
	lw := &LungoWidget{
		Content: Content{
			Name:   name,
			Body:   body,
			Title:  "",
			startX: lineWidth,
			startY: lineHeight,
		},
	}
	w, h := lw.Size()
	lw.endX = lw.startX + w
	lw.endY = lw.startY + h
	return lw
}

func (lw *LungoWidget) Layout(g *gocui.Gui) error {
	v, err := g.SetView(lw.Name, lw.startX, lw.startY, lw.endX, lw.endY)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprint(v, lw.Body)
	}
	return nil
}

type OffWidget struct {
	Content
}

func NewOffWidget(name, body string) *OffWidget {
	const lineHeight = 10
	const lineWidth = 30
	ow := &OffWidget{
		Content: Content{
			Name:   name,
			Body:   body,
			Title:  "",
			startX: lineWidth,
			startY: lineHeight,
		},
	}
	w, h := ow.Size()
	ow.endX = ow.startX + w
	ow.endY = ow.startY + h
	return ow
}

func (ow *OffWidget) Layout(g *gocui.Gui) error {
	v, err := g.SetView(ow.Name, ow.startX, ow.startY, ow.endX, ow.endY)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprint(v, ow.Body)
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

	// lm := NewLayoutManager()
	header := NewHeaderWidget(
		"header", strings.ToUpper("Office Coffee Machine"),
	)
	descr := NewDescriptionWidget("descr",
		`Tab: move betwen buttons
	Enter: push button
	Num Cell: enter 1-10
	^C, Exit Btn: Exit`)
	status := NewStatusWidget("status", "make a coffee")
	water := NewWaterWidget("water", "010")
	beans := NewBeansWidget("beans", "010")
	espresso := NewEspressoWidget("espresso", "Espresso")
	lungo := NewLungoWidget("lungo", "Lungo")
	off := NewOffWidget("off", "Off")
	g.SetManager(header, descr, status, water, beans, espresso, lungo, off)

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

func enterBeans(g *gocui.Gui, v *gocui.View) error {
	v.Editable = true
	return nil
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
