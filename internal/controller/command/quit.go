package command

type CommandQuit struct{}

func NewCommandQuit() *CommandQuit {
	return &CommandQuit{}
}

func (c *CommandQuit) Execute() error { return nil }
func (c *CommandQuit) Undo() error    { return nil }
