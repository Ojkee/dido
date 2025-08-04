package command

type CommandNone struct{}

func NewCommandNone() CommandNone {
	return CommandNone{}
}

func (cn *CommandNone) Execute() error { return nil }
func (cn *CommandNone) Undo() error    { return nil }
