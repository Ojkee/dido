package command

import (
	cursor_api "dido/internal/cursor"
	"dido/internal/textstorage"
)

type CommandInsert struct {
	r      rune
	idx    int
	text   *textstorage.TextStorage
	cursor *cursor_api.Cursor
}

func NewCommandInsert(
	text *textstorage.TextStorage,
	r rune,
	cursor *cursor_api.Cursor,
) *CommandInsert {
	return &CommandInsert{
		r:      r,
		idx:    cursor.CurrentPos(),
		text:   text,
		cursor: cursor,
	}
}

func (c *CommandInsert) Execute() error {
	err := (*c.text).Insert(c.r, c.idx)
	if err != nil {
		return err
	}
	c.cursor.MoveRight()
	return nil
}

func (c *CommandInsert) Undo() error {
	err := (*c.text).Delete(c.idx)
	if err != nil {
		return err
	}
	c.cursor.Move(c.idx)
	return nil
}
