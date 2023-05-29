package main

import (
	"context"
	"machine"

	keyboard "github.com/sago35/tinygo-keyboard"
	"github.com/sago35/tinygo-keyboard/keycodes/jp"
)

func main() {
	d := keyboard.New()

	gpioPins := []machine.Pin{
		machine.WIO_KEY_A,
		machine.WIO_KEY_B,
		machine.WIO_KEY_C,
		machine.WIO_5S_UP,
		machine.WIO_5S_LEFT,
		machine.WIO_5S_RIGHT,
		machine.WIO_5S_DOWN,
		machine.WIO_5S_PRESS,
	}

	for c := range gpioPins {
		gpioPins[c].Configure(machine.PinConfig{Mode: machine.PinInput})
	}

	// KeyMediaXXX will be supported starting with tinygo-0.28.
	d.AddGpioKeyboard(gpioPins, [][][]keyboard.Keycode{
		{
			{
				jp.KeyA,
				jp.KeyB,
				jp.KeyC,
				jp.KeyMediaVolumeInc,
				jp.KeyLeft,
				jp.KeyRight,
				jp.KeyMediaVolumeDec,
				jp.KeyEnter,
			},
		},
	})

	d.Loop(context.Background())
}
