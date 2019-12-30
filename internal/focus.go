package internal

import (
	"syscall/js"
	"time"
)

func RequestFocusEvent(id string) {
	//we try for a total of 1.2s (1200ms)
	for i := 0; i < 15; i++ {
		time.Sleep(10 * time.Duration(i) * time.Millisecond)
		d := js.Global().Get("document").Call("getElementById", id)
		if js.Null() != d.JSValue() {
			d.Call("focus")
		}
		active := js.Global().Get("document").Get("activeElement")
		if active != js.Null() && d == active {
			break
		}
	}
}
