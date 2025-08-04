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

func (ci *CommandInsert) Execute() error {
	err := (*ci.text).Insert(ci.r, ci.idx)
	return err
}

func (ci *CommandInsert) Undo() error {
	err := (*ci.text).Delete(ci.idx)
	return err
}
