package controller

import (
	"bytes"

	"github.com/veandco/go-sdl2/sdl"

	"dido/internal/controller/command"
	cursor_api "dido/internal/cursor"
	"dido/internal/textstorage"
)

type Controller struct{}

func NewController() Controller {
	return Controller{}
}

func (*Controller) GetCommand(
	event sdl.Event,
	text textstorage.TextStorage,
	cursor *cursor_api.Cursor,
) command.Command {
	var cmd command.Command
	switch e := event.(type) {
	case *sdl.QuitEvent:
		cmd = command.NewCommandQuit()
	case *sdl.TextInputEvent:
		letter := bytes.Runes(e.Text[:1])[0]
		cmd = command.NewCommandInsert(letter, &text, cursor)
	default:
		cmd = command.NewCommandNone()
	}

	return cmd
}
