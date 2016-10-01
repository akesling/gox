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

type Display C.Display
type Window C.Window
type Event C.XEvent

func Sync(display *Display, discard bool) int {
	if discard {
		return int(C.XSync((*C.Display)(display), C.Bool(1)))
	} else {
		return int(C.XSync((*C.Display)(display), C.Bool(0)))
	}
}

func NextEvent(display *Display, event *Event) int {
	return int(C.XNextEvent((*C.Display)(display), (*C.XEvent)(event)))
}

func OpenDisplay(displayName *string) *Display {
	var cDisplayName *C.char;
	if (displayName == nil) {
		cDisplayName = nil 
	} else {
		cDisplayName = C.CString(*displayName)
	}

	return (*Display)(C.XOpenDisplay(cDisplayName));
}

func CloseDisplay(display *Display) int {
	return int(C.XCloseDisplay(display));
}

func SupportsLocale() int {
	return int(C.XSupportsLocale())
}

func SelectInput(display *Display, window Window, eventMask int64) int {
	return int(C.XSelectInput(
		(*C.Display)(display), (C.Window)(window), C.long(eventMask)))
}
