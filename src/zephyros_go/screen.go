package zephyros_go

import (
	"encoding/json"
)

type Screen float64

func (self Screen) FrameIncludingDockAndMenu() Rect {
	var buf Rect
	bytes := send(float64(self), nil, false, "frame_including_dock_and_menu")
	json.Unmarshal(bytes, &buf)
	return buf
}

func (self Screen) FrameWithoutDockOrMenu() Rect {
	var buf Rect
	bytes := send(float64(self), nil, false, "frame_without_dock_or_menu")
	json.Unmarshal(bytes, &buf)
	return buf
}

func (self Screen) NextScreen() Screen {
	var buf Screen
	bytes := send(float64(self), nil, false, "next_screen")
	json.Unmarshal(bytes, &buf)
	return buf
}

func (self Screen) PreviousScreen() Screen {
	var buf Screen
	bytes := send(float64(self), nil, false, "previous_screen")
	json.Unmarshal(bytes, &buf)
	return buf
}

// Can only be 0, 90, 180, or 270
func (self Screen) RotateTo(degrees int) {
	send(float64(self), nil, false, "rotate_to", degrees)
}
