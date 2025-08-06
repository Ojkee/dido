package command

import "dido/internal/textstorage"

type CommandInsert struct {
	r    rune
	idx  int
	text *textstorage.TextStorage
}

func NewCommandInsert(text *textstorage.TextStorage, r rune, idx int) CommandInsert {
	return CommandInsert{
		r:    r,
		idx:  idx,
		text: text,
	}
}

func (c *CommandInsert) Execute() error {
	err := (*c.text).Insert(c.r, c.idx)
	return err
}

func (c *CommandInsert) Undo() error {
	err := (*c.text).Delete(c.idx)
	return err
}
