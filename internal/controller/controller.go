package controller

import (
	"bytes"

	"github.com/veandco/go-sdl2/sdl"

	"dido/internal/context"
	"dido/internal/controller/command"
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
	ctx *context.Context,
	event sdl.Event,
) command.Command {
	switch e := event.(type) {
	case *sdl.QuitEvent:
		return command.NewQuit()
	case *sdl.KeyboardEvent:
		if cmd := c.specialSignCommand(ctx, e); cmd != nil {
			return cmd
		}
	case *sdl.TextInputEvent:
		return command.NewInsert(ctx, runeOfBytes(e.Text))
	default:
		return command.NewNone()
	}

	return command.NewNone()
}

func runeOfBytes(b [32]byte) rune {
	return bytes.Runes(b[:3])[0] // rune <-> `3` UTF8 bytes
}

func (c *Controller) specialSignCommand(
	ctx *context.Context,
	event *sdl.KeyboardEvent,
) command.Command {
	commandMap := map[sdl.Keycode]command.Command{
		sdl.K_RETURN:    command.NewInsert(ctx, '\n'),
		sdl.K_TAB:       command.NewInsert(ctx, '\t'),
		sdl.K_BACKSPACE: command.NewDelete(ctx),
		c.DEBUG_KEY:     command.NewLog(*ctx.Buffer.Get()),
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
