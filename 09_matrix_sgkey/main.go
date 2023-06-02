package main

import (
	"machine"
	"machine/usb/hid/keyboard"
	"time"
)

var (
	c1 = machine.D8
	c2 = machine.D9
	c3 = machine.D10
	r1 = machine.D1
	r2 = machine.D2
)

func main() {
	kb := keyboard.Port()

	c1.Configure(machine.PinConfig{Mode: machine.PinOutput})
	c2.Configure(machine.PinConfig{Mode: machine.PinOutput})
	c3.Configure(machine.PinConfig{Mode: machine.PinOutput})
	r1.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	r2.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	for {
		// SW1 / SW2 を読み込むため、 c1 を H、 c2 と c3 を L に
		c1.High()
		c2.Low()
		c3.Low()
		time.Sleep(1 * time.Millisecond) // 少し待つ

		// SW1
		if r1.Get() {
			kb.Down(keyboard.KeyA)
		} else {
			kb.Up(keyboard.KeyA)
		}

		// SW2
		if r2.Get() {
			kb.Down(keyboard.KeyB)
		} else {
			kb.Up(keyboard.KeyB)
		}

		// SW3 / SW4 を読み込むため、 c2 を H、 c1 と c3 を L に
		c1.Low()
		c2.High()
		c3.Low()
		time.Sleep(1 * time.Millisecond) // 少し待つ

		// SW3
		if r1.Get() {
			kb.Down(keyboard.KeyC)
		} else {
			kb.Up(keyboard.KeyC)
		}

		// SW4
		if r2.Get() {
			kb.Down(keyboard.KeyD)
		} else {
			kb.Up(keyboard.KeyD)
		}

		// SW5 / SW6 を読み込むため、 c3 を H、 c1 と c2 を L に
		c1.Low()
		c2.Low()
		c3.High()
		time.Sleep(1 * time.Millisecond) // 少し待つ

		// SW5
		if r1.Get() {
			kb.Down(keyboard.KeyE)
		} else {
			kb.Up(keyboard.KeyE)
		}

		// SW6
		if r2.Get() {
			kb.Down(keyboard.KeyF)
		} else {
			kb.Up(keyboard.KeyF)
		}

		time.Sleep(16 * time.Millisecond)
	}
}
