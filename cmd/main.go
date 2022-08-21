package main

import (
	"blinky"
	"blinky/hardware"
	"machine"
)

func main() {
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	blinky.Blink(hardware.LEDAdapter{Pin: led}, blinky.DefaultWait)
}
