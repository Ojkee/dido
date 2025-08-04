package textstorage

type Buffer struct {
	content []rune
}

func NewBuffer(content []rune) Buffer {
	return Buffer{
		content: content,
	}
}

func (b *Buffer) Insert(r rune, idx int) error {
	return nil
}

func (b *Buffer) Delete(idx int) error {
	panic("TODO")
}

func (b *Buffer) At(idx int) (rune, error) {
	panic("TODO")
}

func (b *Buffer) Get() *[]rune {
	return &b.content
}
