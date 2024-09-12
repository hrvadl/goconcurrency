package main

func NewTrafficLight() *TrafficLight {
	var (
		red    = newColor(redName)
		yellow = newColor(yellowName)
		green  = newColor(greenName)
	)

	red.setNext(yellow)
	yellow.setNext(green)
	green.setNext(red)

	return &TrafficLight{
		activeColor: red,
	}
}

type TrafficLight struct {
	activeColor *Color
}

func (t *TrafficLight) IsAllowedToRide() bool {
	return t.GetActiveColor().IsAllowedToRide()
}

func (t *TrafficLight) GetActiveColor() *Color {
	return t.activeColor
}

func (t *TrafficLight) NextColor() *Color {
	t.activeColor = t.activeColor.next
	return t.GetActiveColor()
}
