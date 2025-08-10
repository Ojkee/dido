package textstorage

import "fmt"

type TextStorage interface {
	Insert(rune, int) error
	Delete(int) error
	At(int) (rune, error)
	Get() *[]rune
	AsLines() *[]string
}

type ErrorOutOfRange struct {
	msg string
}

func NewErrorOutOfRange(msg string) ErrorOutOfRange {
	return ErrorOutOfRange{
		msg: msg,
	}
}

func (e ErrorOutOfRange) Error() string {
	return fmt.Sprintf("Out of range at: %s", e.msg)
}
