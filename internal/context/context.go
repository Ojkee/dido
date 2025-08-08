package context

import (
	cursor_api "dido/internal/cursor"
	"dido/internal/textstorage"
)

type Context struct {
	Buffer textstorage.Buffer
	Cursor cursor_api.Cursor
}

func NewContext() Context {
	return Context{
		Buffer: textstorage.NewBuffer([]rune{}),
		Cursor: cursor_api.NewCursor(),
	}
}
