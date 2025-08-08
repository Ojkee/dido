package command

import (
	"dido/internal/context"
)

type Insert struct {
	idx int
	r   rune
	ctx *context.Context
}

func NewInsert(
	ctx *context.Context,
	r rune,
) *Insert {
	return &Insert{
		idx: ctx.Cursor.CurrentPos(),
		r:   r,
		ctx: ctx,
	}
}

func (c *Insert) Execute() error {
	err := c.ctx.Buffer.Insert(c.r, c.idx)
	if err != nil {
		return err
	}
	c.ctx.Cursor.MoveRight()
	return nil
}

func (c *Insert) Undo() error {
	err := c.ctx.Buffer.Delete(c.idx)
	if err != nil {
		return err
	}
	c.ctx.Cursor.Move(c.idx)
	return nil
}
