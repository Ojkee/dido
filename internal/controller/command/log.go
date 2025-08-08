package command

import (
	"log"
)

type Log struct {
	content any
}

func NewLog(content ...any) *Log {
	return &Log{
		content: content,
	}
}

func (c *Log) Execute() error {
	log.Println(c.content)
	return nil
}

func (c *Log) Undo() error { return nil }
