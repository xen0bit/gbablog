package main

import (
	"image/color"
	"machine"
	"runtime/interrupt"
	"runtime/volatile"
	"unsafe"

	"tinygo.org/x/tinydraw"
	"tinygo.org/x/tinyfont"
)

var (
	//display from machine
	display = machine.Display
	//screen resolution
	screenW, screenH = display.Size()

	//Register for display
	regDISPSTAT = (*volatile.Register16)(unsafe.Pointer(uintptr(0x4000004)))
	//Register for Keypad
	regKEYPAD = (*volatile.Register16)(unsafe.Pointer(uintptr(0x04000130)))

	//KeyCodes
	keyDOWN      = uint16(895)
	keyUP        = uint16(959)
	keyLEFT      = uint16(991)
	keyRIGHT     = uint16(1007)
	keyLSHOULDER = uint16(511)
	keyRSHOULDER = uint16(767)
	keyA         = uint16(1022)
	keyB         = uint16(1021)
	keySTART     = uint16(1015)
	keySELECT    = uint16(1019)

	//Sample Colors
	black = color.RGBA{}
	red   = color.RGBA{R: 255}
	green = color.RGBA{G: 255}
)

func clearScreen() {
	tinydraw.FilledRectangle(
		display,
		int16(0), int16(0),
		screenW, screenH,
		black,
	)
}

func writeTextCentered(buttonName string) {
	tinyfont.WriteLine(
		display, &tinyfont.TomThumb,
		screenW/2, screenH/2,
		buttonName, green,
	)

}

func update(interrupt.Interrupt) {
	//Read uint16 from register that represents the state of current buttons pressed
	switch keyValue := regKEYPAD.Get(); keyValue {
	case keyDOWN:
		writeTextCentered("Down D-Pad")
	case keyUP:
		writeTextCentered("Up D-Pad")
	case keyLEFT:
		writeTextCentered("Left D-Pad")
	case keyRIGHT:
		writeTextCentered("Right D-Pad")
	case keyLSHOULDER:
		writeTextCentered("Left Shoulder Button")
	case keyRSHOULDER:
		writeTextCentered("Right Shoulder Button")
	case keyA:
		writeTextCentered("Button A")
	case keyB:
		writeTextCentered("Button B")
	case keySTART:
		writeTextCentered("Button START")
	case keySELECT:
		writeTextCentered("Button SELECT")
	//Default is to clear the screen for the next round
	default:
		clearScreen()
	}
}

func main() {
	// Use video mode 3 (in BG2, a 16bpp bitmap in VRAM)
	// Enable BG2 (BG0 = 1, BG1 = 2, BG2 = 4, ...)
	display.Configure()
	regDISPSTAT.SetBits(1<<3 | 1<<4)

	//Register IRQ_VBLANK iterrupt that will call "update" every time it's fired
	interrupt.New(machine.IRQ_VBLANK, update).Enable()

	//Prevent game from exiting
	for {
	}
}
