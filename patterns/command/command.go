package command

import "fmt"

type button struct {
	command command
}

func (b *button) press() {
	b.command.execute()
}

type device interface {
	on()
	off()
}

type tv struct {
	isRunning bool
}

func (t *tv) on() {
	t.isRunning = true
	fmt.Println("Turning tv on")
}

func (t *tv) off() {
	t.isRunning = false
	fmt.Println("Turning tv off")
}

type command interface {
	execute()
}

type OnCommand struct {
	device device
}

func (c *OnCommand) execute() {
	c.device.on()
}

type OffCommand struct {
	device device
}

func (c *OffCommand) execute() {
	c.device.off()
}

func main() {
	tv := &tv{}

	onCommand := &OnCommand{
		device: tv,
	}

	offCommand := &OffCommand{
		device: tv,
	}

	onButton := &button{
		command: onCommand,
	}
	onButton.press()

	offButton := &button{
		command: offCommand,
	}
	offButton.press()
}
