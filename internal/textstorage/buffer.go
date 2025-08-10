package textstorage

import (
	"fmt"
	"slices"
)

type Buffer struct {
	content    []rune
	newLineIdx []int
	changed    bool
	linesCache []string
}

func NewBuffer(content []rune) Buffer {
	return Buffer{
		content:    content,
		newLineIdx: make([]int, 0),
		changed:    false,
		linesCache: make([]string, 0),
	}
}

func (b *Buffer) Insert(r rune, idx int) error {
	b.content = slices.Insert(b.content, idx, r)
	if r == rune('\n') {
		b.newLineIdx = append(b.newLineIdx, idx)
	}
	b.changed = true
	return nil
}

func (b *Buffer) Delete(idx int) error {
	if len(b.content) == 0 {
		return nil
	}
	b.content = slices.Delete(b.content, idx, idx+1)
	b.newLineIdx = slices.DeleteFunc(b.newLineIdx, func(i int) bool { return i == idx })
	b.changed = true
	return nil
}

func (b *Buffer) At(idx int) (rune, error) {
	if idx < 0 || len(b.content) < idx {
		msg := fmt.Sprintf("Buffer.At(%d)", idx)
		return 0, NewErrorOutOfRange(msg)
	}
	value := b.content[idx]
	return value, nil
}

func (b *Buffer) AsLines() *[]string {
	if !b.changed {
		return &b.linesCache
	}

	value := make([]string, 0)

	last := 0
	for _, idx := range b.newLineIdx {
		value = append(value, string(b.content[last:idx]))
		last = idx + 1
	}
	value = append(value, string(b.content[last:]))
	b.changed = false
	b.linesCache = value
	return &b.linesCache
}
