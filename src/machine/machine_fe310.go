// +build fe310

package machine

import (
	"device/sifive"
)

type PinMode uint8

const (
	PinInput PinMode = iota
	PinOutput
)

const (
	P00 Pin = 0
	P01 Pin = 1
	P02 Pin = 2
	P03 Pin = 3
	P04 Pin = 4
	P05 Pin = 5
	P06 Pin = 6
	P07 Pin = 7
	P08 Pin = 8
	P09 Pin = 9
	P10 Pin = 10
	P11 Pin = 11
	P12 Pin = 12
	P13 Pin = 13
	P14 Pin = 14
	P15 Pin = 15
	P16 Pin = 16
	P17 Pin = 17
	P18 Pin = 18
	P19 Pin = 19
	P20 Pin = 20
	P21 Pin = 21
	P22 Pin = 22
	P23 Pin = 23
	P24 Pin = 24
	P25 Pin = 25
	P26 Pin = 26
	P27 Pin = 27
	P28 Pin = 28
	P29 Pin = 29
	P30 Pin = 30
	P31 Pin = 31
)

// Configure this pin with the given configuration.
func (p Pin) Configure(config PinConfig) {
	sifive.GPIO0.INPUT_EN.SetBits(1 << uint8(p))
	if config.Mode == PinOutput {
		sifive.GPIO0.OUTPUT_EN.SetBits(1 << uint8(p))
	}
}

// Set the pin to high or low.
func (p Pin) Set(high bool) {
	if high {
		sifive.GPIO0.PORT.SetBits(1 << uint8(p))
	} else {
		sifive.GPIO0.PORT.ClearBits(1 << uint8(p))
	}
}

type UART struct {
	Bus    *sifive.UART_Type
	Buffer *RingBuffer
}

var (
	UART0 = UART{Bus: sifive.UART0, Buffer: NewRingBuffer()}
)

func (uart UART) Configure(config UARTConfig) {
	// Assuming a 16Mhz Crystal (which is Y1 on the HiFive1), the divisor for a
	// 115200 baud rate is 138.
	sifive.UART0.DIV.Set(138)
	sifive.UART0.TXCTRL.Set(sifive.UART_TXCTRL_ENABLE)
}

func (uart UART) WriteByte(c byte) {
	for sifive.UART0.TXDATA.Get()&sifive.UART_TXDATA_FULL != 0 {
	}

	sifive.UART0.TXDATA.Set(uint32(c))
}
