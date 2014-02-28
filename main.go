package main

import (
	. "zephyros_go"
)

func main() {
	Bind("D", []string{"Cmd", "Shift"}, func() {
		win := FocusedWindow()
		frame := win.Frame()
		frame.X += 10
		win.SetFrame(frame)
	})

	ListenForCallbacks()
}
