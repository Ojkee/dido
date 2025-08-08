package controller

import (
	"bytes"

	"github.com/veandco/go-sdl2/sdl"

	"dido/internal/controller/command"
	cursor_api "dido/internal/cursor"
	"dido/internal/textstorage"
)

type Controller struct {
	DEBUG_KEY sdl.Keycode
}

func NewController() Controller {
	return Controller{
		DEBUG_KEY: sdl.K_F4,
	}
}

func (c *Controller) GetCommand(
	event sdl.Event,
	text textstorage.TextStorage,
	cursor *cursor_api.Cursor,
) command.Command {
	switch e := event.(type) {
	case *sdl.QuitEvent:
		return command.NewQuit()
	case *sdl.KeyboardEvent:
		if cmd := c.specialSignCommand(e, text, cursor); cmd != nil {
			return cmd
		}
	case *sdl.TextInputEvent:
		return command.NewInsert(runeOfBytes(e.Text), &text, cursor)
	default:
		return command.NewNone()
	}

	return command.NewNone()
}

func runeOfBytes(b [32]byte) rune {
	return bytes.Runes(b[:3])[0] // rune <-> `3` UTF8 bytes
}

func (c *Controller) specialSignCommand(
	event *sdl.KeyboardEvent,
	text textstorage.TextStorage,
	cursor *cursor_api.Cursor,
) command.Command {
	commandMap := map[sdl.Keycode]command.Command{
		sdl.K_RETURN:    command.NewInsert('\n', &text, cursor),
		sdl.K_TAB:       command.NewInsert('\t', &text, cursor),
		sdl.K_BACKSPACE: command.NewDelete(&text, cursor),
		c.DEBUG_KEY:     command.NewLog(*text.Get()),
	}

	switch event.GetType() {
	case sdl.KEYDOWN:
		key := event.Keysym.Sym
		if cmd, ok := commandMap[key]; ok {
			return cmd
		}
	}
	return nil
}
