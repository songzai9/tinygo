// +build !avr,!nrf,!sam,!stm32

package machine

// Dummy machine package, filled with no-ops.

var (
	SPI0 = SPI{0}
)

type PinMode uint8

const (
	PinInput PinMode = iota
	PinOutput
)

func (p Pin) Configure(config PinConfig) {
	gpioConfigure(p, config)
}

func (p Pin) Set(value bool) {
	gpioSet(p, value)
}

func (p Pin) Get() bool {
	return gpioGet(p)
}

//go:export __tinygo_gpio_configure
func gpioConfigure(pin Pin, config PinConfig)

//go:export __tinygo_gpio_set
func gpioSet(pin Pin, value bool)

//go:export __tinygo_gpio_get
func gpioGet(pin Pin) bool

type SPI struct {
	Bus uint8
}

type SPIConfig struct {
	Frequency uint32
	SCK       Pin
	MOSI      Pin
	MISO      Pin
	Mode      uint8
}

func (spi SPI) Configure(config SPIConfig) {
	if config.SCK == 0 {
		config.SCK = SPI0_SCK_PIN
	}
	if config.MOSI == 0 {
		config.MOSI = SPI0_MOSI_PIN
	}
	if config.MISO == 0 {
		config.MISO = SPI0_MISO_PIN
	}
	spiConfigure(spi.Bus, config.SCK, config.MOSI, config.MISO)
}

// Transfer writes/reads a single byte using the SPI interface.
func (spi SPI) Transfer(w byte) (byte, error) {
	return spiTransfer(spi.Bus, w), nil
}

//go:export __tinygo_spi_configure
func spiConfigure(bus uint8, sck Pin, mosi Pin, miso Pin)

//go:export __tinygo_spi_transfer
func spiTransfer(bus uint8, w uint8) uint8
