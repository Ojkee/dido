package controller

import (
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
	case *sdl.KeyboardEvent:
		cmd = GetKeyCommand(e, text, cursor)
	default:
		cmd = command.NewCommandNone()
	}

	return cmd
}

func GetKeyCommand(
	event *sdl.KeyboardEvent,
	text textstorage.TextStorage,
	cursor *cursor_api.Cursor,
) command.Command {
	switch event.GetType() {
	case sdl.KEYDOWN:
		key := event.Keysym.Sym
		if 'a' <= key && key <= 'z' {
			return command.NewCommandInsert(&text, rune(key), cursor)
		}
	}
	return command.NewCommandNone()
}
