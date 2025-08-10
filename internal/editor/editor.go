package editor

import (
	"github.com/veandco/go-sdl2/sdl"

	config_api "dido/internal/config"
	"dido/internal/context"
	controller_api "dido/internal/controller"
	"dido/internal/controller/command"
	view_api "dido/internal/view"
)

type Editor struct {
	config     *config_api.Config
	view       view_api.View
	controller controller_api.Controller
	run        bool
	ctx        context.Context
}

func NewEditor() Editor {
	config := config_api.NewConfig()
	return Editor{
		config:     config,
		view:       view_api.NewView(config),
		controller: controller_api.NewController(),
		run:        true,
		ctx:        context.NewContext(),
	}
}

func (e *Editor) Close() {
	e.view.Close()
}

func (e *Editor) Run() {
	for e.run {
		event := sdl.WaitEvent()

		cmd := e.controller.Command(&e.ctx, event)
		switch cmd.(type) {
		case *command.Quit:
			e.run = false
		case *command.None:
			break
		default:
			cmd.Execute()
		}

		e.view.Draw(&e.ctx)
	}
}
