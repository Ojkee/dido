package command

import (
	"dido/internal/context"
)

type Delete struct {
	idx int
	ctx *context.Context
	r   rune
}

func NewDelete(ctx *context.Context) *Delete {
	return &Delete{
		idx: ctx.Cursor.CurrentPos() - 1,
		ctx: ctx,
	}
}

func (c *Delete) Execute() error {
	r, err := c.ctx.Buffer.At(c.idx)
	if err != nil {
		return err
	}
	c.r = r

	err = c.ctx.Buffer.Delete(c.idx)
	if err != nil {
		return err
	}
	c.ctx.Cursor.MoveLeft()
	return nil
}

func (c *Delete) Undo() error {
	err := c.ctx.Buffer.Insert(c.r, c.idx)
	if err != nil {
		return err
	}
	c.ctx.Cursor.Move(c.idx)
	return nil
}
