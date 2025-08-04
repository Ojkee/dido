package textstorage

type TextStorage interface {
	Insert(rune, int) error
	Delete(int) error
	At(int) (rune, error)
	Get() *[]rune
}
