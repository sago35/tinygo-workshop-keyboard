package main

import (
	"machine"
	"machine/usb/hid/keyboard"
	"time"
)

var (
	c1 = machine.BCM5
	c2 = machine.BCM6
	r1 = machine.BCM13 // machine.PA04
	r2 = machine.BCM19
	r3 = machine.BCM26
)

func main() {
	kb := keyboard.Port()

	c1.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	c2.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	r1.Configure(machine.PinConfig{Mode: machine.PinOutput})
	r2.Configure(machine.PinConfig{Mode: machine.PinOutput})
	r3.Configure(machine.PinConfig{Mode: machine.PinOutput})
	for {
		// SW1 / SW2 を読み込むため、 r1 を H、 r2 と r3 を L に
		r1.High()
		r2.Low()
		r3.Low()
		time.Sleep(1 * time.Millisecond) // 少し待つ

		// SW1
		if c1.Get() {
			kb.Down(keyboard.KeyA)
		} else {
			kb.Up(keyboard.KeyA)
		}

		// SW2
		if c2.Get() {
			kb.Down(keyboard.KeyB)
		} else {
			kb.Up(keyboard.KeyB)
		}

		// SW3 / SW4 を読み込むため、 r2 を H、 r1 と r3 を L に
		r1.Low()
		r2.High()
		r3.Low()
		time.Sleep(1 * time.Millisecond) // 少し待つ

		// SW3
		if c1.Get() {
			kb.Down(keyboard.KeyC)
		} else {
			kb.Up(keyboard.KeyC)
		}

		// SW4
		if c2.Get() {
			kb.Down(keyboard.KeyD)
		} else {
			kb.Up(keyboard.KeyD)
		}

		// SW5 / SW6 を読み込むため、 r3 を H、 r1 と r2 を L に
		r1.Low()
		r2.Low()
		r3.High()
		time.Sleep(1 * time.Millisecond) // 少し待つ

		// SW5
		if c1.Get() {
			kb.Down(keyboard.KeyE)
		} else {
			kb.Up(keyboard.KeyE)
		}

		// SW6
		if c2.Get() {
			kb.Down(keyboard.KeyF)
		} else {
			kb.Up(keyboard.KeyF)
		}

		time.Sleep(16 * time.Millisecond)
	}
}
