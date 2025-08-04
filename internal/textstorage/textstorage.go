package textstorage

import "fmt"

type TextStorage interface {
	Insert(rune, int) error
	Delete(int) error
	At(int) (rune, error)
	Get() *[]rune
}

type ErrorOutOfRange struct {
	ctx string
}

func NewErrorOutOfRange(ctx string) ErrorOutOfRange {
	return ErrorOutOfRange{
		ctx: ctx,
	}
}

func (e ErrorOutOfRange) Error() string {
	return fmt.Sprintf("Out of range at: %s", e.ctx)
}
