package textstorage

import (
	"fmt"
	"slices"
)

type Buffer struct {
	content []rune
}

func NewBuffer(content []rune) Buffer {
	return Buffer{
		content: content,
	}
}

func (b *Buffer) Insert(r rune, idx int) error {
	b.content = slices.Insert(b.content, idx, r)
	return nil
}

func (b *Buffer) Delete(idx int) error {
	b.content = slices.Delete(b.content, idx, idx)
	return nil
}

func (b *Buffer) At(idx int) (rune, error) {
	if idx < 0 || len(b.content) < idx {
		ctx := fmt.Sprintf("Buffer.At(%d)", idx)
		return 0, NewErrorOutOfRange(ctx)
	}
	value := b.content[idx]
	return value, nil
}

func (b *Buffer) Get() *[]rune {
	return &b.content
}
