package controller

import (
	"bytes"

	"github.com/veandco/go-sdl2/sdl"

	"dido/internal/context"
	"dido/internal/controller/command"
)

type Controller struct {
	debugKey sdl.Keycode
}

func NewController() Controller {
	return Controller{
		debugKey: sdl.K_F4,
	}
}

func (c *Controller) Command(
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
	switch event.GetType() {
	case sdl.KEYDOWN:
		switch event.Keysym.Sym {
		case sdl.K_RETURN:
			return command.NewInsert(ctx, '\n')
		case sdl.K_TAB:
			return command.NewInsert(ctx, '\t')
		case sdl.K_BACKSPACE:
			return command.NewDelete(ctx)
		case c.debugKey:
			return command.NewLog(*ctx.Buffer.AsLines())
		}
	}
	return nil
}
