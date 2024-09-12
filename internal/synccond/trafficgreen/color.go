package main

const (
	redName    = "red"
	yellowName = "yellow"
	greenName  = "green"
)

func newColor(name string) *Color {
	return &Color{name: name}
}

func (c *Color) setNext(next *Color) {
	c.next = next
}

type Color struct {
	name string
	next *Color
}

func (c *Color) IsAllowedToRide() bool {
	return c.name == greenName
}
