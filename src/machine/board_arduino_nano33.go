// +build sam,atsamd21,arduino_nano33

// This contains the pin mappings for the Arduino Nano33 IoT board.
//
// For more information, see: https://store.arduino.cc/nano-33-iot
//
package machine

import "device/sam"

// GPIO Pins
const (
	RX0 = PA11 // UART0 RX
	TX1 = PA10 // UART0 TX
	D2  = PA14
	D3  = PA09 // PWM available
	D4  = PA08
	D5  = PA15 // PWM available
	D6  = PA20 // PWM available
	D7  = PA21
	D8  = PA06
	D9  = PA07 // PWM available
	D10 = PA18 // PWM available
	D11 = PA16 // PWM available
	D12 = PA19 // PWM available
	D13 = PA17 // PWM available
)

// Analog pins
const (
	A0 = PA02 // ADC/AIN[0]
	A1 = PB08 // ADC/AIN[2]
	A2 = PB09 // ADC/AIN[3]
	A3 = PA04 // ADC/AIN[4]
	A4 = PA05 // ADC/AIN[5]
	A5 = PB02 // ADC/AIN[10]
)

const (
	LED = D13
)

// NINA-W102 Pins

const (
	NINA_MOSI   = PA12
	NINA_MISO   = PA13
	NINA_CS     = PA14
	NINA_SCK    = PA15
	NINA_GPIO0  = PA27
	NINA_RESETN = PA08
	NINA_ACK    = PA28
)

// UART0 aka USBCDC pins
const (
	USBCDC_DM_PIN = PA24
	USBCDC_DP_PIN = PA25
)

// UART1 on the Arduino Nano 33 connects to the onboard NINA-W102 WiFi chip.
var (
	UART1 = UART{Bus: sam.SERCOM5_USART,
		Buffer: NewRingBuffer(),
		Mode:   PinSERCOMAlt,
		IRQVal: sam.IRQ_SERCOM5,
	}
)

// UART1 pins
const (
	UART_TX_PIN = PA22
	UART_RX_PIN = PA23
)

//go:export SERCOM5_IRQHandler
func handleUART1() {
	defaultUART1Handler()
}

// I2C pins
const (
	SDA_PIN = PB08 // SDA: SERCOM4/PAD[0]
	SCL_PIN = PB09 // SCL: SERCOM4/PAD[1]
)

// I2C on the Arduino Nano 33.
var (
	I2C0 = I2C{Bus: sam.SERCOM4_I2CM,
		SDA:     SDA_PIN,
		SCL:     SCL_PIN,
		PinMode: PinSERCOMAlt}
)

// SPI pins
const (
	SPI0_SCK_PIN  = PB11 // SCK: SERCOM4/PAD[3]
	SPI0_MOSI_PIN = PB10 // MOSI: SERCOM4/PAD[2]
	SPI0_MISO_PIN = PA12 // MISO: SERCOM4/PAD[0]
)

// SPI on the Arduino Nano 33.
var (
	SPI0 = SPI{Bus: sam.SERCOM1_SPI}
)

// I2S pins
const (
	I2S_SCK_PIN = PA10
	I2S_SD_PIN  = PA08
	I2S_WS_PIN  = NoPin // TODO: figure out what this is on Arduino Nano 33.
)

// I2S on the Arduino Nano 33.
var (
	I2S0 = I2S{Bus: sam.I2S}
)
