package main

import (
	"time"

	"golang.org/x/sys/windows"
)

const (
	KEYEVENTF_KEYDOWN     = 0
	KEYEVENTF_KEYUP       = 2
	KEYEVENTF_EXTENDEDKEY = 1
	VK_CAPITAL            = 0x14
)

var (
	user32         = windows.NewLazySystemDLL("user32.dll")
	procKeybdEvent = user32.NewProc("keybd_event")
)

func keybd_event(bVk byte, bScan byte, dwFlags uint32, dwExtraInfo uintptr) {
	procKeybdEvent.Call(uintptr(bVk), uintptr(bScan), uintptr(dwFlags), dwExtraInfo)
}

func pressWindowsSpace() {
	// Press Windows Key
	keybd_event(0x5B, 0, KEYEVENTF_KEYDOWN, 0)

	// Pause for a short time (you may adjust this if needed)
	time.Sleep(50 * time.Millisecond)

	// Press Spacebar
	keybd_event(0x20, 0, KEYEVENTF_KEYDOWN, 0)

	// Release Spacebar
	keybd_event(0x20, 0, KEYEVENTF_KEYUP, 0)

	// Release Windows Key
	keybd_event(0x5B, 0, KEYEVENTF_KEYUP, 0)
}

func toggleCapsLock() {
	// Simulate Caps Lock key press and release
	keybd_event(VK_CAPITAL, 0, KEYEVENTF_KEYDOWN, 0)
	keybd_event(VK_CAPITAL, 0, KEYEVENTF_KEYUP, 0)
}

var user32_dll = windows.NewLazyDLL("user32.dll")
var GetKeyState = user32_dll.NewProc("GetKeyState")

func isCapsLockOn() bool {
	state, _, _ := GetKeyState.Call(VK_CAPITAL)
	return state&1 == 1
}

// func loop() {
// }

func main() {
	for {
		// Do something...
		if isCapsLockOn() {
			pressWindowsSpace()
			toggleCapsLock()
			//break
		}

		// Do something...
		time.Sleep(time.Millisecond * 100)
	}
	// loop()
}
