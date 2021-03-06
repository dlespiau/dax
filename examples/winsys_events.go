package main

import (
	"fmt"

	"github.com/dlespiau/dax"
)

type eventsBasic struct {
	dax.Scene
}

func (s *eventsBasic) OnMouseMoved(x, y float32) {
	fmt.Printf("Mouse moved at (%.0f, %.0f)\n", x, y)
}

func (s *eventsBasic) OnMouseButtonPressed(b dax.MouseButton, x, y float32) {
	fmt.Printf("Button %v pressed at (%.0f, %.0f)\n", b, x, y)
}

func (s *eventsBasic) OnMouseButtonReleased(b dax.MouseButton, x, y float32) {
	fmt.Printf("Button %v released at (%.0f, %.0f)\n", b, x, y)
}

func (s *eventsBasic) OnRuneEntered(r rune) {
	fmt.Printf("Rune entered '%c'\n", r)
}

var winsysEventsExample = Example{
	Category:    CategoryWinsys,
	Name:        "Events",
	Description: "Print events received by the application on stdout",
	Scene:       &eventsBasic{},
}
