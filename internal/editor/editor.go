package editor

import (
	"github.com/veandco/go-sdl2/sdl"

	controller_api "dido/internal/controller"
	"dido/internal/controller/command"
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
		run:        true,
	}
}

func (e *Editor) Close() {
	e.window.Destroy()
}

func (e *Editor) Run() {
	for e.run {
		event := sdl.PollEvent()
		cmd := e.controller.GetCommand(event)
		switch cmd.(type) {
		case *command.CommandQuit:
			e.run = false
		default:
			cmd.Execute()
		}
	}
}
