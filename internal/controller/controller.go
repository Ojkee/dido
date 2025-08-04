package controller

import (
	"github.com/veandco/go-sdl2/sdl"

	"dido/internal/controller/command"
)

type Controller struct{}

func NewController() Controller {
	return Controller{}
}

func (*Controller) GetCommand(event *sdl.Event) command.Command {
	var cmd command.Command
	cmd = &command.CommandNone{}

	return cmd
}
