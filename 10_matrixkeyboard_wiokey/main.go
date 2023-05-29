package main

import (
	"context"
	"machine"

	keyboard "github.com/sago35/tinygo-keyboard"
	"github.com/sago35/tinygo-keyboard/keycodes/jp"
)

func main() {
	d := keyboard.New()

	colPins := []machine.Pin{
		machine.BCM5,
		machine.BCM6,
	}

	rowPins := []machine.Pin{
		machine.BCM13, // machine.PA04
		machine.BCM19,
		machine.BCM26,
	}

	d.AddMatrixKeyboard(colPins, rowPins, [][][]keyboard.Keycode{
		{
			{jp.KeyT, jp.KeyI},
			{jp.KeyN, jp.KeyY},
			{jp.KeyG, jp.KeyO},
		},
	}, keyboard.InvertDiode(true))

	d.Loop(context.Background())
}
