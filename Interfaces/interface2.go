package main

import "fmt"

type Device interface {
	TurnOn()
	TurnOff()
}

type Lamp struct {
	name string
}

func (l *Lamp) TurnOn() {
	fmt.Printf("Lamp %s is turned on!\n", l.name)
}

func (l *Lamp) TurnOff() {
	fmt.Printf("Lamp %s is turned off!\n", l.name)
}

type Fan struct {
	model string
}

func (f *Fan) TurnOn() {
	fmt.Printf("Fan %s is turned on!\n", f.model)
}

func (f *Fan) TurnOff() {
	fmt.Printf("Fan %s is turned off!\n", f.model)
}

type MockDevice struct{}

func (m *MockDevice) TurnOn() {
	fmt.Print("Mock Device turned on!\n")
}

func (m *MockDevice) TurnOff() {
	fmt.Print("Mock Device turned off!\n")
}

func TurnOnOff(d Device) {
	d.TurnOn()
	d.TurnOff()
}

func main() {
	TurnOnOff(&Lamp{"Лампа Джина"})
	TurnOnOff(&Fan{"Зил2109"})
	TurnOnOff(&MockDevice{})
}
