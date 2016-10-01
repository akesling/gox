package X

/*
#cgo CFLAGS: -std=c99 -pedantic -Wno-deprecated-declarations -Wno-unused -Wno-unused-parameter
#cgo CFLAGS: -Os -I/usr/X11R6/include
#cgo CFLAGS: -D_BSD_SOURCE -D_POSIX_C_SOURCE=2 -DXINERAMA

#cgo LDFLAGS: -s -L/usr/X11R6/lib -lX11 -lXinerama

#include <X11/cursorfont.h>
#include <X11/keysym.h>
#include <X11/Xatom.h>
#include <X11/Xlib.h>
#include <X11/Xproto.h>
#include <X11/Xutil.h>
*/
import "C"

import (
	"bytes"
	"encoding/binary"
	"unsafe"
)

type Display C.Display
type Window C.Window
type Event C.XEvent

func OpenDisplay(displayName *string) *Display {
	var cDisplayName *C.char;
	if (displayName == nil) {
		cDisplayName = nil 
	} else {
		cDisplayName = C.CString(*displayName)
	}

	return (*Display)(C.XOpenDisplay(cDisplayName));
}

func (display *Display) CloseDisplay() int {
	return int(C.XCloseDisplay(display));
}

func (display *Display) Sync(discard bool) int {
	if discard {
		return int(C.XSync((*C.Display)(display), C.Bool(1)))
	} else {
		return int(C.XSync((*C.Display)(display), C.Bool(0)))
	}
}

func (display *Display) NextEvent(event *Event) int {
	return int(C.XNextEvent((*C.Display)(display), (*C.XEvent)(event)))
}

func (display *Display) SelectInput(window Window, eventMask int64) int {
	return int(C.XSelectInput(
		(*C.Display)(display), (C.Window)(window), C.long(eventMask)))
}

func (event Event) EventType() int {
	var eventType C.int
	binary.Read(
		bytes.NewBuffer(event[:unsafe.Sizeof(eventType)]),
		binary.LittleEndian, &eventType)
	return int(eventType)
}

func SupportsLocale() int {
	return int(C.XSupportsLocale())
}
