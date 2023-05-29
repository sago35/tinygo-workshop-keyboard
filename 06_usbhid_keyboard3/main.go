package main

import (
	"machine"
	"machine/usb/hid/keyboard"
	"time"
)

func main() {
	// ボタンの定義はマイコン毎に変更する必要あり
	b1 := machine.BUTTON_1 // wioterminal の場合
	b2 := machine.BUTTON_2 // wioterminal の場合
	//b1 := machine.D0 // xiao-rp2040 の場合
	//b2 := machine.D1 // xiao-rp2040 の場合
	// b1 := machine.GP0 // pico の場合
	// b2 := machine.GP1 // pico の場合

	b1.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	b2.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

	kb := keyboard.Port()
	for {
		if b1.Get() {
			kb.Up(keyboard.KeyT)
		} else {
			kb.Down(keyboard.KeyT)
		}
		if b2.Get() {
			kb.Up(keyboard.KeyLeftShift)
		} else {
			kb.Down(keyboard.KeyLeftShift)
		}

		time.Sleep(time.Millisecond * 10)
	}
}
