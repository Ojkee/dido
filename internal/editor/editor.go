package editor

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"

	controller_api "dido/internal/controller"
	"dido/internal/controller/command"
	cursor_api "dido/internal/cursor"
	"dido/internal/textstorage"
	view_api "dido/internal/view"
)

type Editor struct {
	view       view_api.View
	controller controller_api.Controller
	run        bool
	buffer     textstorage.Buffer
	cursor     cursor_api.Cursor
}

func NewEditor() Editor {
	return Editor{
		view:       view_api.NewView(),
		controller: controller_api.NewController(),
		run:        true,
		buffer:     textstorage.NewBuffer([]rune{}),
		cursor:     cursor_api.NewCursor(),
	}
}

func (e *Editor) Close() {
	e.view.Close()
}

func (e *Editor) Run() {
	for e.run {
		event := sdl.WaitEvent()

		cmd := e.controller.GetCommand(event, &e.buffer, &e.cursor)
		switch cmd.(type) {
		case *command.CommandQuit:
			e.run = false
		case *command.CommandNone:
			break
		default:
			cmd.Execute()
			fmt.Println(e.buffer) // TODO: remove
		}

		e.view.Draw()
		e.view.Update()
	}
}
