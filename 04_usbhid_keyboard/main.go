package main

import (
	"machine/usb/hid/keyboard"
	"time"
)

func main() {
	kb := keyboard.Port()

	time.Sleep(5 * time.Second)
	kb.Write([]byte("tinygo")) // 指定したキー列を入力
	time.Sleep(1 * time.Second)
	kb.Press(keyboard.KeySpace) // 単一キー入力
	kb.Press(keyboard.KeyV)
	kb.Press(keyboard.KeyE)
	kb.Press(keyboard.KeyR)
	kb.Press(keyboard.KeyS)
	kb.Press(keyboard.KeyI)
	kb.Press(keyboard.KeyO)
	kb.Press(keyboard.KeyN)
	kb.Down(keyboard.KeyEnter) // 押下
	kb.Up(keyboard.KeyEnter)   // 離す

	select {}
}
