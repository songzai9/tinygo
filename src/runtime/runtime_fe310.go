// +build fe310

// This file implements target-specific things for the FE310 chip as used in the
// HiFive1.

package runtime

import (
	"machine"
	"unsafe"

	"device/riscv"
	"device/sifive"
)

type timeUnit int64

const tickMicros = 1024 * 32

//go:extern _sbss
var _sbss unsafe.Pointer

//go:extern _ebss
var _ebss unsafe.Pointer

//go:extern _sdata
var _sdata unsafe.Pointer

//go:extern _sidata
var _sidata unsafe.Pointer

//go:extern _edata
var _edata unsafe.Pointer

//go:export main
func main() {
	preinit()
	initAll()
	callMain()
	abort()
}

func init() {
	pric_init()
	machine.UART0.Configure(machine.UARTConfig{})
}

func pric_init() {
	// Make sure the HFROSC is on
	sifive.PRIC.HFROSCCFG.SetBits(sifive.PRIC_HFROSCCFG_ENABLE)

	// Run off 16 MHz Crystal for accuracy.
	sifive.PRIC.PLLCFG.SetBits(sifive.PRIC_PLLCFG_REFSEL | sifive.PRIC_PLLCFG_BYPASS)
	sifive.PRIC.PLLCFG.SetBits(sifive.PRIC_PLLCFG_SEL)

	// Turn off HFROSC to save power
	sifive.PRIC.HFROSCCFG.ClearBits(sifive.PRIC_HFROSCCFG_ENABLE)
}

func preinit() {
	// Initialize .bss: zero-initialized global variables.
	ptr := uintptr(unsafe.Pointer(&_sbss))
	for ptr != uintptr(unsafe.Pointer(&_ebss)) {
		*(*uint32)(unsafe.Pointer(ptr)) = 0
		ptr += 4
	}

	// Initialize .data: global variables initialized from flash.
	src := uintptr(unsafe.Pointer(&_sidata))
	dst := uintptr(unsafe.Pointer(&_sdata))
	for dst != uintptr(unsafe.Pointer(&_edata)) {
		*(*uint32)(unsafe.Pointer(dst)) = *(*uint32)(unsafe.Pointer(src))
		dst += 4
		src += 4
	}
}

func putchar(c byte) {
	machine.UART0.WriteByte(c)
}

func ticks() timeUnit {
	return 0
}

const asyncScheduler = false

func sleepTicks(d timeUnit) {
}

func abort() {
	// lock up forever
	for {
		riscv.Asm("wfi")
	}
}
