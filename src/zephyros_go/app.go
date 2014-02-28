package zephyros_go

import (
	"encoding/json"
)

type App float64

func (self App) AllWindows() []Window {
	var buf []Window
	bytes := send(float64(self), nil, false, "all_windows")
	json.Unmarshal(bytes, &buf)
	return buf
}

func (self App) VisibleWindows() []Window {
	var buf []Window
	bytes := send(float64(self), nil, false, "visible_windows")
	json.Unmarshal(bytes, &buf)
	return buf
}

func (self App) Title() string {
	var buf string
	bytes := send(float64(self), nil, false, "title")
	json.Unmarshal(bytes, &buf)
	return buf
}

func (self App) IsHidden() bool {
	var buf bool
	bytes := send(float64(self), nil, false, "hidden?")
	json.Unmarshal(bytes, &buf)
	return buf
}

func (self App) Show() {
	send(float64(self), nil, false, "show")
}

func (self App) Hide() {
	send(float64(self), nil, false, "hide")
}

func (self App) Kill() {
	send(float64(self), nil, false, "kill")
}

func (self App) Kill9() {
	send(float64(self), nil, false, "kill9")
}
