package main

import (
	"fmt"
	"testing"
)

func TestAddCenter(t *testing.T) {
	e := NewElement("")
	e.x = 4
	e.y = 5
	l := NewLineLayout()
	defer l.Close()
	l.AddCenter(e)

	endX, endY := l.Size()
	fmt.Println(endX, endY)

}
