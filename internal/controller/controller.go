package controller

import (
	"github.com/veandco/go-sdl2/sdl"

	"dido/internal/controller/command"
)

type Controller struct{}

func NewController() Controller {
	return Controller{}
}

func (*Controller) GetCommand(event sdl.Event) (cmd command.Command) {
	switch event.(type) {
	case *sdl.QuitEvent:
		cmd = &command.CommandQuit{}
	default:
		cmd = &command.CommandNone{}
	}

	return
}
