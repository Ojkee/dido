package command

type Command interface {
	Execute() error
	Undo() error
}
