package main

import "fmt"

// Invoker
type Button struct {
	command Command
}

func (b *Button) press() {
	b.command.execute()
}

// Command interface
type Command interface {
	execute()
}

// Concrete command
type OnCommand struct {
	device Device
}

func (c *OnCommand) execute() {
	c.device.on()
}

type OffCommand struct {
	device Device
}

func (c *OffCommand) execute() {
	c.device.off()
}

// Receiver
type Device interface {
	on()
	off()
}

// Concrete Receiver
type Tv struct {
	isRunning bool
}

func (t *Tv) on() {
	t.isRunning = true
	fmt.Println("Turning tv on")
}
func (t *Tv) off() {
	t.isRunning = false
	fmt.Println("Turning tv off")
}

// Client
func main() {
	tv := &Tv{}

	onCommand := &OnCommand{
		device: tv,
	}

	offCommand := &OffCommand{
		device: tv,
	}

	onButton := &Button{
		command: onCommand,
	}

	offButton := &Button{
		command: offCommand,
	}

	onButton.press()
	offButton.press()
}
