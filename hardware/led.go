package hardware

import "machine"

type LEDAdapter struct {
	machine.Pin
}

func (l LEDAdapter) Off() {
	l.Low()
}

func (l LEDAdapter) On() {
	l.High()
}
