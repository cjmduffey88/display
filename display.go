package display

import (
	"github.com/gonutz/prototype/draw"
)

type Display struct {
	buffer     []draw.Color
	width      int
	height     int
	multiplier int
	fullscreen bool
	refresh    bool
}

func NewDisplay(width, height, multiplier uint, fullscreen bool) *Display {
	display := new(Display)
	display.width = int(width)
	display.height = int(height)
	display.multiplier = int(multiplier)
	display.fullscreen = fullscreen
	display.buffer = make([]draw.Color, width*height)
	return display
}

func (display *Display) Run(title string) error {
	return draw.RunWindow(title, display.width*display.multiplier, display.height*display.multiplier, display.update)
}

func (display *Display) SetPixelAt(x, y int, color draw.Color) {
	display.buffer[x+y*display.width] = color
	display.refresh = true
}

func (display *Display) update(window draw.Window) {
	if display.fullscreen {
		window.SetFullscreen(true)
	}
	if display.refresh {
		for i := 0; i < len(display.buffer); i++ {
			x := (i % display.width) * display.multiplier
			y := (i / display.width) * display.multiplier
			window.FillRect(x, y, display.multiplier, display.multiplier, display.buffer[i])
		}
		display.refresh = false
	}
}
