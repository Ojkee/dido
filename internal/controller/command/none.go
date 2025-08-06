package command

type CommandNone struct{}

func NewCommandNone() *CommandNone {
	return &CommandNone{}
}

func (c *CommandNone) Execute() error { return nil }
func (c *CommandNone) Undo() error    { return nil }
