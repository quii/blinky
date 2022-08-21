package blinky

import "time"

type Light interface {
	Off()
	On()
}

type Wait func()

var (
	DefaultWait = func() { time.Sleep(1 * time.Second) }
)

func Blink(light Light, wait Wait) {
	for {
		light.Off()
		wait()
		light.On()
		wait()
	}
}
