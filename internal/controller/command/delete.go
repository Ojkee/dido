package command

import (
	cursor_api "dido/internal/cursor"
	"dido/internal/textstorage"
)

type Delete struct {
	idx    int
	text   *textstorage.TextStorage
	cursor *cursor_api.Cursor
	r      rune
}

func NewDelete(
	text *textstorage.TextStorage,
	cursor *cursor_api.Cursor,
) *Delete {
	return &Delete{
		idx:    cursor.CurrentPos(),
		text:   text,
		cursor: cursor,
	}
}

func (c *Delete) Execute() error {
	r, err := (*c.text).At(c.idx)
	if err != nil {
		return err
	}
	c.r = r

	err = (*c.text).Delete(c.idx)
	if err != nil {
		return err
	}
	c.cursor.MoveLeft()
	return nil
}

func (c *Delete) Undo() error {
	err := (*c.text).Insert(c.r, c.idx)
	if err != nil {
		return err
	}
	c.cursor.Move(c.idx)
	return nil
}
