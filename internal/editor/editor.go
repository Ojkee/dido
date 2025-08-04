package editor

import (
	"github.com/veandco/go-sdl2/sdl"

	controller_api "dido/internal/controller"
)

type Editor struct {
	window     *sdl.Window
	controller controller_api.Controller
	run        bool
}

func NewEditor() Editor {
	window, err := sdl.CreateWindow(
		"Dido",
		sdl.WINDOWPOS_CENTERED,
		sdl.WINDOWPOS_CENTERED,
		800, 600,
		sdl.WINDOW_SHOWN,
	)
	if err != nil {
		panic(err)
	}
	return Editor{
		window:     window,
		controller: controller_api.NewController(),
	}
}

func (e *Editor) Close() {
	e.window.Destroy()
}

func (e *Editor) Display() {
	for e.run {
		event := sdl.PollEvent()
		_ = e.controller.GetCommand(&event) // _ stands for command
	}
}
