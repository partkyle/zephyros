package zephyros_go

import (
	"encoding/json"
)

type Window float64

func (self Window) Title() string {
	var buf string
	bytes := send(float64(self), nil, false, "title")
	json.Unmarshal(bytes, &buf)
	return buf
}

func (self Window) Frame() Rect {
	var buf Rect
	bytes := send(float64(self), nil, false, "frame")
	json.Unmarshal(bytes, &buf)
	return buf
}

func (self Window) Size() Size {
	var buf Size
	bytes := send(float64(self), nil, false, "size")
	json.Unmarshal(bytes, &buf)
	return buf
}

func (self Window) TopLeft() TopLeft {
	var buf TopLeft
	bytes := send(float64(self), nil, false, "top_left")
	json.Unmarshal(bytes, &buf)
	return buf
}

func (self Window) SetFrame(f Rect) {
	send(float64(self), nil, false, "set_frame", f)
}

func (self Window) SetSize(f Size) {
	send(float64(self), nil, false, "set_size", f)
}

func (self Window) SetTopLeft(f TopLeft) {
	send(float64(self), nil, false, "set_top_left", f)
}

func (self Window) Maximize() {
	send(float64(self), nil, false, "maximize")
}

func (self Window) Minimize() {
	send(float64(self), nil, false, "minimize")
}

func (self Window) UnMinimize() {
	send(float64(self), nil, false, "un_minimize")
}

func (self Window) App() App {
	var buf App
	bytes := send(float64(self), nil, false, "app")
	json.Unmarshal(bytes, &buf)
	return buf
}

func (self Window) Screen() Screen {
	var buf Screen
	bytes := send(float64(self), nil, false, "screen")
	json.Unmarshal(bytes, &buf)
	return buf
}

func (self Window) FocusWindow() bool {
	var buf bool
	bytes := send(float64(self), nil, false, "focus_window")
	json.Unmarshal(bytes, &buf)
	return buf
}

func (self Window) FocusWindowDown() {
	send(float64(self), nil, false, "focus_window_down")
}

func (self Window) FocusWindowUp() {
	send(float64(self), nil, false, "focus_window_up")
}

func (self Window) FocusWindowLeft() {
	send(float64(self), nil, false, "focus_window_left")
}

func (self Window) FocusWindowRight() {
	send(float64(self), nil, false, "focus_window_right")
}

func (self Window) IsNormal() bool {
	var buf bool
	bytes := send(float64(self), nil, false, "normal_window?")
	json.Unmarshal(bytes, &buf)
	return buf
}

func (self Window) IsMinimized() bool {
	var buf bool
	bytes := send(float64(self), nil, false, "minimized?")
	json.Unmarshal(bytes, &buf)
	return buf
}

func (self Window) OtherWindowsOnSameScreen() []Window {
	var buf []Window
	bytes := send(float64(self), nil, false, "other_windows_on_same_screen")
	json.Unmarshal(bytes, &buf)
	return buf
}

func (self Window) OtherWindowsOnAllScreens() []Window {
	var buf []Window
	bytes := send(float64(self), nil, false, "other_windows_on_all_screens")
	json.Unmarshal(bytes, &buf)
	return buf
}

func (self Window) WindowsToNorth() []Window {
	var buf []Window
	bytes := send(float64(self), nil, false, "windows_to_north")
	json.Unmarshal(bytes, &buf)
	return buf
}

func (self Window) WindowsToSouth() []Window {
	var buf []Window
	bytes := send(float64(self), nil, false, "windows_to_south")
	json.Unmarshal(bytes, &buf)
	return buf
}

func (self Window) WindowsToWest() []Window {
	var buf []Window
	bytes := send(float64(self), nil, false, "windows_to_west")
	json.Unmarshal(bytes, &buf)
	return buf
}

func (self Window) WindowsToEast() []Window {
	var buf []Window
	bytes := send(float64(self), nil, false, "windows_to_east")
	json.Unmarshal(bytes, &buf)
	return buf
}
